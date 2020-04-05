package WebsocketHandlers

import (
	"gothello/Game"
)

type GameHub struct {
	// Registered clients.
	clients map[*GameClient]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *GameClient

	// Unregister requests from clients.
	unregister chan *GameClient
}

func NewGameHub() *GameHub {
	return &GameHub{
		broadcast:  make(chan []byte),
		register:   make(chan *GameClient),
		unregister: make(chan *GameClient),
		clients:    make(map[*GameClient]bool),
	}
}

func (h *GameHub) MessageClientByPlayerId(playerId int, message []byte) {
	for client, _ := range h.clients {
		if client.playerId == playerId {
			client.send <- message
			break
		}
	}
}

func (h *GameHub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				Game.GetGothelloServerState().DisconnectPlayer(client.playerId)
			}
		}
	}
}