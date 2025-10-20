
# SherlockClusterV3Create


## Properties

Name | Type
------------ | -------------
`address` | string
`azureSubscription` | string
`base` | string
`googleProject` | string
`helmfileRef` | string
`location` | string
`name` | string
`provider` | string
`requiredRole` | string
`requiresSuitability` | boolean

## Example

```typescript
import type { SherlockClusterV3Create } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "address": null,
  "azureSubscription": null,
  "base": null,
  "googleProject": null,
  "helmfileRef": null,
  "location": null,
  "name": null,
  "provider": null,
  "requiredRole": null,
  "requiresSuitability": null,
} satisfies SherlockClusterV3Create

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockClusterV3Create
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


