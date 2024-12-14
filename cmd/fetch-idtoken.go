package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jsp/config"
	"net/http"
)

func fetchIdToken(refreshToken string) (string, error) {
	url := fmt.Sprintf("%s/token/auth_refresh?refreshtoken=%s", config.BASE_URL, refreshToken)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return "", fmt.Errorf("failed to fetch refresh token: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	var result struct {
		IdToken string `json:"idToken"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to parse response JSON: %v", err)
	}

	return result.IdToken, nil
}
