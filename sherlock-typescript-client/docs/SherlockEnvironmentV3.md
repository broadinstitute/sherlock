
# SherlockEnvironmentV3


## Properties

Name | Type
------------ | -------------
`autoPopulateChartReleases` | boolean
`base` | string
`baseDomain` | string
`ciIdentifier` | [SherlockCiIdentifierV3](SherlockCiIdentifierV3.md)
`createdAt` | Date
`defaultCluster` | string
`defaultClusterInfo` | [SherlockClusterV3](SherlockClusterV3.md)
`defaultNamespace` | string
`deleteAfter` | Date
`description` | string
`enableJanitor` | boolean
`helmfileRef` | string
`id` | number
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
`ownerInfo` | [SherlockUserV3](SherlockUserV3.md)
`pactIdentifier` | string
`pagerdutyIntegration` | string
`pagerdutyIntegrationInfo` | [SherlockPagerdutyIntegrationV3](SherlockPagerdutyIntegrationV3.md)
`preventDeletion` | boolean
`requiredRole` | string
`requiredRoleInfo` | [SherlockRoleV3](SherlockRoleV3.md)
`requiresSuitability` | boolean
`serviceBannerBucket` | string
`templateEnvironment` | string
`templateEnvironmentInfo` | object
`uniqueResourcePrefix` | string
`updatedAt` | Date
`valuesName` | string

## Example

```typescript
import type { SherlockEnvironmentV3 } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "autoPopulateChartReleases": null,
  "base": null,
  "baseDomain": null,
  "ciIdentifier": null,
  "createdAt": null,
  "defaultCluster": null,
  "defaultClusterInfo": null,
  "defaultNamespace": null,
  "deleteAfter": null,
  "description": null,
  "enableJanitor": null,
  "helmfileRef": null,
  "id": null,
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
  "ownerInfo": null,
  "pactIdentifier": null,
  "pagerdutyIntegration": null,
  "pagerdutyIntegrationInfo": null,
  "preventDeletion": null,
  "requiredRole": null,
  "requiredRoleInfo": null,
  "requiresSuitability": null,
  "serviceBannerBucket": null,
  "templateEnvironment": null,
  "templateEnvironmentInfo": null,
  "uniqueResourcePrefix": null,
  "updatedAt": null,
  "valuesName": null,
} satisfies SherlockEnvironmentV3

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockEnvironmentV3
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


