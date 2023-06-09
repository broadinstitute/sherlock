package sherlock_webhook_proxy

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/broadinstitute/sherlock/sherlock-webhook-proxy/pkg"
)

func init() {
	functions.HTTP("HandleWebhook", pkg.HandleWebhook)
}
