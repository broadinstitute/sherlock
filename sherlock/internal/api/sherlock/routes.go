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
	apiRouter.GET("app-versions/v3/*selector", appVersionsV3Get)
	apiRouter.GET("app-versions/v3", appVersionsV3List)
	apiRouter.PATCH("app-versions/v3/*selector", appVersionsV3Edit)
	apiRouter.PUT("app-versions/v3", appVersionsV3Upsert)
	apiRouter.GET("app-versions/procedures/v3/changelog", appVersionsProceduresV3Changelog)

	apiRouter.GET("charts/v3/*selector", chartsV3Get)
	apiRouter.POST("charts/v3", chartsV3Create)
	apiRouter.DELETE("charts/v3/*selector", chartsV3Delete)
	apiRouter.GET("charts/v3", chartsV3List)
	apiRouter.PATCH("charts/v3/*selector", chartsV3Edit)

	apiRouter.GET("chart-versions/v3/*selector", chartVersionsV3Get)
	apiRouter.GET("chart-versions/v3", chartVersionsV3List)
	apiRouter.PATCH("chart-versions/v3/*selector", chartVersionsV3Edit)
	apiRouter.PUT("chart-versions/v3", chartVersionsV3Upsert)
	apiRouter.GET("chart-versions/procedures/v3/changelog", chartVersionsProceduresV3Changelog)

	apiRouter.GET("clusters/v3/*selector", clustersV3Get)
	apiRouter.POST("clusters/v3", clustersV3Create)
	apiRouter.DELETE("clusters/v3/*selector", clustersV3Delete)
	apiRouter.GET("clusters/v3", clustersV3List)
	apiRouter.PATCH("clusters/v3/*selector", clustersV3Edit)

	apiRouter.GET("ci-identifiers/v3", ciIdentifiersV3List)
	apiRouter.GET("ci-identifiers/v3/*selector", ciIdentifiersV3Get)

	apiRouter.GET("ci-runs/v3", ciRunsV3List)
	apiRouter.GET("ci-runs/v3/*selector", ciRunsV3Get)
	apiRouter.PUT("ci-runs/v3", ciRunsV3Upsert)
	apiRouter.GET("ci-runs/procedures/v3/github-info", ciRunsProceduresV3GithubInfoList)

	apiRouter.GET("users/v3", usersV3List)
	apiRouter.GET("users/v3/*selector", usersV3Get)
	apiRouter.PUT("users/v3", usersV3Upsert)

	apiRouter.GET("deploy-hooks/slack/v3", slackDeployHooksV3List)
	apiRouter.GET("deploy-hooks/slack/v3/*selector", slackDeployHooksV3Get)
	apiRouter.POST("deploy-hooks/slack/v3", slackDeployHooksV3Create)
	apiRouter.PATCH("deploy-hooks/slack/v3/*selector", slackDeployHooksV3Edit)
	apiRouter.DELETE("deploy-hooks/slack/v3/*selector", slackDeployHooksV3Delete)
	apiRouter.POST("deploy-hooks/slack/procedures/v3/test/*selector", slackDeployHooksV3TestRun)
	apiRouter.GET("deploy-hooks/github-actions/v3", githubActionsDeployHooksV3List)
	apiRouter.GET("deploy-hooks/github-actions/v3/*selector", githubActionsDeployHooksV3Get)
	apiRouter.POST("deploy-hooks/github-actions/v3", githubActionsDeployHooksV3Create)
	apiRouter.PATCH("deploy-hooks/github-actions/v3/*selector", githubActionsDeployHooksV3Edit)
	apiRouter.DELETE("deploy-hooks/github-actions/v3/*selector", githubActionsDeployHooksV3Delete)
	apiRouter.POST("deploy-hooks/github-actions/procedures/v3/test/*selector", githubActionsDeployHooksV3TestRun)

	apiRouter.PUT("git-commits/v3", gitCommitsV3Upsert)

	apiRouter.GET("incidents/v3/*selector", incidentsV3Get)
	apiRouter.POST("incidents/v3", incidentsV3Create)
	apiRouter.DELETE("incidents/v3/*selector", incidentsV3Delete)
	apiRouter.GET("incidents/v3", incidentsV3List)
	apiRouter.PATCH("incidents/v3/*selector", incidentsV3Edit)

	apiRouter.GET("pagerduty-integrations/v3/*selector", pagerdutyIntegrationsV3Get)
	apiRouter.GET("pagerduty-integrations/v3", pagerdutyIntegrationsV3List)
	apiRouter.PATCH("pagerduty-integrations/v3/*selector", pagerdutyIntegrationsV3Edit)
	apiRouter.DELETE("pagerduty-integrations/v3/*selector", pagerdutyIntegrationsV3Delete)
	apiRouter.POST("pagerduty-integrations/v3", pagerdutyIntegrationsV3Create)

	apiRouter.GET("environments/v3/*selector", environmentsV3Get)
	apiRouter.POST("environments/v3", environmentsV3Create)
	apiRouter.DELETE("environments/v3/*selector", environmentsV3Delete)
	apiRouter.GET("environments/v3", environmentsV3List)
	apiRouter.PATCH("environments/v3/*selector", environmentsV3Edit)

	apiRouter.GET("chart-releases/v3/*selector", chartReleasesV3Get)
	apiRouter.POST("chart-releases/v3", chartReleasesV3Create)
	apiRouter.DELETE("chart-releases/v3/*selector", chartReleasesV3Delete)
	apiRouter.GET("chart-releases/v3", chartReleasesV3List)
	apiRouter.PATCH("chart-releases/v3/*selector", chartReleasesV3Edit)

	apiRouter.GET("database-instances/v3/*selector", databaseInstancesV3Get)
	apiRouter.POST("database-instances/v3", databaseInstancesV3Create)
	apiRouter.DELETE("database-instances/v3/*selector", databaseInstancesV3Delete)
	apiRouter.GET("database-instances/v3", databaseInstancesV3List)
	apiRouter.PATCH("database-instances/v3/*selector", databaseInstancesV3Edit)
}
