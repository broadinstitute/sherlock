# Sherlock

[![codecov](https://codecov.io/gh/broadinstitute/sherlock/branch/main/graph/badge.svg?token=kk4gi8Wa3a)](https://codecov.io/gh/broadinstitute/sherlock)  [![Go Report Card](https://goreportcard.com/badge/github.com/broadinstitute/sherlock)](https://goreportcard.com/report/github.com/broadinstitute/sherlock)  ![latest build](https://github.com/broadinstitute/sherlock/actions/workflows/build.yaml/badge.svg?branch=main)

Sherlock is a devops built and maintained service which serves will maintain a historical record of deployments to all Terra services across all environments. 
Additionally it will extract delivery metadata from updates to this historical deployment log. In the future it will also act as the control layer for preview environments.
The eventual replacement for fiabs. 

## Warning
This tool is currently in a POC state and under active development. Frequent and potentially breaking changes are to be expected at the moment.


## Developing Locally

Sherlock consists of 3 core components. a CLI client, a web server, and a postgres database. To facilitate local development this repo provides utilities to easily spin
up a local development environment consisting of the webserver and the database. The cli client can then be used with this local server and db over `localhost`.
Under the hood this process uses docker-compose to build a sherlock server image with your local changes and then spin it up and connect it to a postgres13 container.
When running the server locally database changelogs will automatically be applied the server on start up so that the database is setup and ready to go.

Spinning up the local environment is as simple as running `make local-up`
Additionally `make local-stop` will stop all the containers and `make local-down` where tear down your local environment completely.
