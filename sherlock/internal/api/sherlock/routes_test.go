package sherlock

import (
	"embed"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/api/gin_mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"regexp"
	"strings"
	"testing"
)

//go:embed *.go
var filesToLookForSwaggerCommentsIn embed.FS

var regexToParseSwaggerRouteDeclarations = regexp.MustCompile(`//\s+@router\s+(.+)`)
var regexToHelpConvertGinRouteConfigurations = regexp.MustCompile(`[*:]([^/]+)`)

func TestConfigureRoutes_swaggerComments(t *testing.T) {
	// Map of route declarations and if they were found
	expectedRouteDeclarations := make(map[string]bool)

	// Make a mock router to help us assemble the above map
	mockRouter := gin_mocks.NewMockIRoutes(t)
	for _, method := range []string{"GET", "POST", "DELETE", "PATCH", "PUT", "OPTIONS", "HEAD"} {
		redeclaredMethodToAvoidPointerShenanigans := method
		mockRouter.
			On(redeclaredMethodToAvoidPointerShenanigans, mock.AnythingOfType("string"), mock.Anything).
			Run(func(args mock.Arguments) {
				routeDeclaration := fmt.Sprintf("%s [%s]",
					transformPathToSwaggerFormat(args.Get(0).(string)),
					strings.ToLower(redeclaredMethodToAvoidPointerShenanigans))
				if _, alreadyExists := expectedRouteDeclarations[routeDeclaration]; alreadyExists {
					assert.Fail(t, fmt.Sprintf("routes.go's ConfigureRoutes seems to configure `%s` twice", routeDeclaration))
				} else {
					expectedRouteDeclarations[routeDeclaration] = false
				}
			}).
			Return(nil).
			Maybe()
	}

	// Run our mock through the configuration function
	assert.NotPanics(t, func() {
		ConfigureRoutes(mockRouter)
	})

	// Iterate through the source files, looking for declarations
	sourceFiles, err := filesToLookForSwaggerCommentsIn.ReadDir(".")
	assert.NoError(t, err)
	for _, sourceFile := range sourceFiles {
		assert.Falsef(t, sourceFile.IsDir(), "%s is a directory", sourceFile.Name())
		sourceFileContents, err := filesToLookForSwaggerCommentsIn.ReadFile(sourceFile.Name())
		assert.NoError(t, err)
		for _, routeDeclarationMatch := range regexToParseSwaggerRouteDeclarations.FindAllSubmatch(sourceFileContents, -1) {
			if assert.Lenf(t, routeDeclarationMatch, 2, "route declaration regex matches should only have two items in the array, the whole line and the declaration itself") {
				routeDeclaration := string(routeDeclarationMatch[1])
				alreadyDeclared, shouldBeDeclared := expectedRouteDeclarations[routeDeclaration]
				assert.Truef(t, shouldBeDeclared, "%s was declared in %s's Swagger comments but isn't configured in routes.go's ConfigureRoutes", routeDeclaration, sourceFile.Name())
				assert.Falsef(t, alreadyDeclared, "%s was declared twice in Swagger comments (duplicate detected in %s)", routeDeclaration, sourceFile.Name())
				expectedRouteDeclarations[routeDeclaration] = true
			}
		}
	}

	// If there's any declarations that are still false, we didn't find them and should say so
	for expectedRouteDeclaration, wasDeclared := range expectedRouteDeclarations {
		assert.Truef(t, wasDeclared, "%s was configured in routes.go's ConfigureRoutes but wasn't declared in any Swagger comment", expectedRouteDeclaration)
	}
}

func transformPathToSwaggerFormat(path string) string {
	return fmt.Sprintf("/api/%s", regexToHelpConvertGinRouteConfigurations.ReplaceAllString(path, "{$1}"))
}

func Test_transformPathToSwaggerFormat(t *testing.T) {
	assert.Equal(t, "/api/foo/bar", transformPathToSwaggerFormat("foo/bar"))
	assert.Equal(t, "/api/foo/bar/{baz}", transformPathToSwaggerFormat("foo/bar/*baz"))
	assert.Equal(t, "/api/foo/bar/{baz}", transformPathToSwaggerFormat("foo/bar/:baz"))
	assert.Equal(t, "/api/foo/{bar}/{baz}", transformPathToSwaggerFormat("foo/:bar/*baz"))
}
