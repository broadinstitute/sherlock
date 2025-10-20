
# SherlockAppVersionV3Create


## Properties

Name | Type
------------ | -------------
`appVersion` | string
`chart` | string
`description` | string
`gitBranch` | string
`gitCommit` | string
`parentAppVersion` | string

## Example

```typescript
import type { SherlockAppVersionV3Create } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "appVersion": null,
  "chart": null,
  "description": null,
  "gitBranch": null,
  "gitCommit": null,
  "parentAppVersion": null,
} satisfies SherlockAppVersionV3Create

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockAppVersionV3Create
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


