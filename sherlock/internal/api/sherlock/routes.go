package sherlock

import "github.com/gin-gonic/gin"

// ConfigureRoutes defines Sherlock's primary API surface. It expects to be given a *gin.RouterGroup on `/api` with
// authentication.UserMiddleware and authentication.DbMiddleware already applied. It doesn't do that itself so that
// the caller can share middleware and routes with endpoints defined by other packages.
//
// API routes here should adhere to the following guidelines:
//
//  1. CRUD endpoints should be offered at `<type>/<version>`.
//
//     - <type> should be plural and in skewer-case.
//
//     - <version> should be "v" followed by an integer.
//
//  2. Versions of all endpoints here should start at 3 to avoid confusion with older APIs.
//
//  3. Endpoints selecting an individual entry (e.g. editing, deletion) should consume the entirety of the rest of the
//     URI as their selector, like `<type>/<version>/*selector` rather than `<type>/<version>/:selector`.
//
//     - Many types permit compound or associative selectors that can include slashes. To avoid confusion, types should
//     always handle this and simply error if the selector isn't valid.
//
//     - An example of this behavior is the chart-release type, where valid selectors include name ("leonardo-dev"),
//     environment and chart ("dev/leonardo"), and cluster, namespace, and chart ("terra-dev/terra-dev/leonardo").
//
//  4. Non-CRUD endpoints should be offered at `<type>/procedures/<version>/<behavior>`.
//
//     - This differentiates CRUD and non-CRUD endpoints.
//
//  5. Endpoints should categorize themselves by their type on the Swagger page.
//
//     - This should be the plural PascalCase of the type.
//
//     - This should be done with the @tags directive in the route documentation; see
//     https://github.com/swaggo/swag#api-operation.
//
//  6. If endpoints based on different types are uniquely similar or associated to each other, they may prefix their
//     route with a shared group name, resulting in `<group>/<type>/<version>` and
//     `<group>/<type>/procedures/<version>/<behavior>`.
//
//     - If they do this, the endpoints should categorize themselves by their shared group, instead of their type.
func ConfigureRoutes(apiRouter gin.IRoutes) {
	apiRouter.GET("app-versions/procedures/v3/changelog", appVersionsProceduresV3Changelog)
	apiRouter.GET("app-versions/v3", appVersionsV3List)
	apiRouter.PUT("app-versions/v3", appVersionsV3Upsert)
	apiRouter.PATCH("app-versions/v3/*selector", appVersionsV3Edit)
	apiRouter.GET("app-versions/v3/*selector", appVersionsV3Get)

	apiRouter.POST("changesets/procedures/v3/apply", changesetsProceduresV3Apply)
	apiRouter.GET("changesets/procedures/v3/chart-release-history/*chart-release", changesetsProceduresV3ChartReleaseHistory)
	apiRouter.POST("changesets/procedures/v3/plan", changesetsProceduresV3Plan)
	apiRouter.POST("changesets/procedures/v3/plan-and-apply", changesetsProceduresV3PlanAndApply)
	apiRouter.GET("changesets/procedures/v3/version-history/:version-type/:chart/:version", changesetsProceduresV3VersionHistory)
	apiRouter.GET("changesets/v3", changesetsV3List)
	apiRouter.GET("changesets/v3/:id", changesetsV3Get)

	apiRouter.POST("chart-releases/v3", chartReleasesV3Create)
	apiRouter.GET("chart-releases/v3", chartReleasesV3List)
	apiRouter.DELETE("chart-releases/v3/*selector", chartReleasesV3Delete)
	apiRouter.PATCH("chart-releases/v3/*selector", chartReleasesV3Edit)
	apiRouter.GET("chart-releases/v3/*selector", chartReleasesV3Get)

	apiRouter.GET("chart-versions/procedures/v3/changelog", chartVersionsProceduresV3Changelog)
	apiRouter.GET("chart-versions/v3", chartVersionsV3List)
	apiRouter.PUT("chart-versions/v3", chartVersionsV3Upsert)
	apiRouter.PATCH("chart-versions/v3/*selector", chartVersionsV3Edit)
	apiRouter.GET("chart-versions/v3/*selector", chartVersionsV3Get)

	apiRouter.POST("charts/v3", chartsV3Create)
	apiRouter.GET("charts/v3", chartsV3List)
	apiRouter.DELETE("charts/v3/*selector", chartsV3Delete)
	apiRouter.PATCH("charts/v3/*selector", chartsV3Edit)
	apiRouter.GET("charts/v3/*selector", chartsV3Get)

	apiRouter.GET("ci-identifiers/v3", ciIdentifiersV3List)
	apiRouter.GET("ci-identifiers/v3/*selector", ciIdentifiersV3Get)

	apiRouter.GET("ci-runs/procedures/v3/github-info", ciRunsProceduresV3GithubInfoList)
	apiRouter.GET("ci-runs/v3", ciRunsV3List)
	apiRouter.PUT("ci-runs/v3", ciRunsV3Upsert)
	apiRouter.GET("ci-runs/v3/*selector", ciRunsV3Get)

	apiRouter.POST("clusters/v3", clustersV3Create)
	apiRouter.GET("clusters/v3", clustersV3List)
	apiRouter.DELETE("clusters/v3/*selector", clustersV3Delete)
	apiRouter.PATCH("clusters/v3/*selector", clustersV3Edit)
	apiRouter.GET("clusters/v3/*selector", clustersV3Get)

	apiRouter.POST("database-instances/v3", databaseInstancesV3Create)
	apiRouter.GET("database-instances/v3", databaseInstancesV3List)
	apiRouter.PUT("database-instances/v3", databaseInstancesV3Upsert)
	apiRouter.DELETE("database-instances/v3/*selector", databaseInstancesV3Delete)
	apiRouter.PATCH("database-instances/v3/*selector", databaseInstancesV3Edit)
	apiRouter.GET("database-instances/v3/*selector", databaseInstancesV3Get)

	apiRouter.POST("deploy-hooks/github-actions/procedures/v3/test/*selector", githubActionsDeployHooksV3TestRun)
	apiRouter.POST("deploy-hooks/github-actions/v3", githubActionsDeployHooksV3Create)
	apiRouter.GET("deploy-hooks/github-actions/v3", githubActionsDeployHooksV3List)
	apiRouter.DELETE("deploy-hooks/github-actions/v3/*selector", githubActionsDeployHooksV3Delete)
	apiRouter.PATCH("deploy-hooks/github-actions/v3/*selector", githubActionsDeployHooksV3Edit)
	apiRouter.GET("deploy-hooks/github-actions/v3/*selector", githubActionsDeployHooksV3Get)
	apiRouter.POST("deploy-hooks/slack/procedures/v3/test/*selector", slackDeployHooksV3TestRun)
	apiRouter.POST("deploy-hooks/slack/v3", slackDeployHooksV3Create)
	apiRouter.GET("deploy-hooks/slack/v3", slackDeployHooksV3List)
	apiRouter.DELETE("deploy-hooks/slack/v3/*selector", slackDeployHooksV3Delete)
	apiRouter.PATCH("deploy-hooks/slack/v3/*selector", slackDeployHooksV3Edit)
	apiRouter.GET("deploy-hooks/slack/v3/*selector", slackDeployHooksV3Get)

	apiRouter.POST("environments/v3", environmentsV3Create)
	apiRouter.GET("environments/v3", environmentsV3List)
	apiRouter.DELETE("environments/v3/*selector", environmentsV3Delete)
	apiRouter.PATCH("environments/v3/*selector", environmentsV3Edit)
	apiRouter.GET("environments/v3/*selector", environmentsV3Get)

	apiRouter.GET("github-actions-jobs/v3", githubActionsJobsV3List)
	apiRouter.GET("github-actions-jobs/v3/*selector", githubActionsJobsV3Get)
	apiRouter.PUT("github-actions-jobs/v3", githubActionsJobsV3Upsert)

	apiRouter.PUT("git-commits/v3", gitCommitsV3Upsert)

	apiRouter.POST("incidents/v3", incidentsV3Create)
	apiRouter.GET("incidents/v3", incidentsV3List)
	apiRouter.DELETE("incidents/v3/*selector", incidentsV3Delete)
	apiRouter.PATCH("incidents/v3/*selector", incidentsV3Edit)
	apiRouter.GET("incidents/v3/*selector", incidentsV3Get)

	apiRouter.POST("role-assignments/v3/:role-selector/*user-selector", roleAssignmentsV3Create)
	apiRouter.GET("role-assignments/v3", roleAssignmentsV3List)
	apiRouter.DELETE("role-assignments/v3/:role-selector/*user-selector", roleAssignmentsV3Delete)
	apiRouter.PATCH("role-assignments/v3/:role-selector/*user-selector", roleAssignmentsV3Edit)
	apiRouter.GET("role-assignments/v3/:role-selector/*user-selector", roleAssignmentsV3Get)

	apiRouter.POST("roles/v3", rolesV3Create)
	apiRouter.GET("roles/v3", rolesV3List)
	apiRouter.DELETE("roles/v3/*selector", rolesV3Delete)
	apiRouter.PATCH("roles/v3/*selector", rolesV3Edit)
	apiRouter.GET("roles/v3/*selector", rolesV3Get)

	apiRouter.POST("pagerduty-integrations/procedures/v3/trigger-incident/*selector", pagerdutyIntegrationsProceduresV3TriggerIncident)
	apiRouter.POST("pagerduty-integrations/v3", pagerdutyIntegrationsV3Create)
	apiRouter.GET("pagerduty-integrations/v3", pagerdutyIntegrationsV3List)
	apiRouter.DELETE("pagerduty-integrations/v3/*selector", pagerdutyIntegrationsV3Delete)
	apiRouter.PATCH("pagerduty-integrations/v3/*selector", pagerdutyIntegrationsV3Edit)
	apiRouter.GET("pagerduty-integrations/v3/*selector", pagerdutyIntegrationsV3Get)

	apiRouter.GET("users/v3", usersV3List)
	apiRouter.PUT("users/v3", usersV3Upsert)
	apiRouter.GET("users/v3/*selector", usersV3Get)
}
