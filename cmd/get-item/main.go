package main

import (
	"arizona-server/internal/db"
	"arizona-server/internal/service"
	"net/http"
)

var gItemService *service.Service

func init() {
	dbClient := db.ConnectionToYDB()
	gItemService = service.NewService(dbClient)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	gItemService.GetItem(ctx)
}
func main() {

}
