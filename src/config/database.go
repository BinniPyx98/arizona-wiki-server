package config

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func ConnectionToYDB() *firestore.Client {
	//Connect to db
	ctx := context.Background()

	projectID := "arizona-wiki"
	pathToKey := "C:\\Users\\aleks\\Desktop\\GO_Projects\\Arizona_wiki_server\\src\\config\\firebase_key.json"

	opt := option.WithAuthCredentialsFile(option.ServiceAccount, pathToKey)

	client, err := firestore.NewClient(ctx, projectID, opt)

	if err != nil {
		log.Fatal("error creating client:", err)
	}
	//defer client.Close()
	log.Println("Connected to YDB")
	return client
}
