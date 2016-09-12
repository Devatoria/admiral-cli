package user

import (
	"fmt"
	"net/http"

	"github.com/Devatoria/admiral-cli/client"
	"github.com/Devatoria/admiral-cli/errors"

	"github.com/Devatoria/admiral/api"
	"github.com/spf13/cobra"
)

var create = &cobra.Command{
	Use:   "create <username> <password>",
	Short: "Creates a new user",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			errors.WrongUsage()
		}

		user := api.User{
			Username: args[0],
			Password: args[1],
		}
		data, code, err := client.Request("PUT", "/user", user)
		if err != nil {
			errors.Error(err)
		}

		switch code {
		case http.StatusOK:
			fmt.Println("User is created")
		case http.StatusConflict:
			fmt.Println("User already exists")
		case http.StatusBadRequest:
			errors.PrintBadRequestError(data)
		}
	},
}
