
# SherlockChangesetV3


## Properties

Name | Type
------------ | -------------
`appliedAt` | Date
`appliedBy` | string
`appliedByInfo` | [SherlockUserV3](SherlockUserV3.md)
`chartRelease` | string
`chartReleaseInfo` | [SherlockChartReleaseV3](SherlockChartReleaseV3.md)
`ciIdentifier` | [SherlockCiIdentifierV3](SherlockCiIdentifierV3.md)
`createdAt` | Date
`fromAppVersionBranch` | string
`fromAppVersionCommit` | string
`fromAppVersionExact` | string
`fromAppVersionFollowChartRelease` | string
`fromAppVersionReference` | string
`fromAppVersionResolver` | string
`fromChartVersionExact` | string
`fromChartVersionFollowChartRelease` | string
`fromChartVersionReference` | string
`fromChartVersionResolver` | string
`fromHelmfileRef` | string
`fromHelmfileRefEnabled` | boolean
`fromResolvedAt` | Date
`id` | number
`newAppVersions` | [Array&lt;SherlockAppVersionV3&gt;](SherlockAppVersionV3.md)
`newChartVersions` | [Array&lt;SherlockChartVersionV3&gt;](SherlockChartVersionV3.md)
`plannedBy` | string
`plannedByInfo` | [SherlockUserV3](SherlockUserV3.md)
`supersededAt` | Date
`toAppVersionBranch` | string
`toAppVersionCommit` | string
`toAppVersionExact` | string
`toAppVersionFollowChartRelease` | string
`toAppVersionReference` | string
`toAppVersionResolver` | string
`toChartVersionExact` | string
`toChartVersionFollowChartRelease` | string
`toChartVersionReference` | string
`toChartVersionResolver` | string
`toHelmfileRef` | string
`toHelmfileRefEnabled` | boolean
`toResolvedAt` | Date
`updatedAt` | Date

## Example

```typescript
import type { SherlockChangesetV3 } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "appliedAt": null,
  "appliedBy": null,
  "appliedByInfo": null,
  "chartRelease": null,
  "chartReleaseInfo": null,
  "ciIdentifier": null,
  "createdAt": null,
  "fromAppVersionBranch": null,
  "fromAppVersionCommit": null,
  "fromAppVersionExact": null,
  "fromAppVersionFollowChartRelease": null,
  "fromAppVersionReference": null,
  "fromAppVersionResolver": null,
  "fromChartVersionExact": null,
  "fromChartVersionFollowChartRelease": null,
  "fromChartVersionReference": null,
  "fromChartVersionResolver": null,
  "fromHelmfileRef": null,
  "fromHelmfileRefEnabled": null,
  "fromResolvedAt": null,
  "id": null,
  "newAppVersions": null,
  "newChartVersions": null,
  "plannedBy": null,
  "plannedByInfo": null,
  "supersededAt": null,
  "toAppVersionBranch": null,
  "toAppVersionCommit": null,
  "toAppVersionExact": null,
  "toAppVersionFollowChartRelease": null,
  "toAppVersionReference": null,
  "toAppVersionResolver": null,
  "toChartVersionExact": null,
  "toChartVersionFollowChartRelease": null,
  "toChartVersionReference": null,
  "toChartVersionResolver": null,
  "toHelmfileRef": null,
  "toHelmfileRefEnabled": null,
  "toResolvedAt": null,
  "updatedAt": null,
} satisfies SherlockChangesetV3

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockChangesetV3
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


