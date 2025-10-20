
# SherlockGithubActionsDeployHookV3Edit


## Properties

Name | Type
------------ | -------------
`githubActionsDefaultRef` | string
`githubActionsOwner` | string
`githubActionsRefBehavior` | string
`githubActionsRepo` | string
`githubActionsWorkflowInputs` | object
`githubActionsWorkflowPath` | string
`onFailure` | boolean
`onSuccess` | boolean

## Example

```typescript
import type { SherlockGithubActionsDeployHookV3Edit } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "githubActionsDefaultRef": null,
  "githubActionsOwner": null,
  "githubActionsRefBehavior": null,
  "githubActionsRepo": null,
  "githubActionsWorkflowInputs": null,
  "githubActionsWorkflowPath": null,
  "onFailure": null,
  "onSuccess": null,
} satisfies SherlockGithubActionsDeployHookV3Edit

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockGithubActionsDeployHookV3Edit
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


