// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/broadinstitute/sherlock/sherlock-go-client/client/app_versions"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client/changesets"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client/chart_releases"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client/chart_versions"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client/charts"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client/ci_identifiers"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client/ci_runs"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client/clusters"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client/database_instances"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client/deploy_hooks"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client/environments"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client/git_commits"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client/github_actions_jobs"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client/incidents"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client/misc"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client/pagerduty_integrations"
	"github.com/broadinstitute/sherlock/sherlock-go-client/client/users"
)

// Default sherlock HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "sherlock.dsp-devops.broadinstitute.org"
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
	cli.CiIdentifiers = ci_identifiers.New(transport, formats)
	cli.CiRuns = ci_runs.New(transport, formats)
	cli.Clusters = clusters.New(transport, formats)
	cli.DatabaseInstances = database_instances.New(transport, formats)
	cli.DeployHooks = deploy_hooks.New(transport, formats)
	cli.Environments = environments.New(transport, formats)
	cli.GitCommits = git_commits.New(transport, formats)
	cli.GithubActionsJobs = github_actions_jobs.New(transport, formats)
	cli.Incidents = incidents.New(transport, formats)
	cli.Misc = misc.New(transport, formats)
	cli.PagerdutyIntegrations = pagerduty_integrations.New(transport, formats)
	cli.Users = users.New(transport, formats)
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

	CiIdentifiers ci_identifiers.ClientService

	CiRuns ci_runs.ClientService

	Clusters clusters.ClientService

	DatabaseInstances database_instances.ClientService

	DeployHooks deploy_hooks.ClientService

	Environments environments.ClientService

	GitCommits git_commits.ClientService

	GithubActionsJobs github_actions_jobs.ClientService

	Incidents incidents.ClientService

	Misc misc.ClientService

	PagerdutyIntegrations pagerduty_integrations.ClientService

	Users users.ClientService

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
	c.CiIdentifiers.SetTransport(transport)
	c.CiRuns.SetTransport(transport)
	c.Clusters.SetTransport(transport)
	c.DatabaseInstances.SetTransport(transport)
	c.DeployHooks.SetTransport(transport)
	c.Environments.SetTransport(transport)
	c.GitCommits.SetTransport(transport)
	c.GithubActionsJobs.SetTransport(transport)
	c.Incidents.SetTransport(transport)
	c.Misc.SetTransport(transport)
	c.PagerdutyIntegrations.SetTransport(transport)
	c.Users.SetTransport(transport)
}
