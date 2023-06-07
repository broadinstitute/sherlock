package pagerduty

import "github.com/PagerDuty/go-pagerduty"

// We don't actually need to set any sort of auth on the pagerduty client, but we do need it to send events.
//
// There's technically a static function to send alert events without a client, but there's no such function
// for sending change events.
var client = pagerduty.NewClient("")
