
# ErrorsErrorResponse


## Properties

Name | Type
------------ | -------------
`message` | string
`toBlame` | string
`type` | string

## Example

```typescript
import type { ErrorsErrorResponse } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "message": null,
  "toBlame": null,
  "type": null,
} satisfies ErrorsErrorResponse

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ErrorsErrorResponse
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


