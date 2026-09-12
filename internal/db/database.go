package db

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func ConnectionToYDB() (*firestore.Client, error) {
	//Connect to db
	ctx := context.Background()

	projectID := "arizona-wiki"
	// Path to key for local test
	// pathToKey := "C:\\Users\\aleks\\Desktop\\GO_Projects\\Arizona_wiki_server\\internal\\db\\firebase_key.json"
	//opt := option.WithAuthCredentialsFile(option.ServiceAccount, pathToKey)

	keyBase64 := os.Getenv("DATABASE_KEY")
	keyJSON, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opt := option.WithCredentialsJSON([]byte(keyJSON))

	client, err := firestore.NewClient(ctx, projectID, opt)

	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
	}
	log.Println("Connected to YDB")
	return client, nil
}
