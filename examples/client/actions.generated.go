package client

import (
	"time"

	"github.com/google/uuid"
	"github.com/jobergner/backent-cli/examples/action"
	"github.com/jobergner/backent-cli/examples/logging"
	"github.com/rs/zerolog/log"
)

func (c *Client) AddItemToPlayer(params action.AddItemToPlayerParams) (action.AddItemToPlayerResponse, error) {
	c.mu.Lock()
	c.actions.AddItemToPlayer.Broadcast(params, c.engine, "", c.id)
	c.mu.Unlock()

	msgContent, err := params.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(action.MessageKindAction_addItemToPlayer)).Msg("failed marshalling parameters")
		return action.AddItemToPlayerResponse{}, err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(action.MessageKindAction_addItemToPlayer)).Msg("failed generating message ID")
		return action.AddItemToPlayerResponse{}, err
	}

	idString := id.String()

	msg := Message{idString, action.MessageKindAction_addItemToPlayer, msgContent}

	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageID, msg.ID).Str(logging.Message, string(msgBytes)).Str(logging.MessageKind, string(action.MessageKindAction_addItemToPlayer)).Msg("failed marshalling message")
		return action.AddItemToPlayerResponse{}, err
	}

	responseChan := make(chan []byte)
	c.router.add(idString, responseChan)
	defer c.router.remove(idString)

	c.messageChannel <- msgBytes

	select {
	case <-time.After(2 * time.Second):
		log.Err(ErrResponseTimeout).Str(logging.MessageID, msg.ID).Msg("timed out waiting for response")
		return action.AddItemToPlayerResponse{}, ErrResponseTimeout
	case responseBytes := <-responseChan:
		var res action.AddItemToPlayerResponse

		err := res.UnmarshalJSON(responseBytes)
		if err != nil {
			log.Err(err).Str(logging.MessageID, msg.ID).Str(logging.MessageKind, string(action.MessageKindAction_addItemToPlayer)).Msg("failed unmarshalling response")
			return action.AddItemToPlayerResponse{}, err
		}

		return res, nil
	}
}

func (c *Client) MovePlayer(params action.MovePlayerParams) error {
	c.mu.Lock()
	c.actions.MovePlayer.Broadcast(params, c.engine, "", c.id)
	c.mu.Unlock()

	msgContent, err := params.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(action.MessageKindAction_movePlayer)).Msg("failed marshalling parameters")
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(action.MessageKindAction_movePlayer)).Msg("failed generating message ID")
		return err
	}

	idString := id.String()

	msg := Message{idString, action.MessageKindAction_movePlayer, msgContent}

	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageID, msg.ID).Str(logging.Message, string(msgBytes)).Str(logging.MessageKind, string(action.MessageKindAction_movePlayer)).Msg("failed marshalling message")
		return err
	}

	c.messageChannel <- msgBytes

	return nil
}

func (c *Client) SpawnZoneItems(params action.SpawnZoneItemsParams) (action.SpawnZoneItemsResponse, error) {
	c.mu.Lock()
	c.actions.SpawnZoneItems.Broadcast(params, c.engine, "", c.id)
	c.mu.Unlock()

	msgContent, err := params.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(action.MessageKindAction_spawnZoneItems)).Msg("failed marshalling parameters")
		return action.SpawnZoneItemsResponse{}, err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(action.MessageKindAction_spawnZoneItems)).Msg("failed generating message ID")
		return action.SpawnZoneItemsResponse{}, err
	}

	idString := id.String()

	msg := Message{idString, action.MessageKindAction_spawnZoneItems, msgContent}

	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageID, msg.ID).Str(logging.Message, string(msgBytes)).Str(logging.MessageKind, string(action.MessageKindAction_spawnZoneItems)).Msg("failed marshalling message")
		return action.SpawnZoneItemsResponse{}, err
	}

	responseChan := make(chan []byte)
	c.router.add(idString, responseChan)
	defer c.router.remove(idString)

	c.messageChannel <- msgBytes

	select {
	case <-time.After(2 * time.Second):
		log.Err(ErrResponseTimeout).Str(logging.MessageID, msg.ID).Msg("timed out waiting for response")
		return action.SpawnZoneItemsResponse{}, ErrResponseTimeout
	case responseBytes := <-responseChan:
		var res action.SpawnZoneItemsResponse

		err := res.UnmarshalJSON(responseBytes)
		if err != nil {
			log.Err(err).Str(logging.MessageID, msg.ID).Str(logging.MessageKind, string(action.MessageKindAction_spawnZoneItems)).Msg("failed unmarshalling response")
			return action.SpawnZoneItemsResponse{}, err
		}

		return res, nil
	}
}
