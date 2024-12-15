package cmd

import (
	"fmt"
	"io"
	"net/http"

	"jsp/config"
)

func fetchStockPrice(stockCode string, idToken string) (string, error) {
	url := fmt.Sprintf("%s/prices/daily_quotes?code=%s", config.BASE_URL, stockCode)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return "", err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", idToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return "", err
	}
	return string(body), nil
}
