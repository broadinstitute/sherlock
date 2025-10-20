
# SherlockCiRunV3


## Properties

Name | Type
------------ | -------------
`argoWorkflowsName` | string
`argoWorkflowsNamespace` | string
`argoWorkflowsTemplate` | string
`createdAt` | Date
`githubActionsAttemptNumber` | number
`githubActionsOwner` | string
`githubActionsRepo` | string
`githubActionsRunID` | number
`githubActionsWorkflowPath` | string
`id` | number
`notifySlackChannelsUponFailure` | Array&lt;string&gt;
`notifySlackChannelsUponRetry` | Array&lt;string&gt;
`notifySlackChannelsUponSuccess` | Array&lt;string&gt;
`notifySlackCustomIcon` | string
`platform` | string
`relatedResources` | [Array&lt;SherlockCiIdentifierV3&gt;](SherlockCiIdentifierV3.md)
`resourceStatus` | string
`startedAt` | string
`status` | string
`terminalAt` | string
`terminationHooksDispatchedAt` | Date
`updatedAt` | Date

## Example

```typescript
import type { SherlockCiRunV3 } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "argoWorkflowsName": null,
  "argoWorkflowsNamespace": null,
  "argoWorkflowsTemplate": null,
  "createdAt": null,
  "githubActionsAttemptNumber": null,
  "githubActionsOwner": null,
  "githubActionsRepo": null,
  "githubActionsRunID": null,
  "githubActionsWorkflowPath": null,
  "id": null,
  "notifySlackChannelsUponFailure": null,
  "notifySlackChannelsUponRetry": null,
  "notifySlackChannelsUponSuccess": null,
  "notifySlackCustomIcon": null,
  "platform": null,
  "relatedResources": null,
  "resourceStatus": null,
  "startedAt": null,
  "status": null,
  "terminalAt": null,
  "terminationHooksDispatchedAt": null,
  "updatedAt": null,
} satisfies SherlockCiRunV3

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockCiRunV3
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


