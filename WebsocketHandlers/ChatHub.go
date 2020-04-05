package WebsocketHandlers

type ChatHub struct {
	// Registered clients.
	clients map[*ChatClient]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *ChatClient

	// Unregister requests from clients.
	unregister chan *ChatClient
}

func NewChatHub() *ChatHub {
	return &ChatHub{
		broadcast:  make(chan []byte),
		register:   make(chan *ChatClient),
		unregister: make(chan *ChatClient),
		clients:    make(map[*ChatClient]bool),
	}
}

func (h *ChatHub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
