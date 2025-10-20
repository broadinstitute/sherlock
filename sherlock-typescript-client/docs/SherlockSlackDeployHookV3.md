
# SherlockSlackDeployHookV3


## Properties

Name | Type
------------ | -------------
`createdAt` | Date
`id` | number
`mentionPeople` | boolean
`onChartRelease` | string
`onEnvironment` | string
`onFailure` | boolean
`onSuccess` | boolean
`slackChannel` | string
`updatedAt` | Date

## Example

```typescript
import type { SherlockSlackDeployHookV3 } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "createdAt": null,
  "id": null,
  "mentionPeople": null,
  "onChartRelease": null,
  "onEnvironment": null,
  "onFailure": null,
  "onSuccess": null,
  "slackChannel": null,
  "updatedAt": null,
} satisfies SherlockSlackDeployHookV3

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockSlackDeployHookV3
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


