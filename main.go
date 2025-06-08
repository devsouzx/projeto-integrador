package main

import (
	"fmt"
	"log"

	"github.com/devsouzx/projeto-integrador/src/configuration/database"
	"github.com/devsouzx/projeto-integrador/src/dependency"
	"github.com/devsouzx/projeto-integrador/src/routes"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

func main() {
	const (
		host = "smtp.gmail.com"
		port = 587
		username = "joaoemanuel2215@gmail.com"
		passwword = "efnypeuaxbcdkzsf"
	)

	dialer := gomail.NewDialer(host, port, username, passwword)
	msg := gomail.NewMessage()
	msg.SetHeader("From", username)
	msg.SetHeader("To", "joaoemanuelmarinhosousa@gmail.com")
	msg.SetHeader("Subject", "Enviando email")
	msg.SetBody("text/plain", "Email teste usando gomail ksgdhkd")
	if err := dialer.DialAndSend(msg); err != nil {
		panic(err)
	}

	fmt.Println("Mensagem enviada com sucesso")
	router := gin.Default()
	
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
