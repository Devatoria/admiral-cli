package image

import (
	"github.com/spf13/cobra"
)

func init() {
	Command.AddCommand(delete)
	Command.AddCommand(list)
	Command.AddCommand(setPublic)
	Command.AddCommand(setPrivate)
}

var Command = &cobra.Command{
	Use:   "image",
	Short: "Image command",
}
