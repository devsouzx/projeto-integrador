package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/devsouzx/projeto-integrador/src/configuration/database"
	"github.com/devsouzx/projeto-integrador/src/dependency"
	"github.com/devsouzx/projeto-integrador/src/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	loc, _ := time.LoadLocation("America/Sao_Paulo")

// Simule uma entrada, por exemplo do formulário
entrada := "2025-07-12 14:00"

// Parse já no fuso correto
t, err := time.ParseInLocation("2006-01-02 15:04", entrada, loc)
if err != nil {
	fmt.Println("Erro ao converter:", err)
	return
}

fmt.Println("Data final:", t)             // 2025-07-12 14:00 -03:00
fmt.Println("Em UTC:", t.UTC())  
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"formatDate": func(dateString string) (string, error) {
			layout := "2006-01-02T15:04:05Z"
			t, err := time.Parse(layout, dateString)
			if err != nil {
				return "", err
			}
			return t.Format("02/01/2006"), nil
		},
	})

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	db := database.SetupDB()
	database.ExecuteSQLMigrations(db)
	defer db.Close()

	router.LoadHTMLGlob("src/static/**/*.html")

	container := dependency.InitContainer(db)

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	routes.InitRoutes(router.Group("/"), container)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
