
# SherlockGithubActionsJobV3


## Properties

Name | Type
------------ | -------------
`createdAt` | Date
`githubActionsAttemptNumber` | number
`githubActionsJobID` | number
`githubActionsOwner` | string
`githubActionsRepo` | string
`githubActionsRunID` | number
`id` | number
`jobCreatedAt` | Date
`jobStartedAt` | Date
`jobTerminalAt` | Date
`status` | string
`updatedAt` | Date

## Example

```typescript
import type { SherlockGithubActionsJobV3 } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "createdAt": null,
  "githubActionsAttemptNumber": null,
  "githubActionsJobID": null,
  "githubActionsOwner": null,
  "githubActionsRepo": null,
  "githubActionsRunID": null,
  "id": null,
  "jobCreatedAt": null,
  "jobStartedAt": null,
  "jobTerminalAt": null,
  "status": null,
  "updatedAt": null,
} satisfies SherlockGithubActionsJobV3

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockGithubActionsJobV3
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


