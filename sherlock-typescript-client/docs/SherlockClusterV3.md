
# SherlockClusterV3


## Properties

Name | Type
------------ | -------------
`address` | string
`azureSubscription` | string
`base` | string
`ciIdentifier` | [SherlockCiIdentifierV3](SherlockCiIdentifierV3.md)
`createdAt` | Date
`googleProject` | string
`helmfileRef` | string
`id` | number
`location` | string
`name` | string
`provider` | string
`requiredRole` | string
`requiredRoleInfo` | [SherlockRoleV3](SherlockRoleV3.md)
`requiresSuitability` | boolean
`updatedAt` | Date

## Example

```typescript
import type { SherlockClusterV3 } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "address": null,
  "azureSubscription": null,
  "base": null,
  "ciIdentifier": null,
  "createdAt": null,
  "googleProject": null,
  "helmfileRef": null,
  "id": null,
  "location": null,
  "name": null,
  "provider": null,
  "requiredRole": null,
  "requiredRoleInfo": null,
  "requiresSuitability": null,
  "updatedAt": null,
} satisfies SherlockClusterV3

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockClusterV3
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


