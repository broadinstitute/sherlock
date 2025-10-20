
# SherlockRoleAssignmentV3


## Properties

Name | Type
------------ | -------------
`expiresAt` | Date
`expiresIn` | string
`roleInfo` | object
`suspended` | boolean
`userInfo` | object

## Example

```typescript
import type { SherlockRoleAssignmentV3 } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "expiresAt": null,
  "expiresIn": null,
  "roleInfo": null,
  "suspended": null,
  "userInfo": null,
} satisfies SherlockRoleAssignmentV3

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockRoleAssignmentV3
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


