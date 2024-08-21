package main

import (
	"github.com/gorilla/mux"
	"github.com/maicongavino/database"
	"log"
	"net/http"
)

func main() {
	// Conectar ao banco de dados
	if err := database.Connect(); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Configurar o roteador
	r := mux.NewRouter()

	// Definir rotas
	r.HandleFunc("/cadastro", handlers.CadastroHandler).Methods("GET", "POST")
	r.HandleFunc("/voto", handlers.VotoHandler).Methods("GET", "POST")
	r.HandleFunc("/result", handlers.ResultHandler).Methods("GET")

	// Iniciar o servidor
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
