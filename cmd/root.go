package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringP("address", "a", "localhost", "Admiral address")
	RootCmd.PersistentFlags().IntP("port", "p", 3000, "Admiral port")
	RootCmd.PersistentFlags().StringP("username", "U", "", "Admiral username")
	RootCmd.PersistentFlags().StringP("password", "P", "", "Admiral password")

	viper.BindPFlag("address", RootCmd.PersistentFlags().Lookup("address"))
	viper.BindPFlag("port", RootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("username", RootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", RootCmd.PersistentFlags().Lookup("password"))
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/admiral-cli")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Warning: unable to read config (%s)\n", err)
	}
}

var RootCmd = &cobra.Command{
	Use:   "admiral-cli",
	Short: "Admiral CLI is the Admiral commander",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello you :)")
	},
}
