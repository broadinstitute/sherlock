package pkg

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client/ci_runs"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-playground/webhooks/v6/github"
	"hash"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	// sherlockUrlEnvVar should be set like "https://sherlock.dsp-devops.broadinstitute.org" or "http://localhost:8080"
	sherlockUrlEnvVar = "SHERLOCK_URL"
	// iapTokenOverrideEnvVar overrides IAP-token generating behavior, to allow local dev
	iapTokenOverrideEnvVar = "IAP_TOKEN"
	// iapAudienceEnvVar should be set to the Client ID of the IAP OAuth credentials
	iapAudienceEnvVar = "IAP_AUDIENCE"
	// githubWebhookSecretEnvVar should be set to the secret set in the GitHub Webhook config
	githubWebhookSecretEnvVar = "GITHUB_WEBHOOK_SECRET"
	// allowedGithubOrgsEnvVar should be a space-separated list of GitHub orgs that this cloud function should pay attention to.
	// This is necessary because GitHub Apps can theoretically be installed by anyone. That doesn't affect us except for webhooks,
	// where it technically allows arbitrary people to lob requests at this endpoint. That's not new, this endpoint is public,
	// but these requests will come from GitHub, so we should filter.
	allowedGithubOrgsEnvVar = "ALLOWED_GITHUB_ORGS"
	// idTokenUrl is the URL to use to get an IAP ID token. Will be used with ?audience=${iapAudienceEnvVar}
	idTokenUrl = "http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/identity"
)

var (
	sherlockHostname, sherlockScheme, iapAudience, githubWebhookSecret string
	allowedGithubOrgs                                                  []string
)

// init does first-time initialization
func init() {
	sherlockUrl, present := os.LookupEnv(sherlockUrlEnvVar)
	if !present {
		log.Fatalf("os.LookupEnv(%s): present=false\n", sherlockUrlEnvVar)
	} else if sherlockUrl == "" {
		log.Fatalf("os.LookupEnv(%s): sherlockUrl=''\n", sherlockUrlEnvVar)
	}
	parsedSherlockUrl, err := url.Parse(sherlockUrl)
	if err != nil {
		log.Fatalf("url.Parse(%s): %v\n", sherlockUrl, err)
	}
	sherlockHostname = parsedSherlockUrl.Hostname()
	if parsedSherlockUrl.Port() != "" {
		sherlockHostname += ":"
		sherlockHostname += parsedSherlockUrl.Port()
	}
	sherlockScheme = parsedSherlockUrl.Scheme

	if _, present = os.LookupEnv(iapTokenOverrideEnvVar); !present {
		iapAudience, present = os.LookupEnv(iapAudienceEnvVar)
		if !present {
			log.Fatalf("os.LookupEnv(%s): present=false\n", iapAudienceEnvVar)
		} else if iapAudience == "" {
			log.Fatalf("os.LookupEnv(%s): iapAudience=''\n", iapAudienceEnvVar)
		}
	}

	if githubWebhookSecret, present = os.LookupEnv(githubWebhookSecretEnvVar); !present {
		log.Fatalf("os.LookupEnv(%s): present=false\n", githubWebhookSecretEnvVar)
	} else if githubWebhookSecret == "" {
		log.Fatalf("os.LookupEnv(%s): githubWebhookSecret=''\n", githubWebhookSecretEnvVar)
	}

	var allowedGithubOrgsString string
	if allowedGithubOrgsString, present = os.LookupEnv(allowedGithubOrgsEnvVar); !present {
		log.Fatalf("os.LookupEnv(%s): present=false\n", allowedGithubOrgsEnvVar)
	} else if allowedGithubOrgsString == "" {
		log.Fatalf("os.LookupEnv(%s): allowedGithubOrgsString=''\n", allowedGithubOrgsEnvVar)
	} else if allowedGithubOrgs = strings.Split(allowedGithubOrgsString, " "); len(allowedGithubOrgs) == 0 {
		log.Fatalf("len(strings.Split(\"%s\", \" \"))=0\n", allowedGithubOrgsString)
	}
}

// setup creates instanced things per-request
func setup() (mac hash.Hash, hook *github.Webhook, transport *httptransport.Runtime) {
	transport = httptransport.New(sherlockHostname, "", []string{sherlockScheme})

	mac = hmac.New(sha256.New, []byte(githubWebhookSecret))

	hook, err := github.New(github.Options.Secret(githubWebhookSecret))
	if err != nil {
		log.Fatalf("github.New: %v\n", err)
	}

	return mac, hook, transport
}

// HandleWebhook is what actually does the handling, running once per request
func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	mac, hook, transport := setup()

	switch r.RequestURI {
	case "/webhook":
		// Require that the signature header is present
		signature := strings.TrimPrefix(r.Header.Get("X-Hub-Signature-256"), "sha256=")
		if signature == "" {
			w.WriteHeader(http.StatusUnauthorized)
			log.Printf("request was missing X-Hub-Signature-256 header\n")
			return
		}

		// As r.Body is read, additionally synchronously write it into mac
		r.Body = struct {
			io.Reader
			io.Closer
		}{
			Reader: io.TeeReader(r.Body, mac),
			Closer: r.Body,
		}

		// Call the library and handle its errors (it does try to check signature, but using a more insecure method)
		rawPayload, err := hook.Parse(r, github.WorkflowRunEvent, github.PingEvent)
		if err != nil {
			switch {
			case errors.Is(err, github.ErrMissingHubSignatureHeader):
				w.WriteHeader(http.StatusUnauthorized)
				log.Printf("library said hook was unauthorized\n")
				return
			case errors.Is(err, github.ErrHMACVerificationFailed):
				w.WriteHeader(http.StatusForbidden)
				log.Printf("library said hook signature was invalid\n")
				return
			case errors.Is(err, github.ErrEventNotFound):
				w.WriteHeader(http.StatusNotFound)
				log.Printf("library said hook was a type it wasn't configured to receive\n")
				return
			case errors.Is(err, github.ErrInvalidHTTPMethod):
				w.WriteHeader(http.StatusMethodNotAllowed)
				log.Printf("library said hook was sent with an invalid method\n")
				return
			default:
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("library had an unknown error handling the hook: %v\n", err)
				return
			}
		}

		// Check HMAC + SHA256 signature
		calculatedSignature := hex.EncodeToString(mac.Sum(nil))
		if signature != calculatedSignature {
			w.WriteHeader(http.StatusForbidden)
			log.Printf("HMAC + SHA 256 signature was %s but calculated was %s, rejecting\n", signature, calculatedSignature)
			return
		}

		// Handle parsed payloads from the library
		switch payload := rawPayload.(type) {

		// ping issued upon the webhook being added to a new repo; might as well respond with 200
		case github.PingPayload:
			if !utils.Contains(allowedGithubOrgs, payload.Repository.Owner.Login) {
				w.WriteHeader(http.StatusForbidden)
			} else {
				w.WriteHeader(http.StatusOK)
			}
			log.Printf("received ping from repo %s", payload.Repository.FullName)

		// workflow_run issued upon workflow request, running, and completion
		case github.WorkflowRunPayload:
			if !utils.Contains(allowedGithubOrgs, payload.Repository.Owner.Login) {
				w.WriteHeader(http.StatusForbidden)
				log.Printf("bailing out, workflow run from %s", payload.Repository.FullName)
				return
			}

			if token, present := os.LookupEnv(iapTokenOverrideEnvVar); present {
				// If we have a token, just use that
				transport.DefaultAuthentication = httptransport.BearerToken(token)
			} else {
				// Otherwise, do the dance to get it from the metadata server
				formedIdTokenUrl := fmt.Sprintf("%s?audience=%s", idTokenUrl, iapAudience)
				req, err := http.NewRequest(http.MethodGet, formedIdTokenUrl, nil)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					log.Printf("http.NewRequest(%s, %s): %v\n", http.MethodGet, formedIdTokenUrl, err)
					return
				}
				req.Header.Set("Metadata-Flavor", "Google")
				resp, err := (&http.Client{}).Do(req)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					log.Printf("(&http.Client{}).Do(%s): %v\n", formedIdTokenUrl, err)
					return
				} else if resp.StatusCode != http.StatusOK {
					w.WriteHeader(http.StatusInternalServerError)
					log.Printf("(&http.Client{}).Do(%s): non-200: %d", formedIdTokenUrl, resp.StatusCode)
					return
				}
				idToken, err := io.ReadAll(resp.Body)
				_ = resp.Body.Close()
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					log.Printf("io.ReadAll(resp.Body): %v\n", err)
					return
				}
				transport.DefaultAuthentication = httptransport.BearerToken(string(idToken[:]))
			}

			sherlockClient := client.New(transport, strfmt.Default)

			// Convert webhook fields into what we'll store in Sherlock
			var startedAt, status, terminalAt string
			if !payload.WorkflowRun.RunStartedAt.IsZero() {
				// Seems like this field will always be present, but maybe it'll be zero if it hasn't actually started yet?
				// Seems possible :shrug:
				startedAt = payload.WorkflowRun.RunStartedAt.Format(time.RFC3339)
			}
			if payload.WorkflowRun.Conclusion != "" {
				status = payload.WorkflowRun.Conclusion
				terminalAt = payload.WorkflowRun.UpdatedAt.Format(time.RFC3339)
			} else {
				status = payload.WorkflowRun.Status
			}

			// PUT to Sherlock; will create if the selector isn't found and edit otherwise
			created, err := sherlockClient.CiRuns.PutAPICiRunsV3(&ci_runs.PutAPICiRunsV3Params{
				Context: context.Background(),
				CiRun: &models.SherlockCiRunV3Upsert{
					Platform:                   "github-actions",
					GithubActionsOwner:         payload.Repository.Owner.Login,
					GithubActionsRepo:          payload.Repository.Name,
					GithubActionsRunID:         payload.WorkflowRun.ID,
					GithubActionsAttemptNumber: payload.WorkflowRun.RunAttempt,
					GithubActionsWorkflowPath:  payload.Workflow.Path,
					StartedAt:                  startedAt,
					TerminalAt:                 terminalAt,
					Status:                     status,
				},
			})

			// Handle response cases
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("sherlockClient.CiRuns.PutAPICiRunsV3(): error %v", err)
			} else if created != nil {
				w.WriteHeader(http.StatusCreated)
				log.Printf("sherlockClient.CiRuns.PutAPICiRunsV3(): upserted CiRun %d, '%s'", created.Payload.ID, created.Payload.Status)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("sherlockClient.CiRuns.PutAPICiRunsV3(): error and response both nil")
			}

		// Some payload we don't handle
		default:
			log.Printf("unknown payload type?\n")
			w.WriteHeader(http.StatusNotFound)
		}
	default:
		// If we got a body but not at /webhooks, return a 400 so GitHub shows the webhook as failed
		body, err := io.ReadAll(io.LimitReader(r.Body, 1))
		switch {
		case err != nil:
			w.WriteHeader(http.StatusUnprocessableEntity)
		case len(body) > 0:
			w.WriteHeader(http.StatusBadRequest)
		}
		log.Printf("received request not to /webhook, body present=%v\n", len(body) > 0)
		_, _ = fmt.Fprintln(w, "cloud function operational; direct webhooks to /webhook")
	}
}
