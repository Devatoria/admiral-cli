package team

import (
	"encoding/json"
	"net/http"

	"github.com/Devatoria/admiral-cli/client"
	"github.com/Devatoria/admiral-cli/errors"
	"github.com/Devatoria/admiral-cli/table"

	"github.com/Devatoria/admiral/models"
	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "List user teams",
	Run: func(cmd *cobra.Command, args []string) {
		data, code, err := client.Request("GET", "/teams", nil)
		if err != nil {
			errors.Error(err)
		}

		switch code {
		case http.StatusOK:
			var teams []models.Team
			err = json.Unmarshal(data, &teams)
			if err != nil {
				panic(err)
			}

			t := table.NewTable([]string{"Name", "Created at"})
			for _, team := range teams {
				t.Append([]string{
					team.Name,
					team.CreatedAt.String(),
				})
			}

			t.Render()
		case http.StatusUnauthorized:
			errors.Unauthorized()
		}
	},
}
