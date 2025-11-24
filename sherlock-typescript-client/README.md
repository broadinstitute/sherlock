# @sherlock-js-client/sherlock@v1.6.71

A TypeScript SDK client for the sherlock.dsp-devops-prod.broadinstitute.org API.

## Usage

First, install the SDK from npm.

```bash
npm install @sherlock-js-client/sherlock --save
```

Next, try it out.


```ts
import {
  Configuration,
  AppVersionsApi,
} from '@sherlock-js-client/sherlock';
import type { ApiAppVersionsProceduresV3ChangelogGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new AppVersionsApi();

  const body = {
    // string | The selector of the newer AppVersion for the changelog
    child: child_example,
    // string | The selector of the older AppVersion for the changelog
    parent: parent_example,
  } satisfies ApiAppVersionsProceduresV3ChangelogGetRequest;

  try {
    const data = await api.apiAppVersionsProceduresV3ChangelogGet(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```


## Documentation

### API Endpoints

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Class | Method | HTTP request | Description
| ----- | ------ | ------------ | -------------
*AppVersionsApi* | [**apiAppVersionsProceduresV3ChangelogGet**](docs/AppVersionsApi.md#apiappversionsproceduresv3changelogget) | **GET** /api/app-versions/procedures/v3/changelog | Get a changelog between two AppVersions
*AppVersionsApi* | [**apiAppVersionsV3Get**](docs/AppVersionsApi.md#apiappversionsv3get) | **GET** /api/app-versions/v3 | List AppVersions matching a filter
*AppVersionsApi* | [**apiAppVersionsV3Put**](docs/AppVersionsApi.md#apiappversionsv3put) | **PUT** /api/app-versions/v3 | Upsert a AppVersion
*AppVersionsApi* | [**apiAppVersionsV3SelectorGet**](docs/AppVersionsApi.md#apiappversionsv3selectorget) | **GET** /api/app-versions/v3/{selector} | Get an individual AppVersion
*AppVersionsApi* | [**apiAppVersionsV3SelectorPatch**](docs/AppVersionsApi.md#apiappversionsv3selectorpatch) | **PATCH** /api/app-versions/v3/{selector} | Edit an individual AppVersion
*ChangesetsApi* | [**apiChangesetsProceduresV3ApplyPost**](docs/ChangesetsApi.md#apichangesetsproceduresv3applypost) | **POST** /api/changesets/procedures/v3/apply | Apply previously planned version changes to Chart Releases
*ChangesetsApi* | [**apiChangesetsProceduresV3ChartReleaseHistoryChartReleaseGet**](docs/ChangesetsApi.md#apichangesetsproceduresv3chartreleasehistorychartreleaseget) | **GET** /api/changesets/procedures/v3/chart-release-history/{chart-release} | List applied Changesets for a Chart Release
*ChangesetsApi* | [**apiChangesetsProceduresV3PlanAndApplyPost**](docs/ChangesetsApi.md#apichangesetsproceduresv3planandapplypost) | **POST** /api/changesets/procedures/v3/plan-and-apply | Plan and apply version changes in one step
*ChangesetsApi* | [**apiChangesetsProceduresV3PlanPost**](docs/ChangesetsApi.md#apichangesetsproceduresv3planpost) | **POST** /api/changesets/procedures/v3/plan | Plan--but do not apply--version changes to Chart Releases
*ChangesetsApi* | [**apiChangesetsProceduresV3VersionHistoryVersionTypeChartVersionGet**](docs/ChangesetsApi.md#apichangesetsproceduresv3versionhistoryversiontypechartversionget) | **GET** /api/changesets/procedures/v3/version-history/{version-type}/{chart}/{version} | List applied Changesets for an App or Chart Version
*ChangesetsApi* | [**apiChangesetsV3Get**](docs/ChangesetsApi.md#apichangesetsv3get) | **GET** /api/changesets/v3 | List Changesets matching a filter
*ChangesetsApi* | [**apiChangesetsV3IdGet**](docs/ChangesetsApi.md#apichangesetsv3idget) | **GET** /api/changesets/v3/{id} | Get an individual Changeset
*ChartReleasesApi* | [**apiChartReleasesV3Get**](docs/ChartReleasesApi.md#apichartreleasesv3get) | **GET** /api/chart-releases/v3 | List ChartReleases matching a filter
*ChartReleasesApi* | [**apiChartReleasesV3Post**](docs/ChartReleasesApi.md#apichartreleasesv3post) | **POST** /api/chart-releases/v3 | Create a ChartRelease
*ChartReleasesApi* | [**apiChartReleasesV3SelectorDelete**](docs/ChartReleasesApi.md#apichartreleasesv3selectordelete) | **DELETE** /api/chart-releases/v3/{selector} | Delete an individual ChartRelease
*ChartReleasesApi* | [**apiChartReleasesV3SelectorGet**](docs/ChartReleasesApi.md#apichartreleasesv3selectorget) | **GET** /api/chart-releases/v3/{selector} | Get an individual ChartRelease
*ChartReleasesApi* | [**apiChartReleasesV3SelectorPatch**](docs/ChartReleasesApi.md#apichartreleasesv3selectorpatch) | **PATCH** /api/chart-releases/v3/{selector} | Edit an individual ChartRelease
*ChartVersionsApi* | [**apiChartVersionsProceduresV3ChangelogGet**](docs/ChartVersionsApi.md#apichartversionsproceduresv3changelogget) | **GET** /api/chart-versions/procedures/v3/changelog | Get a changelog between two ChartVersions
*ChartVersionsApi* | [**apiChartVersionsV3Get**](docs/ChartVersionsApi.md#apichartversionsv3get) | **GET** /api/chart-versions/v3 | List ChartVersions matching a filter
*ChartVersionsApi* | [**apiChartVersionsV3Put**](docs/ChartVersionsApi.md#apichartversionsv3put) | **PUT** /api/chart-versions/v3 | Upsert a ChartVersion
*ChartVersionsApi* | [**apiChartVersionsV3SelectorGet**](docs/ChartVersionsApi.md#apichartversionsv3selectorget) | **GET** /api/chart-versions/v3/{selector} | Get an individual ChartVersion
*ChartVersionsApi* | [**apiChartVersionsV3SelectorPatch**](docs/ChartVersionsApi.md#apichartversionsv3selectorpatch) | **PATCH** /api/chart-versions/v3/{selector} | Edit an individual ChartVersion
*ChartsApi* | [**apiChartsV3Get**](docs/ChartsApi.md#apichartsv3get) | **GET** /api/charts/v3 | List Charts matching a filter
*ChartsApi* | [**apiChartsV3Post**](docs/ChartsApi.md#apichartsv3post) | **POST** /api/charts/v3 | Create a Chart
*ChartsApi* | [**apiChartsV3SelectorDelete**](docs/ChartsApi.md#apichartsv3selectordelete) | **DELETE** /api/charts/v3/{selector} | Delete an individual Chart
*ChartsApi* | [**apiChartsV3SelectorGet**](docs/ChartsApi.md#apichartsv3selectorget) | **GET** /api/charts/v3/{selector} | Get an individual Chart
*ChartsApi* | [**apiChartsV3SelectorPatch**](docs/ChartsApi.md#apichartsv3selectorpatch) | **PATCH** /api/charts/v3/{selector} | Edit an individual Chart
*CiIdentifiersApi* | [**apiCiIdentifiersV3Get**](docs/CiIdentifiersApi.md#apiciidentifiersv3get) | **GET** /api/ci-identifiers/v3 | List CiIdentifiers matching a filter
*CiIdentifiersApi* | [**apiCiIdentifiersV3SelectorGet**](docs/CiIdentifiersApi.md#apiciidentifiersv3selectorget) | **GET** /api/ci-identifiers/v3/{selector} | Get CiRuns for a resource by its CiIdentifier
*CiRunsApi* | [**apiCiRunsProceduresV3GithubInfoGet**](docs/CiRunsApi.md#apicirunsproceduresv3githubinfoget) | **GET** /api/ci-runs/procedures/v3/github-info | List GitHub info gleaned from CiRuns
*CiRunsApi* | [**apiCiRunsV3Get**](docs/CiRunsApi.md#apicirunsv3get) | **GET** /api/ci-runs/v3 | List CiRuns matching a filter
*CiRunsApi* | [**apiCiRunsV3Put**](docs/CiRunsApi.md#apicirunsv3put) | **PUT** /api/ci-runs/v3 | Create or update a CiRun
*CiRunsApi* | [**apiCiRunsV3SelectorGet**](docs/CiRunsApi.md#apicirunsv3selectorget) | **GET** /api/ci-runs/v3/{selector} | Get a CiRun, including CiIdentifiers for related resources
*ClustersApi* | [**apiClustersV3Get**](docs/ClustersApi.md#apiclustersv3get) | **GET** /api/clusters/v3 | List Clusters matching a filter
*ClustersApi* | [**apiClustersV3Post**](docs/ClustersApi.md#apiclustersv3post) | **POST** /api/clusters/v3 | Create a Cluster
*ClustersApi* | [**apiClustersV3SelectorDelete**](docs/ClustersApi.md#apiclustersv3selectordelete) | **DELETE** /api/clusters/v3/{selector} | Delete an individual Cluster
*ClustersApi* | [**apiClustersV3SelectorGet**](docs/ClustersApi.md#apiclustersv3selectorget) | **GET** /api/clusters/v3/{selector} | Get an individual Cluster
*ClustersApi* | [**apiClustersV3SelectorPatch**](docs/ClustersApi.md#apiclustersv3selectorpatch) | **PATCH** /api/clusters/v3/{selector} | Edit an individual Cluster
*DatabaseInstancesApi* | [**apiDatabaseInstancesV3Get**](docs/DatabaseInstancesApi.md#apidatabaseinstancesv3get) | **GET** /api/database-instances/v3 | List DatabaseInstances matching a filter
*DatabaseInstancesApi* | [**apiDatabaseInstancesV3Post**](docs/DatabaseInstancesApi.md#apidatabaseinstancesv3post) | **POST** /api/database-instances/v3 | Create a DatabaseInstance
*DatabaseInstancesApi* | [**apiDatabaseInstancesV3Put**](docs/DatabaseInstancesApi.md#apidatabaseinstancesv3put) | **PUT** /api/database-instances/v3 | Create or edit a DatabaseInstance
*DatabaseInstancesApi* | [**apiDatabaseInstancesV3SelectorDelete**](docs/DatabaseInstancesApi.md#apidatabaseinstancesv3selectordelete) | **DELETE** /api/database-instances/v3/{selector} | Delete an individual DatabaseInstance
*DatabaseInstancesApi* | [**apiDatabaseInstancesV3SelectorGet**](docs/DatabaseInstancesApi.md#apidatabaseinstancesv3selectorget) | **GET** /api/database-instances/v3/{selector} | Get an individual DatabaseInstance
*DatabaseInstancesApi* | [**apiDatabaseInstancesV3SelectorPatch**](docs/DatabaseInstancesApi.md#apidatabaseinstancesv3selectorpatch) | **PATCH** /api/database-instances/v3/{selector} | Edit an individual DatabaseInstance
*DeployHooksApi* | [**apiDeployHooksGithubActionsProceduresV3TestSelectorPost**](docs/DeployHooksApi.md#apideployhooksgithubactionsproceduresv3testselectorpost) | **POST** /api/deploy-hooks/github-actions/procedures/v3/test/{selector} | Test a GithubActionsDeployHook
*DeployHooksApi* | [**apiDeployHooksGithubActionsV3Get**](docs/DeployHooksApi.md#apideployhooksgithubactionsv3get) | **GET** /api/deploy-hooks/github-actions/v3 | List GithubActionsDeployHooks matching a filter
*DeployHooksApi* | [**apiDeployHooksGithubActionsV3Post**](docs/DeployHooksApi.md#apideployhooksgithubactionsv3post) | **POST** /api/deploy-hooks/github-actions/v3 | Create a GithubActionsDeployHook
*DeployHooksApi* | [**apiDeployHooksGithubActionsV3SelectorDelete**](docs/DeployHooksApi.md#apideployhooksgithubactionsv3selectordelete) | **DELETE** /api/deploy-hooks/github-actions/v3/{selector} | Delete an individual GithubActionsDeployHook
*DeployHooksApi* | [**apiDeployHooksGithubActionsV3SelectorGet**](docs/DeployHooksApi.md#apideployhooksgithubactionsv3selectorget) | **GET** /api/deploy-hooks/github-actions/v3/{selector} | Get an individual GithubActionsDeployHook
*DeployHooksApi* | [**apiDeployHooksGithubActionsV3SelectorPatch**](docs/DeployHooksApi.md#apideployhooksgithubactionsv3selectorpatch) | **PATCH** /api/deploy-hooks/github-actions/v3/{selector} | Edit an individual GithubActionsDeployHook
*DeployHooksApi* | [**apiDeployHooksSlackProceduresV3TestSelectorPost**](docs/DeployHooksApi.md#apideployhooksslackproceduresv3testselectorpost) | **POST** /api/deploy-hooks/slack/procedures/v3/test/{selector} | Test a SlackDeployHook
*DeployHooksApi* | [**apiDeployHooksSlackV3Get**](docs/DeployHooksApi.md#apideployhooksslackv3get) | **GET** /api/deploy-hooks/slack/v3 | List SlackDeployHooks matching a filter
*DeployHooksApi* | [**apiDeployHooksSlackV3Post**](docs/DeployHooksApi.md#apideployhooksslackv3post) | **POST** /api/deploy-hooks/slack/v3 | Create a SlackDeployHook
*DeployHooksApi* | [**apiDeployHooksSlackV3SelectorDelete**](docs/DeployHooksApi.md#apideployhooksslackv3selectordelete) | **DELETE** /api/deploy-hooks/slack/v3/{selector} | Delete an individual SlackDeployHook
*DeployHooksApi* | [**apiDeployHooksSlackV3SelectorGet**](docs/DeployHooksApi.md#apideployhooksslackv3selectorget) | **GET** /api/deploy-hooks/slack/v3/{selector} | Get an individual SlackDeployHook
*DeployHooksApi* | [**apiDeployHooksSlackV3SelectorPatch**](docs/DeployHooksApi.md#apideployhooksslackv3selectorpatch) | **PATCH** /api/deploy-hooks/slack/v3/{selector} | Edit an individual SlackDeployHook
*EnvironmentsApi* | [**apiEnvironmentsV3Get**](docs/EnvironmentsApi.md#apienvironmentsv3get) | **GET** /api/environments/v3 | List Environments matching a filter
*EnvironmentsApi* | [**apiEnvironmentsV3Post**](docs/EnvironmentsApi.md#apienvironmentsv3post) | **POST** /api/environments/v3 | Create a Environment
*EnvironmentsApi* | [**apiEnvironmentsV3SelectorDelete**](docs/EnvironmentsApi.md#apienvironmentsv3selectordelete) | **DELETE** /api/environments/v3/{selector} | Delete an individual Environment
*EnvironmentsApi* | [**apiEnvironmentsV3SelectorGet**](docs/EnvironmentsApi.md#apienvironmentsv3selectorget) | **GET** /api/environments/v3/{selector} | Get an individual Environment
*EnvironmentsApi* | [**apiEnvironmentsV3SelectorPatch**](docs/EnvironmentsApi.md#apienvironmentsv3selectorpatch) | **PATCH** /api/environments/v3/{selector} | Edit an individual Environment
*GitCommitsApi* | [**apiGitCommitsV3Put**](docs/GitCommitsApi.md#apigitcommitsv3put) | **PUT** /api/git-commits/v3 | Upsert a GitCommit
*GithubActionsJobsApi* | [**apiGithubActionsJobsV3Get**](docs/GithubActionsJobsApi.md#apigithubactionsjobsv3get) | **GET** /api/github-actions-jobs/v3 | List GithubActionsJobs matching a filter
*GithubActionsJobsApi* | [**apiGithubActionsJobsV3Put**](docs/GithubActionsJobsApi.md#apigithubactionsjobsv3put) | **PUT** /api/github-actions-jobs/v3 | Upsert GithubActionsJob
*GithubActionsJobsApi* | [**apiGithubActionsJobsV3SelectorGet**](docs/GithubActionsJobsApi.md#apigithubactionsjobsv3selectorget) | **GET** /api/github-actions-jobs/v3/{selector} | Get an individual GithubActionsJob
*IncidentsApi* | [**apiIncidentsV3Get**](docs/IncidentsApi.md#apiincidentsv3get) | **GET** /api/incidents/v3 | List Incidents matching a filter
*IncidentsApi* | [**apiIncidentsV3Post**](docs/IncidentsApi.md#apiincidentsv3post) | **POST** /api/incidents/v3 | Create a Incident
*IncidentsApi* | [**apiIncidentsV3SelectorDelete**](docs/IncidentsApi.md#apiincidentsv3selectordelete) | **DELETE** /api/incidents/v3/{selector} | Delete an individual Incident
*IncidentsApi* | [**apiIncidentsV3SelectorGet**](docs/IncidentsApi.md#apiincidentsv3selectorget) | **GET** /api/incidents/v3/{selector} | Get an individual Incident
*IncidentsApi* | [**apiIncidentsV3SelectorPatch**](docs/IncidentsApi.md#apiincidentsv3selectorpatch) | **PATCH** /api/incidents/v3/{selector} | Edit an individual Incident
*MiscApi* | [**connectionCheckGet**](docs/MiscApi.md#connectioncheckget) | **GET** /connection-check | Test the client\&#39;s connection to Sherlock
*MiscApi* | [**statusGet**](docs/MiscApi.md#statusget) | **GET** /status | Get Sherlock\&#39;s current status
*MiscApi* | [**versionGet**](docs/MiscApi.md#versionget) | **GET** /version | Get Sherlock\&#39;s own current version
*PagerdutyIntegrationsApi* | [**apiPagerdutyIntegrationsProceduresV3TriggerIncidentSelectorPost**](docs/PagerdutyIntegrationsApi.md#apipagerdutyintegrationsproceduresv3triggerincidentselectorpost) | **POST** /api/pagerduty-integrations/procedures/v3/trigger-incident/{selector} | Get an individual PagerdutyIntegration
*PagerdutyIntegrationsApi* | [**apiPagerdutyIntegrationsV3Get**](docs/PagerdutyIntegrationsApi.md#apipagerdutyintegrationsv3get) | **GET** /api/pagerduty-integrations/v3 | List PagerdutyIntegrations matching a filter
*PagerdutyIntegrationsApi* | [**apiPagerdutyIntegrationsV3Post**](docs/PagerdutyIntegrationsApi.md#apipagerdutyintegrationsv3post) | **POST** /api/pagerduty-integrations/v3 | Create a PagerdutyIntegration
*PagerdutyIntegrationsApi* | [**apiPagerdutyIntegrationsV3SelectorDelete**](docs/PagerdutyIntegrationsApi.md#apipagerdutyintegrationsv3selectordelete) | **DELETE** /api/pagerduty-integrations/v3/{selector} | Delete an individual PagerdutyIntegration
*PagerdutyIntegrationsApi* | [**apiPagerdutyIntegrationsV3SelectorGet**](docs/PagerdutyIntegrationsApi.md#apipagerdutyintegrationsv3selectorget) | **GET** /api/pagerduty-integrations/v3/{selector} | Get an individual PagerdutyIntegration
*PagerdutyIntegrationsApi* | [**apiPagerdutyIntegrationsV3SelectorPatch**](docs/PagerdutyIntegrationsApi.md#apipagerdutyintegrationsv3selectorpatch) | **PATCH** /api/pagerduty-integrations/v3/{selector} | Edit an individual PagerdutyIntegration
*RoleAssignmentsApi* | [**apiRoleAssignmentsV3Get**](docs/RoleAssignmentsApi.md#apiroleassignmentsv3get) | **GET** /api/role-assignments/v3 | List RoleAssignments matching a filter
*RoleAssignmentsApi* | [**apiRoleAssignmentsV3RoleSelectorUserSelectorDelete**](docs/RoleAssignmentsApi.md#apiroleassignmentsv3roleselectoruserselectordelete) | **DELETE** /api/role-assignments/v3/{role-selector}/{user-selector} | Delete a RoleAssignment
*RoleAssignmentsApi* | [**apiRoleAssignmentsV3RoleSelectorUserSelectorGet**](docs/RoleAssignmentsApi.md#apiroleassignmentsv3roleselectoruserselectorget) | **GET** /api/role-assignments/v3/{role-selector}/{user-selector} | Get a RoleAssignment
*RoleAssignmentsApi* | [**apiRoleAssignmentsV3RoleSelectorUserSelectorPatch**](docs/RoleAssignmentsApi.md#apiroleassignmentsv3roleselectoruserselectorpatch) | **PATCH** /api/role-assignments/v3/{role-selector}/{user-selector} | Edit a RoleAssignment
*RoleAssignmentsApi* | [**apiRoleAssignmentsV3RoleSelectorUserSelectorPost**](docs/RoleAssignmentsApi.md#apiroleassignmentsv3roleselectoruserselectorpost) | **POST** /api/role-assignments/v3/{role-selector}/{user-selector} | Create a RoleAssignment
*RolesApi* | [**apiRolesV3Get**](docs/RolesApi.md#apirolesv3get) | **GET** /api/roles/v3 | List Roles matching a filter
*RolesApi* | [**apiRolesV3Post**](docs/RolesApi.md#apirolesv3post) | **POST** /api/roles/v3 | Create a Role
*RolesApi* | [**apiRolesV3SelectorDelete**](docs/RolesApi.md#apirolesv3selectordelete) | **DELETE** /api/roles/v3/{selector} | Delete a Role
*RolesApi* | [**apiRolesV3SelectorGet**](docs/RolesApi.md#apirolesv3selectorget) | **GET** /api/roles/v3/{selector} | Get a Role
*RolesApi* | [**apiRolesV3SelectorPatch**](docs/RolesApi.md#apirolesv3selectorpatch) | **PATCH** /api/roles/v3/{selector} | Edit a Role
*ServiceAlertApi* | [**apiServiceAlertsProceduresV3SyncPost**](docs/ServiceAlertApi.md#apiservicealertsproceduresv3syncpost) | **POST** /api/service-alerts/procedures/v3/sync | Sync service alerts
*ServiceAlertApi* | [**apiServiceAlertsV3Get**](docs/ServiceAlertApi.md#apiservicealertsv3get) | **GET** /api/service-alerts/v3 | List ServiceAlerts matching a filter
*ServiceAlertApi* | [**apiServiceAlertsV3Post**](docs/ServiceAlertApi.md#apiservicealertsv3post) | **POST** /api/service-alerts/v3 | Create a service alert
*ServiceAlertApi* | [**apiServiceAlertsV3SelectorDelete**](docs/ServiceAlertApi.md#apiservicealertsv3selectordelete) | **DELETE** /api/service-alerts/v3/{selector} | Delete a ServiceAlert
*ServiceAlertApi* | [**apiServiceAlertsV3SelectorGet**](docs/ServiceAlertApi.md#apiservicealertsv3selectorget) | **GET** /api/service-alerts/v3/{selector} | Get a Service Alert
*ServiceAlertApi* | [**apiServiceAlertsV3SelectorPatch**](docs/ServiceAlertApi.md#apiservicealertsv3selectorpatch) | **PATCH** /api/service-alerts/v3/{selector} | Edit a service alert
*UsersApi* | [**apiUsersProceduresV3DeactivatePost**](docs/UsersApi.md#apiusersproceduresv3deactivatepost) | **POST** /api/users/procedures/v3/deactivate | Deactivate Users
*UsersApi* | [**apiUsersV3Get**](docs/UsersApi.md#apiusersv3get) | **GET** /api/users/v3 | List Users matching a filter
*UsersApi* | [**apiUsersV3Put**](docs/UsersApi.md#apiusersv3put) | **PUT** /api/users/v3 | Update the calling User\&#39;s information
*UsersApi* | [**apiUsersV3SelectorGet**](docs/UsersApi.md#apiusersv3selectorget) | **GET** /api/users/v3/{selector} | Get an individual User


### Models

- [ErrorsErrorResponse](docs/ErrorsErrorResponse.md)
- [MiscConnectionCheckResponse](docs/MiscConnectionCheckResponse.md)
- [MiscStatusResponse](docs/MiscStatusResponse.md)
- [MiscVersionResponse](docs/MiscVersionResponse.md)
- [PagerdutyAlertSummary](docs/PagerdutyAlertSummary.md)
- [PagerdutySendAlertResponse](docs/PagerdutySendAlertResponse.md)
- [SherlockAppVersionV3](docs/SherlockAppVersionV3.md)
- [SherlockAppVersionV3ChangelogResponse](docs/SherlockAppVersionV3ChangelogResponse.md)
- [SherlockAppVersionV3Create](docs/SherlockAppVersionV3Create.md)
- [SherlockAppVersionV3Edit](docs/SherlockAppVersionV3Edit.md)
- [SherlockChangesetV3](docs/SherlockChangesetV3.md)
- [SherlockChangesetV3PlanRequest](docs/SherlockChangesetV3PlanRequest.md)
- [SherlockChangesetV3PlanRequestChartReleaseEntry](docs/SherlockChangesetV3PlanRequestChartReleaseEntry.md)
- [SherlockChangesetV3PlanRequestEnvironmentEntry](docs/SherlockChangesetV3PlanRequestEnvironmentEntry.md)
- [SherlockChartReleaseV3](docs/SherlockChartReleaseV3.md)
- [SherlockChartReleaseV3Create](docs/SherlockChartReleaseV3Create.md)
- [SherlockChartReleaseV3Edit](docs/SherlockChartReleaseV3Edit.md)
- [SherlockChartV3](docs/SherlockChartV3.md)
- [SherlockChartV3Create](docs/SherlockChartV3Create.md)
- [SherlockChartV3Edit](docs/SherlockChartV3Edit.md)
- [SherlockChartVersionV3](docs/SherlockChartVersionV3.md)
- [SherlockChartVersionV3ChangelogResponse](docs/SherlockChartVersionV3ChangelogResponse.md)
- [SherlockChartVersionV3Create](docs/SherlockChartVersionV3Create.md)
- [SherlockChartVersionV3Edit](docs/SherlockChartVersionV3Edit.md)
- [SherlockCiIdentifierV3](docs/SherlockCiIdentifierV3.md)
- [SherlockCiRunV3](docs/SherlockCiRunV3.md)
- [SherlockCiRunV3Upsert](docs/SherlockCiRunV3Upsert.md)
- [SherlockClusterV3](docs/SherlockClusterV3.md)
- [SherlockClusterV3Create](docs/SherlockClusterV3Create.md)
- [SherlockClusterV3Edit](docs/SherlockClusterV3Edit.md)
- [SherlockDatabaseInstanceV3](docs/SherlockDatabaseInstanceV3.md)
- [SherlockDatabaseInstanceV3Create](docs/SherlockDatabaseInstanceV3Create.md)
- [SherlockDatabaseInstanceV3Edit](docs/SherlockDatabaseInstanceV3Edit.md)
- [SherlockEnvironmentV3](docs/SherlockEnvironmentV3.md)
- [SherlockEnvironmentV3Create](docs/SherlockEnvironmentV3Create.md)
- [SherlockEnvironmentV3Edit](docs/SherlockEnvironmentV3Edit.md)
- [SherlockGitCommitV3](docs/SherlockGitCommitV3.md)
- [SherlockGitCommitV3Upsert](docs/SherlockGitCommitV3Upsert.md)
- [SherlockGithubActionsDeployHookTestRunRequest](docs/SherlockGithubActionsDeployHookTestRunRequest.md)
- [SherlockGithubActionsDeployHookTestRunResponse](docs/SherlockGithubActionsDeployHookTestRunResponse.md)
- [SherlockGithubActionsDeployHookV3](docs/SherlockGithubActionsDeployHookV3.md)
- [SherlockGithubActionsDeployHookV3Create](docs/SherlockGithubActionsDeployHookV3Create.md)
- [SherlockGithubActionsDeployHookV3Edit](docs/SherlockGithubActionsDeployHookV3Edit.md)
- [SherlockGithubActionsJobV3](docs/SherlockGithubActionsJobV3.md)
- [SherlockGithubActionsJobV3Create](docs/SherlockGithubActionsJobV3Create.md)
- [SherlockIncidentV3](docs/SherlockIncidentV3.md)
- [SherlockIncidentV3Create](docs/SherlockIncidentV3Create.md)
- [SherlockIncidentV3Edit](docs/SherlockIncidentV3Edit.md)
- [SherlockPagerdutyIntegrationV3](docs/SherlockPagerdutyIntegrationV3.md)
- [SherlockPagerdutyIntegrationV3Create](docs/SherlockPagerdutyIntegrationV3Create.md)
- [SherlockPagerdutyIntegrationV3Edit](docs/SherlockPagerdutyIntegrationV3Edit.md)
- [SherlockRoleAssignmentV3](docs/SherlockRoleAssignmentV3.md)
- [SherlockRoleAssignmentV3Edit](docs/SherlockRoleAssignmentV3Edit.md)
- [SherlockRoleV3](docs/SherlockRoleV3.md)
- [SherlockRoleV3Edit](docs/SherlockRoleV3Edit.md)
- [SherlockServiceAlertV3](docs/SherlockServiceAlertV3.md)
- [SherlockServiceAlertV3Create](docs/SherlockServiceAlertV3Create.md)
- [SherlockServiceAlertV3EditableFields](docs/SherlockServiceAlertV3EditableFields.md)
- [SherlockServiceAlertV3SyncRequest](docs/SherlockServiceAlertV3SyncRequest.md)
- [SherlockSlackDeployHookTestRunRequest](docs/SherlockSlackDeployHookTestRunRequest.md)
- [SherlockSlackDeployHookTestRunResponse](docs/SherlockSlackDeployHookTestRunResponse.md)
- [SherlockSlackDeployHookV3](docs/SherlockSlackDeployHookV3.md)
- [SherlockSlackDeployHookV3Create](docs/SherlockSlackDeployHookV3Create.md)
- [SherlockSlackDeployHookV3Edit](docs/SherlockSlackDeployHookV3Edit.md)
- [SherlockUserV3](docs/SherlockUserV3.md)
- [SherlockUserV3DeactivateRequest](docs/SherlockUserV3DeactivateRequest.md)
- [SherlockUserV3DeactivateResponse](docs/SherlockUserV3DeactivateResponse.md)
- [SherlockUserV3Upsert](docs/SherlockUserV3Upsert.md)

### Authorization

Endpoints do not require authorization.


## About

This TypeScript SDK client supports the [Fetch API](https://fetch.spec.whatwg.org/)
and is automatically generated by the
[OpenAPI Generator](https://openapi-generator.tech) project:

- API version: `development`
- Package version: `v1.6.71`
- Generator version: `7.18.0-SNAPSHOT`
- Build package: `org.openapitools.codegen.languages.TypeScriptFetchClientCodegen`

The generated npm module supports the following:

- Environments
  * Node.js
  * Webpack
  * Browserify
- Language levels
  * ES5 - you must have a Promises/A+ library installed
  * ES6
- Module systems
  * CommonJS
  * ES6 module system


## Development

### Building

To build the TypeScript source code, you need to have Node.js and npm installed.
After cloning the repository, navigate to the project directory and run:

```bash
npm install
npm run build
```

### Publishing

Once you've built the package, you can publish it to npm:

```bash
npm publish
```

## License

[BSD-3-Clause](https://github.com/broadinstitute/sherlock/blob/main/LICENSE.txt)
