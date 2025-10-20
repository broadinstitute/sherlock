
# SherlockSlackDeployHookV3Edit


## Properties

Name | Type
------------ | -------------
`mentionPeople` | boolean
`onFailure` | boolean
`onSuccess` | boolean
`slackChannel` | string

## Example

```typescript
import type { SherlockSlackDeployHookV3Edit } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "mentionPeople": null,
  "onFailure": null,
  "onSuccess": null,
  "slackChannel": null,
} satisfies SherlockSlackDeployHookV3Edit

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockSlackDeployHookV3Edit
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


