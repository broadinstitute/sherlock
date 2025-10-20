
# SherlockUserV3DeactivateRequest


## Properties

Name | Type
------------ | -------------
`suspendEmailHandlesAcrossGoogleWorkspaceDomains` | Array&lt;string&gt;
`userEmailHomeDomain` | string
`userEmails` | Array&lt;string&gt;

## Example

```typescript
import type { SherlockUserV3DeactivateRequest } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "suspendEmailHandlesAcrossGoogleWorkspaceDomains": null,
  "userEmailHomeDomain": null,
  "userEmails": null,
} satisfies SherlockUserV3DeactivateRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockUserV3DeactivateRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


