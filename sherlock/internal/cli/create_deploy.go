package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/broadinstitute/sherlock/internal/handlers/v1handlers"
	"github.com/broadinstitute/sherlock/internal/serializers/v1serializers"

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
	newDeployRequest := v1handlers.CreateDeployRequestBody{
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

func dispatchCreateDeployRequest(newDeploy v1handlers.CreateDeployRequestBody, environment, service string) (*v1serializers.DeploysResponse, []byte, error) {
	var (
		req *resty.Request
		err error
	)

	client := resty.New()
	urlPath := fmt.Sprintf("%s/api/v1/deploys/%s/%s", sherlockServerURL, environment, service)
	req = client.R()
	// set authorization headers when running cli via automated workflows
	if useServiceAccountAuth {
		req, err = setAuthHeader(req)
		if err != nil {
			return nil, []byte{}, fmt.Errorf("error setting auth header: %v", err)
		}
	}
	resp, err := req.
		SetHeader("Content-Type", "application/json").
		SetBody(newDeploy).
		Post(urlPath)
	if err != nil {
		return nil, []byte{}, fmt.Errorf("ERROR sending POST deploy request: %v", err)
	}

	var result v1serializers.DeploysResponse
	responseBodyBytes := bytes.NewBuffer(resp.Body())
	if err := json.NewDecoder(responseBodyBytes).Decode(&result); err != nil {
		return nil, []byte{}, fmt.Errorf("error parsing create deploy response: %v", err)
	}

	return &result, resp.Body(), nil
}
