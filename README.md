# Sherlock

[![codecov](https://codecov.io/gh/broadinstitute/sherlock/branch/main/graph/badge.svg?token=kk4gi8Wa3a)](https://codecov.io/gh/broadinstitute/sherlock) [![Go Report Card](https://goreportcard.com/badge/github.com/broadinstitute/sherlock)](https://goreportcard.com/report/github.com/broadinstitute/sherlock) ![latest build](https://github.com/broadinstitute/sherlock/actions/workflows/sherlock-build.yaml/badge.svg?branch=main) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=broadinstitute_sherlock&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=broadinstitute_sherlock)
## DSP DevOps's Source-of-Truth Service

Sherlock stores information about our Kubernetes-based deployments, including Helm Chart versions and application versions. 
Sherlock doesn't do the deploying itself--it offers an API that other tools can use to understand our infrastructure.

The primary clients are Beehive, a UI for changing the information stored in Sherlock, and Thelma, a CLI that combines Sherlock's knowledge with Helm, ArgoCD, and Kubernetes APIs to directly manage infrastructure.

### Project Structure

Sherlock is a Golang server relying on a Postgres database.

An overview of interactions:

| Sherlock                  |                                                                                                                |
|---------------------------|----------------------------------------------------------------------------------------------------------------|
| API Endpoint              | :8080/api                                                                                                      |
| Swagger Endpoint          | :8080/swagger/index.html                                                                                       |
| Prometheus Endpoint       | :8080/metrics                                                                                                  |
| Go Client Library         | [./sherlock-go-client](./sherlock-go-client)                                                                   |
| TypeScript Client Library | [./sherlock-typescript-client](./sherlock-typescript-client)                                                   |
| CLI                       | Via [Thelma](https://github.com/broadinstitute/thelma)                                                         |
| GitHub Actions            | Via [Thelma](https://github.com/broadinstitute/thelma) and [./.github/workflows/client-*](./.github/workflows) |
| UI                        | Via [Beehive](https://github.com/broadinstitute/beehive)                                                       |

Sherlock is meant to be deployed behind [Google Cloud's Identity-Aware Proxy](https://cloud.google.com/iap). 
It connects to Google Workspace's Admin API to evaluate permissions of the calling users.

### Developing Locally

GoLand should be able to understand the monorepo structure and should download dependencies appropriately.

Some key [makefile](./makefile) commands:

- `make install-pact` will ask for sudo credentials to install the Pact FFI library necessary for running those tests
- `make pg-up` will run a blank local database so that GoLand can run tests for you
  - `make pg-down` will tear it down if it gets in a bad state
- `make local-up` will run Sherlock locally (different database from `make pg-up`'s, so you can add state there)
  - `make local-stop` will shut down Sherlock so you can rebuild it **without wiping your database state**
  - `make local-down` **will wipe your database state**

There's more in the [makefile](./makefile).
