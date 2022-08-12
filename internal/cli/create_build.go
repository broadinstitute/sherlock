package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/broadinstitute/sherlock/internal/controllers/v1controllers"
	"github.com/broadinstitute/sherlock/internal/serializers/v1serializers"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
)

const (
	versionStringHelpText string = "unique identifier for this build. should be a full image repository path and tag"
	commitShaHelpText     string = "git commit sha associated with a particular build"
	buildURLHelpText      string = "url for the job run that created this build ie a jenkins job or github action log url OPTIONAL"
	repoHelpText          string = "url for the repo containing code for the service being build OPTIONAL"
)

var (
	// ErrInvalidArgs is returned when builds create command receives no or too many arguments
	serviceRepo    string
	versionString  string
	commitSha      string
	buildURL       string
	errInvalidArgs = errors.New("builds create usage: sherlock builds create SERVICE_NAME [flags]")
	createBuildCmd = &cobra.Command{
		Use:   "create",
		Short: "create a new build",
		Long:  `creates a new build of service which will be tracked by sherlock.`,

		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errInvalidArgs
			}
			return nil
		},
		RunE: createBuild,
	}
)

func init() {
	createBuildCmd.Flags().StringVar(&versionString, "version-string", "", versionStringHelpText)
	_ = createBuildCmd.MarkFlagRequired("version-string")

	createBuildCmd.Flags().StringVar(&commitSha, "commit-sha", "", commitShaHelpText)
	_ = createBuildCmd.MarkFlagRequired("commit-sha")

	createBuildCmd.Flags().StringVar(&buildURL, "build-url", "", buildURLHelpText)
	createBuildCmd.Flags().StringVar(&serviceRepo, "repo-url", "", repoHelpText)

	buildCmd.AddCommand(createBuildCmd)
}

func createBuild(cmd *cobra.Command, args []string) error {
	serviceName := args[0]
	newBuild := v1controllers.CreateBuildRequest{
		VersionString: versionString,
		CommitSha:     commitSha,
		ServiceName:   serviceName,
		BuiltAt:       time.Now(),
		BuildURL:      buildURL,
		ServiceRepo:   serviceRepo,
	}

	result, rawResponseBody, err := dispatchCreateBuildRequest(newBuild)
	if err != nil {
		return fmt.Errorf("ERROR: %v", err)
	}

	// check for errors returned in response
	if result.Error != "" {
		return fmt.Errorf("ERROR: %v", result.Error)
	}

	// pretty print the sherlock api response
	var prettyResult bytes.Buffer
	if err := json.Indent(&prettyResult, rawResponseBody, "", "  "); err != nil {
		return fmt.Errorf("error pretty formatting response body: %v", err)
	}

	fmt.Fprint(cmd.OutOrStdout(), prettyResult.String())
	return nil
}

func dispatchCreateBuildRequest(newBuild v1controllers.CreateBuildRequest) (*v1serializers.BuildsResponse, []byte, error) {
	var (
		req *resty.Request
		err error
	)

	client := resty.New()
	req = client.R()
	// set authorization headers when running cli via automated workflows
	if useServiceAccountAuth {
		req, err = setAuthHeader(req)
		if err != nil {
			return nil, []byte{}, fmt.Errorf("error setting auth header: %v", err)
		}
	}

	resp, err := req.SetHeader("Content-Type", "application/json").
		SetBody(newBuild).
		Post(fmt.Sprintf("%s/api/v1/builds", sherlockServerURL))
	if err != nil {
		return nil, []byte{}, fmt.Errorf("ERROR sending post /api/v1/builds request: %v", err)
	}

	var result v1serializers.BuildsResponse
	responseBodyBytes := bytes.NewBuffer(resp.Body())
	if err := json.NewDecoder(responseBodyBytes).Decode(&result); err != nil {
		return nil, []byte{}, fmt.Errorf("error parsing create build response %v. Status code: %d", err, resp.StatusCode())
	}
	return &result, resp.Body(), nil
}

func setAuthHeader(req *resty.Request) (*resty.Request, error) {
	ctx := context.Background()
	audience := os.Getenv("SHERLOCK_OAUTH_AUDIENCE")
	tokenGenerator, err := idtoken.NewTokenSource(ctx, audience, option.WithCredentialsFile(clientCredentials))
	if err != nil {
		return nil, fmt.Errorf("error creating token generator: %v", err)
	}
	token, err := tokenGenerator.Token()
	if err != nil {
		return nil, fmt.Errorf("error generating oidc token: %v", err)
	}

	return req.SetAuthToken(token.AccessToken), nil
}
