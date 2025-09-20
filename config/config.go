package config

import (
	"fmt"
)

type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

func LoadConfig() (*Config, error) {
	config := &Config{}

	fmt.Print("Enter Client ID: ")
	_, err := fmt.Scanln(&config.ClientID)
	if err != nil {
		return nil, fmt.Errorf("failed to read client ID: %w", err)
	}

	fmt.Print("Enter Client Secret: ")
	_, err = fmt.Scanln(&config.ClientSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to read client secret: %w", err)
	}

	fmt.Print("Enter Redirect URI: ")
	_, err = fmt.Scanln(&config.RedirectURI)
	if err != nil {
		return nil, fmt.Errorf("failed to read redirect URI: %w", err)
	}

	if config.ClientID == "" || config.ClientSecret == "" || config.RedirectURI == "" {
		return nil, fmt.Errorf("all credentials are required")
	}

	return config, nil
}
