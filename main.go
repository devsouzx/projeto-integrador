package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/devsouzx/projeto-integrador/src/configuration/database"
	"github.com/devsouzx/projeto-integrador/src/dependency"
	"github.com/devsouzx/projeto-integrador/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	from := "+18166404985"
	to := "+5562981743829"
	message := "Seu código de recuperação é 123456."
	
	twilioURL := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
	
	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", from)
	msgData.Set("Body", message)
	msgDataReader := *strings.NewReader(msgData.Encode())
	
	req, _ := http.NewRequest("POST", twilioURL, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Erro ao enviar SMS:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Resposta:", string(body))

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
