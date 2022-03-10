package state

import (
	"log"
	"sync"
	"time"
)

type RoomMode int

const (
	RoomModeIdle RoomMode = iota
	RoomModeRunning
	RoomModeTerminating
)

type Room struct {
	name            string
	mu              sync.Mutex
	clients         map[*Client]struct{}
	incomingClients map[*Client]struct{}
	state           *Engine
	actions         Actions
	mode            RoomMode
}

func newRoom(actions Actions) *Room {
	return &Room{
		clients:         make(map[*Client]struct{}),
		incomingClients: make(map[*Client]struct{}),
		state:           newEngine(),
		actions:         actions,
	}
}

func (r *Room) Name() string {
	return r.name
}

func (r *Room) RemoveClient(client *Client) {
	r.removeClient(client)
}

func (r *Room) AddClient(client *Client) {
	r.addClient(client)
}

func (r *Room) AlterState(fn func(*Engine)) {
	r.mu.Lock()
	defer r.mu.Unlock()
	fn(r.state)
}

func (r *Room) RangeClients(fn func(client *Client)) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for c := range r.incomingClients {
		fn(c)
	}
	for c := range r.clients {
		fn(c)
	}
}

func (r *Room) processMessageSync(msg Message) {
	r.mu.Lock()
	defer r.mu.Unlock()

	response, err := r.processClientMessage(msg)
	if err != nil {
		return
	}

	// actions may not have a response
	// which means `response` is empty here
	// and can be skipped
	if response.client == nil {
		return
	}

	responseBytes, err := response.MarshalJSON()
	if err != nil {
		log.Printf("error marshalling pending response message: %s", err)
		return
	}

	select {
	case response.client.messageChannel <- responseBytes:
	default:
		log.Printf("client's message buffer full -> dropping client %s", response.client.id)
		r.unregisterClientAsync(response.client)
	}
}

func (r *Room) addClient(client *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()

	client.room = r
	r.incomingClients[client] = struct{}{}
}

func (r *Room) removeClient(client *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.clients, client)
	delete(r.incomingClients, client)
}

// removes a client from a room and closes the connection
func (r *Room) unregisterClientSync(client *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.unregisterClientAsync(client)
}

func (r *Room) run(sideEffects SideEffects, fps int) {
	ticker := time.NewTicker(time.Second / time.Duration(fps))

	for {
		<-ticker.C
		r.tickSync(sideEffects)
		if r.mode == RoomModeTerminating {
			log.Printf("shutting down room %s", r.name)
			return
		}
	}
}

func (r *Room) Deploy(sideEffects SideEffects, fps int) {
	if sideEffects.OnDeploy != nil {
		r.mu.Lock()
		sideEffects.OnDeploy(r.state)
		r.mu.Unlock()
	}

	go r.run(sideEffects, fps)
}
