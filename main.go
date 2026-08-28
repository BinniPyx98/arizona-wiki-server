package main

import (
	"arizona-server/cmd/controller"
	"arizona-server/internal/db"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

var DBClient *firestore.Client

func init() {
	DBClient = db.ConnectionToYDB()

	//config.GetEnv()

	//config.DB.AutoMigrate(entity.GameItem{})

}

func main() {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		fmt.Println("1")
		ctx.JSON(200, "https://avatarko.ru/img/kartinka/14/zhivotnye_kot_13379.jpg")
	})

	router.POST("/addItem", func(context *gin.Context) {
		controller.AddGameItem(context, DBClient)
		controller.GetItem(context, DBClient)
	})

	router.GET("getItem", func(context *gin.Context) {
		controller.GetItem(context, DBClient)
	})

	router.Run()
}
