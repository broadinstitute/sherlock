package cli

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/broadinstitute/sherlock/internal/deploys"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

const (
	deployEnvironmentNameHelpText = "the name of the environment being deployed to"
	deployServiceNameHelpText     = "the name of the service that is being deployed"
)

var (
	deployEnvironmentName string
	deployServiceName     string
	createDeployCmd       = &cobra.Command{
		Use:   "create",
		Short: "create a new deploy",
		Long:  `creates a new deploy of service which will be tracked by sherlock.`,

		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errInvalidArgs
			}
			return nil
		},
		RunE: createDeploy,
	}
)

func init() {
	createDeployCmd.Flags().StringVar(&deployEnvironmentName, "environment", "", deployEnvironmentNameHelpText)
	_ = createDeployCmd.MarkFlagRequired("environment")

	createDeployCmd.Flags().StringVar(&deployServiceName, "service", "", deployServiceNameHelpText)
	_ = createDeployCmd.MarkFlagRequired("service")

	deployCmd.AddCommand(createDeployCmd)
}

func createDeploy(cmd *cobra.Command, args []string) error {
	// version string ie docker image url and tag is passed as first
	// positional arg
	versionString := args[0]
	newDeployRequest := deploys.CreateDeployRequestBody{
		VersionString: versionString,
	}

	result, rawResponseBody, err := dispatchCreateDeployRequest(newDeployRequest, deployEnvironmentName, deployServiceName)
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
func dispatchCreateDeployRequest(newDeploy deploys.CreateDeployRequestBody, environment, service string) (*deploys.Response, []byte, error) {
	client := resty.New()
	urlPath := fmt.Sprintf("%s/deploys/%s/%s", sherlockServerURL, environment, service)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(newDeploy).
		Post(urlPath)
	if err != nil {
		return nil, []byte{}, fmt.Errorf("ERROR sending POST deploy request: %v", err)
	}

	var result deploys.Response
	responseBodyBytes := bytes.NewBuffer(resp.Body())
	if err := json.NewDecoder(responseBodyBytes).Decode(&result); err != nil {
		return nil, []byte{}, fmt.Errorf("error parsing create deploy response: %v", err)
	}

	return &result, resp.Body(), nil
}
