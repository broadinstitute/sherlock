
# SherlockChangesetV3PlanRequestEnvironmentEntry


## Properties

Name | Type
------------ | -------------
`environment` | string
`excludeCharts` | Array&lt;string&gt;
`filterToMatchingBranches` | boolean
`followVersionsFromOtherEnvironment` | string
`includeCharts` | Array&lt;string&gt;
`useExactVersionsFromOtherEnvironment` | string

## Example

```typescript
import type { SherlockChangesetV3PlanRequestEnvironmentEntry } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "environment": null,
  "excludeCharts": null,
  "filterToMatchingBranches": null,
  "followVersionsFromOtherEnvironment": null,
  "includeCharts": null,
  "useExactVersionsFromOtherEnvironment": null,
} satisfies SherlockChangesetV3PlanRequestEnvironmentEntry

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockChangesetV3PlanRequestEnvironmentEntry
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


