
# SherlockAppVersionV3


## Properties

Name | Type
------------ | -------------
`appVersion` | string
`authoredBy` | string
`authoredByInfo` | [SherlockUserV3](SherlockUserV3.md)
`chart` | string
`chartInfo` | [SherlockChartV3](SherlockChartV3.md)
`ciIdentifier` | [SherlockCiIdentifierV3](SherlockCiIdentifierV3.md)
`createdAt` | Date
`description` | string
`gitBranch` | string
`gitCommit` | string
`id` | number
`parentAppVersion` | string
`parentAppVersionInfo` | object
`updatedAt` | Date

## Example

```typescript
import type { SherlockAppVersionV3 } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "appVersion": null,
  "authoredBy": null,
  "authoredByInfo": null,
  "chart": null,
  "chartInfo": null,
  "ciIdentifier": null,
  "createdAt": null,
  "description": null,
  "gitBranch": null,
  "gitCommit": null,
  "id": null,
  "parentAppVersion": null,
  "parentAppVersionInfo": null,
  "updatedAt": null,
} satisfies SherlockAppVersionV3

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockAppVersionV3
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


