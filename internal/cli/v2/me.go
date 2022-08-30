package v2

import (
	"fmt"

	"github.com/broadinstitute/sherlock/clients/go/client/misc"
	"github.com/spf13/cobra"
)

var (
	meCmd = &cobra.Command{
		Use:   "me",
		Short: "me returns info about the calling user",
		Long:  "me returns information about the calling user which sherlock has access to",

		RunE: getMeInfo,
	}
)

// getMeInfo executes a request to sherlock's /my-user endpoint and returns
// info about the calling user including suitability
func getMeInfo(cmd *cobra.Command, args []string) error {
	meParams := misc.NewGetMyUserParams()
	me, err := app.client.Misc.GetMyUser(meParams)
	if err != nil {
		return fmt.Errorf("error retrieving current user info: %v", err)
	}

	response := formatMeResponse(me.Payload.Email, me.Payload.Suitability)
	fmt.Fprint(cmd.OutOrStdout(), response)
	return nil
}

func formatMeResponse(user, suitability string) string {
	return fmt.Sprintf("User: %s\nSuitability: %s\n", user, suitability)
}
