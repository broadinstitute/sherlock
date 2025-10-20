
# SherlockChangesetV3PlanRequestChartReleaseEntry


## Properties

Name | Type
------------ | -------------
`chartRelease` | string
`followVersionsFromOtherChartRelease` | string
`toAppVersionBranch` | string
`toAppVersionCommit` | string
`toAppVersionExact` | string
`toAppVersionFollowChartRelease` | string
`toAppVersionResolver` | string
`toChartVersionExact` | string
`toChartVersionFollowChartRelease` | string
`toChartVersionResolver` | string
`toHelmfileRef` | string
`toHelmfileRefEnabled` | boolean
`useExactVersionsFromOtherChartRelease` | string

## Example

```typescript
import type { SherlockChangesetV3PlanRequestChartReleaseEntry } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "chartRelease": null,
  "followVersionsFromOtherChartRelease": null,
  "toAppVersionBranch": null,
  "toAppVersionCommit": null,
  "toAppVersionExact": null,
  "toAppVersionFollowChartRelease": null,
  "toAppVersionResolver": null,
  "toChartVersionExact": null,
  "toChartVersionFollowChartRelease": null,
  "toChartVersionResolver": null,
  "toHelmfileRef": null,
  "toHelmfileRefEnabled": null,
  "useExactVersionsFromOtherChartRelease": null,
} satisfies SherlockChangesetV3PlanRequestChartReleaseEntry

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockChangesetV3PlanRequestChartReleaseEntry
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


