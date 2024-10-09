package authentication

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_forbidDeactivatedUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.LoadTestConfig()
	tests := []struct {
		name        string
		user        *models.User
		expectError bool
	}{
		{
			name:        "nil",
			user:        nil,
			expectError: true,
		},
		{
			name:        "active",
			user:        &models.User{},
			expectError: false,
		},
		{
			name: "deactivated",
			user: &models.User{
				DeactivatedAt: utils.PointerTo(time.Now()),
			},
			expectError: true,
		},
		{
			name: "deleted",
			user: &models.User{
				Model: gorm.Model{DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: true}},
			},
			expectError: true,
		},
		{
			name: "via active",
			user: &models.User{
				Via: &models.User{},
			},
		},
		{
			name: "via deactivated",
			user: &models.User{
				Via: &models.User{
					DeactivatedAt: utils.PointerTo(time.Now()),
					Via:           &models.User{},
				},
			},
			expectError: true,
		},
		{
			name: "via deleted",
			user: &models.User{
				Via: &models.User{
					Model: gorm.Model{DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: true}},
					Via:   &models.User{},
				},
			},
			expectError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Set(ctxUserFieldName, tt.user)
			forbidDeactivatedUsers()(ctx)
			assert.Equal(t, tt.expectError, ctx.IsAborted())
		})
	}
}
