
# SherlockUserV3DeactivateResponse


## Properties

Name | Type
------------ | -------------
`alreadyDeactivatedEmails` | Array&lt;string&gt;
`newlyDeactivatedEmails` | Array&lt;string&gt;
`notFoundEmails` | Array&lt;string&gt;

## Example

```typescript
import type { SherlockUserV3DeactivateResponse } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "alreadyDeactivatedEmails": null,
  "newlyDeactivatedEmails": null,
  "notFoundEmails": null,
} satisfies SherlockUserV3DeactivateResponse

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockUserV3DeactivateResponse
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


