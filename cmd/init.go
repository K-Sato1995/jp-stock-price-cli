package cmd

import (
	"fmt"
	"os"
)

const (
	JQUANTS_EMAIL_ENV    = "JQUANTS_EMAIL"
	JQUANTS_PASSWORD_ENV = "JQUANTS_PASSWORD"
)

func initFunc() (string, error) {
	email := os.Getenv(JQUANTS_EMAIL_ENV)
	if email == "" {
		return "", fmt.Errorf("JQUANTS_EMAIL environment variable is not set")
	}
	pass := os.Getenv(JQUANTS_PASSWORD_ENV)
	if email == "" {
		return "", fmt.Errorf("JQUANTS_PASSWORD environment variable is not set")
	}
	refToken, err := fetchRefreshToken(
		email,
		pass,
	)
	if err != nil {
		return "", err
	}
	idToken, err := fetchIdToken(refToken)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return idToken, nil
}
