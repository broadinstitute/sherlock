
# SherlockEnvironmentV3Edit


## Properties

Name | Type
------------ | -------------
`baseDomain` | string
`defaultCluster` | string
`deleteAfter` | Date
`description` | string
`enableJanitor` | boolean
`helmfileRef` | string
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

## Example

```typescript
import type { SherlockEnvironmentV3Edit } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "baseDomain": null,
  "defaultCluster": null,
  "deleteAfter": null,
  "description": null,
  "enableJanitor": null,
  "helmfileRef": null,
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
} satisfies SherlockEnvironmentV3Edit

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockEnvironmentV3Edit
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


