package main

import (
	"context"
	"fmt"
	"log"
	"twitchium/auth"
	"twitchium/config"
	"twitchium/util"

	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	client := auth.NewAuthClient(cfg.ClientID, cfg.ClientSecret, cfg.RedirectURI)
	state := uuid.New().String()

	fmt.Println("Please open the following URL in your browser:")
	fmt.Println(client.GetAuthURL(state, util.Scopes))
	fmt.Println("\nAfter authorizing, you will be redirected. Please paste the full redirect URL here:")

	var redirectURL string
	if _, err = fmt.Scanln(&redirectURL); err != nil {
		panic(fmt.Sprintf("Failed to read redirect URL: %v", err))
	}

	code, err := util.ExtractCodeFromURL(redirectURL)
	if err != nil {
		log.Fatalf("Failed to extract code: %v", err)
	}

	tokens, err := client.ExchangeCode(ctx, code)
	if err != nil {
		log.Fatalf("Failed to exchange code: %v", err)
	}

	fmt.Printf("\nRefresh Token: %s\n", tokens.RefreshToken)
	fmt.Printf("Access Token: %s\n", tokens.AccessToken)
	fmt.Printf("Expires In: %d seconds\n", tokens.ExpiresIn)
}
