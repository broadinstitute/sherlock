package v2

import (
	"context"
	"fmt"
	"os"

	"github.com/broadinstitute/sherlock/clients/go/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
)

type sherlockClient struct {
	client *client.Sherlock
}

type sherlockClientOptions struct {
	hostURL               string
	schemes               []string
	credentialsPath       string
	useServiceAccountAuth bool
}

func NewSherlockClient(options sherlockClientOptions) (*sherlockClient, error) {
	transport := httptransport.New(options.hostURL, "", options.schemes)
	if options.useServiceAccountAuth {
		idToken, err := getIapTokenFromSA(options.credentialsPath)
		if err != nil {
			return nil, err
		}

		transport.DefaultAuthentication = httptransport.BearerToken(idToken)
	}

	client := client.New(transport, strfmt.Default)
	sherlock := &sherlockClient{client}

	return sherlock, nil
}

func getIapTokenFromSA(credentialsPath string) (string, error) {
	ctx := context.Background()
	audience, ok := os.LookupEnv("SHERLOCK_OAUTH_AUDIENCE")
	if !ok {
		return "", fmt.Errorf("unable to determine oauth audience")
	}

	tokenSource, err := idtoken.NewTokenSource(ctx, audience, option.WithCredentialsFile(credentialsPath))
	if err != nil {
		return "", fmt.Errorf("unable to generate oauth token source from credentials: %v", err)
	}

	accessToken, err := tokenSource.Token()
	if err != nil {
		return "", fmt.Errorf("unable to fetch token from source: %v", err)
	}

	idToken, ok := accessToken.Extra("id_token").(string)
	if !ok {
		return "", fmt.Errorf("unable to parse id_token")
	}

	return idToken, nil
}
