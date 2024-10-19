package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file")
	}

	router := http.NewServeMux()

	err = http.ListenAndServe(":8000", router)
	fmt.Println("Server running on localhost:8000")
	
	if err != nil {
		log.Fatalf("An error occurred: " + err.Error())
	}
}
