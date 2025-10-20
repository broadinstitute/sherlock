
# SherlockChartV3


## Properties

Name | Type
------------ | -------------
`appImageGitMainBranch` | string
`appImageGitRepo` | string
`chartExposesEndpoint` | boolean
`chartRepo` | string
`ciIdentifier` | [SherlockCiIdentifierV3](SherlockCiIdentifierV3.md)
`createdAt` | Date
`defaultPort` | number
`defaultProtocol` | string
`defaultSubdomain` | string
`description` | string
`id` | number
`name` | string
`pactParticipant` | boolean
`playbookURL` | string
`updatedAt` | Date

## Example

```typescript
import type { SherlockChartV3 } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "appImageGitMainBranch": null,
  "appImageGitRepo": null,
  "chartExposesEndpoint": null,
  "chartRepo": null,
  "ciIdentifier": null,
  "createdAt": null,
  "defaultPort": null,
  "defaultProtocol": null,
  "defaultSubdomain": null,
  "description": null,
  "id": null,
  "name": null,
  "pactParticipant": null,
  "playbookURL": null,
  "updatedAt": null,
} satisfies SherlockChartV3

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockChartV3
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


