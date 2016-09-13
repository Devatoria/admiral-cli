package namespace

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
	Short: "List user namespaces",
	Run: func(cmd *cobra.Command, args []string) {
		data, code, err := client.Request("GET", "/namespaces", nil)
		if err != nil {
			errors.Error(err)
		}

		switch code {
		case http.StatusOK:
			var namespaces []models.Namespace
			err = json.Unmarshal(data, &namespaces)
			if err != nil {
				panic(err)
			}

			t := table.NewTable([]string{"Name", "Created at"})
			for _, namespace := range namespaces {
				t.Append([]string{
					namespace.Name,
					namespace.CreatedAt.String(),
				})
			}

			t.Render()
		case http.StatusUnauthorized:
			errors.Unauthorized()
		}
	},
}
