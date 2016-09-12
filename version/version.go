package version

import (
	"fmt"
	"os"

	"github.com/Devatoria/admiral-cli/client"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Admiral",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := client.Request("GET", "/version")
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		fmt.Println(string(data))
	},
}
