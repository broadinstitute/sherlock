// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/clients/go/client/app_versions"
	"github.com/broadinstitute/sherlock/clients/go/client/changesets"
	"github.com/broadinstitute/sherlock/clients/go/client/chart_releases"
	"github.com/broadinstitute/sherlock/clients/go/client/chart_versions"
	"github.com/broadinstitute/sherlock/clients/go/client/charts"
	"github.com/broadinstitute/sherlock/clients/go/client/clusters"
	"github.com/broadinstitute/sherlock/clients/go/client/database_instances"
	"github.com/broadinstitute/sherlock/clients/go/client/environments"
	"github.com/broadinstitute/sherlock/clients/go/client/misc"
	"github.com/broadinstitute/sherlock/clients/go/client/pagerduty_integrations"
)

// Default sherlock HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new sherlock HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Sherlock {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new sherlock HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Sherlock {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new sherlock client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Sherlock {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Sherlock)
	cli.Transport = transport
	cli.AppVersions = app_versions.New(transport, formats)
	cli.Changesets = changesets.New(transport, formats)
	cli.ChartReleases = chart_releases.New(transport, formats)
	cli.ChartVersions = chart_versions.New(transport, formats)
	cli.Charts = charts.New(transport, formats)
	cli.Clusters = clusters.New(transport, formats)
	cli.DatabaseInstances = database_instances.New(transport, formats)
	cli.Environments = environments.New(transport, formats)
	cli.Misc = misc.New(transport, formats)
	cli.PagerdutyIntegrations = pagerduty_integrations.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Sherlock is a client for sherlock
type Sherlock struct {
	AppVersions app_versions.ClientService

	Changesets changesets.ClientService

	ChartReleases chart_releases.ClientService

	ChartVersions chart_versions.ClientService

	Charts charts.ClientService

	Clusters clusters.ClientService

	DatabaseInstances database_instances.ClientService

	Environments environments.ClientService

	Misc misc.ClientService

	PagerdutyIntegrations pagerduty_integrations.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Sherlock) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.AppVersions.SetTransport(transport)
	c.Changesets.SetTransport(transport)
	c.ChartReleases.SetTransport(transport)
	c.ChartVersions.SetTransport(transport)
	c.Charts.SetTransport(transport)
	c.Clusters.SetTransport(transport)
	c.DatabaseInstances.SetTransport(transport)
	c.Environments.SetTransport(transport)
	c.Misc.SetTransport(transport)
	c.PagerdutyIntegrations.SetTransport(transport)
}
