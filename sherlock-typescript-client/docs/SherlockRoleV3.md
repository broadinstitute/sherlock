
# SherlockRoleV3


## Properties

Name | Type
------------ | -------------
`assignments` | [Array&lt;SherlockRoleAssignmentV3&gt;](SherlockRoleAssignmentV3.md)
`autoAssignAllUsers` | boolean
`canBeGlassBrokenByRole` | number
`canBeGlassBrokenByRoleInfo` | object
`createdAt` | Date
`defaultGlassBreakDuration` | string
`grantsBroadInstituteGroup` | string
`grantsDevAzureAccount` | boolean
`grantsDevAzureDirectoryRoles` | boolean
`grantsDevAzureGroup` | string
`grantsDevFirecloudFolderOwner` | string
`grantsDevFirecloudGroup` | string
`grantsProdAzureAccount` | boolean
`grantsProdAzureDirectoryRoles` | boolean
`grantsProdAzureGroup` | string
`grantsProdFirecloudFolderOwner` | string
`grantsProdFirecloudGroup` | string
`grantsQaFirecloudFolderOwner` | string
`grantsQaFirecloudGroup` | string
`grantsSherlockSuperAdmin` | boolean
`id` | number
`name` | string
`suspendNonSuitableUsers` | boolean
`updatedAt` | Date

## Example

```typescript
import type { SherlockRoleV3 } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "assignments": null,
  "autoAssignAllUsers": null,
  "canBeGlassBrokenByRole": null,
  "canBeGlassBrokenByRoleInfo": null,
  "createdAt": null,
  "defaultGlassBreakDuration": null,
  "grantsBroadInstituteGroup": null,
  "grantsDevAzureAccount": null,
  "grantsDevAzureDirectoryRoles": null,
  "grantsDevAzureGroup": null,
  "grantsDevFirecloudFolderOwner": null,
  "grantsDevFirecloudGroup": null,
  "grantsProdAzureAccount": null,
  "grantsProdAzureDirectoryRoles": null,
  "grantsProdAzureGroup": null,
  "grantsProdFirecloudFolderOwner": null,
  "grantsProdFirecloudGroup": null,
  "grantsQaFirecloudFolderOwner": null,
  "grantsQaFirecloudGroup": null,
  "grantsSherlockSuperAdmin": null,
  "id": null,
  "name": null,
  "suspendNonSuitableUsers": null,
  "updatedAt": null,
} satisfies SherlockRoleV3

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockRoleV3
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


