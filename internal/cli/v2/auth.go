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

type SherlockClient struct {
	client client.Sherlock
}

func NewSherlockClient() (SherlockClient, error) {
	transport := httptransport.New("sherlock.dsp-devops.broadinstitute.org", "", []string{"https"})
	transport.SetDebug(true)
	// transport.DefaultAuthentication = httptransport.BearerToken()

	_ = client.New(transport, strfmt.Default)
}

func getIapTokenFromSA(credentialsPath string) (string, error) {
	ctx := context.Background()
	audience, ok := os.LookupEnv("SHERLOCK_OAUTH_AUDIENCE")
	if !ok {
		return "", fmt.Errorf("Unable to determine oauth audience")
	}

	tokenSource, err := idtoken.NewTokenSource(ctx, audience, option.WithCredentialsFile(credentialsPath))
	if err != nil {
		return "", fmt.Errorf("Unable to generate oauth token source from credentials: %v", err)
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
