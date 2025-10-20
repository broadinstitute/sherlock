
# MiscVersionResponse


## Properties

Name | Type
------------ | -------------
`buildInfo` | { [key: string]: string; }
`goVersion` | string
`version` | string

## Example

```typescript
import type { MiscVersionResponse } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "buildInfo": null,
  "goVersion": null,
  "version": null,
} satisfies MiscVersionResponse

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as MiscVersionResponse
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


