package image

import (
	"fmt"
	"net/http"

	"github.com/Devatoria/admiral-cli/client"
	"github.com/Devatoria/admiral-cli/errors"

	"github.com/spf13/cobra"
)

var setPublic = &cobra.Command{
	Use:   "set-public <image>",
	Short: "Set the given image as public (pull only)",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			errors.WrongUsage()
		}

		_, code, err := client.Request("PATCH", fmt.Sprintf("/image/public/%s", args[0]), nil)
		if err != nil {
			errors.Error(err)
		}

		switch code {
		case http.StatusNotFound:
			errors.NotFound()
		}
	},
}

var setPrivate = &cobra.Command{
	Use:   "set-private <image>",
	Short: "Set the given image as private",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			errors.WrongUsage()
		}

		_, code, err := client.Request("PATCH", fmt.Sprintf("/image/private/%s", args[0]), nil)
		if err != nil {
			errors.Error(err)
		}

		switch code {
		case http.StatusNotFound:
			errors.NotFound()
		}
	},
}
