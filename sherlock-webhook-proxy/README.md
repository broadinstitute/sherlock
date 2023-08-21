# sherlock-webhook-proxy

It receives webhooks, validates them, and turns them into IAP-authenticated calls to Sherlock's normal API.

### Running locally

```
make local-up
```
```
cd sherlock-webhook-proxy && FUNCTION_TARGET="HandleWebhook" IAP_TOKEN=$(thelma auth iap --echo) SHERLOCK_URL=http://localhost:8080 GITHUB_WEBHOOK_SECRET=foobar ALLOWED_GITHUB_ORGS=broadinstitute go run cmd/main.go
```
```
gh webhook forward --repo=broadinstitute/sherlock --events=workflow_run --url=http://localhost:8090/webhook --secret=foobar
```

You'll need to vary those commands to do things like talk to sherlock-dev or test different IAP behavior.