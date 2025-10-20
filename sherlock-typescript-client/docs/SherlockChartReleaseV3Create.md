
# SherlockChartReleaseV3Create


## Properties

Name | Type
------------ | -------------
`appVersionBranch` | string
`appVersionCommit` | string
`appVersionExact` | string
`appVersionFollowChartRelease` | string
`appVersionResolver` | string
`chart` | string
`chartVersionExact` | string
`chartVersionFollowChartRelease` | string
`chartVersionResolver` | string
`cluster` | string
`environment` | string
`helmfileRef` | string
`helmfileRefEnabled` | boolean
`includedInBulkChangesets` | boolean
`name` | string
`namespace` | string
`pagerdutyIntegration` | string
`port` | number
`protocol` | string
`subdomain` | string

## Example

```typescript
import type { SherlockChartReleaseV3Create } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "appVersionBranch": null,
  "appVersionCommit": null,
  "appVersionExact": null,
  "appVersionFollowChartRelease": null,
  "appVersionResolver": null,
  "chart": null,
  "chartVersionExact": null,
  "chartVersionFollowChartRelease": null,
  "chartVersionResolver": null,
  "cluster": null,
  "environment": null,
  "helmfileRef": null,
  "helmfileRefEnabled": null,
  "includedInBulkChangesets": null,
  "name": null,
  "namespace": null,
  "pagerdutyIntegration": null,
  "port": null,
  "protocol": null,
  "subdomain": null,
} satisfies SherlockChartReleaseV3Create

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockChartReleaseV3Create
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


