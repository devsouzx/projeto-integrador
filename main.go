package main

import (
	"log"
	"time"

	"github.com/devsouzx/projeto-integrador/src/configuration/database"
	"github.com/devsouzx/projeto-integrador/src/dependency"
	"github.com/devsouzx/projeto-integrador/src/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:8080"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

	db := database.SetupDB()
	database.ExecuteSQLMigrations(db)
	defer db.Close()

	router.LoadHTMLGlob("src/static/**/*.html")

	container := dependency.InitContainer(db)

	routes.InitRoutes(router.Group("/"), container)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
