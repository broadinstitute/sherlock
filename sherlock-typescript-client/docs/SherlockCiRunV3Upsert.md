
# SherlockCiRunV3Upsert


## Properties

Name | Type
------------ | -------------
`appVersions` | Array&lt;string&gt;
`argoWorkflowsName` | string
`argoWorkflowsNamespace` | string
`argoWorkflowsTemplate` | string
`changesets` | Array&lt;string&gt;
`chartReleaseStatuses` | { [key: string]: string; }
`chartReleases` | Array&lt;string&gt;
`chartVersions` | Array&lt;string&gt;
`charts` | Array&lt;string&gt;
`clusters` | Array&lt;string&gt;
`environments` | Array&lt;string&gt;
`githubActionsAttemptNumber` | number
`githubActionsOwner` | string
`githubActionsRepo` | string
`githubActionsRunID` | number
`githubActionsWorkflowPath` | string
`ignoreBadSelectors` | boolean
`notifySlackChannelsUponFailure` | Array&lt;string&gt;
`notifySlackChannelsUponRetry` | Array&lt;string&gt;
`notifySlackChannelsUponSuccess` | Array&lt;string&gt;
`notifySlackCustomIcon` | string
`platform` | string
`relateToChangesetNewVersions` | string
`startedAt` | string
`status` | string
`terminalAt` | string

## Example

```typescript
import type { SherlockCiRunV3Upsert } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "appVersions": null,
  "argoWorkflowsName": null,
  "argoWorkflowsNamespace": null,
  "argoWorkflowsTemplate": null,
  "changesets": null,
  "chartReleaseStatuses": null,
  "chartReleases": null,
  "chartVersions": null,
  "charts": null,
  "clusters": null,
  "environments": null,
  "githubActionsAttemptNumber": null,
  "githubActionsOwner": null,
  "githubActionsRepo": null,
  "githubActionsRunID": null,
  "githubActionsWorkflowPath": null,
  "ignoreBadSelectors": null,
  "notifySlackChannelsUponFailure": null,
  "notifySlackChannelsUponRetry": null,
  "notifySlackChannelsUponSuccess": null,
  "notifySlackCustomIcon": null,
  "platform": null,
  "relateToChangesetNewVersions": null,
  "startedAt": null,
  "status": null,
  "terminalAt": null,
} satisfies SherlockCiRunV3Upsert

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockCiRunV3Upsert
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


