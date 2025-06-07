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

	teste := model.NewPaciente("teste@discente.ufg.br", "@Teste123", "Teste", "08153160389")
	teste.EncryptPassword()

	_, err := db.Exec(`
	INSERT INTO paciente (email, nomecompleto, senha, cpf)
	SELECT $1, $2, $3, $4
`, teste.GetEmail(), teste.GetName(), teste.GetPassword(), teste.GetCPF())

	if err != nil {
		panic("Erro ao inserir paciente teste: " + err.Error())
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
