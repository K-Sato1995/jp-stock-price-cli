package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var email string
var password string

var rootCmd = &cobra.Command{
	Use:   "jsp",
	Short: "CLI tool for getting Japanese Stock Price",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of jp-stock-price",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.Flags().StringP("flagname", "code", "0000", "The stock code listed in TYO")
	initCmd.PersistentFlags().StringVarP(&email, "email", "e", "", "Your J-Quants API email")
	initCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Your J-Quants API password")
	viper.BindPFlag("email", initCmd.PersistentFlags().Lookup("email"))
	viper.BindPFlag("password", initCmd.PersistentFlags().Lookup("password"))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
