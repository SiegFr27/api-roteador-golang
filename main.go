package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"Roteador/Routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	port := os.Getenv("PORT")

	router := Routes.SetupRouter()

	fmt.Println("Servidor iniciado com sucesso", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
