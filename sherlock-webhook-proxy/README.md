# sherlock-webhook-proxy

It receives webhooks, validates them, and turns them into IAP-authenticated calls to Sherlock's normal API.

### Running locally

```
go run cmd/main.go
```

Then you can curl `localhost:8090`.