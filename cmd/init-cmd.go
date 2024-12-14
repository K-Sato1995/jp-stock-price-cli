package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the cli",
	Run: func(cmd *cobra.Command, args []string) {
		email := viper.GetString("email")
		pass := viper.GetString("password")

		if len(email) == 0 || len(pass) == 0 {
			fmt.Println("Please provide email and password")
			return
		}
		refToken, err := fetchRefreshToken(
			email,
			pass,
		)
		if err != nil {
			fmt.Println(err)
			return
		}
		idToken, err := fetchIdToken(refToken)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(idToken)
	},
}
