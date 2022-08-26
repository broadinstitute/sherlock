package v2

import (
	"fmt"

	"github.com/broadinstitute/sherlock/clients/go/client/misc"
	"github.com/rs/zerolog/log"
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

func getMeInfo(cmd *cobra.Command, args []string) error {
	meParams := misc.NewGetMyUserParams()
	me, err := app.client.Misc.GetMyUser(meParams)
	if err != nil {
		return fmt.Errorf("error retrieving current user info: %v", err)
	}

	showMeResponse(me.Payload.Email, me.Payload.Suitability)
	return nil
}

func showMeResponse(user, suitability string) {
	log.Info().Msgf("retrieving current user info\nUser: %s\nSuitability: %s\n", user, suitability)
}
