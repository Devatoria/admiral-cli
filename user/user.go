package user

import (
	"github.com/spf13/cobra"
)

func init() {
	Command.AddCommand(create)
}

var Command = &cobra.Command{
	Use:   "user",
	Short: "User command",
}
