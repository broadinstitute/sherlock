
# SherlockServiceAlertV3Create


## Properties

Name | Type
------------ | -------------
`link` | string
`message` | string
`onEnvironment` | string
`severity` | string
`title` | string

## Example

```typescript
import type { SherlockServiceAlertV3Create } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "link": null,
  "message": null,
  "onEnvironment": null,
  "severity": null,
  "title": null,
} satisfies SherlockServiceAlertV3Create

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockServiceAlertV3Create
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


