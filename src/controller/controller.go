package controller

import (
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AddGameItem(ctx *gin.Context, dbClient *firestore.Client) {
	gameItem := map[string]any{
		"first":  "Ada",
		"last":   "Lovelace",
		"born":   1815,
		"active": true,
	}
	ref, result, err := dbClient.Collection("items").Add(ctx, gameItem)
	if err != nil {
		log.Fatal("Error while adding gameItem to firestore:", err)
	}
	log.Printf("Successfully added gameItem to firestore: %v", ref.ID, result.UpdateTime)
}
func GetItem(ctx *gin.Context, dbClient *firestore.Client) {

	type Item struct {
	}
	result, err := dbClient.Collection("items").Doc("X3tSYvMIoA39LxZTC0ne").Get(ctx)

	if err != nil {
		if status.Code(err) == codes.NotFound {
			log.Fatalf("Could not find X3tSYvMIoA39LxZTC0ne")
		} else {
			log.Fatalf("Error getting X3tSYvMIoA39LxZTC0ne: %v", err)
		}
	}
	dataMap := result.Data()

	for key, value := range dataMap {
		fmt.Println(key, value, value)
	}

}
