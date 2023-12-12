package pkg

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"io"
	"log"
	"net/http"
	"os"
)

func authenticateSherlockClient(w http.ResponseWriter, transport *httptransport.Runtime) (_client *client.Sherlock, _ok bool) {
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
			return nil, false
		}
		req.Header.Set("Metadata-Flavor", "Google")
		resp, err := (&http.Client{}).Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("(&http.Client{}).Do(%s): %v\n", formedIdTokenUrl, err)
			return nil, false
		} else if resp.StatusCode != http.StatusOK {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("(&http.Client{}).Do(%s): non-200: %d", formedIdTokenUrl, resp.StatusCode)
			return nil, false
		}
		idToken, err := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("io.ReadAll(resp.Body): %v\n", err)
			return nil, false
		}
		transport.DefaultAuthentication = httptransport.BearerToken(string(idToken[:]))
	}

	return client.New(transport, strfmt.Default), true
}
