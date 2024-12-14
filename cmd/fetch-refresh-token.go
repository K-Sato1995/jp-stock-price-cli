package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jsp/config"
	"net/http"
)

func fetchRefreshToken(emailAddress string, password string) (string, error) {
	requestBody := struct {
		MailAddress string `json:"mailaddress"`
		Password    string `json:"password"`
	}{
		MailAddress: emailAddress,
		Password:    password,
	}
	// Convert request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request body: %v", err)
	}
	url := fmt.Sprintf("%s/token/auth_user", config.BASE_URL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}
	var response struct {
		RefreshToken string `json:"refreshToken"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("error parsing response: %v", err)
	}

	return response.RefreshToken, nil
}
