package client

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jobergner/backent-cli/examples/action"
	"github.com/jobergner/backent-cli/examples/connect"
	"github.com/jobergner/backent-cli/examples/message"
	"github.com/jobergner/backent-cli/examples/state"
	"nhooyr.io/websocket"
)

// easyjson:skip
type Client struct {
	id             string
	mu             sync.Mutex
	actions        action.Actions
	engine         *state.Engine
	conn           connect.Connector
	router         *responseRouter
	idSignal       chan string
	messageChannel chan []byte
	patchChannel   chan []byte
}

func NewClient(actions action.Actions) (*Client, context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)

	c, _, err := websocket.Dial(ctx, "http://localhost:8080/ws", nil)
	if err != nil {
		return nil, cancel, err
	}

	client := Client{
		actions: actions,
		engine:  state.NewEngine(),
		conn:    connect.NewConnection(c, ctx),
		router: &responseRouter{
			pending: make(map[string]chan []byte),
		},
		idSignal:       make(chan string, 1),
		messageChannel: make(chan []byte),
		patchChannel:   make(chan []byte),
	}

	go client.runReadMessages()
	go client.runWriteMessages()

	client.id = <-client.idSignal

	return &client, cancel, nil
}

func (c *Client) tickSync() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.engine.Patch.IsEmpty() {
		return
	}

	patchBytes, err := c.engine.Patch.MarshalJSON()
	if err != nil {
		log.Printf("error marshalling patch: %s", err)
	}

	c.patchChannel <- patchBytes
}

func (c *Client) ReadUpdate() []byte {
	return <-c.patchChannel
}

func (c *Client) emitPatches() {
	ticker := time.NewTicker(time.Second)

	for {
		<-ticker.C

		c.tickSync()
	}
}

func (c *Client) ID() string {
	return c.id
}

func (c *Client) handleInernalError() {
	c.conn.Close()
}

func (c *Client) runReadMessages() {
	defer c.handleInernalError()

	for {
		_, msgBytes, err := c.conn.ReadMessage()

		if err != nil {
			log.Printf("unregistering client due to error while reading connection: %s", err)
			break
		}

		var msg Message
		err = msg.UnmarshalJSON(msgBytes)

		fmt.Println(string(msg.Content))

		if err != nil {
			log.Printf("error parsing message \"%s\" with error %s", string(msgBytes), err)
			continue
		}

		c.processMessageSync(msg)
	}
}

func (c *Client) runWriteMessages() {
	defer c.handleInernalError()

	for {
		msg, ok := <-c.messageChannel

		if !ok {
			log.Printf("messageChannel of client %s has been closed", c.id)
			break
		}

		c.conn.WriteMessage(msg)
	}
}

func (c *Client) processMessageSync(msg Message) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	switch msg.Kind {
	case message.MessageKindID:
		c.idSignal <- string(msg.Content)
	case message.MessageKindUpdate, message.MessageKindCurrentState:
		var patch state.State
		err := patch.UnmarshalJSON(msg.Content)
		if err != nil {
			return err
		}
		c.engine.ImportPatch(&patch)
	default:
		c.router.route(msg)
	}

	return nil
}
