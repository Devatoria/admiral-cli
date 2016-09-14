package image

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Devatoria/admiral-cli/client"
	"github.com/Devatoria/admiral-cli/table"

	"github.com/Devatoria/admiral/models"
	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "List images available in my namespace",
	Run: func(cmd *cobra.Command, args []string) {
		data, _, err := client.Request("GET", "/images", nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		var images []models.Image
		err = json.Unmarshal(data, &images)
		if err != nil {
			panic(err)
		}

		t := table.NewTable([]string{"Name", "Created at"})
		for _, image := range images {
			t.Append([]string{
				image.Name,
				image.CreatedAt.String(),
			})
		}

		t.Render()
	},
}