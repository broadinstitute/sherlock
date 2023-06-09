package pkg

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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
	// idTokenUrl is the URL to use to get an IAP ID token. Will be used with ?audience=${iapAudienceEnvVar}
	idTokenUrl = "http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/identity"
)

var (
	_mac       hash.Hash
	_hook      *github.Webhook
	_transport *httptransport.Runtime
)

// init does first-time setup
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
	hostname := parsedSherlockUrl.Hostname()
	if parsedSherlockUrl.Port() != "" {
		hostname += ":"
		hostname += parsedSherlockUrl.Port()
	}
	_transport = httptransport.New(hostname, "", []string{parsedSherlockUrl.Scheme})

	if _, present := os.LookupEnv(iapTokenOverrideEnvVar); !present {
		iapAudience, present := os.LookupEnv(iapAudienceEnvVar)
		if !present {
			log.Fatalf("os.LookupEnv(%s): present=false\n", iapAudienceEnvVar)
		} else if iapAudience == "" {
			log.Fatalf("os.LookupEnv(%s): sherlockUrl=''\n", iapAudienceEnvVar)
		}
	}

	secret, present := os.LookupEnv(githubWebhookSecretEnvVar)
	if !present {
		log.Fatalf("os.LookupEnv(%s): present=false\n", githubWebhookSecretEnvVar)
	} else if secret == "" {
		log.Fatalf("os.LookupEnv(%s): secret=''\n", githubWebhookSecretEnvVar)
	}
	_mac = hmac.New(sha256.New, []byte(secret))

	hook, err := github.New(github.Options.Secret(secret))
	if err != nil {
		log.Fatalf("github.New: %v\n", err)
	}
	_hook = hook
}

// HandleWebhook is what actually does the handling, running once per request
func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	switch r.RequestURI {
	case "/webhook":
		// Require that the signature header is present
		signature := strings.TrimPrefix(r.Header.Get("X-Hub-Signature-256"), "sha256=")
		if signature == "" {
			w.WriteHeader(http.StatusUnauthorized)
			log.Printf("request was missing X-Hub-Signature-256 header\n")
			return
		}

		// As r.Body is read, additionally synchronously write it into _mac
		_mac.Reset()
		r.Body = struct {
			io.Reader
			io.Closer
		}{
			Reader: io.TeeReader(r.Body, _mac),
			Closer: r.Body,
		}

		// Call the library and handle its errors (it does try to check signature, but using a more insecure method)
		rawPayload, err := _hook.Parse(r, github.WorkflowRunEvent, github.PingEvent)
		if err != nil {
			switch err {
			case github.ErrMissingHubSignatureHeader:
				w.WriteHeader(http.StatusUnauthorized)
				log.Printf("library said hook was unauthorized\n")
				return
			case github.ErrHMACVerificationFailed:
				w.WriteHeader(http.StatusForbidden)
				log.Printf("library said hook signature was invalid\n")
				return
			case github.ErrEventNotFound:
				w.WriteHeader(http.StatusNotFound)
				log.Printf("library said hook was a type it wasn't configured to receive\n")
				return
			case github.ErrInvalidHTTPMethod:
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
		calculatedSignature := hex.EncodeToString(_mac.Sum(nil))
		if signature != calculatedSignature {
			w.WriteHeader(http.StatusForbidden)
			log.Printf("HMAC + SHA 256 signature was %s but calculated was %s, rejecting\n", signature, calculatedSignature)
			return
		}

		// Handle parsed payloads from the library
		switch payload := rawPayload.(type) {

		// ping issued upon the webhook being added to a new repo; might as well respond with 200
		case github.PingPayload:
			w.WriteHeader(http.StatusOK)
			log.Printf("received ping from repo %s", payload.Repository.FullName)

		// workflow_run issued upon workflow request, running, and completion
		case github.WorkflowRunPayload:
			if token, present := os.LookupEnv(iapTokenOverrideEnvVar); present {
				// If we have a token, just use that
				_transport.DefaultAuthentication = httptransport.BearerToken(token)
			} else {
				// Otherwise, do the dance to get it from the metadata server
				formedIdTokenUrl := fmt.Sprintf("%s?audience=%s", idTokenUrl, os.Getenv(iapAudienceEnvVar))
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
				_transport.DefaultAuthentication = httptransport.BearerToken(string(idToken[:]))
			}

			sherlockClient := client.New(_transport, strfmt.Default)

			// Convert webhook fields into what we'll store in Sherlock
			var status, terminalAt string
			if payload.WorkflowRun.Conclusion != "" {
				status = payload.WorkflowRun.Conclusion
				terminalAt = payload.WorkflowRun.UpdatedAt.Format(time.RFC3339)
			} else {
				status = payload.WorkflowRun.Status
			}

			// PUT to Sherlock; will create if the selector isn't found and edit otherwise
			edited, created, err := sherlockClient.CiRuns.PutAPIV2CiRunsSelector(&ci_runs.PutAPIV2CiRunsSelectorParams{
				Context: context.Background(),
				Selector: fmt.Sprintf("github-actions/%s/%s/%d/%d",
					payload.Repository.Owner.Login,
					payload.Repository.Name,
					payload.WorkflowRun.ID,
					payload.WorkflowRun.RunAttempt),
				CiRun: &models.V2controllersCreatableCiRun{
					Platform:                   "github-actions",
					GithubActionsOwner:         payload.Repository.Owner.Login,
					GithubActionsRepo:          payload.Repository.Name,
					GithubActionsRunID:         payload.WorkflowRun.ID,
					GithubActionsAttemptNumber: payload.WorkflowRun.RunAttempt,
					GithubActionsWorkflowPath:  payload.Workflow.Path,
					TerminalAt:                 terminalAt,
					Status:                     status,
				},
			})

			// Handle response cases
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("sherlockClient.CiRuns.PutAPIV2CiRunsSelector(): error %v", err)
			} else if edited != nil {
				w.WriteHeader(http.StatusOK)
				log.Printf("sherlockClient.CiRuns.PutAPIV2CiRunsSelector(): edited CiRun %d, '%s'", edited.Payload.ID, edited.Payload.Status)
			} else if created != nil {
				w.WriteHeader(http.StatusCreated)
				log.Printf("sherlockClient.CiRuns.PutAPIV2CiRunsSelector(): created CiRun %d, '%s'", created.Payload.ID, created.Payload.Status)
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
		log.Printf("received request not to /webhook, body present=%v\n", len(body))
		_, _ = fmt.Fprintln(w, "cloud function operational; direct webhooks to /webhook")
	}
}
