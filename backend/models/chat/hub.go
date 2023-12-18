package chat

type Hub struct {
	Clients      map[*Client]bool
	RegisterCh   chan *Client
	UnRegisterCh chan *Client
	BroadcastCh  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		Clients:      make(map[*Client]bool),
		RegisterCh:   make(chan *Client),
		UnRegisterCh: make(chan *Client),
		BroadcastCh:  make(chan []byte),
	}
}

func (h *Hub) Register(c *Client) {
	h.Clients[c] = true
}

func (h *Hub) Unregister(c *Client) {
	delete(h.Clients, c)
}

func (h *Hub) BroadCastToAllClient(msg []byte) {
	for c := range h.Clients {
		c.SendCh <- msg
	}
}
