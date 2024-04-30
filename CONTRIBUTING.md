# Contributing

> ### The bottom line is to reach out to DevOps in [#dsp-devops-discussions](https://broadinstitute.enterprise.slack.com/archives/C029LTN5L80) before contributing.
>
> This is a small project with a very short list of contributors and it's not as easy to contribute to as many other codebases.

> This document is still under construction -- it's written a bit out of order to quickly meet compliance requirements.

## Testing

Sherlock is entirely unit/functional tested. 

There are no connected or integration tests. Instead, client libraries are used whenever possible and those libraries are mocked for testing. Internal packages should tie themselves into Sherlock's boot system to validate their connections on startup (since such connections wouldn't be tested, halting an incremental rollout on any issues reduces risk).

Sherlock does have capability to run contract tests as a provider but this is disabled since it has no consumers and Pact requires system-level dependencies.

## Backwards Compatibility

### APIs

Endpoints are versioned in groups according [our internal API versioning standards](https://docs.google.com/document/d/1qXNHTijdPn9ApYrznSkTFnxkt0g-o-Uh0SjqQlYd-ZA/edit), which are similar to [those used by Google](https://cloud.google.com/apis/design/versioning).

This means that a given API endpoint must be backwards compatible for its entire available lifetime. An endpoint may only be removed once metrics indicate that it is entirely out of use.

### Database Schema

Sherlock is a multi-replica system, so we can't make database changes that would break existing replicas during incremental rollout. As this is internal software, we do allow breaking changes when the change would not be breaking for any instances we have deployed. Current deployments can be found at https://beehive.dsp-devops-prod.broadinstitute.org/charts/sherlock/chart-releases.

## Submitting Changes

PRs require a ticket in the title with CLIA risk and security impact description filled. The description must include a summary of the changes, an explanation of what testing was done (screenshots recommended), and a note on the risk of the change.
