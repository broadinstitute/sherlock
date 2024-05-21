package slack

import (
	"context"
	"github.com/slack-go/slack"
)

func GetUser(ctx context.Context, email string) (slackID string, username string, name string, err error) {
	if isEnabled() && email != "" {
		var user *slack.User
		if user, err = client.GetUserByEmailContext(ctx, email); err == nil && user != nil {
			slackID = user.ID
			username = user.Name
			name = user.RealName
		}
	}
	return
}
