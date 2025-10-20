
# SherlockEnvironmentV3Create


## Properties

Name | Type
------------ | -------------
`autoPopulateChartReleases` | boolean
`base` | string
`baseDomain` | string
`defaultCluster` | string
`defaultNamespace` | string
`deleteAfter` | Date
`description` | string
`enableJanitor` | boolean
`helmfileRef` | string
`lifecycle` | string
`name` | string
`namePrefixesDomain` | boolean
`offline` | boolean
`offlineScheduleBeginEnabled` | boolean
`offlineScheduleBeginTime` | Date
`offlineScheduleEndEnabled` | boolean
`offlineScheduleEndTime` | Date
`offlineScheduleEndWeekends` | boolean
`owner` | string
`pactIdentifier` | string
`pagerdutyIntegration` | string
`preventDeletion` | boolean
`requiredRole` | string
`requiresSuitability` | boolean
`serviceBannerBucket` | string
`templateEnvironment` | string
`uniqueResourcePrefix` | string
`valuesName` | string

## Example

```typescript
import type { SherlockEnvironmentV3Create } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "autoPopulateChartReleases": null,
  "base": null,
  "baseDomain": null,
  "defaultCluster": null,
  "defaultNamespace": null,
  "deleteAfter": null,
  "description": null,
  "enableJanitor": null,
  "helmfileRef": null,
  "lifecycle": null,
  "name": null,
  "namePrefixesDomain": null,
  "offline": null,
  "offlineScheduleBeginEnabled": null,
  "offlineScheduleBeginTime": null,
  "offlineScheduleEndEnabled": null,
  "offlineScheduleEndTime": null,
  "offlineScheduleEndWeekends": null,
  "owner": null,
  "pactIdentifier": null,
  "pagerdutyIntegration": null,
  "preventDeletion": null,
  "requiredRole": null,
  "requiresSuitability": null,
  "serviceBannerBucket": null,
  "templateEnvironment": null,
  "uniqueResourcePrefix": null,
  "valuesName": null,
} satisfies SherlockEnvironmentV3Create

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockEnvironmentV3Create
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


