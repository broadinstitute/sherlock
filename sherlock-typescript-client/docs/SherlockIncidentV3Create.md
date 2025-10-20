
# SherlockIncidentV3Create


## Properties

Name | Type
------------ | -------------
`description` | string
`remediatedAt` | string
`reviewCompletedAt` | string
`startedAt` | string
`ticket` | string

## Example

```typescript
import type { SherlockIncidentV3Create } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "description": null,
  "remediatedAt": null,
  "reviewCompletedAt": null,
  "startedAt": null,
  "ticket": null,
} satisfies SherlockIncidentV3Create

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockIncidentV3Create
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


