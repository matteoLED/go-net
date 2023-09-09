package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	"net/http"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	
)

// Server représente le serveur Go-net.
type Server struct {
	Port        int
	FlowTrack   *FlowTracker
	PacketCount  func(*sync.WaitGroup, chan int)
}

// NewServer crée une nouvelle instance de serveur.
func NewServer(port int) *Server {
  
  return &Server{
        Port:        port,
        FlowTrack:   NewFlowTracker(),
		PacketCount: packetCount,
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
	http.HandleFunc("/packet", s.HandlePacketRequest)

}

type PacketInfo struct {
	EthernetSrcMAC  string
	EthernetDstMAC  string
	SourceIP        string
	DestinationIP   string
	SourcePort      uint16
	DestinationPort uint16
	// Ajoutez d'autres informations des paquets ici si nécessaire
	// Exemple : Informations sur la couche IP
	TTL         uint8
	Protocol    uint8
	TotalLength uint16
	// Exemple : Informations sur la couche TCP
	SequenceNumber       uint32
	AcknowledgmentNumber uint32
	// Ajoutez d'autres champs au besoin
}

func (s *Server) HandlePacketRequest(w http.ResponseWriter, r *http.Request) {
	// Lisez le paquet de la requête HTTP (vous devez mettre en place la logique pour le faire)
	packet, err := ReadPacketFromHTTPRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Créez une structure PacketInfo pour stocker les informations du paquet
	packetInfo := PacketInfo{}
	fmt.Printf("Ethernet Source MAC: %s, Destination MAC: %s\n", packetInfo.EthernetSrcMAC, packetInfo.EthernetDstMAC)
	fmt.Printf("Source IP: %s, Destination IP: %s\n", packetInfo.SourceIP, packetInfo.DestinationIP)
	fmt.Printf("Source Port: %d, Destination Port: %d\n", packetInfo.SourcePort, packetInfo.DestinationPort)

	// Utilisez la fonction HandlePacket pour collecter les informations du paquet
	HandlePacket(packet, &packetInfo)

	// Encodez les informations du paquet en JSON
	packetInfoJSON, err := json.Marshal(packetInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Définissez le type de contenu dans l'en-tête de la réponse HTTP
	w.Header().Set("Content-Type", "application/json")

	// Écrivez le JSON dans la réponse HTTP
	w.Write(packetInfoJSON)
}

// Start démarre le serveur.
func (s *Server) Start() error {
	s.setupRoutes() // Configure les routes HTTP

	fmt.Printf("Server is running on port %d...\n", s.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil)
}

func ReadPacketFromHTTPRequest(r *http.Request) (gopacket.Packet, error) {
	// Lisez les données binaires du corps de la requête
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// Créez un paquet à partir des données binaires
	packet := gopacket.NewPacket(body, layers.LayerTypeEthernet, gopacket.Default)
	return packet, nil
}
