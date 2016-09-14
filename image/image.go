package image

import (
	"github.com/spf13/cobra"
)

func init() {
	Command.AddCommand(list)
}

var Command = &cobra.Command{
	Use:   "image",
	Short: "Image command",
}
