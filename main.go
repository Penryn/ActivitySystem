package main

import (
	"activitySystem/internal/global"
	"activitySystem/internal/middleware"
	"activitySystem/internal/pkg/database"
	"activitySystem/internal/router"
	"activitySystem/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Init()
	service.Init(db)
	r := gin.Default()
	r.NoMethod(middleware.HandleNotFound)
	r.NoRoute(middleware.HandleNotFound)
	r.Static("public", "./public")
	router.Init(r)
	err := r.Run(":" + global.Config.GetString("server.port"))
	if err != nil {
		log.Fatal(err)
	}
}
