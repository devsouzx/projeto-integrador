package main

import (
	"log"

	"github.com/devsouzx/projeto-integrador/src/configuration/database"
	"github.com/devsouzx/projeto-integrador/src/controller"
	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/repository"
	"github.com/devsouzx/projeto-integrador/src/routes"
	"github.com/devsouzx/projeto-integrador/src/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := database.SetupDB()

	database.ExecuteSQLMigrations(db)

	_, err := db.Exec(
		` 
		DROP TABLE IF EXISTS users;

		CREATE EXTENSION IF NOT EXISTS pgcrypto; 
		
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			email VARCHAR(255) NOT NULL UNIQUE,
			name VARCHAR(255),
			user_password TEXT NOT NULL,
			role VARCHAR(10) NOT NULL,
			cpf VARCHAR(14) UNIQUE NOT NULL
		);
	`,
	)
	if err != nil {
		log.Fatal(err)
	}

	teste := model.NewUserDomain("teste@discente.ufg.br", "08153160389", "@Teste123", "Teste", "paciente")
	teste.EncryptPassword()

	_, err = db.Exec(`
	INSERT INTO users (email, cpf, name, user_password, role)
	SELECT $1, $2, $3, $4,$5
`, teste.GetEmail(), teste.GetCPF(), teste.GetName(), teste.GetPassword(), teste.GetRole())

	if err != nil {
		panic("Erro ao inserir usuario teste: " + err.Error())
	}

	defer db.Close()

	router.LoadHTMLGlob("src/static/**/*.html")

	repository := repository.NewUserRepository(db)
	service := service.NewUserDomainService(repository)
	routes.InitRoutes(router.Group("/"), controller.NewUserController(service))

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
