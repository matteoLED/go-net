// gonet-web.go

package view

import (
	"fmt"
	"net/http"
)

func Home() {
	// Gestionnaire pour la racine de l'application web
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Go-net Web Interface!")
	})

	// Démarrez le serveur web sur le port 8080
	port := 8080
	fmt.Printf("Web server is running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
