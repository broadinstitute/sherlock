
# SherlockDatabaseInstanceV3Create


## Properties

Name | Type
------------ | -------------
`chartRelease` | string
`defaultDatabase` | string
`googleProject` | string
`instanceName` | string
`platform` | string

## Example

```typescript
import type { SherlockDatabaseInstanceV3Create } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "chartRelease": null,
  "defaultDatabase": null,
  "googleProject": null,
  "instanceName": null,
  "platform": null,
} satisfies SherlockDatabaseInstanceV3Create

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockDatabaseInstanceV3Create
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


