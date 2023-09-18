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
func ConfigureRoutes(apiRouter *gin.RouterGroup) {
	apiRouter.GET("charts/v3/*selector", chartsV3Get)
	apiRouter.POST("charts/v3", chartsV3Create)

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
	apiRouter.GET("deploy-hooks/github-actions/v3", githubActionsDeployHooksV3List)
	apiRouter.GET("deploy-hooks/github-actions/v3/*selector", githubActionsDeployHooksV3Get)
	apiRouter.POST("deploy-hooks/github-actions/v3", githubActionsDeployHooksV3Create)
	apiRouter.PATCH("deploy-hooks/github-actions/v3/*selector", githubActionsDeployHooksV3Edit)
	apiRouter.DELETE("deploy-hooks/github-actions/v3/*selector", githubActionsDeployHooksV3Delete)
}
