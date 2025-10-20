
# SherlockGitCommitV3


## Properties

Name | Type
------------ | -------------
`committedAt` | string
`createdAt` | Date
`gitBranch` | string
`gitCommit` | string
`gitRepo` | string
`id` | number
`isMainBranch` | boolean
`secSincePrev` | number
`updatedAt` | Date

## Example

```typescript
import type { SherlockGitCommitV3 } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "committedAt": null,
  "createdAt": null,
  "gitBranch": null,
  "gitCommit": null,
  "gitRepo": null,
  "id": null,
  "isMainBranch": null,
  "secSincePrev": null,
  "updatedAt": null,
} satisfies SherlockGitCommitV3

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockGitCommitV3
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


