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
	// Obtener el puerto de la variable de entorno de Heroku
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Si no está configurado, usa el puerto 8080 por defecto
	}

	// Registrar el handler para la ruta principal
	http.HandleFunc("/", handler)

	fmt.Println("Servidor escuchando en el puerto " + port + "...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
