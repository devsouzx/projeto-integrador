package main

import (
	"log"

	"github.com/devsouzx/projeto-integrador/src/configuration/database"
	"github.com/devsouzx/projeto-integrador/src/controller"
	"github.com/devsouzx/projeto-integrador/src/repository"
	"github.com/devsouzx/projeto-integrador/src/routes"
	"github.com/devsouzx/projeto-integrador/src/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := database.SetupDB()

	database.ExecuteSQLMigrations(db)

	defer db.Close()

	router.LoadHTMLGlob("src/static/**/*.html")

	repository := repository.NewUserRepository(db)
	service := service.NewUserDomainService(repository)
	routes.InitRoutes(router.Group("/"), controller.NewUserController(service))

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
