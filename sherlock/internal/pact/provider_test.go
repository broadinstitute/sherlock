package pact

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/version"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/boot"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
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

// TestProvider validates all Sherlock provider pacts. Do any of those exist? Not right now.
//
// Between Beehive and Sherlock, frankly a better check would be if a new Sherlock's client
// library caused any new type errors in Beehive... but we'll cross that bridge if we need
// to.
//
// This test exists mostly as dogfooding for Pact. The goal is meaningful dogfooding, too:
// we're not mocking just a few endpoints. We're wiring up Sherlock's entire API surface,
// backed by an ephemeral in-memory database identical to how Sherlock's own functional
// tests are run, plus full access to all models.TestData.
//
// In theory, here's the kind of interaction that Sherlock is enabling  (forgive the
// formatting, we're awkwardly embedding indented Go code into comments):
//
// ```golang
//
//	pact.
//	AddInteraction().
//	Given("ChartRelease_LeonardoProd exists").
//	Given("caller is non-suitable").
//	UponReceiving("a delete request for leonardo-prod").
//	WithRequest("DELETE", S("/api/chart-releases/v3/leonardo-prod")).
//	WillRespondWith(403).
//	WithJSONBody(Like(...))
//
// ```
func TestProvider(t *testing.T) {
	// "Suite"-level one-time setup
	gin.SetMode(gin.TestMode)
	config.LoadTestConfig()
	oldPactDoNotTrackSetting := os.Getenv("PACT_DO_NOT_TRACK")
	assert.NoError(t, os.Setenv("PACT_DO_NOT_TRACK", "true")) // I'd care less about their tracking if it didn't log so much
	pactversion.CheckVersion()
	// We're smarter than the linter here. GoLand will complain that this is always false because
	// version.BuildVersion == version.DevelopmentVersionString, but we actually replace the former from the linker
	// at compile time.
	//goland:noinspection GoBoolExpressions
	publishPacts := version.BuildVersion != version.DevelopmentVersionString &&
		((os.Getenv("PACT_BROKER_USERNAME") != "" && os.Getenv("PACT_BROKER_PASSWORD") != "") ||
			(os.Getenv("PACT_BROKER_TOKEN") != ""))
	ctx := context.Background()
	testUserHelper := test_users.TestUserHelper{}
	modelTestSuiteHelper := models.TestSuiteHelper{}
	modelTestSuiteHelper.SetupSuite()

	// Prepare the state handlers available for interactions
	var explicitlySetSuitableUser, explicitlySetNonSuitableUser bool
	stateHandlers := make(map[string]pactmodels.StateHandler)
	// 1. We expose test_users.TestUserHelper's header helpers here. We just set some local state and actually call
	//    out to test_users.TestUserHelper later when handling each request.
	stateHandlers["caller is suitable"] = func(setup bool, _ pactmodels.ProviderState) (pactmodels.ProviderStateResponse, error) {
		if setup {
			if explicitlySetNonSuitableUser {
				return nil, fmt.Errorf("both 'caller is suitable' and 'caller is non-suitable' set")
			}
			explicitlySetSuitableUser = true
		}
		return nil, nil
	}
	stateHandlers["caller is non-suitable"] = func(setup bool, _ pactmodels.ProviderState) (pactmodels.ProviderStateResponse, error) {
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
			stateHandlers[fmt.Sprintf("%s exists", methodName)] = func(setup bool, _ pactmodels.ProviderState) (pactmodels.ProviderStateResponse, error) {
				if setup {
					// When running the state handler, reflect on whatever TestData exists then and call the
					// method we want on it
					reflect.ValueOf(modelTestSuiteHelper.TestData).MethodByName(methodName).Call([]reflect.Value{})
				}
				return nil, nil
			}
		}
	}
	modelTestSuiteHelper.TearDownTest()
	stateHandlerNames := make([]string, 0, len(stateHandlers))
	for name := range stateHandlers {
		stateHandlerNames = append(stateHandlerNames, name)
	}
	// Log without a level so that it shows up amidst test output (info etc. gets filtered)
	log.Log().Strs("state-handlers", stateHandlerNames).Int("state-handler-count", len(stateHandlers)).Msg("pact provider state handlers computed")

	var sherlockRouter *gin.Engine
	verifier := provider.NewVerifier()
	err := verifier.VerifyProvider(t, provider.VerifyRequest{
		BrokerURL: config.Config.String("pactbroker.url"),
		Provider:  "sherlock",
		// ProviderBaseURL is set to an empty string intentionally, because we aren't actually going to run Sherlock's
		// server on localhost. We'll handle the requests directly in the RequestFilter middleware.
		ProviderBaseURL:            "",
		ProviderVersion:            version.BuildVersion,
		PublishVerificationResults: publishPacts,
		// FailIfNoPactsFound is false right now because there aren't any Pact consumers of Sherlock
		FailIfNoPactsFound: false,
		// BeforeEach interaction, run "test"-level each-time setup
		BeforeEach: func() error {
			explicitlySetSuitableUser = false
			explicitlySetNonSuitableUser = false
			// Heavy-lifting here: modelTestSuiteHelper.DB is now an active-transaction database reference we
			// can build a disposable router from. Sherlock's own API test suite does something very similar
			// to this, the only real difference is that we build a much more complete router here to make
			// Sherlock's entire HTTP surface available.
			modelTestSuiteHelper.SetupTest()
			sherlockRouter = boot.BuildRouter(ctx, modelTestSuiteHelper.DB)
			return nil
		},
		// StateHandlers are executed between BeforeEach and RequestFilter
		StateHandlers: stateHandlers,
		// RequestFilter directly handles requests, ignoring the http.Handler argument that would normally lob the
		// request at the ProviderBaseURL. Gin also directly implements http.Handler, so we just call it instead.
		RequestFilter: func(_ http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if explicitlySetSuitableUser {
					testUserHelper.UseSuitableUserFor(r)
				} else if explicitlySetNonSuitableUser {
					testUserHelper.UseNonSuitableUserFor(r)
				}
				sherlockRouter.ServeHTTP(w, r)
			})
		},
		// AfterEach interaction, run "test"-level each-time teardown
		AfterEach: func() error {
			modelTestSuiteHelper.TearDownTest()
			return nil
		},
	})
	assert.NoError(t, err)

	// "Suite"-level one-time teardown
	modelTestSuiteHelper.TearDownSuite()
	assert.NoError(t, os.Setenv("PACT_DO_NOT_TRACK", oldPactDoNotTrackSetting))
}
