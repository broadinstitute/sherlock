
# SherlockIncidentV3


## Properties

Name | Type
------------ | -------------
`createdAt` | Date
`description` | string
`id` | number
`remediatedAt` | string
`reviewCompletedAt` | string
`startedAt` | string
`ticket` | string
`updatedAt` | Date

## Example

```typescript
import type { SherlockIncidentV3 } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "createdAt": null,
  "description": null,
  "id": null,
  "remediatedAt": null,
  "reviewCompletedAt": null,
  "startedAt": null,
  "ticket": null,
  "updatedAt": null,
} satisfies SherlockIncidentV3

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockIncidentV3
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


