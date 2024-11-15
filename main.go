package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Hola Mundo!")
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Servidor escuchando en el puerto " + port + "...")
	http.ListenAndServe(":"+port, nil)
}
