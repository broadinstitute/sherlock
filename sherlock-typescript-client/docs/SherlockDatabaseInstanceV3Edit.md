
# SherlockDatabaseInstanceV3Edit


## Properties

Name | Type
------------ | -------------
`defaultDatabase` | string
`googleProject` | string
`instanceName` | string
`platform` | string

## Example

```typescript
import type { SherlockDatabaseInstanceV3Edit } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "defaultDatabase": null,
  "googleProject": null,
  "instanceName": null,
  "platform": null,
} satisfies SherlockDatabaseInstanceV3Edit

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockDatabaseInstanceV3Edit
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


