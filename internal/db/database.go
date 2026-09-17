package db

import (
	"context"
	"encoding/base64"
	"os"
	"strings"

	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

// Connect to db
func ConnectionToYDB() (*firestore.Client, error) {

	ctx := context.Background()

	projectID := "arizona-wiki"
	keyBase64 := os.Getenv("DATABASE_KEY")
	cleanedKey := strings.NewReplacer(" ", "", "\n", "", "\r", "").Replace(keyBase64)

	keyJSON, err := base64.StdEncoding.DecodeString(cleanedKey)
	log.Printf("Base64 decode error: %v", err)

	if err != nil {
		log.Printf("Base64 decode error: %v", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opt := option.WithCredentialsJSON(keyJSON)

	client, err := firestore.NewClient(ctx, projectID, opt)

	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
	}
	log.Println("Connected to YDB")
	return client, nil
}
