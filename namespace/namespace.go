package namespace

import (
	"github.com/spf13/cobra"
)

func init() {
	Command.AddCommand(list)
}

var Command = &cobra.Command{
	Use:   "namespace",
	Short: "Namespace command",
}
