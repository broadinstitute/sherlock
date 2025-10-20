
# SherlockGithubActionsDeployHookV3


## Properties

Name | Type
------------ | -------------
`createdAt` | Date
`githubActionsDefaultRef` | string
`githubActionsOwner` | string
`githubActionsRefBehavior` | string
`githubActionsRepo` | string
`githubActionsWorkflowInputs` | object
`githubActionsWorkflowPath` | string
`id` | number
`onChartRelease` | string
`onEnvironment` | string
`onFailure` | boolean
`onSuccess` | boolean
`updatedAt` | Date

## Example

```typescript
import type { SherlockGithubActionsDeployHookV3 } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "createdAt": null,
  "githubActionsDefaultRef": null,
  "githubActionsOwner": null,
  "githubActionsRefBehavior": null,
  "githubActionsRepo": null,
  "githubActionsWorkflowInputs": null,
  "githubActionsWorkflowPath": null,
  "id": null,
  "onChartRelease": null,
  "onEnvironment": null,
  "onFailure": null,
  "onSuccess": null,
  "updatedAt": null,
} satisfies SherlockGithubActionsDeployHookV3

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockGithubActionsDeployHookV3
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


