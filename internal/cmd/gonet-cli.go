// gonet-cli.go

package view

import (
	"flag"
	"fmt"
)

func View() {
	// Définition des drapeaux de ligne de commande
	port := flag.Int("port", 8082, "Port number for the Go-net server")
	verbose := flag.Bool("verbose", false, "Enable verbose mode")

	// Analyse des drapeaux de ligne de commande
	flag.Parse()

	// Affichage des valeurs des drapeaux
	fmt.Printf("Go-net CLI\n")
	fmt.Printf("Port: %d\n", *port)
	fmt.Printf("Verbose mode: %v\n", *verbose)

	// Vous pouvez ajouter ici le code pour interagir avec votre projet Go-net en utilisant les drapeaux définis
	// par exemple, pour démarrer votre serveur Go-net avec le port spécifié
	// server := server.NewServer(*port)
	// err := server.Start()
	// if err != nil {
	//     fmt.Printf("Error: %v\n", err)
	// }
}
