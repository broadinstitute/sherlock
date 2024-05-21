package pact

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/version"
	"github.com/broadinstitute/sherlock/sherlock/internal/boot"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	pactmodels "github.com/pact-foundation/pact-go/v2/models"
	"github.com/pact-foundation/pact-go/v2/provider"
	pactversion "github.com/pact-foundation/pact-go/v2/version"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"reflect"
	"testing"
)

// TestProvider validates all Sherlock provider pacts.
//
// Sherlock's entire HTTP surface is available for consumers to test against.
//
// Provider states can be composed and are documented in ./README.md. They include control over caller permissions
// courtesy of test_users.TestUserHelper and creation of sample data courtesy of models.TestData.
//
// In theory, here's the kind of interaction that Sherlock is enabling  (forgive the
// formatting, we're awkwardly embedding indented Go code into comments):
//
// ```go
//
//	pact.
//	AddInteraction().
//	Given("ChartRelease_LeonardoProd exists").
//	Given("caller is non-suitable").
//	UponReceiving("a delete request for leonardo-prod").
//	WithRequest("DELETE", "/api/chart-releases/v3/leonardo-prod").
//	WillRespondWith(403).
//	WithJSONBody(...)
//
// ```
//
// Sherlock doesn't currently have any consumers, making this test a bit of a dogfooding exercise.
func TestProvider(t *testing.T) {
	// Pact broker requires authentication. Sherlock's tests don't otherwise require authentication, so we skip this.
	if (os.Getenv("PACT_BROKER_USERNAME") == "" || os.Getenv("PACT_BROKER_PASSWORD") == "") && os.Getenv("PACT_BROKER_TOKEN") == "" {
		t.Skip("No authenticated connection to Pact, skipping Pact contract testing")
	}

	// Quiet down logging and load config
	gin.SetMode(gin.TestMode)
	config.LoadTestConfig()

	// If you don't disable tracking, the library will log a lot of messages about it
	oldPactDoNotTrackSetting := os.Getenv("PACT_DO_NOT_TRACK")
	assert.NoError(t, os.Setenv("PACT_DO_NOT_TRACK", "true"))

	// This does way more than just check the version, it'll actually try to repair installations, but it does
	// some helpful compatibility checks for the FFI library.
	pactversion.CheckVersion()

	// Decide whether to publish results back to the broker, "if we have a real version."
	// We're being smarter than the linter here -- we change version.BuildVersion at link-time, so this won't
	// short-circuit to false.
	//goland:noinspection GoBoolExpressions
	publishPacts := version.BuildVersion != version.DevelopmentVersionString

	// No usage of testify's suite.Suite here, so we initialize helpers inline.
	ctx := context.Background()
	testUserHelper := test_users.TestUserHelper{}
	modelTestSuiteHelper := models.TestSuiteHelper{}
	modelTestSuiteHelper.SetupSuite()

	// Prepare the state handlers available for interactions
	var explicitlySetSuitableUser, explicitlySetNonSuitableUser bool
	stateHandlers := make(map[string]pactmodels.StateHandler)

	// 1. We expose test_users.TestUserHelper's header helpers here. We just set some local state and actually call
	//    out to test_users.TestUserHelper later when handling each request.
	stateHandlers["caller is suitable"] =
		func(setup bool, _ pactmodels.ProviderState) (pactmodels.ProviderStateResponse, error) {
			if setup {
				if explicitlySetNonSuitableUser {
					return nil, fmt.Errorf("both 'caller is suitable' and 'caller is non-suitable' set")
				}
				explicitlySetSuitableUser = true
			}
			return nil, nil
		}
	stateHandlers["caller is non-suitable"] =
		func(setup bool, _ pactmodels.ProviderState) (pactmodels.ProviderStateResponse, error) {
			if setup {
				if explicitlySetSuitableUser {
					return nil, fmt.Errorf("both 'caller is suitable' and 'caller is non-suitable' set")
				}
				explicitlySetNonSuitableUser = true
			}
			return nil, nil
		}

	// 2. We expose models.TestSuiteHelper's TestData helpers here. We blank-fire a test, reflect on the TestData,
	//    and iterate over its methods, making a state handler for each one. Ex: "Chart_Leonardo exists"
	modelTestSuiteHelper.SetupTest()
	testDataTemporaryValue := reflect.ValueOf(modelTestSuiteHelper.TestData)
	for i := 0; i < testDataTemporaryValue.NumMethod(); i++ {
		methodTypeValue := testDataTemporaryValue.Type().Method(i)
		// Method receivers count as an argument, so we filter for methods with 1, not 0, arguments
		if methodTypeValue.Type.NumIn() == 1 {
			methodName := methodTypeValue.Name
			stateHandlers[fmt.Sprintf("%s exists", methodName)] =
				func(setup bool, _ pactmodels.ProviderState) (pactmodels.ProviderStateResponse, error) {
					if setup {
						// When running the state handler, reflect on whatever TestData exists then and call the
						// method we want on it
						reflect.ValueOf(modelTestSuiteHelper.TestData).MethodByName(methodName).Call([]reflect.Value{})
					}
					return nil, nil
				}
		}
	}
	// We have to tear down the test we blank-fired -- this was just the cleanest way to safely list the methods of
	// a TestData
	modelTestSuiteHelper.TearDownTest()

	// Log without a level so that it shows up amidst test output (info etc. gets filtered)
	stateHandlerNames := make([]string, 0, len(stateHandlers))
	for name := range stateHandlers {
		stateHandlerNames = append(stateHandlerNames, name)
	}
	log.Log().Strs("state-handlers", stateHandlerNames).Int("state-handler-count", len(stateHandlers)).Msg("pact provider state handlers computed")

	var sherlockRouter *gin.Engine
	err := provider.NewVerifier().VerifyProvider(t, provider.VerifyRequest{
		BrokerURL: config.Config.String("pactbroker.url"),
		Provider:  "sherlock",
		// ProviderBaseURL is set to an empty string intentionally, because we aren't actually going to run Sherlock's
		// server on localhost. We'll handle the requests directly in the RequestFilter middleware.
		ProviderBaseURL:            "",
		ProviderVersion:            version.BuildVersion,
		PublishVerificationResults: publishPacts,
		// FailIfNoPactsFound is false right now because there aren't any Pact consumers of Sherlock
		FailIfNoPactsFound: false,
		// BeforeEach interaction, run test-level each-time setup
		BeforeEach: func() error {
			explicitlySetSuitableUser = false
			explicitlySetNonSuitableUser = false
			// Heavy-lifting here: modelTestSuiteHelper.DB is now an active-transaction database reference we
			// can build a disposable router from. Sherlock's own API test suite does something very similar
			// to this, the only real difference is that we build a much more complete router here to make
			// Sherlock's entire HTTP surface available (down to redirects, Swagger page, everything).
			modelTestSuiteHelper.SetupTest()
			sherlockRouter = boot.BuildRouter(ctx, modelTestSuiteHelper.DB)
			return nil
		},
		// StateHandlers are executed between BeforeEach and RequestFilter
		StateHandlers: stateHandlers,
		// RequestFilter directly handles requests, ignoring the http.Handler argument that would normally lob the
		// request at the ProviderBaseURL. The Gin server library also directly implements http.Handler, so we just
		// hand the request to the router directly.
		RequestFilter: func(_ http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Sherlock's test-time user middleware will default to treating requests as suitable, but it exposes
				// explicit controls for it to make tests clear and stable so we respect that here.
				if explicitlySetSuitableUser {
					testUserHelper.UseSuitableUserFor(r)
				} else if explicitlySetNonSuitableUser {
					testUserHelper.UseNonSuitableUserFor(r)
				}

				// Heavy-lifting here: you're sorta not supposed to do this, but we short-circuit Pact's middleware
				// chain right near the end here by handling the request in-memory directly with Sherlock's router.
				// We don't have to worry about actually running a server and this follows the pattern of how the
				// rest of Sherlock's tests are run.
				sherlockRouter.ServeHTTP(w, r)
			})
		},
		// AfterEach interaction, run test-level each-time teardown
		AfterEach: func() error {
			sherlockRouter = nil
			modelTestSuiteHelper.TearDownTest()
			return nil
		},
	})
	assert.NoError(t, err)

	// We run any suite-level one-time teardown functions inline, since there's no actual suite to do so for us
	modelTestSuiteHelper.TearDownSuite()

	// Set PACT_DO_NOT_TRACK back to whatever it was before to minimize the test's side effects
	assert.NoError(t, os.Setenv("PACT_DO_NOT_TRACK", oldPactDoNotTrackSetting))
}
