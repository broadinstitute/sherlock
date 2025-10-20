
# SherlockServiceAlertV3


## Properties

Name | Type
------------ | -------------
`createdAt` | Date
`createdBy` | string
`deletedBy` | string
`id` | number
`link` | string
`message` | string
`onEnvironment` | string
`severity` | string
`title` | string
`updatedAt` | Date
`updatedBy` | string
`uuid` | string

## Example

```typescript
import type { SherlockServiceAlertV3 } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "createdAt": null,
  "createdBy": null,
  "deletedBy": null,
  "id": null,
  "link": null,
  "message": null,
  "onEnvironment": null,
  "severity": null,
  "title": null,
  "updatedAt": null,
  "updatedBy": null,
  "uuid": null,
} satisfies SherlockServiceAlertV3

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockServiceAlertV3
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


