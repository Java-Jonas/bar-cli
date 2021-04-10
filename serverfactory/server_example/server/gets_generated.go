package state

import (
	"errors"
)

const (
	messageKindAction_MovePlayer messageKind = iota + 1
	messageKindAction_addItemToPlayer
	messageKindAction_spawnZoneItems
)

type _MovePlayerParams struct {
	PlayerID PlayerID `json:"playerID"`
	ChangeX  float64  `json:"changeX"`
	ChangeY  float64  `json:"changeY"`
}

type _addItemToPlayerParams struct {
	Item     tItem    `json:"item"`
	PlayerID PlayerID `json:"playerID"`
}

type _spawnZoneItemsParams struct {
	Items []tItem `json:"items"`
}

type actions struct {
	movePlayer           func(PlayerID, float64, float64, *Engine)
	addItemToPlayer      func(tItem, PlayerID, *Engine)
	spawnZoneItemsParams func([]tItem, *Engine)
}

func (r *Room) processClientMessage(msg message) error {
	switch messageKind(msg.Kind) {
	case messageKindAction_addItemToPlayer:
		var params _addItemToPlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return err
		}
		r.actions.addItemToPlayer(params.Item, params.PlayerID, r.state)
	case messageKindAction_MovePlayer:
		var params _MovePlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return err
		}
		r.actions.movePlayer(params.PlayerID, params.ChangeX, params.ChangeY, r.state)
	case messageKindAction_spawnZoneItems:
		var params _spawnZoneItemsParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return err
		}
		r.actions.spawnZoneItemsParams(params.Items, r.state)
	default:
		return errors.New("unknown message kind")
	}

	return nil
}
