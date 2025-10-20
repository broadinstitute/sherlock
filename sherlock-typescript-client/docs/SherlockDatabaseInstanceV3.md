
# SherlockDatabaseInstanceV3


## Properties

Name | Type
------------ | -------------
`chartRelease` | string
`chartReleaseInfo` | [SherlockChartReleaseV3](SherlockChartReleaseV3.md)
`createdAt` | Date
`defaultDatabase` | string
`googleProject` | string
`id` | number
`instanceName` | string
`platform` | string
`updatedAt` | Date

## Example

```typescript
import type { SherlockDatabaseInstanceV3 } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "chartRelease": null,
  "chartReleaseInfo": null,
  "createdAt": null,
  "defaultDatabase": null,
  "googleProject": null,
  "id": null,
  "instanceName": null,
  "platform": null,
  "updatedAt": null,
} satisfies SherlockDatabaseInstanceV3

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockDatabaseInstanceV3
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


