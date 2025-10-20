
# SherlockChartReleaseV3


## Properties

Name | Type
------------ | -------------
`appVersionBranch` | string
`appVersionCommit` | string
`appVersionExact` | string
`appVersionFollowChartRelease` | string
`appVersionInfo` | [SherlockAppVersionV3](SherlockAppVersionV3.md)
`appVersionReference` | string
`appVersionResolver` | string
`chart` | string
`chartInfo` | [SherlockChartV3](SherlockChartV3.md)
`chartVersionExact` | string
`chartVersionFollowChartRelease` | string
`chartVersionInfo` | [SherlockChartVersionV3](SherlockChartVersionV3.md)
`chartVersionReference` | string
`chartVersionResolver` | string
`ciIdentifier` | [SherlockCiIdentifierV3](SherlockCiIdentifierV3.md)
`cluster` | string
`clusterInfo` | [SherlockClusterV3](SherlockClusterV3.md)
`createdAt` | Date
`destinationType` | string
`environment` | string
`environmentInfo` | [SherlockEnvironmentV3](SherlockEnvironmentV3.md)
`helmfileRef` | string
`helmfileRefEnabled` | boolean
`id` | number
`includedInBulkChangesets` | boolean
`name` | string
`namespace` | string
`pagerdutyIntegration` | string
`pagerdutyIntegrationInfo` | [SherlockPagerdutyIntegrationV3](SherlockPagerdutyIntegrationV3.md)
`port` | number
`protocol` | string
`resolvedAt` | Date
`subdomain` | string
`updatedAt` | Date

## Example

```typescript
import type { SherlockChartReleaseV3 } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "appVersionBranch": null,
  "appVersionCommit": null,
  "appVersionExact": null,
  "appVersionFollowChartRelease": null,
  "appVersionInfo": null,
  "appVersionReference": null,
  "appVersionResolver": null,
  "chart": null,
  "chartInfo": null,
  "chartVersionExact": null,
  "chartVersionFollowChartRelease": null,
  "chartVersionInfo": null,
  "chartVersionReference": null,
  "chartVersionResolver": null,
  "ciIdentifier": null,
  "cluster": null,
  "clusterInfo": null,
  "createdAt": null,
  "destinationType": null,
  "environment": null,
  "environmentInfo": null,
  "helmfileRef": null,
  "helmfileRefEnabled": null,
  "id": null,
  "includedInBulkChangesets": null,
  "name": null,
  "namespace": null,
  "pagerdutyIntegration": null,
  "pagerdutyIntegrationInfo": null,
  "port": null,
  "protocol": null,
  "resolvedAt": null,
  "subdomain": null,
  "updatedAt": null,
} satisfies SherlockChartReleaseV3

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockChartReleaseV3
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


