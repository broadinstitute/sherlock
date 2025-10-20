
# SherlockChartV3Edit


## Properties

Name | Type
------------ | -------------
`appImageGitMainBranch` | string
`appImageGitRepo` | string
`chartExposesEndpoint` | boolean
`chartRepo` | string
`defaultPort` | number
`defaultProtocol` | string
`defaultSubdomain` | string
`description` | string
`pactParticipant` | boolean
`playbookURL` | string

## Example

```typescript
import type { SherlockChartV3Edit } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "appImageGitMainBranch": null,
  "appImageGitRepo": null,
  "chartExposesEndpoint": null,
  "chartRepo": null,
  "defaultPort": null,
  "defaultProtocol": null,
  "defaultSubdomain": null,
  "description": null,
  "pactParticipant": null,
  "playbookURL": null,
} satisfies SherlockChartV3Edit

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockChartV3Edit
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


