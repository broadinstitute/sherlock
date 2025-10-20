
# SherlockGitCommitV3Upsert


## Properties

Name | Type
------------ | -------------
`committedAt` | string
`gitBranch` | string
`gitCommit` | string
`gitRepo` | string
`isMainBranch` | boolean

## Example

```typescript
import type { SherlockGitCommitV3Upsert } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "committedAt": null,
  "gitBranch": null,
  "gitCommit": null,
  "gitRepo": null,
  "isMainBranch": null,
} satisfies SherlockGitCommitV3Upsert

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockGitCommitV3Upsert
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


