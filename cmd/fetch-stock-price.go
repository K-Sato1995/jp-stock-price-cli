package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"jsp/config"
)

func FetchStockPrice() {
	symbol := "AMZN"
	apiKey := "YOUR_API_KEY"
	url := fmt.Sprintf(config.BASE_URL, symbol, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching stock price: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	if quote, ok := result["Global Quote"].(map[string]interface{}); ok {
		if price, exists := quote["05. price"]; exists {
			fmt.Printf("Current stock price for %s: $%v\n", symbol, price)
		}
	}
}
