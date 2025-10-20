
# SherlockUserV3


## Properties

Name | Type
------------ | -------------
`assignments` | [Array&lt;SherlockRoleAssignmentV3&gt;](SherlockRoleAssignmentV3.md)
`createdAt` | Date
`deactivatedAt` | string
`email` | string
`githubID` | string
`githubUsername` | string
`googleID` | string
`id` | number
`name` | string
`nameFrom` | string
`slackID` | string
`slackUsername` | string
`suitabilityDescription` | string
`suitable` | boolean
`updatedAt` | Date

## Example

```typescript
import type { SherlockUserV3 } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "assignments": null,
  "createdAt": null,
  "deactivatedAt": null,
  "email": null,
  "githubID": null,
  "githubUsername": null,
  "googleID": null,
  "id": null,
  "name": null,
  "nameFrom": null,
  "slackID": null,
  "slackUsername": null,
  "suitabilityDescription": null,
  "suitable": null,
  "updatedAt": null,
} satisfies SherlockUserV3

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockUserV3
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


