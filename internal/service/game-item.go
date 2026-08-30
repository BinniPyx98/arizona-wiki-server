package service

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	dbClient *firestore.Client
}

func NewService(dbClient *firestore.Client) *Service {
	return &Service{dbClient: dbClient}
}

func (s *Service) AddGameItem(ctx context.Context) {
	gameItem := map[string]any{
		"first":  "Ada",
		"last":   "Lovelace",
		"born":   1815,
		"active": true,
	}
	ref, result, err := s.dbClient.Collection("items").Add(ctx, gameItem)
	if err != nil {
		log.Fatal("Error while adding gameItem to firestore:", err)
	}
	log.Printf("Successfully added gameItem to firestore: %v", ref.ID, result.UpdateTime)
}

func (s *Service) GetItem(ctx context.Context) {

	type Item struct {
	}
	result, err := s.dbClient.Collection("items").Doc("X3tSYvMIoA39LxZTC0ne").Get(ctx)

	if err != nil {
		if status.Code(err) == codes.NotFound {
			log.Fatalf("Could not find X3tSYvMIoA39LxZTC0ne")
		} else {
			log.Fatalf("Error getting X3tSYvMIoA39LxZTC0ne: %v", err)
		}
	}
	dataMap := result.Data()

	for key, value := range dataMap {
		log.Print(key, value, value)
	}

}
