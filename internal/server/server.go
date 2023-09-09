package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// Server représente le serveur Go-net.
type Server struct {
	Port      int
	FlowTrack *FlowTracker
}

// NewServer crée une nouvelle instance de serveur.
func NewServer(port int) *Server {
	return &Server{
		Port:      port,
		FlowTrack: NewFlowTracker(),
	}
}

func (s *Server) setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Encodez le packet en JSON
		packetJSON, err := json.Marshal("Welcome to Go-net !!")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Définissez le type de contenu dans l'en-tête de la réponse HTTP
		w.Header().Set("Content-Type", "application/json")

		// Écrivez le JSON dans la réponse HTTP
		w.Write(packetJSON)
	})

	http.HandleFunc("/packet", func(w http.ResponseWriter, r *http.Request) {
		// Créez une liste de PacketInfo pour stocker les informations des paquets
		packets := []Flow{}

		// Ajoutez les informations de chaque paquet à la liste
		packet1 := Flow{
			SourceIP:        r.RemoteAddr,
			DestinationIP:   "destinationIP",
			SourcePort:      uint16(12345),
			DestinationPort: uint16(54321),
		}
		packets = append(packets, packet1)

		packet2 := Flow{
			SourceIP:        r.RemoteAddr,
			DestinationIP:   "destinationIP",
			SourcePort:      uint16(23456),
			DestinationPort: uint16(65432),
		}
		packets = append(packets, packet2)

		// Encodez la liste de paquets en JSON
		packetsJSON, err := json.Marshal(packets)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Définissez le type de contenu dans l'en-tête de la réponse HTTP
		w.Header().Set("Content-Type", "application/json")

		// Écrivez le JSON dans la réponse HTTP
		w.Write(packetsJSON)
	})

}

func (s *Server) handlePacketRequest(w http.ResponseWriter, r *http.Request) {
	// Vous pouvez mettre ici la logique pour traiter les paquets réseau
	// Utilisez la méthode HandlePacket pour traiter les paquets, par exemple :
	packet := gopacket.NewPacket([]byte{}, layers.LayerTypeEthernet, gopacket.Default)
	HandlePacket(packet)

	// Vous pouvez envoyer une réponse au client ici si nécessaire
	fmt.Fprintf(w, "Packet handling complete")
}

// Start démarre le serveur.
func (s *Server) Start() error {
	s.setupRoutes() // Configure les routes HTTP

	fmt.Printf("Server is running on port %d...\n", s.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil)
}
func NetworkHandler(packet gopacket.Packet) {

	HandlePacket(packet)

}
