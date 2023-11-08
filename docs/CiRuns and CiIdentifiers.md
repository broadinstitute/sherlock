# CiRuns and CiIdentifiers

## How It Started

Sherlock isn't a CI (continuous integration) system, like GitHub Actions, Jenkins, or Argo Workflows.

But... between Sherlock and Beehive, there's a lot of use cases where we kinda want our internal platform to quack like one.

Beehive started us down this path when it launched. Beehive could update Sherlock with new versions or a new BEE environment, but to make anything happen in our infrastructure, something needed to talk to ArgoCD.

Talking to ArgoCD isn't an easy feat, either. It's multiple steps, requires polling, and complicated auth that we can't make click quite like we can Beehive and Sherlock. This is one of the key things that Thelma is good at... but Thelma is a CLI, and again, Sherlock isn't a CI system.

So Beehive borrows user GitHub auth and fires a workflow dispatch off to GitHub, saying just "sync X please."

## The Problem With GitHub's API

There's one issue with the above. GitHub Actions is super event-based in GitHub's backend. It's _so_ event-based that the response to a workflow dispatch is... nothing. Queued workflows still have IDs and attempt numbers but the workflow dispatch response doesn't include any of that info.

Most folks solve this problem by polling. Except it can actually take a while for the workflow to appear in the list! And it can appear out of order! And you can't access the inputs to the workflow at all, so you might identify the wrong workflow! It's horrible.

Why do we care?

Well, if there's a deployment rolling out to `sam-dev`, we don't really want someone deploying something to `sam-dev` again simultaneously. Honestly maybe not even to all of `dev`, we'd at least want folks to know. Same goes maybe for things other than deployments, like a post-deployment test or something. Somehow, we need to know if there's any CI workflows running against different Sherlock-controlled resources. We want to show a loading spinner in Beehive... and GitHub doesn't really help us with that.

## Full-Send

What GitHub does offer is a webhook system. This doesn't offer any more information than the API, but we can still use this. 

We know we'll need to make a smarter query from Beehive than what GitHub supports ("what GitHub Actions are currently running against the sam-dev Sherlock resource"), so the request will be from Beehive to Sherlock. That means that Sherlock is going to need an understanding of what workflows are currently running, and we don't want Sherlock to get rate limited polling GitHub Actions across a bunch of different repos. Webhooks won't solve the "against the sam-dev Sherlock resource" part, but they can give Sherlock the "what GitHub Actions are currently running" part, which is at least a start.

We'll need a Sherlock data type to represent these workflows. Since we're here, might as well abstract it over GitHub Actions and Argo Workflows, which we might make more use of in the future. This data type is called a CiRun.

A CiRun can only be created or updated via a PUT to `/api/ci-runs/v3` -- it edits the existing CiRun with new status info if it exists, or creates it if it doesn't.

GitHub obviously won't conform to our input spec, and we don't want to expose Sherlock endpoints directly to the open internet either (it's behind Identity-Aware Proxy), so that's where the [sherlock-webhook-proxy](../sherlock-webhook-proxy) comes in. It's a cloud function that eats GitHub webhook payloads, validates them, massages the JSON, and then lobs them to Sherlock's `/api/ci-runs/v3` with basic IAP auth.

The last piece here is that we need to roll out an appropriate webhook config to a _ton_ of repos. [Terraform to the rescue.](https://github.com/broadinstitute/terraform-ap-deployments/blob/master/github/webhook.tf)

For fun and profit and to make sure that this mechanism works, Sherlock exposes [a bunch of GitHub Actions metrics from this CiRuns table](https://grafana.dsp-devops.broadinstitute.org/d/u9zqGE_Vk/github-actions?orgId=1).

## Tell Sherlock What Workflows Affect What

Now we have to tackle answering the actual Beehive query -- we need to solve the "against the sam-dev Sherlock resource" part of "what GitHub Actions are currently running against the sam-dev Sherlock resource."

If you look online, there's [some folks who have figured out a nasty workaround](https://stackoverflow.com/questions/69479400/get-run-id-after-triggering-a-github-workflow-dispatch-event) to GitHub not exposing info about the inputs in the API. You can use the GitHub Actions expressions to template in parts of the workflow input into the _names_ of some of the steps, and _step names_ are technically exposed deep in the API.

But that would still require us to poll, and now we have to worry about weird step names and weirder inputs and, ugh.

Let's see if we can do better. The workflow is going to be passed valid selectors of whatever it affects (it'll be told "leonardo-dev" in an input or something), so let's just make it so that the workflow can tell that to Sherlock, somehow associate it to its own CiRun.

Enter the concept of a CiIdentifier. 

CiIdentifiers have a ["many-to-many" relationship](https://gorm.io/docs/many_to_many.html) with CiRuns. Each CiRun can be related to many CiIdentifiers, and each CiIdentifier can be related to many CiRuns. 

A CiIdentifier has a polymorphic ["has one" relationship](https://gorm.io/docs/has_one.html#Polymorphism-Association) with a bunch of other Sherlock data types. Every CiIdentifier refers to some other Sherlock resource, by that resource's type and ID. Those other resources aren't guaranteed to have a CiIdentifier, it's just possible for them to.

Here's how this works: if you load `/api/ci-identifiers/v3/chart-release/sam-dev`, you'll get either a 404 (if there's no CiIdentifier for sam-dev), or you'll get a CiIdentifier _with a list of associated CiRuns_. Similarly, if you get an individual CiRun from the API, it'll have a list of CiIdentifiers of all the related resources. Those CiIdentifiers don't include much info, just resource type and resource ID, but that's enough to get a list of, say, the IDs of any chart releases related to the workflow.

> What's with this polymorphism? Well, if we didn't do that, we'd need to have a many-to-many relationship from CiRun to at least a half dozen different Sherlock data types. Each of those means a new join table, and when this was originally built, Sherlock used generics with Gorm in a way that would make it very difficult to limit the amount of data loaded. There was a real risk of "request an environment, get 200MB of JSON including every CiRun that ever related to it."
> 
> Instead, when you request an environment, suppose with ID 123, the response possibly includes a field like this:
> ```json
> {
>   "ciIdentifier": {
>     "id": 12345,
>     "resource-type": "environment",
>     "resource-id": 123
>   }
> }
> ```
> 
> This indicates that if you were to request `/api/ci-identifiers/v3/12345` or `/api/ci-identifiers/v3/environment/123` you'd then get a list of the CiRuns. In practice, it's usually easier to just take advantage of selectors and hit it like `/api/ci-identifiers/v3/environment/{name}` and handle the 404, but the important thing is that this is a separate request from the actual environment itself. That means it can be streamed in post-page-load in Beehive or ignored entirely from Thelma.
> 
> Similarly, this means we can request CiRuns in the frontend without suddenly getting a massive bunch of details on everything they affect -- presumably, we already know that if we're requesting it. The indirection here is actually fairly ergonomic, especially when you're being careful to not load extraneous things. Remember that Go doesn't actually support abstract classes or anything, so a `[]CiIdentifier` with a switch statement over a resource type field isn't out of place.

The way we close the loop is we have a [client-report-workflow reusable workflow step](../.github/workflows/client-report-workflow.yaml) made available by this repo. Other workflows in other repos can use it so say "I affect the sam-dev chart release" or whatever else. That step turns around and calls the PUT `/api/ci-runs/v3`, except rather than noting anything about status (fun fact, something you can't read from within an action), it notes those related resources. `/api/ci-runs/v3` will turn around, upsert CiIdentifiers for related resources as necessary, and build the many-to-many association.

That's a lot, but here's what it looks like in practice:

1. GitHub webhook to Sherlock, "this run exists and is queued"
2. GitHub webhook to Sherlock, "this run exists and is running", Sherlock updates the CiRun from step 1
3. Curl from the GitHub Action itself to Sherlock, "this run exists and relates to sam-dev", Sherlock updates the CiRun from step 1 and adds an association to the CiIdentifier for sam-dev
4. GitHub webhook to Sherlock, "this run exists and is terminated at time X with outcome Y", Sherlock updates the CiRun from step 1

Meanwhile, suppose a user has Beehive open, looking at sam-dev. Beehive will hit Sherlock's `/api/ci-identifiers/v3/chart-release/sam-dev` every few seconds, and it'll return the new CiRun as soon as step 3 completes. That means that for basically the entire runtime of the action (save for a few seconds at the start), Beehive can display the loading wheel in the right place, and it'll update quickly enough that it makes sense to users.

## What About Deploy Hooks?

Turns out deploy hooks kinda fall out of this mechanism. If you've got a single bit of code in Sherlock that runs whenever a GitHub Action completes... pretty easy to make it say "does this happen to be terra-github-workflows's sync-release.yaml? If so, run any deploy hooks found on environments or chart releases in this CiRun's related resources."

Deploy hooks are otherwise represented in Sherlock's API are pretty mundane data types. There's a bit of database fun going on to handle the fact that a deploy hook can be either a Slack notification or a GitHub Action (it's a tiny and much simpler version of the polymorphism that CiIdentifiers have).

## What About Slack Notifications?

Slack notifications fall out of this mechanism too. There's some added fields on CiRuns that the client-report-workflow will help set, notifySlackChannelsUponSuccess and notifySlackChannelsUponFailure. Right next to the code that runs deploy hooks there's a bit that says "does this happen to be a success with any Slack channels to notify" and "does this happen to be a failure with any Slack channels to notify."

_This_ is why you don't need to worry about doing weird things with the "needs" field for the client-report-workflow step, to make it run after all the other jobs/steps. All it does is tell Sherlock "when GitHub tells you this workflow is done, notify these channels." The GitHub Action won't complete until all jobs terminate, and Sherlock processes the ask from client-report-workflow synchronously, so there's no race condition.