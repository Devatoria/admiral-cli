package image

import (
	"fmt"

	"github.com/Devatoria/admiral-cli/client"
	"github.com/Devatoria/admiral-cli/errors"

	"github.com/spf13/cobra"
)

var delete = &cobra.Command{
	Use:   "delete <image>",
	Short: "Delete the given image from the registry",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			errors.WrongUsage()
		}

		_, code, err := client.Request("DELETE", fmt.Sprintf("/image/%s", args[0]), nil)
		if err != nil {
			errors.Error(err)
		}

		switch code {
		case 200:
			fmt.Printf("Image %s has been deleted\n", args[0])
		case 401:
			errors.Unauthorized()
		case 404:
			errors.NotFound()
		}
	},
}
