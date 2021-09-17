package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/broadinstitute/sherlock/internal/builds"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

const (
	versionStringHelpText string = "unique identifier for this build. should be a full image repository path and tag"
	commitShaHelpText     string = "git commit sha associated with a particular build"
	buildURLHelpText      string = "url for the job run that created this build ie a jenkins job or github action log url OPTIONAL"
	repoHelpText          string = "url for the repo containing code for the service peing build OPTIONAL"
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
		Long:  `creates a new build entity in sherlock.`,

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
	createBuildCmd.Flags().StringVar(&serviceRepo, "repo", "", repoHelpText)

	buildCmd.AddCommand(createBuildCmd)
}

func createBuild(cmd *cobra.Command, args []string) error {
	serviceName := args[0]
	newBuild := builds.CreateBuildRequest{
		VersionString: versionString,
		CommitSha:     commitSha,
		ServiceName:   serviceName,
		BuiltAt:       time.Now(),
		BuildURL:      buildURL,
		ServiceRepo:   serviceRepo,
	}

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(newBuild).
		Post(fmt.Sprintf("%s/builds", sherlockServerURL))
	if err != nil {
		return fmt.Errorf("ERROR sending post /builds request: %v", err)
	}
	var result builds.Response
	responseBodyBytes := bytes.NewBuffer(resp.Body())
	if err := json.NewDecoder(responseBodyBytes).Decode(&result); err != nil {
		return fmt.Errorf("error parsing create build response %v. Status code: %d", err, resp.StatusCode())
	}

	// check for errors returned in response body
	if result.Error != "" {
		return fmt.Errorf("ERROR: %v, Code: %d", result.Error, resp.StatusCode())
	}

	// pretty print the sherlock api response
	var prettyResult bytes.Buffer
	if err := json.Indent(&prettyResult, resp.Body(), "", "  "); err != nil {
		return fmt.Errorf("error pretty formatting response body: %v", err)
	}

	fmt.Fprint(cmd.OutOrStdout(), prettyResult.String())
	return nil
}
