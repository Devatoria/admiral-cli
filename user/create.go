package user

import (
	"fmt"

	"github.com/Devatoria/admiral-cli/client"
	"github.com/Devatoria/admiral-cli/errors"

	"github.com/Devatoria/admiral/models"
	"github.com/spf13/cobra"
)

var create = &cobra.Command{
	Use:   "create <username> <password>",
	Short: "Creates a new user if it does not exist",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			errors.WrongUsage()
		}
		user := models.User{
			Username: args[0],
			Password: args[1],
		}

		_, code, err := client.Request("PUT", "/user", user)
		if err != nil {
			errors.Error(err)
		}

		switch code {
		case 200:
			fmt.Println("User has been created!")
		case 409:
			fmt.Println("User already exists.")
		}
	},
}
