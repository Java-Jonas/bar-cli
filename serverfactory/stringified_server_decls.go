// this file was generated by https://github.com/Java-Jonas/decltostring

package serverfactory

const gets_generated_go_import string = `import (
	"errors"
	"log"
	"net/http"
)`

const messageKindAction_MovePlayer_type string = `const (
	messageKindAction_MovePlayer	messageKind	= iota + 1
	messageKindAction_addItemToPlayer
	messageKindAction_spawnZoneItems
)`

const __MovePlayerParams_type string = `type _MovePlayerParams struct {
	ChangeX		float64		` + "`" + `json:"changeX"` + "`" + `
	ChangeY		float64		` + "`" + `json:"changeY"` + "`" + `
	PlayerID	PlayerID	` + "`" + `json:"playerID"` + "`" + `
}`

const __addItemToPlayerParams_type string = `type _addItemToPlayerParams struct {
	Item		Item		` + "`" + `json:"item"` + "`" + `
	PlayerID	PlayerID	` + "`" + `json:"playerID"` + "`" + `
}`

const __spawnZoneItemsParams_type string = `type _spawnZoneItemsParams struct {
	Items []Item ` + "`" + `json:"items"` + "`" + `
}`

const actions_type string = `type actions struct {
	MovePlayer	func(PlayerID, float64, float64, *Engine)
	addItemToPlayer	func(Item, PlayerID, *Engine)
	spawnZoneItems	func([]Item, *Engine)
}`

const processClientMessage_Room_func string = `func (r *Room) processClientMessage(msg message) error {
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
		r.actions.MovePlayer(params.PlayerID, params.ChangeX, params.ChangeY, r.state)
	case messageKindAction_spawnZoneItems:
		var params _spawnZoneItemsParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			return err
		}
		r.actions.spawnZoneItems(params.Items, r.state)
	default:
		return errors.New("unknown message kind")
	}
	return nil
}`

const _Start_func string = `func Start(movePlayer func(PlayerID, float64, float64, *Engine), addItemToPlayer func(Item, PlayerID, *Engine), spawnZoneItemsParams func([]Item, *Engine), onDeploy func(*Engine), onFrameTick func(*Engine)) {
	log.Println("Hello World")
	a := actions{movePlayer, addItemToPlayer, spawnZoneItemsParams}
	setupRoutes(a, onDeploy, onFrameTick)
	log.Fatal(http.ListenAndServe(":8080", nil))
}`
