package sherlock

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/google_workspace"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/suitability_synchronization"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

type UserV3DeactivateRequest struct {
	UserEmails                                      []string `json:"userEmails"`
	UserEmailSubstitutableDomains                   []string `json:"userEmailSubstitutableDomains" default:"[\"broadinstitute.org\"]"` // Domains of UserEmails that should be swapped out to match Google Workspace domains
	SuspendEmailHandlesAcrossGoogleWorkspaceDomains []string `json:"suspendEmailHandlesAcrossGoogleWorkspaceDomains"`
}

type UserV3DeactivateResponse struct {
	NewlyDeactivatedEmails   []string `json:"newlyDeactivatedEmails"`
	AlreadyDeactivatedEmails []string `json:"alreadyDeactivatedEmails"`
	NotFoundEmails           []string `json:"notFoundEmails"`
}

// usersProceduresV3Deactivate godoc
//
//	@summary		Deactivate Users
//	@description	Super-admin only method to deactivate users. Deactivated users will be removed from all roles and can't authenticate to Sherlock.
//	@description	This endpoint can optionally also attempt to suspend the same email handles across given Google Workspace domains, substituting email domains as necessary.
//	@description	It will do so by impersonating the caller in each given domain.
//	@tags			Users
//	@accept			json
//	@produce		json
//	@param			users					body		UserV3DeactivateRequest	true	"Information on the users to deactivate"
//	@success		200						{object}	UserV3DeactivateResponse
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/users/procedures/v3/deactivate [post]
func usersProceduresV3Deactivate(ctx *gin.Context) {
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body UserV3DeactivateRequest
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	if err = defaults.Set(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error setting defaults: %w", err))
		return
	}

	callingUser, err := authentication.MustUseUser(ctx)
	if err != nil {
		return
	}

	if err = callingUser.ErrIfNotSuperAdmin(); err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	if len(body.UserEmails) == 0 {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) no user emails provided", errors.BadRequest))
		return
	}

	// Make a set of the emails to casually de-dupe and then keep track of who we can't find at all
	emailsToAttemptToDeactivate := make(map[string]struct{})
	for _, email := range body.UserEmails {
		if email == "" {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) empty email provided", errors.BadRequest))
			return
		}
		emailsToAttemptToDeactivate[email] = struct{}{}
	}
	// We need a slice again for later, unfortunately
	sliceOfDedupedEmailsToAttemptToDeactivate := make([]string, 0, len(emailsToAttemptToDeactivate))
	for emailToAttemptToDeactivate := range emailsToAttemptToDeactivate {
		sliceOfDedupedEmailsToAttemptToDeactivate = append(sliceOfDedupedEmailsToAttemptToDeactivate, emailToAttemptToDeactivate)
	}

	if _, present := emailsToAttemptToDeactivate[callingUser.Email]; present {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) you cannot deactivate yourself", errors.BadRequest))
		return
	}

	googleWorkspaceDomainsToSuspendIn := make(map[string]google_workspace.WorkspaceClient)
	for _, domain := range body.SuspendEmailHandlesAcrossGoogleWorkspaceDomains {
		if domain == "" {
			errors.AbortRequest(ctx, fmt.Errorf("(%s) empty Google Workspace domain provided", errors.BadRequest))
			return
		}
		impersonateTarget := utils.SubstituteSuffix(callingUser.Email, body.UserEmailSubstitutableDomains, domain)
		if client, err := google_workspace.InitializeRealWorkspaceClient(ctx, impersonateTarget); err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("failed to initialize Google Workspace client as %s for domain %s: %w", impersonateTarget, domain, err))
		} else {
			googleWorkspaceDomainsToSuspendIn[domain] = client
		}
	}

	var foundUsers []models.User
	if err = db.Model(&models.User{}).Where("email IN ?", sliceOfDedupedEmailsToAttemptToDeactivate).Find(&foundUsers).Error; err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) error querying for users: %w", errors.InternalServerError, err))
		return
	}

	var response UserV3DeactivateResponse
	notification := slack.PermissionChangeNotificationInputs{
		Summary: "deactivated users",
	}
	var usersToDeactivate []models.User
	for _, user := range foundUsers {
		if user.DeactivatedAt != nil {
			response.AlreadyDeactivatedEmails = append(response.AlreadyDeactivatedEmails, user.Email)
			notification.Results = append(notification.Results, fmt.Sprintf("User %s was already deactivated in Sherlock", user.Email))
		} else {

			usersToDeactivate = append(usersToDeactivate, user)
		}
		delete(emailsToAttemptToDeactivate, user.Email)
	}
	for notFoundEmail := range emailsToAttemptToDeactivate {
		response.NotFoundEmails = append(response.NotFoundEmails, notFoundEmail)
		notification.Results = append(notification.Results, fmt.Sprintf("User %s not found in Sherlock", notFoundEmail))
	}
	now := time.Now()
	err = db.Transaction(func(tx *gorm.DB) error {
		for _, user := range usersToDeactivate {
			if txErr := tx.Model(&user).Updates(&models.User{DeactivatedAt: &now}).Error; txErr != nil {
				return txErr
			} else {
				response.NewlyDeactivatedEmails = append(response.NewlyDeactivatedEmails, user.Email)
				notification.Results = append(notification.Results, fmt.Sprintf("User %s deactivated in Sherlock", user.Email))
			}
		}
		return nil
	})

	slack.SendPermissionChangeNotification(ctx, callingUser.SlackReference(true), notification)

	if len(googleWorkspaceDomainsToSuspendIn) > 0 {
		var wg sync.WaitGroup
		wg.Add(len(googleWorkspaceDomainsToSuspendIn))
		for unsafeDomain, unsafeClient := range googleWorkspaceDomainsToSuspendIn {
			domain := unsafeDomain
			client := unsafeClient
			go func() {
				defer wg.Done()
				processGoogleWorkspaceSuspensions(callingUser.SlackReference(true), domain, client, sliceOfDedupedEmailsToAttemptToDeactivate, body.UserEmailSubstitutableDomains)
			}()
		}

		// The actual thing that we care about is suspending the Google Workspace accounts. But, after we've done so...
		// We can kick off an attempt at loading the suitability data into the database again. This would run on its own
		// eventually, and the scheduled runs have error reporting, so we don't really worry about that here, this is
		// just a best-effort.
		go func() {
			wg.Wait()
			_ = suitability_synchronization.LoadIntoDB(context.Background(), db)
		}()
	}

	ctx.JSON(http.StatusOK, response)
}

func processGoogleWorkspaceSuspensions(actor string, domain string, client google_workspace.WorkspaceClient, emails []string, substitutableDomains []string) {
	notFoundRegex := regexp.MustCompile(`(?i)not\s*found`)
	asyncCtx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancelFunc()
	notification := slack.PermissionChangeNotificationInputs{
		Summary: fmt.Sprintf("suspended Google Workspace users in domain %s", domain),
	}
	for _, email := range emails {
		target := utils.SubstituteSuffix(email, substitutableDomains, domain)
		if strings.HasSuffix(target, domain) {
			if err := client.SuspendUser(asyncCtx, target); err != nil {
				if notFoundRegex.MatchString(err.Error()) {
					notification.Results = append(notification.Results, fmt.Sprintf("Google Workspace user %s not found", target))
				} else {
					notification.Errors = append(notification.Errors, fmt.Errorf("error suspending Google Workspace user %s: %w", target, err))
				}
			} else {
				notification.Results = append(notification.Results, fmt.Sprintf("Google Workspace user %s suspended", target))
			}
		}
	}
	slack.SendPermissionChangeNotification(asyncCtx, actor, notification)
}
