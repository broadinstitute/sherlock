
# SherlockRoleV3Edit


## Properties

Name | Type
------------ | -------------
`autoAssignAllUsers` | boolean
`canBeGlassBrokenByRole` | number
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
`name` | string
`suspendNonSuitableUsers` | boolean

## Example

```typescript
import type { SherlockRoleV3Edit } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "autoAssignAllUsers": null,
  "canBeGlassBrokenByRole": null,
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
  "name": null,
  "suspendNonSuitableUsers": null,
} satisfies SherlockRoleV3Edit

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockRoleV3Edit
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


