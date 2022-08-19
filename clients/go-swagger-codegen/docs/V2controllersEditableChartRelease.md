# V2controllersEditableChartRelease

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentAppVersionExact** | **string** |  | [optional] [default to null]
**CurrentChartVersionExact** | **string** |  | [optional] [default to null]
**HelmfileRef** | **string** |  | [optional] [default to null]
**TargetAppVersionBranch** | **string** | When creating, will default to the app&#39;s main branch if it has one recorded | [optional] [default to null]
**TargetAppVersionCommit** | **string** |  | [optional] [default to null]
**TargetAppVersionExact** | **string** |  | [optional] [default to null]
**TargetAppVersionUse** | **string** | When creating, will default to referencing any provided target app version field (exact, then commit, then branch) | [optional] [default to null]
**TargetChartVersionExact** | **string** |  | [optional] [default to null]
**TargetChartVersionUse** | **string** | When creating, will default to latest unless an exact target chart version is provided | [optional] [default to null]
**ThelmaMode** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


