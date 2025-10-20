
# SherlockChangesetV3PlanRequest


## Properties

Name | Type
------------ | -------------
`chartReleases` | [Array&lt;SherlockChangesetV3PlanRequestChartReleaseEntry&gt;](SherlockChangesetV3PlanRequestChartReleaseEntry.md)
`environments` | [Array&lt;SherlockChangesetV3PlanRequestEnvironmentEntry&gt;](SherlockChangesetV3PlanRequestEnvironmentEntry.md)
`recreateChangesets` | Array&lt;number&gt;

## Example

```typescript
import type { SherlockChangesetV3PlanRequest } from '@sherlock-js-client/sherlock'

// TODO: Update the object below with actual values
const example = {
  "chartReleases": null,
  "environments": null,
  "recreateChangesets": null,
} satisfies SherlockChangesetV3PlanRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SherlockChangesetV3PlanRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


