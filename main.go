package main

import (
	"cassandraTest/pkg/utils"
	"cassandraTest/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server starting")
	urlRouter := router.UrlRouter()
	url := fmt.Sprintf("%s:%s", utils.GoDotEnvVariable("BASE_URL"), utils.GoDotEnvVariable("USSD_PORT_GP"))
	log.Printf("Failed to Listen and Server: %v", http.ListenAndServe(url, urlRouter))
}