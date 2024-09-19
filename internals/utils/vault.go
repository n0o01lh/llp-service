package utils

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
	vault "github.com/hashicorp/vault/api"
)

func VaultConnection() *vault.Client {

	config := vault.DefaultConfig()

	config.Address = os.Getenv("VAULT_ADDR")

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	// Authenticate
	client.SetToken(os.Getenv("VAULT_TOKEN"))

	return client
}
