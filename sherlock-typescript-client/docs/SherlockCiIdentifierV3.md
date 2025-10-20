
# SherlockCiIdentifierV3


## Properties

Name | Type
------------ | -------------
`ciRuns` | [Array&lt;SherlockCiRunV3&gt;](SherlockCiRunV3.md)
`createdAt` | Date
`id` | number
`resourceID` | number
`resourceStatus` | string
`resourceType` | string
`updatedAt` | Date

## Example

```typescript
import type { SherlockCiIdentifierV3 } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "ciRuns": null,
  "createdAt": null,
  "id": null,
  "resourceID": null,
  "resourceStatus": null,
  "resourceType": null,
  "updatedAt": null,
} satisfies SherlockCiIdentifierV3

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockCiIdentifierV3
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


