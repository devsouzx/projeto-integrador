package main

import (
	"log"

	"github.com/devsouzx/projeto-integrador/src/configuration/database"
	"github.com/devsouzx/projeto-integrador/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := database.SetupDB()
	defer db.Close()

  	router.LoadHTMLGlob("src/static/**/*.html")
	
	routes.InitRoutes(router.Group("/"))

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}