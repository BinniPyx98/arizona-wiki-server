package main

import (
	"arizona-server/internal/db"
	"arizona-server/internal/service"
	"fmt"
	"log"
	"net/http"
)

var gItemService *service.Service

func init() {
	log.Println("Function step 0: init func get item")
	dbClient, err := db.ConnectionToYDB()
	if err != nil {
		fmt.Printf("DB connection error: %v\n", err)
		return
	}
	gItemService = service.NewService(dbClient)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	log.Printf("Function step 1")
	if gItemService == nil {
		fmt.Println("[ERROR] Service not initialized, returning 500")
		http.Error(w, `{"error":"db not initialized"}`, http.StatusInternalServerError)
		return
	}
	gItemService.GetItem(r.Context())
}
