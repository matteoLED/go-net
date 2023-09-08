package server

import (
	"sync"
)

// Flow représente un flux réseau.
type Flow struct {
	SourceIP        string
	DestinationIP   string
	SourcePort      uint16
	DestinationPort uint16
}

// FlowTracker permet de suivre les flux réseau.
type FlowTracker struct {
	mu    sync.Mutex
	Flows map[Flow]int
}

// NewFlowTracker crée une nouvelle instance de FlowTracker.
func NewFlowTracker() *FlowTracker {
	return &FlowTracker{
		Flows: make(map[Flow]int),
	}
}

// AddFlow ajoute un flux à la liste des flux.
func (ft *FlowTracker) AddFlow(flow Flow) {
	ft.mu.Lock()
	defer ft.mu.Unlock()
	ft.Flows[flow]++
}

// GetFlowCount renvoie le nombre de paquets dans un flux donné.
func (ft *FlowTracker) GetFlowCount(flow Flow) int {
	ft.mu.Lock()
	defer ft.mu.Unlock()
	return ft.Flows[flow]
}
