package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var email string
var password string
var idToken string

var rootCmd = &cobra.Command{
	Use:   "jsp",
	Short: "CLI tool for getting Japanese Stock Price",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		idToken, err := initFunc()
		if len(idToken) == 0 {
			fmt.Println(err)
			return
		}
		stockCode, _ := cmd.Flags().GetString("code")
		if len(stockCode) == 0 {
			fmt.Println("Plese set stock code flag. jsp --code ****")
			return
		}
		price, err := fetchStockPrice(stockCode, idToken)
		if err != nil {
			fmt.Println(err)
			return
		}
		// ```json
		//
		//	     {
		//	  "daily_quotes": [
		//	    {
		//	      "Date": "2023-03-24",
		//	      "Code": "86970",
		//	      "Open": 2047,
		//	      "High": 2069,
		//	      "Low": 2035,
		//	      "Close": 2045,
		//	      "UpperLimit": "0",
		//	      "LowerLimit": "0",
		//	      "Volume": 2202500,
		//	      "TurnoverValue": 4507051850,
		//	      "AdjustmentFactor": 1,
		//	      "AdjustmentOpen": 1023.5,
		//	      "AdjustmentHigh": 1034.5,
		//	      "AdjustmentLow": 1017.5,
		//	      "AdjustmentClose": 1022.5,
		//	      "AdjustmentVolume": 4405000
		//	    }
		//	  ]
		//	}
		//
		// ```
		fmt.Println(price)
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
	rootCmd.Flags().StringP("code", "c", "0000", "The stock code listed in TYO")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
