
# SherlockChartVersionV3


## Properties

Name | Type
------------ | -------------
`authoredBy` | string
`authoredByInfo` | [SherlockUserV3](SherlockUserV3.md)
`chart` | string
`chartInfo` | [SherlockChartV3](SherlockChartV3.md)
`chartVersion` | string
`ciIdentifier` | [SherlockCiIdentifierV3](SherlockCiIdentifierV3.md)
`createdAt` | Date
`description` | string
`id` | number
`parentChartVersion` | string
`parentChartVersionInfo` | object
`updatedAt` | Date

## Example

```typescript
import type { SherlockChartVersionV3 } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "authoredBy": null,
  "authoredByInfo": null,
  "chart": null,
  "chartInfo": null,
  "chartVersion": null,
  "ciIdentifier": null,
  "createdAt": null,
  "description": null,
  "id": null,
  "parentChartVersion": null,
  "parentChartVersionInfo": null,
  "updatedAt": null,
} satisfies SherlockChartVersionV3

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockChartVersionV3
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


