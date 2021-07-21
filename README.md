
# golang-project-template

A Repo containing a standard layout and configuration for golang projects.

The project layout defined in this repo is a simplified version of [https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)
with some additional ci/cd setup added in. Some of the patterns defined in that example are overkill for dsp-devops' current needs, however this can easily be updated if needs change.

The basic structure provided in this repo is intended to serve as a starting point when starting a new go project and to avoid boilerplate.

## Usage

Click the `use this template` button in the top right to create a new repo with the desired owner and name using the files and folder structure defined here. Make sure to check the `include all branches` option in order to include the `gh-pages` branch so that code coverage html reports work properly.

Most of the CI/CD worflows included here are intended to be generic with two exceptions.

1. [this line in the dockerfile](https://github.com/broadinstitute/golang-project-template/blob/142d0dc810fa4f3afa68e0a5d37aac03f0c3796f/Dockerfile#L13) which will need to be updated to match the actual name of any executable(s).

## Additional Steps

1. After creating a new repo from the template. Github secrets referenced in the ci/cd jobs need to be created. This can be done automatically using terraform.
   [Instructions here](ttps://docs.google.com/document/d/1JbjV4xjAlSOuZY-2bInatl4av3M-y_LmHQkLYyISYns/edit?usp=sharing). [Example for ci in this repo](https://github.com/broadinstitute/terraform-ap-deployments/blob/master/github/tfvars/broadinstitute-golang-project-template.tfvars)

2. Create an a new image repository in `dsp-artifact-registry` for the ci/cd pipeline to push images to. This can also be done automatically via terraform. [Here is the example for this repo](https://github.com/broadinstitute/terraform-ap-deployments/blob/91715091d935e5f0727d108b371322e8dce19094/dsp-artifact-registry/tfvars/dsp-artifact-registry.tfvars#L11). A similar entry for the new repo just needs to be added to that file
