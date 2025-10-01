package main 

import (
	"context"
	"os"
	"log"


	"github.com/redis/go-redis-entraid"
	"github.com/redis/go-redis-entraid/identity"
	"github.com/redis/go-redis/v9"
	// "github.com/github/go-redis/crdb"
)


func main() {
	Auth Variables Needed for SPN Connection 
	client_id, ok := os.LookupEnv("AZURE_CLIENT_ID")
		if !ok {
			log.Println("AZURE_CLIENT_ID environment variable is required")
		}

	client_secret, ok := os.LookupEnv("AZURE_CLIENT_SECRET")
		if !ok {
			log.Println("AZURE_CLIENT_SECRET environment variable is required")
		}

	tenant_id, ok := os.LookupEnv("AZURE_TENANT_ID")
		if !ok {
			log.Println("AZURE_TENANT_ID environment variable is required")
		}
	
	provider, err := entraid.NewConfidentialCredentialsProvider(
		entraid.ConfidentialCredentialsProviderOptions{
			ConfidentialIdentityProviderOptions: identity.ConfidentialIdentityProviderOptions{
				ClientID:        client_id,
				ClientSecret:    client_secret,
				CredentialsType: identity.ClientSecretCredentialType,
				Authority: identity.AuthorityConfiguration{
					AuthorityType: identity.AuthorityTypeDefault,
					TenantID:     tenant_id,
				},
			},
		},
	)

	if err != nil {
			log.Fatalf("Failed to create credentials provider: %v", err)
		}

	client := redis.NewClient(&redis.Options{
			Addr: "data-architecture-test-amr.eastus.redis.azure.net:10000",
			StreamingCredentialsProvider: provider,
		})
		defer client.Close()

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatal("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis!")
}