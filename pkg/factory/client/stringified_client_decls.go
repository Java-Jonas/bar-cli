// this file was generated by https://github.com/jobergner/decltostring

package client

const actions_generated_go_import string = `import (
	"time"
	"github.com/google/uuid"
	"github.com/jobergner/backent-cli/examples/logging"
	"github.com/jobergner/backent-cli/examples/message"
	"github.com/rs/zerolog/log"
)`

const _AddItemToPlayer_Client_func string = `func (c *Client) AddItemToPlayer(params message.AddItemToPlayerParams) (message.AddItemToPlayerResponse, error) {
	c.mu.Lock()
	c.controller.AddItemToPlayerBroadcast(params, c.engine, "", c.id)
	c.mu.Unlock()
	msgContent, err := params.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_addItemToPlayer)).Msg("failed marshalling parameters")
		return message.AddItemToPlayerResponse{}, err
	}
	id, err := uuid.NewRandom()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_addItemToPlayer)).Msg("failed generating message ID")
		return message.AddItemToPlayerResponse{}, err
	}
	idString := id.String()
	msg := Message{idString, message.MessageKindAction_addItemToPlayer, msgContent}
	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageID, msg.ID).Str(logging.Message, string(msgBytes)).Str(logging.MessageKind, string(message.MessageKindAction_addItemToPlayer)).Msg("failed marshalling message")
		return message.AddItemToPlayerResponse{}, err
	}
	responseChan := make(chan []byte)
	c.router.add(idString, responseChan)
	defer c.router.remove(idString)
	c.messageChannel <- msgBytes
	select {
	case <-time.After(2 * time.Second):
		log.Err(ErrResponseTimeout).Str(logging.MessageID, msg.ID).Msg("timed out waiting for response")
		return message.AddItemToPlayerResponse{}, ErrResponseTimeout
	case responseBytes := <-responseChan:
		var res message.AddItemToPlayerResponse
		err := res.UnmarshalJSON(responseBytes)
		if err != nil {
			log.Err(err).Str(logging.MessageID, msg.ID).Str(logging.MessageKind, string(message.MessageKindAction_addItemToPlayer)).Msg("failed unmarshalling response")
			return message.AddItemToPlayerResponse{}, err
		}
		return res, nil
	}
}`

const _MovePlayer_Client_func string = `func (c *Client) MovePlayer(params message.MovePlayerParams) error {
	c.mu.Lock()
	c.controller.MovePlayerBroadcast(params, c.engine, "", c.id)
	c.mu.Unlock()
	msgContent, err := params.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_movePlayer)).Msg("failed marshalling parameters")
		return err
	}
	id, err := uuid.NewRandom()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_movePlayer)).Msg("failed generating message ID")
		return err
	}
	idString := id.String()
	msg := Message{idString, message.MessageKindAction_movePlayer, msgContent}
	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageID, msg.ID).Str(logging.Message, string(msgBytes)).Str(logging.MessageKind, string(message.MessageKindAction_movePlayer)).Msg("failed marshalling message")
		return err
	}
	c.messageChannel <- msgBytes
	return nil
}`

const _SpawnZoneItems_Client_func string = `func (c *Client) SpawnZoneItems(params message.SpawnZoneItemsParams) (message.SpawnZoneItemsResponse, error) {
	c.mu.Lock()
	c.controller.SpawnZoneItemsBroadcast(params, c.engine, "", c.id)
	c.mu.Unlock()
	msgContent, err := params.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_spawnZoneItems)).Msg("failed marshalling parameters")
		return message.SpawnZoneItemsResponse{}, err
	}
	id, err := uuid.NewRandom()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(message.MessageKindAction_spawnZoneItems)).Msg("failed generating message ID")
		return message.SpawnZoneItemsResponse{}, err
	}
	idString := id.String()
	msg := Message{idString, message.MessageKindAction_spawnZoneItems, msgContent}
	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageID, msg.ID).Str(logging.Message, string(msgBytes)).Str(logging.MessageKind, string(message.MessageKindAction_spawnZoneItems)).Msg("failed marshalling message")
		return message.SpawnZoneItemsResponse{}, err
	}
	responseChan := make(chan []byte)
	c.router.add(idString, responseChan)
	defer c.router.remove(idString)
	c.messageChannel <- msgBytes
	select {
	case <-time.After(2 * time.Second):
		log.Err(ErrResponseTimeout).Str(logging.MessageID, msg.ID).Msg("timed out waiting for response")
		return message.SpawnZoneItemsResponse{}, ErrResponseTimeout
	case responseBytes := <-responseChan:
		var res message.SpawnZoneItemsResponse
		err := res.UnmarshalJSON(responseBytes)
		if err != nil {
			log.Err(err).Str(logging.MessageID, msg.ID).Str(logging.MessageKind, string(message.MessageKindAction_spawnZoneItems)).Msg("failed unmarshalling response")
			return message.SpawnZoneItemsResponse{}, err
		}
		return res, nil
	}
}`

const client_go_import string = `import (
	"context"
	"sync"
	"time"
	"github.com/jobergner/backent-cli/examples/connect"
	"github.com/jobergner/backent-cli/examples/logging"
	"github.com/jobergner/backent-cli/examples/message"
	"github.com/jobergner/backent-cli/examples/state"
	"github.com/rs/zerolog/log"
	"nhooyr.io/websocket"
)`

const _Client_type string = `type Client struct {
	id		string
	mu		sync.Mutex
	controller	Controller
	engine		*state.Engine
	conn		connect.Connector
	router		*responseRouter
	idSignal	chan string
	messageChannel	chan []byte
	patchChannel	chan []byte
}`

const _NewClient_func string = `func NewClient(controller Controller) (*Client, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	c, _, err := websocket.Dial(ctx, "http://localhost:8080/ws", nil)
	if err != nil {
		log.Err(err).Msg("failed creating client while dialing server")
		return nil, cancel, err
	}
	client := Client{controller: controller, engine: state.NewEngine(), conn: connect.NewConnection(c, ctx), router: &responseRouter{pending: make(map[string]chan []byte)}, idSignal: make(chan string, 1), messageChannel: make(chan []byte), patchChannel: make(chan []byte)}
	go client.runReadMessages()
	go client.runWriteMessages()
	client.id = <-client.idSignal
	return &client, cancel, nil
}`

const tickSync_Client_func string = `func (c *Client) tickSync() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.engine.Patch.IsEmpty() {
		return
	}
	patchBytes, err := c.engine.Patch.MarshalJSON()
	if err != nil {
		log.Err(err).Msg("failed marshalling patch")
		return
	}
	c.patchChannel <- patchBytes
}`

const _ReadUpdate_Client_func string = `func (c *Client) ReadUpdate() []byte {
	return <-c.patchChannel
}`

const emitPatches_Client_func string = `func (c *Client) emitPatches() {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		c.tickSync()
	}
}`

const _ID_Client_func string = `func (c *Client) ID() string {
	return c.id
}`

const handleInernalError_Client_func string = `func (c *Client) handleInernalError() {
	c.conn.Close()
}`

const runReadMessages_Client_func string = `func (c *Client) runReadMessages() {
	defer c.handleInernalError()
	for {
		_, msgBytes, err := c.conn.ReadMessage()
		if err != nil {
			log.Err(err).Msg("failed reading message")
			break
		}
		var msg Message
		err = msg.UnmarshalJSON(msgBytes)
		if err != nil {
			log.Err(err).Str(logging.Message, string(msgBytes)).Msg("failed unmarshalling message")
			continue
		}
		c.processMessageSync(msg)
	}
}`

const runWriteMessages_Client_func string = `func (c *Client) runWriteMessages() {
	defer c.handleInernalError()
	for {
		msg, ok := <-c.messageChannel
		if !ok {
			log.Warn().Msg("failed while attempted sending to closed client message channel")
			break
		}
		c.conn.WriteMessage(msg)
	}
}`

const processMessageSync_Client_func string = `func (c *Client) processMessageSync(msg Message) error {
	switch msg.Kind {
	case message.MessageKindID:
		c.idSignal <- string(msg.Content)
	case message.MessageKindUpdate, message.MessageKindCurrentState:
		var patch state.State
		err := patch.UnmarshalJSON(msg.Content)
		if err != nil {
			log.Warn().Msg("failed unmarshalling patch")
			return err
		}
		c.mu.Lock()
		c.engine.ImportPatch(&patch)
		c.mu.Unlock()
	default:
		c.router.route(msg)
	}
	return nil
}`

const controller_generated_go_import string = `import (
	"github.com/jobergner/backent-cli/examples/message"
	"github.com/jobergner/backent-cli/examples/state"
)`

const _Controller_type string = `type Controller interface {
	AddItemToPlayerBroadcast(params message.AddItemToPlayerParams, engine *state.Engine, roomName, clientID string)
	EAddItemToPlayerEmit(params message.AddItemToPlayerParams, engine *state.Engine, roomName, clientID string) message.AddItemToPlayerResponse
	MovePlayerBroadcast(params message.MovePlayerParams, engine *state.Engine, roomName, clientID string)
	MovePlayerEmit(params message.MovePlayerParams, engine *state.Engine, roomName, clientID string)
	SpawnZoneItemsBroadcast(params message.SpawnZoneItemsParams, engine *state.Engine, roomName, clientID string)
	SpawnZoneItemsEmit(params message.SpawnZoneItemsParams, engine *state.Engine, roomName, clientID string) message.SpawnZoneItemsResponse
}`

const error_go_import string = `import "errors"`

const _ErrResponseTimeout_type string = `var (
	ErrResponseTimeout = errors.New("timeout")
)`

const message_go_import string = `import "github.com/jobergner/backent-cli/examples/message"`

const _Message_type string = `type Message struct {
	ID	string		` + "`" + `json:"id"` + "`" + `
	Kind	message.Kind	` + "`" + `json:"kind"` + "`" + `
	Content	[]byte		` + "`" + `json:"content"` + "`" + `
}`

const response_router_go_import string = `import (
	"sync"
	"github.com/jobergner/backent-cli/examples/logging"
	"github.com/rs/zerolog/log"
)`

const responseRouter_type string = `type responseRouter struct {
	pending	map[string]chan []byte
	mu	sync.Mutex
}`

const add_responseRouter_func string = `func (r *responseRouter) add(id string, ch chan []byte) {
	r.mu.Lock()
	log.Debug().Str(logging.MessageID, id).Msg("adding channel to router")
	r.pending[id] = ch
	r.mu.Unlock()
}`

const remove_responseRouter_func string = `func (r *responseRouter) remove(id string) {
	r.mu.Lock()
	log.Debug().Str(logging.MessageID, id).Msg("removing channel to router")
	ch := r.pending[id]
	delete(r.pending, id)
	close(ch)
	r.mu.Unlock()
}`

const route_responseRouter_func string = `func (r *responseRouter) route(response Message) {
	ch, ok := r.pending[response.ID]
	if !ok {
		log.Warn().Str(logging.MessageID, response.ID).Msg("cannot find channel for routing response")
		return
	}
	log.Debug().Str(logging.MessageID, response.ID).Msg("routing response")
	ch <- response.Content
}`
