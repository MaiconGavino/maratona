package handlers

import (
	"net/http"
)

func CadastroHandler(w http.ResponseWriter, r *http.Request) {
	// Aqui você pode adicionar a lógica para lidar com o cadastro
	w.Write([]byte("Cadastro Page"))
}
