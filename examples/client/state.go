package state

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"sync"
)

func (_equipmentSet EquipmentSet) AddEquipment(itemID ItemID) {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	if equipmentSet.equipmentSet.OperationKind == OperationKindDelete {
		return
	}
	if equipmentSet.equipmentSet.engine.Item(itemID).item.OperationKind == OperationKindDelete {
		return
	}
	for _, currentRefID := range equipmentSet.equipmentSet.Equipment {
		currentRef := equipmentSet.equipmentSet.engine.equipmentSetEquipmentRef(currentRefID)

		if currentRef.equipmentSetEquipmentRef.ReferencedElementID == itemID {
			return
		}
	}

	ref := equipmentSet.equipmentSet.engine.createEquipmentSetEquipmentRef(equipmentSet.equipmentSet.Path, equipmentSet_equipmentIdentifier, itemID, equipmentSet.equipmentSet.ID)
	equipmentSet.equipmentSet.Equipment = append(equipmentSet.equipmentSet.Equipment, ref.ID)
	equipmentSet.equipmentSet.OperationKind = OperationKindUpdate
	equipmentSet.equipmentSet.Meta.sign(equipmentSet.equipmentSet.engine.broadcastingClientID)
	equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID] = equipmentSet.equipmentSet

}
func (_player Player) AddAction() AttackEvent {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return AttackEvent{attackEvent: attackEventCore{
			OperationKind: OperationKindDelete,
			engine:        player.player.engine,
		}}
	}

	attackEvent := player.player.engine.createAttackEvent(player.player.Path, player_actionIdentifier)

	player.player.Action = append(player.player.Action, attackEvent.attackEvent.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.Meta.sign(player.player.engine.broadcastingClientID)
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return attackEvent
}
func (_player Player) AddEquipmentSet(equipmentSetID EquipmentSetID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return
	}
	if player.player.engine.EquipmentSet(equipmentSetID).equipmentSet.OperationKind == OperationKindDelete {
		return
	}
	for _, currentRefID := range player.player.EquipmentSets {
		currentRef := player.player.engine.playerEquipmentSetRef(currentRefID)

		if currentRef.playerEquipmentSetRef.ReferencedElementID == equipmentSetID {
			return
		}
	}

	ref := player.player.engine.createPlayerEquipmentSetRef(player.player.Path, player_equipmentSetsIdentifier, equipmentSetID, player.player.ID)
	player.player.EquipmentSets = append(player.player.EquipmentSets, ref.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.Meta.sign(player.player.engine.broadcastingClientID)
	player.player.engine.Patch.Player[player.player.ID] = player.player

}
func (_player Player) AddGuildMember(playerID PlayerID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return
	}
	if player.player.engine.Player(playerID).player.OperationKind == OperationKindDelete {
		return
	}
	for _, currentRefID := range player.player.GuildMembers {
		currentRef := player.player.engine.playerGuildMemberRef(currentRefID)

		if currentRef.playerGuildMemberRef.ReferencedElementID == playerID {
			return
		}
	}

	ref := player.player.engine.createPlayerGuildMemberRef(player.player.Path, player_guildMembersIdentifier, playerID, player.player.ID)
	player.player.GuildMembers = append(player.player.GuildMembers, ref.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.Meta.sign(player.player.engine.broadcastingClientID)
	player.player.engine.Patch.Player[player.player.ID] = player.player

}
func (_player Player) AddItem() Item {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return Item{item: itemCore{
			OperationKind: OperationKindDelete,
			engine:        player.player.engine,
		}}
	}

	item := player.player.engine.createItem(player.player.Path, player_itemsIdentifier)

	player.player.Items = append(player.player.Items, item.item.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.Meta.sign(player.player.engine.broadcastingClientID)
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return item
}
func (_player Player) AddTargetedByPlayer(playerID PlayerID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return
	}
	if player.player.engine.Player(playerID).player.OperationKind == OperationKindDelete {
		return
	}
	for _, currentRefID := range player.player.TargetedBy {
		currentRef := player.player.engine.playerTargetedByRef(currentRefID)
		anyContainer := player.player.engine.anyOfPlayer_ZoneItem(currentRef.playerTargetedByRef.ReferencedElementID)
		if PlayerID(anyContainer.anyOfPlayer_ZoneItem.ChildID) == playerID {
			return
		}
	}

	anyContainer := player.player.engine.createAnyOfPlayer_ZoneItem(int(player.player.ID), int(playerID), ElementKindPlayer, player.player.Path, player_targetedByIdentifier).anyOfPlayer_ZoneItem

	ref := player.player.engine.createPlayerTargetedByRef(player.player.Path, player_targetedByIdentifier, anyContainer.ID, player.player.ID, ElementKindPlayer, int(playerID))
	player.player.TargetedBy = append(player.player.TargetedBy, ref.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.Meta.sign(player.player.engine.broadcastingClientID)
	player.player.engine.Patch.Player[player.player.ID] = player.player

}
func (_player Player) AddTargetedByZoneItem(zoneItemID ZoneItemID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return
	}
	if player.player.engine.ZoneItem(zoneItemID).zoneItem.OperationKind == OperationKindDelete {
		return
	}
	for _, currentRefID := range player.player.TargetedBy {
		currentRef := player.player.engine.playerTargetedByRef(currentRefID)
		anyContainer := player.player.engine.anyOfPlayer_ZoneItem(currentRef.playerTargetedByRef.ReferencedElementID)
		if ZoneItemID(anyContainer.anyOfPlayer_ZoneItem.ChildID) == zoneItemID {
			return
		}
	}

	anyContainer := player.player.engine.createAnyOfPlayer_ZoneItem(int(player.player.ID), int(zoneItemID), ElementKindZoneItem, player.player.Path, player_targetedByIdentifier).anyOfPlayer_ZoneItem

	ref := player.player.engine.createPlayerTargetedByRef(player.player.Path, player_targetedByIdentifier, anyContainer.ID, player.player.ID, ElementKindZoneItem, int(zoneItemID))
	player.player.TargetedBy = append(player.player.TargetedBy, ref.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.Meta.sign(player.player.engine.broadcastingClientID)
	player.player.engine.Patch.Player[player.player.ID] = player.player

}
func (_zone Zone) AddInteractableItem() Item {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return Item{item: itemCore{
			OperationKind: OperationKindDelete,
			engine:        zone.zone.engine,
		}}
	}

	item := zone.zone.engine.createItem(zone.zone.Path, zone_interactablesIdentifier)
	anyContainer := zone.zone.engine.createAnyOfItem_Player_ZoneItem(int(zone.zone.ID), int(item.item.ID), ElementKindItem, zone.zone.Path, zone_interactablesIdentifier).anyOfItem_Player_ZoneItem

	zone.zone.Interactables = append(zone.zone.Interactables, anyContainer.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.Meta.sign(zone.zone.engine.broadcastingClientID)
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return item
}
func (_zone Zone) AddInteractablePlayer() Player {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return Player{player: playerCore{
			OperationKind: OperationKindDelete,
			engine:        zone.zone.engine,
		}}
	}

	player := zone.zone.engine.createPlayer(zone.zone.Path, zone_interactablesIdentifier)
	anyContainer := zone.zone.engine.createAnyOfItem_Player_ZoneItem(int(zone.zone.ID), int(player.player.ID), ElementKindPlayer, zone.zone.Path, zone_interactablesIdentifier).anyOfItem_Player_ZoneItem

	zone.zone.Interactables = append(zone.zone.Interactables, anyContainer.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.Meta.sign(zone.zone.engine.broadcastingClientID)
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return player
}
func (_zone Zone) AddInteractableZoneItem() ZoneItem {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return ZoneItem{zoneItem: zoneItemCore{
			OperationKind: OperationKindDelete,
			engine:        zone.zone.engine,
		}}
	}

	zoneItem := zone.zone.engine.createZoneItem(zone.zone.Path, zone_interactablesIdentifier)
	anyContainer := zone.zone.engine.createAnyOfItem_Player_ZoneItem(int(zone.zone.ID), int(zoneItem.zoneItem.ID), ElementKindZoneItem, zone.zone.Path, zone_interactablesIdentifier).anyOfItem_Player_ZoneItem

	zone.zone.Interactables = append(zone.zone.Interactables, anyContainer.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.Meta.sign(zone.zone.engine.broadcastingClientID)
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return zoneItem
}
func (_zone Zone) AddItem() ZoneItem {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return ZoneItem{zoneItem: zoneItemCore{
			OperationKind: OperationKindDelete,
			engine:        zone.zone.engine,
		}}
	}

	zoneItem := zone.zone.engine.createZoneItem(zone.zone.Path, zone_itemsIdentifier)

	zone.zone.Items = append(zone.zone.Items, zoneItem.zoneItem.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.Meta.sign(zone.zone.engine.broadcastingClientID)
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return zoneItem
}
func (_zone Zone) AddPlayer() Player {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return Player{player: playerCore{
			OperationKind: OperationKindDelete,
			engine:        zone.zone.engine,
		}}
	}

	player := zone.zone.engine.createPlayer(zone.zone.Path, zone_playersIdentifier)

	zone.zone.Players = append(zone.zone.Players, player.player.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.Meta.sign(zone.zone.engine.broadcastingClientID)
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return player
}
func (_zone Zone) AddTag(tag string) {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return
	}

	tagValue := zone.zone.engine.createStringValue(zone.zone.Path, zone_tagsIdentifier, tag)

	zone.zone.Tags = append(zone.zone.Tags, tagValue.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.Meta.sign(zone.zone.engine.broadcastingClientID)
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone

}
func (_any AnyOfPlayer_Position) Kind() ElementKind {
	any := _any.anyOfPlayer_Position.engine.anyOfPlayer_Position(_any.anyOfPlayer_Position.ID)
	return any.anyOfPlayer_Position.ElementKind
}
func (_any AnyOfPlayer_Position) BePlayer() Player {
	any := _any.anyOfPlayer_Position.engine.anyOfPlayer_Position(_any.anyOfPlayer_Position.ID)
	if any.anyOfPlayer_Position.ElementKind == ElementKindPlayer || any.anyOfPlayer_Position.OperationKind == OperationKindDelete {
		return any.Player()
	}
	player := any.anyOfPlayer_Position.engine.createPlayer(any.anyOfPlayer_Position.ParentElementPath, any.anyOfPlayer_Position.FieldIdentifier)
	any.anyOfPlayer_Position.bePlayer(player.ID(), true)
	return player
}
func (_any anyOfPlayer_PositionCore) bePlayer(playerID PlayerID, deleteCurrentChild bool) {
	any := _any.engine.anyOfPlayer_Position(_any.ID).anyOfPlayer_Position
	any.engine.deleteAnyOfPlayer_Position(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfPlayer_Position(any.ID.ParentID, int(playerID), ElementKindPlayer, any.ParentElementPath, any.FieldIdentifier).anyOfPlayer_Position
	switch any.FieldIdentifier {

	case item_originIdentifier:
		item := any.engine.Item(ItemID(any.ID.ParentID)).item
		item.Origin = any.ID
		item.Meta.sign(item.engine.broadcastingClientID)
		item.engine.Patch.Item[item.ID] = item

	}
	any.Meta.sign(any.engine.broadcastingClientID)
	any.engine.Patch.AnyOfPlayer_Position[any.ID] = any
}
func (_any AnyOfPlayer_Position) BePosition() Position {
	any := _any.anyOfPlayer_Position.engine.anyOfPlayer_Position(_any.anyOfPlayer_Position.ID)
	if any.anyOfPlayer_Position.ElementKind == ElementKindPosition || any.anyOfPlayer_Position.OperationKind == OperationKindDelete {
		return any.Position()
	}
	position := any.anyOfPlayer_Position.engine.createPosition(any.anyOfPlayer_Position.ParentElementPath, any.anyOfPlayer_Position.FieldIdentifier)
	any.anyOfPlayer_Position.bePosition(position.ID(), true)
	return position
}
func (_any anyOfPlayer_PositionCore) bePosition(positionID PositionID, deleteCurrentChild bool) {
	any := _any.engine.anyOfPlayer_Position(_any.ID).anyOfPlayer_Position
	any.engine.deleteAnyOfPlayer_Position(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfPlayer_Position(any.ID.ParentID, int(positionID), ElementKindPosition, any.ParentElementPath, any.FieldIdentifier).anyOfPlayer_Position
	switch any.FieldIdentifier {

	case item_originIdentifier:
		item := any.engine.Item(ItemID(any.ID.ParentID)).item
		item.Origin = any.ID
		item.Meta.sign(item.engine.broadcastingClientID)
		item.engine.Patch.Item[item.ID] = item

	}
	any.Meta.sign(any.engine.broadcastingClientID)
	any.engine.Patch.AnyOfPlayer_Position[any.ID] = any
}
func (_any anyOfPlayer_PositionCore) deleteChild() {
	any := _any.engine.anyOfPlayer_Position(_any.ID).anyOfPlayer_Position
	switch any.ElementKind {
	case ElementKindPlayer:
		any.engine.deletePlayer(PlayerID(any.ChildID))
	case ElementKindPosition:
		any.engine.deletePosition(PositionID(any.ChildID))

	}
}
func (_any AnyOfPlayer_ZoneItem) Kind() ElementKind {
	any := _any.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_any.anyOfPlayer_ZoneItem.ID)
	return any.anyOfPlayer_ZoneItem.ElementKind
}
func (_any AnyOfPlayer_ZoneItem) BePlayer() Player {
	any := _any.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_any.anyOfPlayer_ZoneItem.ID)
	if any.anyOfPlayer_ZoneItem.ElementKind == ElementKindPlayer || any.anyOfPlayer_ZoneItem.OperationKind == OperationKindDelete {
		return any.Player()
	}
	player := any.anyOfPlayer_ZoneItem.engine.createPlayer(any.anyOfPlayer_ZoneItem.ParentElementPath, any.anyOfPlayer_ZoneItem.FieldIdentifier)
	any.anyOfPlayer_ZoneItem.bePlayer(player.ID(), true)
	return player
}
func (_any anyOfPlayer_ZoneItemCore) bePlayer(playerID PlayerID, deleteCurrentChild bool) {
	any := _any.engine.anyOfPlayer_ZoneItem(_any.ID).anyOfPlayer_ZoneItem
	any.engine.deleteAnyOfPlayer_ZoneItem(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfPlayer_ZoneItem(any.ID.ParentID, int(playerID), ElementKindPlayer, any.ParentElementPath, any.FieldIdentifier).anyOfPlayer_ZoneItem
	switch any.FieldIdentifier {

	}
	any.Meta.sign(any.engine.broadcastingClientID)
	any.engine.Patch.AnyOfPlayer_ZoneItem[any.ID] = any
}
func (_any AnyOfPlayer_ZoneItem) BeZoneItem() ZoneItem {
	any := _any.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_any.anyOfPlayer_ZoneItem.ID)
	if any.anyOfPlayer_ZoneItem.ElementKind == ElementKindZoneItem || any.anyOfPlayer_ZoneItem.OperationKind == OperationKindDelete {
		return any.ZoneItem()
	}
	zoneItem := any.anyOfPlayer_ZoneItem.engine.createZoneItem(any.anyOfPlayer_ZoneItem.ParentElementPath, any.anyOfPlayer_ZoneItem.FieldIdentifier)
	any.anyOfPlayer_ZoneItem.beZoneItem(zoneItem.ID(), true)
	return zoneItem
}
func (_any anyOfPlayer_ZoneItemCore) beZoneItem(zoneItemID ZoneItemID, deleteCurrentChild bool) {
	any := _any.engine.anyOfPlayer_ZoneItem(_any.ID).anyOfPlayer_ZoneItem
	any.engine.deleteAnyOfPlayer_ZoneItem(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfPlayer_ZoneItem(any.ID.ParentID, int(zoneItemID), ElementKindZoneItem, any.ParentElementPath, any.FieldIdentifier).anyOfPlayer_ZoneItem
	switch any.FieldIdentifier {

	}
	any.Meta.sign(any.engine.broadcastingClientID)
	any.engine.Patch.AnyOfPlayer_ZoneItem[any.ID] = any
}
func (_any anyOfPlayer_ZoneItemCore) deleteChild() {
	any := _any.engine.anyOfPlayer_ZoneItem(_any.ID).anyOfPlayer_ZoneItem
	switch any.ElementKind {
	case ElementKindPlayer:
		any.engine.deletePlayer(PlayerID(any.ChildID))
	case ElementKindZoneItem:
		any.engine.deleteZoneItem(ZoneItemID(any.ChildID))

	}
}
func (_any AnyOfItem_Player_ZoneItem) Kind() ElementKind {
	any := _any.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_any.anyOfItem_Player_ZoneItem.ID)
	return any.anyOfItem_Player_ZoneItem.ElementKind
}
func (_any AnyOfItem_Player_ZoneItem) BeItem() Item {
	any := _any.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_any.anyOfItem_Player_ZoneItem.ID)
	if any.anyOfItem_Player_ZoneItem.ElementKind == ElementKindItem || any.anyOfItem_Player_ZoneItem.OperationKind == OperationKindDelete {
		return any.Item()
	}
	item := any.anyOfItem_Player_ZoneItem.engine.createItem(any.anyOfItem_Player_ZoneItem.ParentElementPath, any.anyOfItem_Player_ZoneItem.FieldIdentifier)
	any.anyOfItem_Player_ZoneItem.beItem(item.ID(), true)
	return item
}
func (_any anyOfItem_Player_ZoneItemCore) beItem(itemID ItemID, deleteCurrentChild bool) {
	any := _any.engine.anyOfItem_Player_ZoneItem(_any.ID).anyOfItem_Player_ZoneItem
	any.engine.deleteAnyOfItem_Player_ZoneItem(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfItem_Player_ZoneItem(any.ID.ParentID, int(itemID), ElementKindItem, any.ParentElementPath, any.FieldIdentifier).anyOfItem_Player_ZoneItem
	switch any.FieldIdentifier {

	}
	any.Meta.sign(any.engine.broadcastingClientID)
	any.engine.Patch.AnyOfItem_Player_ZoneItem[any.ID] = any
}
func (_any AnyOfItem_Player_ZoneItem) BePlayer() Player {
	any := _any.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_any.anyOfItem_Player_ZoneItem.ID)
	if any.anyOfItem_Player_ZoneItem.ElementKind == ElementKindPlayer || any.anyOfItem_Player_ZoneItem.OperationKind == OperationKindDelete {
		return any.Player()
	}
	player := any.anyOfItem_Player_ZoneItem.engine.createPlayer(any.anyOfItem_Player_ZoneItem.ParentElementPath, any.anyOfItem_Player_ZoneItem.FieldIdentifier)
	any.anyOfItem_Player_ZoneItem.bePlayer(player.ID(), true)
	return player
}
func (_any anyOfItem_Player_ZoneItemCore) bePlayer(playerID PlayerID, deleteCurrentChild bool) {
	any := _any.engine.anyOfItem_Player_ZoneItem(_any.ID).anyOfItem_Player_ZoneItem
	any.engine.deleteAnyOfItem_Player_ZoneItem(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfItem_Player_ZoneItem(any.ID.ParentID, int(playerID), ElementKindPlayer, any.ParentElementPath, any.FieldIdentifier).anyOfItem_Player_ZoneItem
	switch any.FieldIdentifier {

	}
	any.Meta.sign(any.engine.broadcastingClientID)
	any.engine.Patch.AnyOfItem_Player_ZoneItem[any.ID] = any
}
func (_any AnyOfItem_Player_ZoneItem) BeZoneItem() ZoneItem {
	any := _any.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_any.anyOfItem_Player_ZoneItem.ID)
	if any.anyOfItem_Player_ZoneItem.ElementKind == ElementKindZoneItem || any.anyOfItem_Player_ZoneItem.OperationKind == OperationKindDelete {
		return any.ZoneItem()
	}
	zoneItem := any.anyOfItem_Player_ZoneItem.engine.createZoneItem(any.anyOfItem_Player_ZoneItem.ParentElementPath, any.anyOfItem_Player_ZoneItem.FieldIdentifier)
	any.anyOfItem_Player_ZoneItem.beZoneItem(zoneItem.ID(), true)
	return zoneItem
}
func (_any anyOfItem_Player_ZoneItemCore) beZoneItem(zoneItemID ZoneItemID, deleteCurrentChild bool) {
	any := _any.engine.anyOfItem_Player_ZoneItem(_any.ID).anyOfItem_Player_ZoneItem
	any.engine.deleteAnyOfItem_Player_ZoneItem(any.ID, deleteCurrentChild)
	any = any.engine.createAnyOfItem_Player_ZoneItem(any.ID.ParentID, int(zoneItemID), ElementKindZoneItem, any.ParentElementPath, any.FieldIdentifier).anyOfItem_Player_ZoneItem
	switch any.FieldIdentifier {

	}
	any.Meta.sign(any.engine.broadcastingClientID)
	any.engine.Patch.AnyOfItem_Player_ZoneItem[any.ID] = any
}
func (_any anyOfItem_Player_ZoneItemCore) deleteChild() {
	any := _any.engine.anyOfItem_Player_ZoneItem(_any.ID).anyOfItem_Player_ZoneItem
	switch any.ElementKind {
	case ElementKindItem:
		any.engine.deleteItem(ItemID(any.ChildID))
	case ElementKindPlayer:
		any.engine.deletePlayer(PlayerID(any.ChildID))
	case ElementKindZoneItem:
		any.engine.deleteZoneItem(ZoneItemID(any.ChildID))

	}
}

type anyOfPlayer_PositionRef struct {
	anyOfPlayer_PositionWrapper AnyOfPlayer_Position
	anyOfPlayer_Position        anyOfPlayer_PositionCore
}

func (_any anyOfPlayer_PositionRef) Kind() ElementKind {
	return _any.anyOfPlayer_PositionWrapper.Kind()
}
func (_any anyOfPlayer_PositionRef) Player() Player {
	return _any.anyOfPlayer_PositionWrapper.Player()
}
func (_any anyOfPlayer_PositionRef) Position() Position {
	return _any.anyOfPlayer_PositionWrapper.Position()
}

type anyOfPlayer_ZoneItemRef struct {
	anyOfPlayer_ZoneItemWrapper AnyOfPlayer_ZoneItem
	anyOfPlayer_ZoneItem        anyOfPlayer_ZoneItemCore
}

func (_any anyOfPlayer_ZoneItemRef) Kind() ElementKind {
	return _any.anyOfPlayer_ZoneItemWrapper.Kind()
}
func (_any anyOfPlayer_ZoneItemRef) Player() Player {
	return _any.anyOfPlayer_ZoneItemWrapper.Player()
}
func (_any anyOfPlayer_ZoneItemRef) ZoneItem() ZoneItem {
	return _any.anyOfPlayer_ZoneItemWrapper.ZoneItem()
}

type anyOfItem_Player_ZoneItemRef struct {
	anyOfItem_Player_ZoneItemWrapper AnyOfItem_Player_ZoneItem
	anyOfItem_Player_ZoneItem        anyOfItem_Player_ZoneItemCore
}

func (_any anyOfItem_Player_ZoneItemRef) Kind() ElementKind {
	return _any.anyOfItem_Player_ZoneItemWrapper.Kind()
}
func (_any anyOfItem_Player_ZoneItemRef) Item() Item {
	return _any.anyOfItem_Player_ZoneItemWrapper.Item()
}
func (_any anyOfItem_Player_ZoneItemRef) Player() Player {
	return _any.anyOfItem_Player_ZoneItemWrapper.Player()
}
func (_any anyOfItem_Player_ZoneItemRef) ZoneItem() ZoneItem {
	return _any.anyOfItem_Player_ZoneItemWrapper.ZoneItem()
}

type assemblePlanner struct {
	updatedPaths          []path
	updatedReferencePaths map[ComplexID]path
	updatedElementPaths   map[int]path
	includedElements      map[int]bool
}

func newAssemblePlanner() *assemblePlanner {
	return &assemblePlanner{
		includedElements:      make(map[int]bool),
		updatedElementPaths:   make(map[int]path),
		updatedPaths:          make([]path, 0),
		updatedReferencePaths: make(map[ComplexID]path),
	}
}
func (a *assemblePlanner) clear() {
	a.updatedPaths = a.updatedPaths[:0]
	for key := range a.updatedElementPaths {
		delete(a.updatedElementPaths, key)
	}
	for key := range a.updatedReferencePaths {
		delete(a.updatedReferencePaths, key)
	}
	for key := range a.includedElements {
		delete(a.includedElements, key)
	}
}
func (ap *assemblePlanner) plan(state, patch *State) {
	for _, boolValue := range patch.BoolValue {
		ap.updatedElementPaths[int(boolValue.ID)] = boolValue.Path
	}
	for _, floatValue := range patch.FloatValue {
		ap.updatedElementPaths[int(floatValue.ID)] = floatValue.Path
	}
	for _, intValue := range patch.IntValue {
		ap.updatedElementPaths[int(intValue.ID)] = intValue.Path
	}
	for _, stringValue := range patch.StringValue {
		ap.updatedElementPaths[int(stringValue.ID)] = stringValue.Path
	}

	for _, attackEvent := range patch.AttackEvent {
		ap.updatedElementPaths[int(attackEvent.ID)] = attackEvent.Path
	}
	for _, equipmentSet := range patch.EquipmentSet {
		ap.updatedElementPaths[int(equipmentSet.ID)] = equipmentSet.Path
	}
	for _, gearScore := range patch.GearScore {
		ap.updatedElementPaths[int(gearScore.ID)] = gearScore.Path
	}
	for _, item := range patch.Item {
		ap.updatedElementPaths[int(item.ID)] = item.Path
	}
	for _, player := range patch.Player {
		ap.updatedElementPaths[int(player.ID)] = player.Path
	}
	for _, position := range patch.Position {
		ap.updatedElementPaths[int(position.ID)] = position.Path
	}
	for _, zone := range patch.Zone {
		ap.updatedElementPaths[int(zone.ID)] = zone.Path
	}
	for _, zoneItem := range patch.ZoneItem {
		ap.updatedElementPaths[int(zoneItem.ID)] = zoneItem.Path
	}

	for _, attackEventTargetRef := range patch.AttackEventTargetRef {
		ap.updatedReferencePaths[ComplexID(attackEventTargetRef.ID)] = attackEventTargetRef.Path
	}
	for _, equipmentSetEquipmentRef := range patch.EquipmentSetEquipmentRef {
		ap.updatedReferencePaths[ComplexID(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.Path
	}
	for _, itemBoundToRef := range patch.ItemBoundToRef {
		ap.updatedReferencePaths[ComplexID(itemBoundToRef.ID)] = itemBoundToRef.Path
	}
	for _, playerEquipmentSetRef := range patch.PlayerEquipmentSetRef {
		ap.updatedReferencePaths[ComplexID(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.Path
	}
	for _, playerGuildMemberRef := range patch.PlayerGuildMemberRef {
		ap.updatedReferencePaths[ComplexID(playerGuildMemberRef.ID)] = playerGuildMemberRef.Path
	}
	for _, playerTargetRef := range patch.PlayerTargetRef {
		ap.updatedReferencePaths[ComplexID(playerTargetRef.ID)] = playerTargetRef.Path
	}
	for _, playerTargetedByRef := range patch.PlayerTargetedByRef {
		ap.updatedReferencePaths[ComplexID(playerTargetedByRef.ID)] = playerTargetedByRef.Path
	}

	previousLen := 0
	for {
		for _, p := range ap.updatedElementPaths {
			for _, seg := range p {
				ap.includedElements[seg.ID] = true
			}
		}
		for _, p := range ap.updatedReferencePaths {
			for _, seg := range p {
				if seg.RefID != (ComplexID{}) {
				} else {
					ap.includedElements[seg.ID] = true
				}
			}
		}
		if previousLen == len(ap.includedElements) {
			break
		}
		previousLen = len(ap.includedElements)
		for _, attackEventTargetRef := range patch.AttackEventTargetRef {
			if _, ok := ap.includedElements[int(attackEventTargetRef.ID.ChildID)]; ok {
				ap.updatedReferencePaths[ComplexID(attackEventTargetRef.ID)] = attackEventTargetRef.Path
			}
		}
		for _, attackEventTargetRef := range state.AttackEventTargetRef {
			if _, ok := ap.updatedReferencePaths[ComplexID(attackEventTargetRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(attackEventTargetRef.ID.ChildID)]; ok {
				ap.updatedReferencePaths[ComplexID(attackEventTargetRef.ID)] = attackEventTargetRef.Path
			}
		}
		for _, equipmentSetEquipmentRef := range patch.EquipmentSetEquipmentRef {
			if _, ok := ap.includedElements[int(equipmentSetEquipmentRef.ID.ChildID)]; ok {
				ap.updatedReferencePaths[ComplexID(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.Path
			}
		}
		for _, equipmentSetEquipmentRef := range state.EquipmentSetEquipmentRef {
			if _, ok := ap.updatedReferencePaths[ComplexID(equipmentSetEquipmentRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(equipmentSetEquipmentRef.ID.ChildID)]; ok {
				ap.updatedReferencePaths[ComplexID(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.Path
			}
		}
		for _, itemBoundToRef := range patch.ItemBoundToRef {
			if _, ok := ap.includedElements[int(itemBoundToRef.ID.ChildID)]; ok {
				ap.updatedReferencePaths[ComplexID(itemBoundToRef.ID)] = itemBoundToRef.Path
			}
		}
		for _, itemBoundToRef := range state.ItemBoundToRef {
			if _, ok := ap.updatedReferencePaths[ComplexID(itemBoundToRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(itemBoundToRef.ID.ChildID)]; ok {
				ap.updatedReferencePaths[ComplexID(itemBoundToRef.ID)] = itemBoundToRef.Path
			}
		}
		for _, playerEquipmentSetRef := range patch.PlayerEquipmentSetRef {
			if _, ok := ap.includedElements[int(playerEquipmentSetRef.ID.ChildID)]; ok {
				ap.updatedReferencePaths[ComplexID(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.Path
			}
		}
		for _, playerEquipmentSetRef := range state.PlayerEquipmentSetRef {
			if _, ok := ap.updatedReferencePaths[ComplexID(playerEquipmentSetRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(playerEquipmentSetRef.ID.ChildID)]; ok {
				ap.updatedReferencePaths[ComplexID(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.Path
			}
		}
		for _, playerGuildMemberRef := range patch.PlayerGuildMemberRef {
			if _, ok := ap.includedElements[int(playerGuildMemberRef.ID.ChildID)]; ok {
				ap.updatedReferencePaths[ComplexID(playerGuildMemberRef.ID)] = playerGuildMemberRef.Path
			}
		}
		for _, playerGuildMemberRef := range state.PlayerGuildMemberRef {
			if _, ok := ap.updatedReferencePaths[ComplexID(playerGuildMemberRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(playerGuildMemberRef.ID.ChildID)]; ok {
				ap.updatedReferencePaths[ComplexID(playerGuildMemberRef.ID)] = playerGuildMemberRef.Path
			}
		}
		for _, playerTargetRef := range patch.PlayerTargetRef {
			if _, ok := ap.includedElements[int(playerTargetRef.ID.ChildID)]; ok {
				ap.updatedReferencePaths[ComplexID(playerTargetRef.ID)] = playerTargetRef.Path
			}
		}
		for _, playerTargetRef := range state.PlayerTargetRef {
			if _, ok := ap.updatedReferencePaths[ComplexID(playerTargetRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(playerTargetRef.ID.ChildID)]; ok {
				ap.updatedReferencePaths[ComplexID(playerTargetRef.ID)] = playerTargetRef.Path
			}
		}
		for _, playerTargetedByRef := range patch.PlayerTargetedByRef {
			if _, ok := ap.includedElements[int(playerTargetedByRef.ID.ChildID)]; ok {
				ap.updatedReferencePaths[ComplexID(playerTargetedByRef.ID)] = playerTargetedByRef.Path
			}
		}
		for _, playerTargetedByRef := range state.PlayerTargetedByRef {
			if _, ok := ap.updatedReferencePaths[ComplexID(playerTargetedByRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(playerTargetedByRef.ID.ChildID)]; ok {
				ap.updatedReferencePaths[ComplexID(playerTargetedByRef.ID)] = playerTargetedByRef.Path
			}
		}

	}
	for _, p := range ap.updatedElementPaths {
		ap.updatedPaths = append(ap.updatedPaths, p)
	}
	for _, p := range ap.updatedReferencePaths {
		ap.updatedPaths = append(ap.updatedPaths, p)
	}
}
func (ap *assemblePlanner) fill(state *State) {
	for _, boolValue := range state.BoolValue {
		ap.updatedElementPaths[int(boolValue.ID)] = boolValue.Path
	}
	for _, floatValue := range state.FloatValue {
		ap.updatedElementPaths[int(floatValue.ID)] = floatValue.Path
	}
	for _, intValue := range state.IntValue {
		ap.updatedElementPaths[int(intValue.ID)] = intValue.Path
	}
	for _, stringValue := range state.StringValue {
		ap.updatedElementPaths[int(stringValue.ID)] = stringValue.Path
	}

	for _, attackEvent := range state.AttackEvent {
		ap.updatedElementPaths[int(attackEvent.ID)] = attackEvent.Path
	}
	for _, equipmentSet := range state.EquipmentSet {
		ap.updatedElementPaths[int(equipmentSet.ID)] = equipmentSet.Path
	}
	for _, gearScore := range state.GearScore {
		ap.updatedElementPaths[int(gearScore.ID)] = gearScore.Path
	}
	for _, item := range state.Item {
		ap.updatedElementPaths[int(item.ID)] = item.Path
	}
	for _, player := range state.Player {
		ap.updatedElementPaths[int(player.ID)] = player.Path
	}
	for _, position := range state.Position {
		ap.updatedElementPaths[int(position.ID)] = position.Path
	}
	for _, zone := range state.Zone {
		ap.updatedElementPaths[int(zone.ID)] = zone.Path
	}
	for _, zoneItem := range state.ZoneItem {
		ap.updatedElementPaths[int(zoneItem.ID)] = zoneItem.Path
	}

	for _, attackEventTargetRef := range state.AttackEventTargetRef {
		ap.updatedReferencePaths[ComplexID(attackEventTargetRef.ID)] = attackEventTargetRef.Path
	}
	for _, equipmentSetEquipmentRef := range state.EquipmentSetEquipmentRef {
		ap.updatedReferencePaths[ComplexID(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.Path
	}
	for _, itemBoundToRef := range state.ItemBoundToRef {
		ap.updatedReferencePaths[ComplexID(itemBoundToRef.ID)] = itemBoundToRef.Path
	}
	for _, playerEquipmentSetRef := range state.PlayerEquipmentSetRef {
		ap.updatedReferencePaths[ComplexID(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.Path
	}
	for _, playerGuildMemberRef := range state.PlayerGuildMemberRef {
		ap.updatedReferencePaths[ComplexID(playerGuildMemberRef.ID)] = playerGuildMemberRef.Path
	}
	for _, playerTargetRef := range state.PlayerTargetRef {
		ap.updatedReferencePaths[ComplexID(playerTargetRef.ID)] = playerTargetRef.Path
	}
	for _, playerTargetedByRef := range state.PlayerTargetedByRef {
		ap.updatedReferencePaths[ComplexID(playerTargetedByRef.ID)] = playerTargetedByRef.Path
	}

	for _, p := range ap.updatedElementPaths {
		ap.updatedPaths = append(ap.updatedPaths, p)
	}
	for _, p := range ap.updatedReferencePaths {
		ap.updatedPaths = append(ap.updatedPaths, p)
	}
}
func (engine *Engine) assembleAttackEventPath(element *attackEvent, p path, pIndex int, includedElements map[int]bool) {
	attackEventData, ok := engine.Patch.AttackEvent[element.ID]
	if !ok {
		attackEventData = engine.State.AttackEvent[element.ID]
	}
	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = attackEventData.OperationKind
	}
	if pIndex+1 == len(p) {
		return
	}

	nextSeg := p[pIndex+1]

	switch nextSeg.Identifier {
	case attackEvent_targetIdentifier:
		ref := engine.attackEventTargetRef(AttackEventTargetRefID(nextSeg.RefID)).attackEventTargetRef
		if element.Target != nil && ref.OperationKind == OperationKindDelete {
			break
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.Player(ref.ReferencedElementID).player
		treeRef := elementReference{OperationKind: ref.OperationKind, ElementID: int(ref.ReferencedElementID), ElementKind: ElementKindPlayer, ReferencedDataStatus: referencedDataStatus, ElementPath: referencedElement.JSONPath}
		element.Target = &treeRef

	}
	_ = attackEventData
}
func (engine *Engine) assembleEquipmentSetPath(element *equipmentSet, p path, pIndex int, includedElements map[int]bool) {
	equipmentSetData, ok := engine.Patch.EquipmentSet[element.ID]
	if !ok {
		equipmentSetData = engine.State.EquipmentSet[element.ID]
	}
	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = equipmentSetData.OperationKind
	}
	if pIndex+1 == len(p) {
		return
	}

	nextSeg := p[pIndex+1]

	switch nextSeg.Identifier {
	case equipmentSet_equipmentIdentifier:
		ref := engine.equipmentSetEquipmentRef(EquipmentSetEquipmentRefID(nextSeg.RefID)).equipmentSetEquipmentRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.Item(ref.ReferencedElementID).item
		treeRef := elementReference{OperationKind: ref.OperationKind, ElementID: int(ref.ReferencedElementID), ElementKind: ElementKindItem, ReferencedDataStatus: referencedDataStatus, ElementPath: referencedElement.JSONPath}
		if element.Equipment == nil {
			element.Equipment = make(map[ItemID]elementReference)
		}
		element.Equipment[referencedElement.ID] = treeRef
	case equipmentSet_nameIdentifier:
		child := engine.stringValue(equipmentSetData.Name)
		element.OperationKind = child.OperationKind
		element.Name = &child.Value

	}
	_ = equipmentSetData
}
func (engine *Engine) assembleGearScorePath(element *gearScore, p path, pIndex int, includedElements map[int]bool) {
	gearScoreData, ok := engine.Patch.GearScore[element.ID]
	if !ok {
		gearScoreData = engine.State.GearScore[element.ID]
	}
	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = gearScoreData.OperationKind
	}
	if pIndex+1 == len(p) {
		return
	}

	nextSeg := p[pIndex+1]

	switch nextSeg.Identifier {
	case gearScore_levelIdentifier:
		child := engine.intValue(gearScoreData.Level)
		element.OperationKind = child.OperationKind
		element.Level = &child.Value
	case gearScore_scoreIdentifier:
		child := engine.intValue(gearScoreData.Score)
		element.OperationKind = child.OperationKind
		element.Score = &child.Value

	}
	_ = gearScoreData
}
func (engine *Engine) assembleItemPath(element *item, p path, pIndex int, includedElements map[int]bool) {
	itemData, ok := engine.Patch.Item[element.ID]
	if !ok {
		itemData = engine.State.Item[element.ID]
	}
	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = itemData.OperationKind
	}
	if pIndex+1 == len(p) {
		return
	}

	nextSeg := p[pIndex+1]

	switch nextSeg.Identifier {
	case item_boundToIdentifier:
		ref := engine.itemBoundToRef(ItemBoundToRefID(nextSeg.RefID)).itemBoundToRef
		if element.BoundTo != nil && ref.OperationKind == OperationKindDelete {
			break
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.Player(ref.ReferencedElementID).player
		treeRef := elementReference{OperationKind: ref.OperationKind, ElementID: int(ref.ReferencedElementID), ElementKind: ElementKindPlayer, ReferencedDataStatus: referencedDataStatus, ElementPath: referencedElement.JSONPath}
		element.BoundTo = &treeRef
	case item_gearScoreIdentifier:
		child := element.GearScore
		if child == nil {
			child = &gearScore{ID: GearScoreID(nextSeg.ID)}
		}
		engine.assembleGearScorePath(child, p, pIndex+1, includedElements)
		element.GearScore = child
	case item_nameIdentifier:
		child := engine.stringValue(itemData.Name)
		element.OperationKind = child.OperationKind
		element.Name = &child.Value
	case item_originIdentifier:
		switch nextSeg.Kind {
		case ElementKindPlayer:
			child, ok := element.Origin.(*player)
			if !ok || child == nil {
				child = &player{ID: PlayerID(nextSeg.ID)}
			}
			engine.assemblePlayerPath(child, p, pIndex+1, includedElements)
			if child.OperationKind == OperationKindDelete && element.Origin != nil {
				break
			}
			element.Origin = child
		case ElementKindPosition:
			child, ok := element.Origin.(*position)
			if !ok || child == nil {
				child = &position{ID: PositionID(nextSeg.ID)}
			}
			engine.assemblePositionPath(child, p, pIndex+1, includedElements)
			if child.OperationKind == OperationKindDelete && element.Origin != nil {
				break
			}
			element.Origin = child

		}

	}
	_ = itemData
}
func (engine *Engine) assemblePlayerPath(element *player, p path, pIndex int, includedElements map[int]bool) {
	playerData, ok := engine.Patch.Player[element.ID]
	if !ok {
		playerData = engine.State.Player[element.ID]
	}
	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = playerData.OperationKind
	}
	if pIndex+1 == len(p) {
		return
	}

	nextSeg := p[pIndex+1]

	switch nextSeg.Identifier {
	case player_actionIdentifier:
		if element.Action == nil {
			element.Action = make(map[AttackEventID]attackEvent)
		}
		child, ok := element.Action[AttackEventID(nextSeg.ID)]
		if !ok {
			child = attackEvent{ID: AttackEventID(nextSeg.ID)}
		}
		engine.assembleAttackEventPath(&child, p, pIndex+1, includedElements)
		element.Action[child.ID] = child
	case player_equipmentSetsIdentifier:
		ref := engine.playerEquipmentSetRef(PlayerEquipmentSetRefID(nextSeg.RefID)).playerEquipmentSetRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.EquipmentSet(ref.ReferencedElementID).equipmentSet
		treeRef := elementReference{OperationKind: ref.OperationKind, ElementID: int(ref.ReferencedElementID), ElementKind: ElementKindEquipmentSet, ReferencedDataStatus: referencedDataStatus, ElementPath: referencedElement.JSONPath}
		if element.EquipmentSets == nil {
			element.EquipmentSets = make(map[EquipmentSetID]elementReference)
		}
		element.EquipmentSets[referencedElement.ID] = treeRef
	case player_gearScoreIdentifier:
		child := element.GearScore
		if child == nil {
			child = &gearScore{ID: GearScoreID(nextSeg.ID)}
		}
		engine.assembleGearScorePath(child, p, pIndex+1, includedElements)
		element.GearScore = child
	case player_guildMembersIdentifier:
		ref := engine.playerGuildMemberRef(PlayerGuildMemberRefID(nextSeg.RefID)).playerGuildMemberRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.Player(ref.ReferencedElementID).player
		treeRef := elementReference{OperationKind: ref.OperationKind, ElementID: int(ref.ReferencedElementID), ElementKind: ElementKindPlayer, ReferencedDataStatus: referencedDataStatus, ElementPath: referencedElement.JSONPath}
		if element.GuildMembers == nil {
			element.GuildMembers = make(map[PlayerID]elementReference)
		}
		element.GuildMembers[referencedElement.ID] = treeRef
	case player_itemsIdentifier:
		if element.Items == nil {
			element.Items = make(map[ItemID]item)
		}
		child, ok := element.Items[ItemID(nextSeg.ID)]
		if !ok {
			child = item{ID: ItemID(nextSeg.ID)}
		}
		engine.assembleItemPath(&child, p, pIndex+1, includedElements)
		element.Items[child.ID] = child
	case player_positionIdentifier:
		child := element.Position
		if child == nil {
			child = &position{ID: PositionID(nextSeg.ID)}
		}
		engine.assemblePositionPath(child, p, pIndex+1, includedElements)
		element.Position = child
	case player_targetIdentifier:
		ref := engine.playerTargetRef(PlayerTargetRefID(nextSeg.RefID)).playerTargetRef
		if element.Target != nil && ref.OperationKind == OperationKindDelete {
			break
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[ref.ReferencedElementID.ChildID]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		switch nextSeg.Kind {
		case ElementKindPlayer:
			referencedElement := engine.Player(PlayerID(ref.ReferencedElementID.ChildID)).player
			treeRef := elementReference{OperationKind: ref.OperationKind, ElementID: ref.ReferencedElementID.ChildID, ElementKind: ElementKindPlayer, ReferencedDataStatus: referencedDataStatus, ElementPath: referencedElement.JSONPath}
			element.Target = &treeRef
		case ElementKindZoneItem:
			referencedElement := engine.ZoneItem(ZoneItemID(ref.ReferencedElementID.ChildID)).zoneItem
			treeRef := elementReference{OperationKind: ref.OperationKind, ElementID: ref.ReferencedElementID.ChildID, ElementKind: ElementKindZoneItem, ReferencedDataStatus: referencedDataStatus, ElementPath: referencedElement.JSONPath}
			element.Target = &treeRef

		}
	case player_targetedByIdentifier:
		if element.TargetedBy == nil {
			element.TargetedBy = make(map[int]elementReference)
		}
		ref := engine.playerTargetedByRef(PlayerTargetedByRefID(nextSeg.RefID)).playerTargetedByRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[ref.ReferencedElementID.ChildID]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		switch nextSeg.Kind {
		case ElementKindPlayer:
			referencedElement := engine.Player(PlayerID(ref.ReferencedElementID.ChildID)).player
			treeRef := elementReference{OperationKind: ref.OperationKind, ElementID: ref.ReferencedElementID.ChildID, ElementKind: ElementKindPlayer, ReferencedDataStatus: referencedDataStatus, ElementPath: referencedElement.JSONPath}
			element.TargetedBy[ref.ReferencedElementID.ChildID] = treeRef
		case ElementKindZoneItem:
			referencedElement := engine.ZoneItem(ZoneItemID(ref.ReferencedElementID.ChildID)).zoneItem
			treeRef := elementReference{OperationKind: ref.OperationKind, ElementID: ref.ReferencedElementID.ChildID, ElementKind: ElementKindZoneItem, ReferencedDataStatus: referencedDataStatus, ElementPath: referencedElement.JSONPath}
			element.TargetedBy[ref.ReferencedElementID.ChildID] = treeRef

		}

	}
	_ = playerData
}
func (engine *Engine) assemblePositionPath(element *position, p path, pIndex int, includedElements map[int]bool) {
	positionData, ok := engine.Patch.Position[element.ID]
	if !ok {
		positionData = engine.State.Position[element.ID]
	}
	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = positionData.OperationKind
	}
	if pIndex+1 == len(p) {
		return
	}

	nextSeg := p[pIndex+1]

	switch nextSeg.Identifier {
	case position_xIdentifier:
		child := engine.floatValue(positionData.X)
		element.OperationKind = child.OperationKind
		element.X = &child.Value
	case position_yIdentifier:
		child := engine.floatValue(positionData.Y)
		element.OperationKind = child.OperationKind
		element.Y = &child.Value

	}
	_ = positionData
}
func (engine *Engine) assembleZonePath(element *zone, p path, pIndex int, includedElements map[int]bool) {
	zoneData, ok := engine.Patch.Zone[element.ID]
	if !ok {
		zoneData = engine.State.Zone[element.ID]
	}
	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = zoneData.OperationKind
	}
	if pIndex+1 == len(p) {
		return
	}

	nextSeg := p[pIndex+1]

	switch nextSeg.Identifier {
	case zone_interactablesIdentifier:
		if element.Interactables == nil {
			element.Interactables = make(map[int]interface{})
		}
		switch nextSeg.Kind {
		case ElementKindItem:
			child, ok := element.Interactables[nextSeg.ID].(item)
			if !ok {
				child = item{ID: ItemID(nextSeg.ID)}
			}
			engine.assembleItemPath(&child, p, pIndex+1, includedElements)
			element.Interactables[nextSeg.ID] = child
		case ElementKindPlayer:
			child, ok := element.Interactables[nextSeg.ID].(player)
			if !ok {
				child = player{ID: PlayerID(nextSeg.ID)}
			}
			engine.assemblePlayerPath(&child, p, pIndex+1, includedElements)
			element.Interactables[nextSeg.ID] = child
		case ElementKindZoneItem:
			child, ok := element.Interactables[nextSeg.ID].(zoneItem)
			if !ok {
				child = zoneItem{ID: ZoneItemID(nextSeg.ID)}
			}
			engine.assembleZoneItemPath(&child, p, pIndex+1, includedElements)
			element.Interactables[nextSeg.ID] = child

		}
	case zone_itemsIdentifier:
		if element.Items == nil {
			element.Items = make(map[ZoneItemID]zoneItem)
		}
		child, ok := element.Items[ZoneItemID(nextSeg.ID)]
		if !ok {
			child = zoneItem{ID: ZoneItemID(nextSeg.ID)}
		}
		engine.assembleZoneItemPath(&child, p, pIndex+1, includedElements)
		element.Items[child.ID] = child
	case zone_playersIdentifier:
		if element.Players == nil {
			element.Players = make(map[PlayerID]player)
		}
		child, ok := element.Players[PlayerID(nextSeg.ID)]
		if !ok {
			child = player{ID: PlayerID(nextSeg.ID)}
		}
		engine.assemblePlayerPath(&child, p, pIndex+1, includedElements)
		element.Players[child.ID] = child
	case zone_tagsIdentifier:
		if element.Tags == nil {
			element.Tags = make([]string, 0, len(zoneData.Tags))
		}
		child := engine.stringValue(StringValueID(nextSeg.ID))
		element.OperationKind = child.OperationKind
		element.Tags = append(element.Tags, child.Value)

	}
	_ = zoneData
}
func (engine *Engine) assembleZoneItemPath(element *zoneItem, p path, pIndex int, includedElements map[int]bool) {
	zoneItemData, ok := engine.Patch.ZoneItem[element.ID]
	if !ok {
		zoneItemData = engine.State.ZoneItem[element.ID]
	}
	if element.OperationKind != OperationKindUpdate && element.OperationKind != OperationKindDelete {
		element.OperationKind = zoneItemData.OperationKind
	}
	if pIndex+1 == len(p) {
		return
	}

	nextSeg := p[pIndex+1]

	switch nextSeg.Identifier {
	case zoneItem_itemIdentifier:
		child := element.Item
		if child == nil {
			child = &item{ID: ItemID(nextSeg.ID)}
		}
		engine.assembleItemPath(child, p, pIndex+1, includedElements)
		element.Item = child
	case zoneItem_positionIdentifier:
		child := element.Position
		if child == nil {
			child = &position{ID: PositionID(nextSeg.ID)}
		}
		engine.assemblePositionPath(child, p, pIndex+1, includedElements)
		element.Position = child

	}
	_ = zoneItemData
}
func (engine *Engine) assembleUpdateTree() {
	engine.planner.clear()
	engine.Tree.clear()
	engine.planner.plan(engine.State, engine.Patch)
	engine.assembleTree()
}
func (engine *Engine) assembleFullTree() {
	engine.planner.clear()
	engine.Tree.clear()
	engine.planner.fill(engine.State)
	engine.assembleTree()
}
func (engine *Engine) assembleTree() {
	for _, p := range engine.planner.updatedPaths {
		switch p[0].Identifier {
		case attackEventIdentifier:
			child, ok := engine.Tree.AttackEvent[AttackEventID(p[0].ID)]
			if !ok {
				child = attackEvent{ID: AttackEventID(p[0].ID)}
			}
			engine.assembleAttackEventPath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.AttackEvent[AttackEventID(p[0].ID)] = child
		case equipmentSetIdentifier:
			child, ok := engine.Tree.EquipmentSet[EquipmentSetID(p[0].ID)]
			if !ok {
				child = equipmentSet{ID: EquipmentSetID(p[0].ID)}
			}
			engine.assembleEquipmentSetPath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.EquipmentSet[EquipmentSetID(p[0].ID)] = child
		case gearScoreIdentifier:
			child, ok := engine.Tree.GearScore[GearScoreID(p[0].ID)]
			if !ok {
				child = gearScore{ID: GearScoreID(p[0].ID)}
			}
			engine.assembleGearScorePath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.GearScore[GearScoreID(p[0].ID)] = child
		case itemIdentifier:
			child, ok := engine.Tree.Item[ItemID(p[0].ID)]
			if !ok {
				child = item{ID: ItemID(p[0].ID)}
			}
			engine.assembleItemPath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.Item[ItemID(p[0].ID)] = child
		case playerIdentifier:
			child, ok := engine.Tree.Player[PlayerID(p[0].ID)]
			if !ok {
				child = player{ID: PlayerID(p[0].ID)}
			}
			engine.assemblePlayerPath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.Player[PlayerID(p[0].ID)] = child
		case positionIdentifier:
			child, ok := engine.Tree.Position[PositionID(p[0].ID)]
			if !ok {
				child = position{ID: PositionID(p[0].ID)}
			}
			engine.assemblePositionPath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.Position[PositionID(p[0].ID)] = child
		case zoneIdentifier:
			child, ok := engine.Tree.Zone[ZoneID(p[0].ID)]
			if !ok {
				child = zone{ID: ZoneID(p[0].ID)}
			}
			engine.assembleZonePath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.Zone[ZoneID(p[0].ID)] = child
		case zoneItemIdentifier:
			child, ok := engine.Tree.ZoneItem[ZoneItemID(p[0].ID)]
			if !ok {
				child = zoneItem{ID: ZoneItemID(p[0].ID)}
			}
			engine.assembleZoneItemPath(&child, p, 0, engine.planner.includedElements)
			engine.Tree.ZoneItem[ZoneItemID(p[0].ID)] = child

		}
	}
}
func (engine *Engine) createBoolValue(p path, fieldIdentifier treeFieldIdentifier, value bool) boolValue {
	var element boolValue
	element.Value = value
	element.engine = engine
	element.ID = BoolValueID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindBoolValue, ComplexID{})
	element.JSONPath = element.Path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	engine.Patch.BoolValue[element.ID] = element
	return element
}
func (engine *Engine) createFloatValue(p path, fieldIdentifier treeFieldIdentifier, value float64) floatValue {
	var element floatValue
	element.Value = value
	element.engine = engine
	element.ID = FloatValueID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindFloatValue, ComplexID{})
	element.JSONPath = element.Path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	engine.Patch.FloatValue[element.ID] = element
	return element
}
func (engine *Engine) createIntValue(p path, fieldIdentifier treeFieldIdentifier, value int64) intValue {
	var element intValue
	element.Value = value
	element.engine = engine
	element.ID = IntValueID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindIntValue, ComplexID{})
	element.JSONPath = element.Path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	engine.Patch.IntValue[element.ID] = element
	return element
}
func (engine *Engine) createStringValue(p path, fieldIdentifier treeFieldIdentifier, value string) stringValue {
	var element stringValue
	element.Value = value
	element.engine = engine
	element.ID = StringValueID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindStringValue, ComplexID{})
	element.JSONPath = element.Path.toJSONPath()
	element.OperationKind = OperationKindUpdate
	engine.Patch.StringValue[element.ID] = element
	return element
}
func (engine *Engine) CreateAttackEvent() AttackEvent {
	return engine.createAttackEvent(newPath(), attackEventIdentifier)
}
func (engine *Engine) createAttackEvent(p path, fieldIdentifier treeFieldIdentifier) AttackEvent {
	var element attackEventCore
	element.engine = engine
	element.ID = AttackEventID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindAttackEvent, ComplexID{})
	element.JSONPath = element.Path.toJSONPath()

	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
	engine.Patch.AttackEvent[element.ID] = element
	return AttackEvent{attackEvent: element}
}
func (engine *Engine) CreateEquipmentSet() EquipmentSet {
	return engine.createEquipmentSet(newPath(), equipmentSetIdentifier)
}
func (engine *Engine) createEquipmentSet(p path, fieldIdentifier treeFieldIdentifier) EquipmentSet {
	var element equipmentSetCore
	element.engine = engine
	element.ID = EquipmentSetID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindEquipmentSet, ComplexID{})
	element.JSONPath = element.Path.toJSONPath()

	elementName := engine.createStringValue(element.Path, equipmentSet_nameIdentifier, "")
	element.Name = elementName.ID

	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
	engine.Patch.EquipmentSet[element.ID] = element
	return EquipmentSet{equipmentSet: element}
}
func (engine *Engine) CreateGearScore() GearScore {
	return engine.createGearScore(newPath(), gearScoreIdentifier)
}
func (engine *Engine) createGearScore(p path, fieldIdentifier treeFieldIdentifier) GearScore {
	var element gearScoreCore
	element.engine = engine
	element.ID = GearScoreID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindGearScore, ComplexID{})
	element.JSONPath = element.Path.toJSONPath()

	elementLevel := engine.createIntValue(element.Path, gearScore_levelIdentifier, 0)
	element.Level = elementLevel.ID

	elementScore := engine.createIntValue(element.Path, gearScore_scoreIdentifier, 0)
	element.Score = elementScore.ID

	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
	engine.Patch.GearScore[element.ID] = element
	return GearScore{gearScore: element}
}
func (engine *Engine) CreateItem() Item {
	return engine.createItem(newPath(), itemIdentifier)
}
func (engine *Engine) createItem(p path, fieldIdentifier treeFieldIdentifier) Item {
	var element itemCore
	element.engine = engine
	element.ID = ItemID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindItem, ComplexID{})
	element.JSONPath = element.Path.toJSONPath()

	elementGearScore := engine.createGearScore(element.Path, item_gearScoreIdentifier)
	element.GearScore = elementGearScore.gearScore.ID

	elementName := engine.createStringValue(element.Path, item_nameIdentifier, "")
	element.Name = elementName.ID
	originElement := engine.createPlayer(element.Path, item_originIdentifier)
	elementOrigin := engine.createAnyOfPlayer_Position(int(element.ID), int(originElement.player.ID), ElementKindPlayer, element.Path, item_originIdentifier)
	element.Origin = elementOrigin.anyOfPlayer_Position.ID

	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
	engine.Patch.Item[element.ID] = element
	return Item{item: element}
}
func (engine *Engine) CreatePlayer() Player {
	return engine.createPlayer(newPath(), playerIdentifier)
}
func (engine *Engine) createPlayer(p path, fieldIdentifier treeFieldIdentifier) Player {
	var element playerCore
	element.engine = engine
	element.ID = PlayerID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindPlayer, ComplexID{})
	element.JSONPath = element.Path.toJSONPath()

	elementGearScore := engine.createGearScore(element.Path, player_gearScoreIdentifier)
	element.GearScore = elementGearScore.gearScore.ID

	elementPosition := engine.createPosition(element.Path, player_positionIdentifier)
	element.Position = elementPosition.position.ID

	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
	engine.Patch.Player[element.ID] = element
	return Player{player: element}
}
func (engine *Engine) CreatePosition() Position {
	return engine.createPosition(newPath(), positionIdentifier)
}
func (engine *Engine) createPosition(p path, fieldIdentifier treeFieldIdentifier) Position {
	var element positionCore
	element.engine = engine
	element.ID = PositionID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindPosition, ComplexID{})
	element.JSONPath = element.Path.toJSONPath()

	elementX := engine.createFloatValue(element.Path, position_xIdentifier, 0.0)
	element.X = elementX.ID

	elementY := engine.createFloatValue(element.Path, position_yIdentifier, 0.0)
	element.Y = elementY.ID

	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
	engine.Patch.Position[element.ID] = element
	return Position{position: element}
}
func (engine *Engine) CreateZone() Zone {
	return engine.createZone(newPath(), zoneIdentifier)
}
func (engine *Engine) createZone(p path, fieldIdentifier treeFieldIdentifier) Zone {
	var element zoneCore
	element.engine = engine
	element.ID = ZoneID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindZone, ComplexID{})
	element.JSONPath = element.Path.toJSONPath()

	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
	engine.Patch.Zone[element.ID] = element
	return Zone{zone: element}
}
func (engine *Engine) CreateZoneItem() ZoneItem {
	return engine.createZoneItem(newPath(), zoneItemIdentifier)
}
func (engine *Engine) createZoneItem(p path, fieldIdentifier treeFieldIdentifier) ZoneItem {
	var element zoneItemCore
	element.engine = engine
	element.ID = ZoneItemID(engine.GenerateID())
	element.Path = p.extendAndCopy(fieldIdentifier, int(element.ID), ElementKindZoneItem, ComplexID{})
	element.JSONPath = element.Path.toJSONPath()

	elementItem := engine.createItem(element.Path, zoneItem_itemIdentifier)
	element.Item = elementItem.item.ID

	elementPosition := engine.createPosition(element.Path, zoneItem_positionIdentifier)
	element.Position = elementPosition.position.ID

	element.OperationKind = OperationKindUpdate
	element.HasParent = len(element.Path) > 1
	engine.Patch.ZoneItem[element.ID] = element
	return ZoneItem{zoneItem: element}
}
func (engine *Engine) createAttackEventTargetRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID PlayerID, parentID AttackEventID) attackEventTargetRefCore {
	var element attackEventTargetRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = AttackEventTargetRefID{fieldIdentifier, int(parentID), int(referencedElementID), false}
	element.Path = p.extendAndCopy(fieldIdentifier, 0, ElementKindPlayer, ComplexID(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.AttackEventTargetRef[element.ID] = element
	return element
}
func (engine *Engine) createEquipmentSetEquipmentRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID ItemID, parentID EquipmentSetID) equipmentSetEquipmentRefCore {
	var element equipmentSetEquipmentRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = EquipmentSetEquipmentRefID{fieldIdentifier, int(parentID), int(referencedElementID), false}
	element.Path = p.extendAndCopy(fieldIdentifier, 0, ElementKindItem, ComplexID(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.EquipmentSetEquipmentRef[element.ID] = element
	return element
}
func (engine *Engine) createItemBoundToRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID PlayerID, parentID ItemID) itemBoundToRefCore {
	var element itemBoundToRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = ItemBoundToRefID{fieldIdentifier, int(parentID), int(referencedElementID), false}
	element.Path = p.extendAndCopy(fieldIdentifier, 0, ElementKindPlayer, ComplexID(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.ItemBoundToRef[element.ID] = element
	return element
}
func (engine *Engine) createPlayerEquipmentSetRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID EquipmentSetID, parentID PlayerID) playerEquipmentSetRefCore {
	var element playerEquipmentSetRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = PlayerEquipmentSetRefID{fieldIdentifier, int(parentID), int(referencedElementID), false}
	element.Path = p.extendAndCopy(fieldIdentifier, 0, ElementKindEquipmentSet, ComplexID(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerEquipmentSetRef[element.ID] = element
	return element
}
func (engine *Engine) createPlayerGuildMemberRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID PlayerID, parentID PlayerID) playerGuildMemberRefCore {
	var element playerGuildMemberRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = PlayerGuildMemberRefID{fieldIdentifier, int(parentID), int(referencedElementID), false}
	element.Path = p.extendAndCopy(fieldIdentifier, 0, ElementKindPlayer, ComplexID(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerGuildMemberRef[element.ID] = element
	return element
}
func (engine *Engine) createPlayerTargetRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID AnyOfPlayer_ZoneItemID, parentID PlayerID, childKind ElementKind, childID int) playerTargetRefCore {
	var element playerTargetRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = PlayerTargetRefID{referencedElementID.Field, referencedElementID.ParentID, referencedElementID.ChildID, true}
	element.Path = p.extendAndCopy(fieldIdentifier, 0, childKind, ComplexID(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerTargetRef[element.ID] = element
	return element
}
func (engine *Engine) createPlayerTargetedByRef(p path, fieldIdentifier treeFieldIdentifier, referencedElementID AnyOfPlayer_ZoneItemID, parentID PlayerID, childKind ElementKind, childID int) playerTargetedByRefCore {
	var element playerTargetedByRefCore
	element.engine = engine
	element.ReferencedElementID = referencedElementID
	element.ParentID = parentID
	element.ID = PlayerTargetedByRefID{referencedElementID.Field, referencedElementID.ParentID, referencedElementID.ChildID, true}
	element.Path = p.extendAndCopy(fieldIdentifier, 0, childKind, ComplexID(element.ID))
	element.OperationKind = OperationKindUpdate
	engine.Patch.PlayerTargetedByRef[element.ID] = element
	return element
}
func (engine *Engine) createAnyOfPlayer_Position(parentID int, childID int, childKind ElementKind, p path, fieldIdentifier treeFieldIdentifier) AnyOfPlayer_Position {
	var element anyOfPlayer_PositionCore
	element.engine = engine
	element.ID = AnyOfPlayer_PositionID{fieldIdentifier, parentID, childID, false}
	element.ChildID = childID
	element.ElementKind = childKind
	element.OperationKind = OperationKindUpdate
	element.ParentElementPath = p
	element.FieldIdentifier = fieldIdentifier
	engine.Patch.AnyOfPlayer_Position[element.ID] = element
	return AnyOfPlayer_Position{anyOfPlayer_Position: element}
}
func (engine *Engine) createAnyOfPlayer_ZoneItem(parentID int, childID int, childKind ElementKind, p path, fieldIdentifier treeFieldIdentifier) AnyOfPlayer_ZoneItem {
	var element anyOfPlayer_ZoneItemCore
	element.engine = engine
	element.ID = AnyOfPlayer_ZoneItemID{fieldIdentifier, parentID, childID, false}
	element.ChildID = childID
	element.ElementKind = childKind
	element.OperationKind = OperationKindUpdate
	element.ParentElementPath = p
	element.FieldIdentifier = fieldIdentifier
	engine.Patch.AnyOfPlayer_ZoneItem[element.ID] = element
	return AnyOfPlayer_ZoneItem{anyOfPlayer_ZoneItem: element}
}
func (engine *Engine) createAnyOfItem_Player_ZoneItem(parentID int, childID int, childKind ElementKind, p path, fieldIdentifier treeFieldIdentifier) AnyOfItem_Player_ZoneItem {
	var element anyOfItem_Player_ZoneItemCore
	element.engine = engine
	element.ID = AnyOfItem_Player_ZoneItemID{fieldIdentifier, parentID, childID, false}
	element.ChildID = childID
	element.ElementKind = childKind
	element.OperationKind = OperationKindUpdate
	element.ParentElementPath = p
	element.FieldIdentifier = fieldIdentifier
	engine.Patch.AnyOfItem_Player_ZoneItem[element.ID] = element
	return AnyOfItem_Player_ZoneItem{anyOfItem_Player_ZoneItem: element}
}
func (engine *Engine) deleteBoolValue(boolValueID BoolValueID) {
	boolValue := engine.boolValue(boolValueID)
	if boolValue.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.BoolValue[boolValueID]; ok {
		boolValue.OperationKind = OperationKindDelete
		boolValue.Meta.sign(boolValue.engine.broadcastingClientID)
		engine.Patch.BoolValue[boolValue.ID] = boolValue
	} else {
		delete(engine.Patch.BoolValue, boolValueID)
	}
}
func (engine *Engine) deleteFloatValue(floatValueID FloatValueID) {
	floatValue := engine.floatValue(floatValueID)
	if floatValue.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.FloatValue[floatValueID]; ok {
		floatValue.OperationKind = OperationKindDelete
		floatValue.Meta.sign(floatValue.engine.broadcastingClientID)
		engine.Patch.FloatValue[floatValue.ID] = floatValue
	} else {
		delete(engine.Patch.FloatValue, floatValueID)
	}
}
func (engine *Engine) deleteIntValue(intValueID IntValueID) {
	intValue := engine.intValue(intValueID)
	if intValue.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.IntValue[intValueID]; ok {
		intValue.OperationKind = OperationKindDelete
		intValue.Meta.sign(intValue.engine.broadcastingClientID)
		engine.Patch.IntValue[intValue.ID] = intValue
	} else {
		delete(engine.Patch.IntValue, intValueID)
	}
}
func (engine *Engine) deleteStringValue(stringValueID StringValueID) {
	stringValue := engine.stringValue(stringValueID)
	if stringValue.OperationKind == OperationKindDelete {
		return
	}
	if _, ok := engine.State.StringValue[stringValueID]; ok {
		stringValue.OperationKind = OperationKindDelete
		stringValue.Meta.sign(stringValue.engine.broadcastingClientID)
		engine.Patch.StringValue[stringValue.ID] = stringValue
	} else {
		delete(engine.Patch.StringValue, stringValueID)
	}
}
func (engine *Engine) DeleteAttackEvent(attackEventID AttackEventID) {
	attackEvent := engine.AttackEvent(attackEventID).attackEvent
	if attackEvent.HasParent {
		return
	}
	engine.deleteAttackEvent(attackEventID)
}
func (engine *Engine) deleteAttackEvent(attackEventID AttackEventID) {
	attackEvent := engine.AttackEvent(attackEventID).attackEvent
	if attackEvent.OperationKind == OperationKindDelete {
		return
	}
	engine.deleteAttackEventTargetRef(attackEvent.Target)

	if _, ok := engine.State.AttackEvent[attackEventID]; ok {
		attackEvent.OperationKind = OperationKindDelete
		attackEvent.Meta.sign(attackEvent.engine.broadcastingClientID)
		engine.Patch.AttackEvent[attackEvent.ID] = attackEvent
	} else {
		delete(engine.Patch.AttackEvent, attackEventID)
	}
}
func (engine *Engine) DeleteEquipmentSet(equipmentSetID EquipmentSetID) {

	engine.deleteEquipmentSet(equipmentSetID)
}
func (engine *Engine) deleteEquipmentSet(equipmentSetID EquipmentSetID) {
	equipmentSet := engine.EquipmentSet(equipmentSetID).equipmentSet
	if equipmentSet.OperationKind == OperationKindDelete {
		return
	}
	engine.dereferencePlayerEquipmentSetRefs(equipmentSetID)

	for _, equipmentID := range equipmentSet.Equipment {
		engine.deleteEquipmentSetEquipmentRef(equipmentID)
	}
	engine.deleteStringValue(equipmentSet.Name)

	if _, ok := engine.State.EquipmentSet[equipmentSetID]; ok {
		equipmentSet.OperationKind = OperationKindDelete
		equipmentSet.Meta.sign(equipmentSet.engine.broadcastingClientID)
		engine.Patch.EquipmentSet[equipmentSet.ID] = equipmentSet
	} else {
		delete(engine.Patch.EquipmentSet, equipmentSetID)
	}
}
func (engine *Engine) DeleteGearScore(gearScoreID GearScoreID) {
	gearScore := engine.GearScore(gearScoreID).gearScore
	if gearScore.HasParent {
		return
	}
	engine.deleteGearScore(gearScoreID)
}
func (engine *Engine) deleteGearScore(gearScoreID GearScoreID) {
	gearScore := engine.GearScore(gearScoreID).gearScore
	if gearScore.OperationKind == OperationKindDelete {
		return
	}
	engine.deleteIntValue(gearScore.Level)
	engine.deleteIntValue(gearScore.Score)

	if _, ok := engine.State.GearScore[gearScoreID]; ok {
		gearScore.OperationKind = OperationKindDelete
		gearScore.Meta.sign(gearScore.engine.broadcastingClientID)
		engine.Patch.GearScore[gearScore.ID] = gearScore
	} else {
		delete(engine.Patch.GearScore, gearScoreID)
	}
}
func (engine *Engine) DeleteItem(itemID ItemID) {
	item := engine.Item(itemID).item
	if item.HasParent {
		return
	}
	engine.deleteItem(itemID)
}
func (engine *Engine) deleteItem(itemID ItemID) {
	item := engine.Item(itemID).item
	if item.OperationKind == OperationKindDelete {
		return
	}
	engine.dereferenceEquipmentSetEquipmentRefs(itemID)

	engine.deleteItemBoundToRef(item.BoundTo)
	engine.deleteGearScore(item.GearScore)
	engine.deleteStringValue(item.Name)
	engine.deleteAnyOfPlayer_Position(item.Origin, true)

	if _, ok := engine.State.Item[itemID]; ok {
		item.OperationKind = OperationKindDelete
		item.Meta.sign(item.engine.broadcastingClientID)
		engine.Patch.Item[item.ID] = item
	} else {
		delete(engine.Patch.Item, itemID)
	}
}
func (engine *Engine) DeletePlayer(playerID PlayerID) {
	player := engine.Player(playerID).player
	if player.HasParent {
		return
	}
	engine.deletePlayer(playerID)
}
func (engine *Engine) deletePlayer(playerID PlayerID) {
	player := engine.Player(playerID).player
	if player.OperationKind == OperationKindDelete {
		return
	}
	engine.dereferenceAttackEventTargetRefs(playerID)
	engine.dereferenceItemBoundToRefs(playerID)
	engine.dereferencePlayerGuildMemberRefs(playerID)
	engine.dereferencePlayerTargetRefsPlayer(playerID)
	engine.dereferencePlayerTargetedByRefsPlayer(playerID)

	for _, actionID := range player.Action {
		engine.deleteAttackEvent(actionID)
	}
	for _, equipmentSetID := range player.EquipmentSets {
		engine.deletePlayerEquipmentSetRef(equipmentSetID)
	}
	engine.deleteGearScore(player.GearScore)
	for _, guildMemberID := range player.GuildMembers {
		engine.deletePlayerGuildMemberRef(guildMemberID)
	}
	for _, itemID := range player.Items {
		engine.deleteItem(itemID)
	}
	engine.deletePosition(player.Position)
	engine.deletePlayerTargetRef(player.Target)
	for _, targetedByID := range player.TargetedBy {
		engine.deletePlayerTargetedByRef(targetedByID)
	}

	if _, ok := engine.State.Player[playerID]; ok {
		player.OperationKind = OperationKindDelete
		player.Meta.sign(player.engine.broadcastingClientID)
		engine.Patch.Player[player.ID] = player
	} else {
		delete(engine.Patch.Player, playerID)
	}
}
func (engine *Engine) DeletePosition(positionID PositionID) {
	position := engine.Position(positionID).position
	if position.HasParent {
		return
	}
	engine.deletePosition(positionID)
}
func (engine *Engine) deletePosition(positionID PositionID) {
	position := engine.Position(positionID).position
	if position.OperationKind == OperationKindDelete {
		return
	}
	engine.deleteFloatValue(position.X)
	engine.deleteFloatValue(position.Y)

	if _, ok := engine.State.Position[positionID]; ok {
		position.OperationKind = OperationKindDelete
		position.Meta.sign(position.engine.broadcastingClientID)
		engine.Patch.Position[position.ID] = position
	} else {
		delete(engine.Patch.Position, positionID)
	}
}
func (engine *Engine) DeleteZone(zoneID ZoneID) {

	engine.deleteZone(zoneID)
}
func (engine *Engine) deleteZone(zoneID ZoneID) {
	zone := engine.Zone(zoneID).zone
	if zone.OperationKind == OperationKindDelete {
		return
	}
	for _, interactableID := range zone.Interactables {
		engine.deleteAnyOfItem_Player_ZoneItem(interactableID, true)
	}
	for _, itemID := range zone.Items {
		engine.deleteZoneItem(itemID)
	}
	for _, playerID := range zone.Players {
		engine.deletePlayer(playerID)
	}
	for _, tagID := range zone.Tags {
		engine.deleteStringValue(tagID)
	}

	if _, ok := engine.State.Zone[zoneID]; ok {
		zone.OperationKind = OperationKindDelete
		zone.Meta.sign(zone.engine.broadcastingClientID)
		engine.Patch.Zone[zone.ID] = zone
	} else {
		delete(engine.Patch.Zone, zoneID)
	}
}
func (engine *Engine) DeleteZoneItem(zoneItemID ZoneItemID) {
	zoneItem := engine.ZoneItem(zoneItemID).zoneItem
	if zoneItem.HasParent {
		return
	}
	engine.deleteZoneItem(zoneItemID)
}
func (engine *Engine) deleteZoneItem(zoneItemID ZoneItemID) {
	zoneItem := engine.ZoneItem(zoneItemID).zoneItem
	if zoneItem.OperationKind == OperationKindDelete {
		return
	}
	engine.dereferencePlayerTargetRefsZoneItem(zoneItemID)
	engine.dereferencePlayerTargetedByRefsZoneItem(zoneItemID)

	engine.deleteItem(zoneItem.Item)
	engine.deletePosition(zoneItem.Position)

	if _, ok := engine.State.ZoneItem[zoneItemID]; ok {
		zoneItem.OperationKind = OperationKindDelete
		zoneItem.Meta.sign(zoneItem.engine.broadcastingClientID)
		engine.Patch.ZoneItem[zoneItem.ID] = zoneItem
	} else {
		delete(engine.Patch.ZoneItem, zoneItemID)
	}
}
func (engine *Engine) deleteAttackEventTargetRef(attackEventTargetRefID AttackEventTargetRefID) {
	attackEventTargetRef := engine.attackEventTargetRef(attackEventTargetRefID).attackEventTargetRef
	if attackEventTargetRef.OperationKind == OperationKindDelete {
		return
	}

	if _, ok := engine.State.AttackEventTargetRef[attackEventTargetRefID]; ok {
		attackEventTargetRef.OperationKind = OperationKindDelete
		attackEventTargetRef.Meta.sign(attackEventTargetRef.engine.broadcastingClientID)
		engine.Patch.AttackEventTargetRef[attackEventTargetRef.ID] = attackEventTargetRef
	} else {
		delete(engine.Patch.AttackEventTargetRef, attackEventTargetRefID)
	}
}
func (engine *Engine) deleteEquipmentSetEquipmentRef(equipmentSetEquipmentRefID EquipmentSetEquipmentRefID) {
	equipmentSetEquipmentRef := engine.equipmentSetEquipmentRef(equipmentSetEquipmentRefID).equipmentSetEquipmentRef
	if equipmentSetEquipmentRef.OperationKind == OperationKindDelete {
		return
	}

	if _, ok := engine.State.EquipmentSetEquipmentRef[equipmentSetEquipmentRefID]; ok {
		equipmentSetEquipmentRef.OperationKind = OperationKindDelete
		equipmentSetEquipmentRef.Meta.sign(equipmentSetEquipmentRef.engine.broadcastingClientID)
		engine.Patch.EquipmentSetEquipmentRef[equipmentSetEquipmentRef.ID] = equipmentSetEquipmentRef
	} else {
		delete(engine.Patch.EquipmentSetEquipmentRef, equipmentSetEquipmentRefID)
	}
}
func (engine *Engine) deleteItemBoundToRef(itemBoundToRefID ItemBoundToRefID) {
	itemBoundToRef := engine.itemBoundToRef(itemBoundToRefID).itemBoundToRef
	if itemBoundToRef.OperationKind == OperationKindDelete {
		return
	}

	if _, ok := engine.State.ItemBoundToRef[itemBoundToRefID]; ok {
		itemBoundToRef.OperationKind = OperationKindDelete
		itemBoundToRef.Meta.sign(itemBoundToRef.engine.broadcastingClientID)
		engine.Patch.ItemBoundToRef[itemBoundToRef.ID] = itemBoundToRef
	} else {
		delete(engine.Patch.ItemBoundToRef, itemBoundToRefID)
	}
}
func (engine *Engine) deletePlayerEquipmentSetRef(playerEquipmentSetRefID PlayerEquipmentSetRefID) {
	playerEquipmentSetRef := engine.playerEquipmentSetRef(playerEquipmentSetRefID).playerEquipmentSetRef
	if playerEquipmentSetRef.OperationKind == OperationKindDelete {
		return
	}

	if _, ok := engine.State.PlayerEquipmentSetRef[playerEquipmentSetRefID]; ok {
		playerEquipmentSetRef.OperationKind = OperationKindDelete
		playerEquipmentSetRef.Meta.sign(playerEquipmentSetRef.engine.broadcastingClientID)
		engine.Patch.PlayerEquipmentSetRef[playerEquipmentSetRef.ID] = playerEquipmentSetRef
	} else {
		delete(engine.Patch.PlayerEquipmentSetRef, playerEquipmentSetRefID)
	}
}
func (engine *Engine) deletePlayerGuildMemberRef(playerGuildMemberRefID PlayerGuildMemberRefID) {
	playerGuildMemberRef := engine.playerGuildMemberRef(playerGuildMemberRefID).playerGuildMemberRef
	if playerGuildMemberRef.OperationKind == OperationKindDelete {
		return
	}

	if _, ok := engine.State.PlayerGuildMemberRef[playerGuildMemberRefID]; ok {
		playerGuildMemberRef.OperationKind = OperationKindDelete
		playerGuildMemberRef.Meta.sign(playerGuildMemberRef.engine.broadcastingClientID)
		engine.Patch.PlayerGuildMemberRef[playerGuildMemberRef.ID] = playerGuildMemberRef
	} else {
		delete(engine.Patch.PlayerGuildMemberRef, playerGuildMemberRefID)
	}
}
func (engine *Engine) deletePlayerTargetRef(playerTargetRefID PlayerTargetRefID) {
	playerTargetRef := engine.playerTargetRef(playerTargetRefID).playerTargetRef
	if playerTargetRef.OperationKind == OperationKindDelete {
		return
	}
	engine.deleteAnyOfPlayer_ZoneItem(playerTargetRef.ReferencedElementID, false)
	if _, ok := engine.State.PlayerTargetRef[playerTargetRefID]; ok {
		playerTargetRef.OperationKind = OperationKindDelete
		playerTargetRef.Meta.sign(playerTargetRef.engine.broadcastingClientID)
		engine.Patch.PlayerTargetRef[playerTargetRef.ID] = playerTargetRef
	} else {
		delete(engine.Patch.PlayerTargetRef, playerTargetRefID)
	}
}
func (engine *Engine) deletePlayerTargetedByRef(playerTargetedByRefID PlayerTargetedByRefID) {
	playerTargetedByRef := engine.playerTargetedByRef(playerTargetedByRefID).playerTargetedByRef
	if playerTargetedByRef.OperationKind == OperationKindDelete {
		return
	}
	engine.deleteAnyOfPlayer_ZoneItem(playerTargetedByRef.ReferencedElementID, false)
	if _, ok := engine.State.PlayerTargetedByRef[playerTargetedByRefID]; ok {
		playerTargetedByRef.OperationKind = OperationKindDelete
		playerTargetedByRef.Meta.sign(playerTargetedByRef.engine.broadcastingClientID)
		engine.Patch.PlayerTargetedByRef[playerTargetedByRef.ID] = playerTargetedByRef
	} else {
		delete(engine.Patch.PlayerTargetedByRef, playerTargetedByRefID)
	}
}
func (engine *Engine) deleteAnyOfPlayer_Position(anyOfPlayer_PositionID AnyOfPlayer_PositionID, deleteChild bool) {
	anyOfPlayer_Position := engine.anyOfPlayer_Position(anyOfPlayer_PositionID).anyOfPlayer_Position
	if anyOfPlayer_Position.OperationKind == OperationKindDelete {
		return
	}
	if deleteChild {
		anyOfPlayer_Position.deleteChild()
	}
	if _, ok := engine.State.AnyOfPlayer_Position[anyOfPlayer_PositionID]; ok {
		anyOfPlayer_Position.OperationKind = OperationKindDelete
		anyOfPlayer_Position.Meta.sign(anyOfPlayer_Position.engine.broadcastingClientID)
		engine.Patch.AnyOfPlayer_Position[anyOfPlayer_Position.ID] = anyOfPlayer_Position
	} else {
		delete(engine.Patch.AnyOfPlayer_Position, anyOfPlayer_PositionID)
	}
}
func (engine *Engine) deleteAnyOfPlayer_ZoneItem(anyOfPlayer_ZoneItemID AnyOfPlayer_ZoneItemID, deleteChild bool) {
	anyOfPlayer_ZoneItem := engine.anyOfPlayer_ZoneItem(anyOfPlayer_ZoneItemID).anyOfPlayer_ZoneItem
	if anyOfPlayer_ZoneItem.OperationKind == OperationKindDelete {
		return
	}
	if deleteChild {
		anyOfPlayer_ZoneItem.deleteChild()
	}
	if _, ok := engine.State.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItemID]; ok {
		anyOfPlayer_ZoneItem.OperationKind = OperationKindDelete
		anyOfPlayer_ZoneItem.Meta.sign(anyOfPlayer_ZoneItem.engine.broadcastingClientID)
		engine.Patch.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItem.ID] = anyOfPlayer_ZoneItem
	} else {
		delete(engine.Patch.AnyOfPlayer_ZoneItem, anyOfPlayer_ZoneItemID)
	}
}
func (engine *Engine) deleteAnyOfItem_Player_ZoneItem(anyOfItem_Player_ZoneItemID AnyOfItem_Player_ZoneItemID, deleteChild bool) {
	anyOfItem_Player_ZoneItem := engine.anyOfItem_Player_ZoneItem(anyOfItem_Player_ZoneItemID).anyOfItem_Player_ZoneItem
	if anyOfItem_Player_ZoneItem.OperationKind == OperationKindDelete {
		return
	}
	if deleteChild {
		anyOfItem_Player_ZoneItem.deleteChild()
	}
	if _, ok := engine.State.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItemID]; ok {
		anyOfItem_Player_ZoneItem.OperationKind = OperationKindDelete
		anyOfItem_Player_ZoneItem.Meta.sign(anyOfItem_Player_ZoneItem.engine.broadcastingClientID)
		engine.Patch.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItem.ID] = anyOfItem_Player_ZoneItem
	} else {
		delete(engine.Patch.AnyOfItem_Player_ZoneItem, anyOfItem_Player_ZoneItemID)
	}
}
func (engine *Engine) boolValue(boolValueID BoolValueID) boolValue {
	patchingBoolValue, ok := engine.Patch.BoolValue[boolValueID]
	if ok {
		return patchingBoolValue
	}
	return engine.State.BoolValue[boolValueID]
}
func (engine *Engine) floatValue(floatValueID FloatValueID) floatValue {
	patchingFloatValue, ok := engine.Patch.FloatValue[floatValueID]
	if ok {
		return patchingFloatValue
	}
	return engine.State.FloatValue[floatValueID]
}
func (engine *Engine) intValue(intValueID IntValueID) intValue {
	patchingIntValue, ok := engine.Patch.IntValue[intValueID]
	if ok {
		return patchingIntValue
	}
	return engine.State.IntValue[intValueID]
}
func (engine *Engine) stringValue(stringValueID StringValueID) stringValue {
	patchingStringValue, ok := engine.Patch.StringValue[stringValueID]
	if ok {
		return patchingStringValue
	}
	return engine.State.StringValue[stringValueID]
}
func (_attackEvent AttackEvent) Exists() (AttackEvent, bool) {
	attackEvent := _attackEvent.attackEvent.engine.AttackEvent(_attackEvent.attackEvent.ID)
	return attackEvent, attackEvent.attackEvent.OperationKind != OperationKindDelete
}
func (engine *Engine) QueryAttackEvents(matcher func(AttackEvent) bool) []AttackEvent {
	attackEventIDs := engine.allAttackEventIDs()
	sort.Slice(attackEventIDs, func(i, j int) bool {
		return attackEventIDs[i] < attackEventIDs[j]
	})
	var attackEvents []AttackEvent
	for _, attackEventID := range attackEventIDs {
		attackEvent := engine.AttackEvent(attackEventID)
		if matcher(attackEvent) {
			attackEvents = append(attackEvents, attackEvent)
		}
	}
	attackEventIDSlicePool.Put(attackEventIDs)
	return attackEvents
}
func (engine *Engine) EveryAttackEvent() []AttackEvent {
	attackEventIDs := engine.allAttackEventIDs()
	sort.Slice(attackEventIDs, func(i, j int) bool {
		return attackEventIDs[i] < attackEventIDs[j]
	})
	var attackEvents []AttackEvent
	for _, attackEventID := range attackEventIDs {
		attackEvent := engine.AttackEvent(attackEventID)
		if attackEvent.attackEvent.HasParent {
			continue
		}
		attackEvents = append(attackEvents, attackEvent)
	}
	attackEventIDSlicePool.Put(attackEventIDs)
	return attackEvents
}
func (_attackEvent AttackEvent) ParentKind() (ElementKind, bool) {
	if !_attackEvent.attackEvent.HasParent {
		return "", false
	}
	return _attackEvent.attackEvent.Path[len(_attackEvent.attackEvent.Path)-2].Kind, true
}
func (engine *Engine) AttackEvent(attackEventID AttackEventID) AttackEvent {
	patchingAttackEvent, ok := engine.Patch.AttackEvent[attackEventID]
	if ok {
		return AttackEvent{attackEvent: patchingAttackEvent}
	}
	currentAttackEvent, ok := engine.State.AttackEvent[attackEventID]
	if ok {
		return AttackEvent{attackEvent: currentAttackEvent}
	}
	return AttackEvent{attackEvent: attackEventCore{
		OperationKind: OperationKindDelete,
		engine:        engine,
	}}
}
func (_attackEvent AttackEvent) ParentPlayer() Player {
	attackEvent := _attackEvent.attackEvent.engine.AttackEvent(_attackEvent.attackEvent.ID)
	if !attackEvent.attackEvent.HasParent {
		return Player{player: playerCore{
			OperationKind: OperationKindDelete,
			engine:        attackEvent.attackEvent.engine,
		}}
	}
	parentSeg := attackEvent.attackEvent.Path[len(attackEvent.attackEvent.Path)-2]
	return attackEvent.attackEvent.engine.Player(PlayerID(parentSeg.ID))
}
func (_attackEvent AttackEvent) ID() AttackEventID {
	return _attackEvent.attackEvent.ID
}
func (_attackEvent AttackEvent) Path() string {
	return _attackEvent.attackEvent.JSONPath
}
func (_attackEvent AttackEvent) Target() AttackEventTargetRef {
	attackEvent := _attackEvent.attackEvent.engine.AttackEvent(_attackEvent.attackEvent.ID)

	return attackEvent.attackEvent.engine.attackEventTargetRef(attackEvent.attackEvent.Target)
}
func (_equipmentSet EquipmentSet) Exists() (EquipmentSet, bool) {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	return equipmentSet, equipmentSet.equipmentSet.OperationKind != OperationKindDelete
}
func (engine *Engine) QueryEquipmentSets(matcher func(EquipmentSet) bool) []EquipmentSet {
	equipmentSetIDs := engine.allEquipmentSetIDs()
	sort.Slice(equipmentSetIDs, func(i, j int) bool {
		return equipmentSetIDs[i] < equipmentSetIDs[j]
	})
	var equipmentSets []EquipmentSet
	for _, equipmentSetID := range equipmentSetIDs {
		equipmentSet := engine.EquipmentSet(equipmentSetID)
		if matcher(equipmentSet) {
			equipmentSets = append(equipmentSets, equipmentSet)
		}
	}
	equipmentSetIDSlicePool.Put(equipmentSetIDs)
	return equipmentSets
}
func (engine *Engine) EveryEquipmentSet() []EquipmentSet {
	equipmentSetIDs := engine.allEquipmentSetIDs()
	sort.Slice(equipmentSetIDs, func(i, j int) bool {
		return equipmentSetIDs[i] < equipmentSetIDs[j]
	})
	var equipmentSets []EquipmentSet
	for _, equipmentSetID := range equipmentSetIDs {
		equipmentSet := engine.EquipmentSet(equipmentSetID)

		equipmentSets = append(equipmentSets, equipmentSet)
	}
	equipmentSetIDSlicePool.Put(equipmentSetIDs)
	return equipmentSets
}
func (engine *Engine) EquipmentSet(equipmentSetID EquipmentSetID) EquipmentSet {
	patchingEquipmentSet, ok := engine.Patch.EquipmentSet[equipmentSetID]
	if ok {
		return EquipmentSet{equipmentSet: patchingEquipmentSet}
	}
	currentEquipmentSet, ok := engine.State.EquipmentSet[equipmentSetID]
	if ok {
		return EquipmentSet{equipmentSet: currentEquipmentSet}
	}
	return EquipmentSet{equipmentSet: equipmentSetCore{
		OperationKind: OperationKindDelete,
		engine:        engine,
	}}
}
func (_equipmentSet EquipmentSet) ID() EquipmentSetID {
	return _equipmentSet.equipmentSet.ID
}
func (_equipmentSet EquipmentSet) Path() string {
	return _equipmentSet.equipmentSet.JSONPath
}
func (_equipmentSet EquipmentSet) Equipment() []EquipmentSetEquipmentRef {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	var equipment []EquipmentSetEquipmentRef
	for _, refID := range equipmentSet.equipmentSet.Equipment {
		equipment = append(equipment, equipmentSet.equipmentSet.engine.equipmentSetEquipmentRef(refID))
	}
	return equipment

}
func (_equipmentSet EquipmentSet) Name() string {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)

	return equipmentSet.equipmentSet.engine.stringValue(equipmentSet.equipmentSet.Name).Value
}
func (_gearScore GearScore) Exists() (GearScore, bool) {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	return gearScore, gearScore.gearScore.OperationKind != OperationKindDelete
}
func (engine *Engine) QueryGearScores(matcher func(GearScore) bool) []GearScore {
	gearScoreIDs := engine.allGearScoreIDs()
	sort.Slice(gearScoreIDs, func(i, j int) bool {
		return gearScoreIDs[i] < gearScoreIDs[j]
	})
	var gearScores []GearScore
	for _, gearScoreID := range gearScoreIDs {
		gearScore := engine.GearScore(gearScoreID)
		if matcher(gearScore) {
			gearScores = append(gearScores, gearScore)
		}
	}
	gearScoreIDSlicePool.Put(gearScoreIDs)
	return gearScores
}
func (engine *Engine) EveryGearScore() []GearScore {
	gearScoreIDs := engine.allGearScoreIDs()
	sort.Slice(gearScoreIDs, func(i, j int) bool {
		return gearScoreIDs[i] < gearScoreIDs[j]
	})
	var gearScores []GearScore
	for _, gearScoreID := range gearScoreIDs {
		gearScore := engine.GearScore(gearScoreID)
		if gearScore.gearScore.HasParent {
			continue
		}
		gearScores = append(gearScores, gearScore)
	}
	gearScoreIDSlicePool.Put(gearScoreIDs)
	return gearScores
}
func (_gearScore GearScore) ParentKind() (ElementKind, bool) {
	if !_gearScore.gearScore.HasParent {
		return "", false
	}
	return _gearScore.gearScore.Path[len(_gearScore.gearScore.Path)-2].Kind, true
}
func (engine *Engine) GearScore(gearScoreID GearScoreID) GearScore {
	patchingGearScore, ok := engine.Patch.GearScore[gearScoreID]
	if ok {
		return GearScore{gearScore: patchingGearScore}
	}
	currentGearScore, ok := engine.State.GearScore[gearScoreID]
	if ok {
		return GearScore{gearScore: currentGearScore}
	}
	return GearScore{gearScore: gearScoreCore{
		OperationKind: OperationKindDelete,
		engine:        engine,
	}}
}
func (_gearScore GearScore) ParentItem() Item {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	if !gearScore.gearScore.HasParent {
		return Item{item: itemCore{
			OperationKind: OperationKindDelete,
			engine:        gearScore.gearScore.engine,
		}}
	}
	parentSeg := gearScore.gearScore.Path[len(gearScore.gearScore.Path)-2]
	return gearScore.gearScore.engine.Item(ItemID(parentSeg.ID))
}
func (_gearScore GearScore) ParentPlayer() Player {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	if !gearScore.gearScore.HasParent {
		return Player{player: playerCore{
			OperationKind: OperationKindDelete,
			engine:        gearScore.gearScore.engine,
		}}
	}
	parentSeg := gearScore.gearScore.Path[len(gearScore.gearScore.Path)-2]
	return gearScore.gearScore.engine.Player(PlayerID(parentSeg.ID))
}
func (_gearScore GearScore) ID() GearScoreID {
	return _gearScore.gearScore.ID
}
func (_gearScore GearScore) Path() string {
	return _gearScore.gearScore.JSONPath
}
func (_gearScore GearScore) Level() int64 {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)

	return gearScore.gearScore.engine.intValue(gearScore.gearScore.Level).Value
}
func (_gearScore GearScore) Score() int64 {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)

	return gearScore.gearScore.engine.intValue(gearScore.gearScore.Score).Value
}
func (_item Item) Exists() (Item, bool) {
	item := _item.item.engine.Item(_item.item.ID)
	return item, item.item.OperationKind != OperationKindDelete
}
func (engine *Engine) QueryItems(matcher func(Item) bool) []Item {
	itemIDs := engine.allItemIDs()
	sort.Slice(itemIDs, func(i, j int) bool {
		return itemIDs[i] < itemIDs[j]
	})
	var items []Item
	for _, itemID := range itemIDs {
		item := engine.Item(itemID)
		if matcher(item) {
			items = append(items, item)
		}
	}
	itemIDSlicePool.Put(itemIDs)
	return items
}
func (engine *Engine) EveryItem() []Item {
	itemIDs := engine.allItemIDs()
	sort.Slice(itemIDs, func(i, j int) bool {
		return itemIDs[i] < itemIDs[j]
	})
	var items []Item
	for _, itemID := range itemIDs {
		item := engine.Item(itemID)
		if item.item.HasParent {
			continue
		}
		items = append(items, item)
	}
	itemIDSlicePool.Put(itemIDs)
	return items
}
func (_item Item) ParentKind() (ElementKind, bool) {
	if !_item.item.HasParent {
		return "", false
	}
	return _item.item.Path[len(_item.item.Path)-2].Kind, true
}
func (engine *Engine) Item(itemID ItemID) Item {
	patchingItem, ok := engine.Patch.Item[itemID]
	if ok {
		return Item{item: patchingItem}
	}
	currentItem, ok := engine.State.Item[itemID]
	if ok {
		return Item{item: currentItem}
	}
	return Item{item: itemCore{
		OperationKind: OperationKindDelete,
		engine:        engine,
	}}
}
func (_item Item) ParentPlayer() Player {
	item := _item.item.engine.Item(_item.item.ID)
	if !item.item.HasParent {
		return Player{player: playerCore{
			OperationKind: OperationKindDelete,
			engine:        item.item.engine,
		}}
	}
	parentSeg := item.item.Path[len(item.item.Path)-2]
	return item.item.engine.Player(PlayerID(parentSeg.ID))
}
func (_item Item) ParentZone() Zone {
	item := _item.item.engine.Item(_item.item.ID)
	if !item.item.HasParent {
		return Zone{zone: zoneCore{
			OperationKind: OperationKindDelete,
			engine:        item.item.engine,
		}}
	}
	parentSeg := item.item.Path[len(item.item.Path)-2]
	return item.item.engine.Zone(ZoneID(parentSeg.ID))
}
func (_item Item) ParentZoneItem() ZoneItem {
	item := _item.item.engine.Item(_item.item.ID)
	if !item.item.HasParent {
		return ZoneItem{zoneItem: zoneItemCore{
			OperationKind: OperationKindDelete,
			engine:        item.item.engine,
		}}
	}
	parentSeg := item.item.Path[len(item.item.Path)-2]
	return item.item.engine.ZoneItem(ZoneItemID(parentSeg.ID))
}
func (_item Item) ID() ItemID {
	return _item.item.ID
}
func (_item Item) Path() string {
	return _item.item.JSONPath
}
func (_item Item) BoundTo() ItemBoundToRef {
	item := _item.item.engine.Item(_item.item.ID)

	return item.item.engine.itemBoundToRef(item.item.BoundTo)
}
func (_item Item) GearScore() GearScore {
	item := _item.item.engine.Item(_item.item.ID)

	return item.item.engine.GearScore(item.item.GearScore)
}
func (_item Item) Name() string {
	item := _item.item.engine.Item(_item.item.ID)

	return item.item.engine.stringValue(item.item.Name).Value
}
func (_item Item) Origin() AnyOfPlayer_Position {
	item := _item.item.engine.Item(_item.item.ID)

	return item.item.engine.anyOfPlayer_Position(item.item.Origin)
}
func (_player Player) Exists() (Player, bool) {
	player := _player.player.engine.Player(_player.player.ID)
	return player, player.player.OperationKind != OperationKindDelete
}
func (engine *Engine) QueryPlayers(matcher func(Player) bool) []Player {
	playerIDs := engine.allPlayerIDs()
	sort.Slice(playerIDs, func(i, j int) bool {
		return playerIDs[i] < playerIDs[j]
	})
	var players []Player
	for _, playerID := range playerIDs {
		player := engine.Player(playerID)
		if matcher(player) {
			players = append(players, player)
		}
	}
	playerIDSlicePool.Put(playerIDs)
	return players
}
func (engine *Engine) EveryPlayer() []Player {
	playerIDs := engine.allPlayerIDs()
	sort.Slice(playerIDs, func(i, j int) bool {
		return playerIDs[i] < playerIDs[j]
	})
	var players []Player
	for _, playerID := range playerIDs {
		player := engine.Player(playerID)
		if player.player.HasParent {
			continue
		}
		players = append(players, player)
	}
	playerIDSlicePool.Put(playerIDs)
	return players
}
func (_player Player) ParentKind() (ElementKind, bool) {
	if !_player.player.HasParent {
		return "", false
	}
	return _player.player.Path[len(_player.player.Path)-2].Kind, true
}
func (engine *Engine) Player(playerID PlayerID) Player {
	patchingPlayer, ok := engine.Patch.Player[playerID]
	if ok {
		return Player{player: patchingPlayer}
	}
	currentPlayer, ok := engine.State.Player[playerID]
	if ok {
		return Player{player: currentPlayer}
	}
	return Player{player: playerCore{
		OperationKind: OperationKindDelete,
		engine:        engine,
	}}
}
func (_player Player) ParentItem() Item {
	player := _player.player.engine.Player(_player.player.ID)
	if !player.player.HasParent {
		return Item{item: itemCore{
			OperationKind: OperationKindDelete,
			engine:        player.player.engine,
		}}
	}
	parentSeg := player.player.Path[len(player.player.Path)-2]
	return player.player.engine.Item(ItemID(parentSeg.ID))
}
func (_player Player) ParentZone() Zone {
	player := _player.player.engine.Player(_player.player.ID)
	if !player.player.HasParent {
		return Zone{zone: zoneCore{
			OperationKind: OperationKindDelete,
			engine:        player.player.engine,
		}}
	}
	parentSeg := player.player.Path[len(player.player.Path)-2]
	return player.player.engine.Zone(ZoneID(parentSeg.ID))
}
func (_player Player) ID() PlayerID {
	return _player.player.ID
}
func (_player Player) Path() string {
	return _player.player.JSONPath
}
func (_player Player) Action() []AttackEvent {
	player := _player.player.engine.Player(_player.player.ID)
	var action []AttackEvent
	for _, attackEventID := range player.player.Action {
		action = append(action, player.player.engine.AttackEvent(attackEventID))
	}
	return action

}
func (_player Player) EquipmentSets() []PlayerEquipmentSetRef {
	player := _player.player.engine.Player(_player.player.ID)
	var equipmentSets []PlayerEquipmentSetRef
	for _, refID := range player.player.EquipmentSets {
		equipmentSets = append(equipmentSets, player.player.engine.playerEquipmentSetRef(refID))
	}
	return equipmentSets

}
func (_player Player) GearScore() GearScore {
	player := _player.player.engine.Player(_player.player.ID)

	return player.player.engine.GearScore(player.player.GearScore)
}
func (_player Player) GuildMembers() []PlayerGuildMemberRef {
	player := _player.player.engine.Player(_player.player.ID)
	var guildMembers []PlayerGuildMemberRef
	for _, refID := range player.player.GuildMembers {
		guildMembers = append(guildMembers, player.player.engine.playerGuildMemberRef(refID))
	}
	return guildMembers

}
func (_player Player) Items() []Item {
	player := _player.player.engine.Player(_player.player.ID)
	var items []Item
	for _, itemID := range player.player.Items {
		items = append(items, player.player.engine.Item(itemID))
	}
	return items

}
func (_player Player) Position() Position {
	player := _player.player.engine.Player(_player.player.ID)

	return player.player.engine.Position(player.player.Position)
}
func (_player Player) Target() PlayerTargetRef {
	player := _player.player.engine.Player(_player.player.ID)

	return player.player.engine.playerTargetRef(player.player.Target)
}
func (_player Player) TargetedBy() []PlayerTargetedByRef {
	player := _player.player.engine.Player(_player.player.ID)
	var targetedBy []PlayerTargetedByRef
	for _, refID := range player.player.TargetedBy {
		targetedBy = append(targetedBy, player.player.engine.playerTargetedByRef(refID))
	}
	return targetedBy

}
func (_position Position) Exists() (Position, bool) {
	position := _position.position.engine.Position(_position.position.ID)
	return position, position.position.OperationKind != OperationKindDelete
}
func (engine *Engine) QueryPositions(matcher func(Position) bool) []Position {
	positionIDs := engine.allPositionIDs()
	sort.Slice(positionIDs, func(i, j int) bool {
		return positionIDs[i] < positionIDs[j]
	})
	var positions []Position
	for _, positionID := range positionIDs {
		position := engine.Position(positionID)
		if matcher(position) {
			positions = append(positions, position)
		}
	}
	positionIDSlicePool.Put(positionIDs)
	return positions
}
func (engine *Engine) EveryPosition() []Position {
	positionIDs := engine.allPositionIDs()
	sort.Slice(positionIDs, func(i, j int) bool {
		return positionIDs[i] < positionIDs[j]
	})
	var positions []Position
	for _, positionID := range positionIDs {
		position := engine.Position(positionID)
		if position.position.HasParent {
			continue
		}
		positions = append(positions, position)
	}
	positionIDSlicePool.Put(positionIDs)
	return positions
}
func (_position Position) ParentKind() (ElementKind, bool) {
	if !_position.position.HasParent {
		return "", false
	}
	return _position.position.Path[len(_position.position.Path)-2].Kind, true
}
func (engine *Engine) Position(positionID PositionID) Position {
	patchingPosition, ok := engine.Patch.Position[positionID]
	if ok {
		return Position{position: patchingPosition}
	}
	currentPosition, ok := engine.State.Position[positionID]
	if ok {
		return Position{position: currentPosition}
	}
	return Position{position: positionCore{
		OperationKind: OperationKindDelete,
		engine:        engine,
	}}
}
func (_position Position) ParentItem() Item {
	position := _position.position.engine.Position(_position.position.ID)
	if !position.position.HasParent {
		return Item{item: itemCore{
			OperationKind: OperationKindDelete,
			engine:        position.position.engine,
		}}
	}
	parentSeg := position.position.Path[len(position.position.Path)-2]
	return position.position.engine.Item(ItemID(parentSeg.ID))
}
func (_position Position) ParentPlayer() Player {
	position := _position.position.engine.Position(_position.position.ID)
	if !position.position.HasParent {
		return Player{player: playerCore{
			OperationKind: OperationKindDelete,
			engine:        position.position.engine,
		}}
	}
	parentSeg := position.position.Path[len(position.position.Path)-2]
	return position.position.engine.Player(PlayerID(parentSeg.ID))
}
func (_position Position) ParentZoneItem() ZoneItem {
	position := _position.position.engine.Position(_position.position.ID)
	if !position.position.HasParent {
		return ZoneItem{zoneItem: zoneItemCore{
			OperationKind: OperationKindDelete,
			engine:        position.position.engine,
		}}
	}
	parentSeg := position.position.Path[len(position.position.Path)-2]
	return position.position.engine.ZoneItem(ZoneItemID(parentSeg.ID))
}
func (_position Position) ID() PositionID {
	return _position.position.ID
}
func (_position Position) Path() string {
	return _position.position.JSONPath
}
func (_position Position) X() float64 {
	position := _position.position.engine.Position(_position.position.ID)

	return position.position.engine.floatValue(position.position.X).Value
}
func (_position Position) Y() float64 {
	position := _position.position.engine.Position(_position.position.ID)

	return position.position.engine.floatValue(position.position.Y).Value
}
func (_zone Zone) Exists() (Zone, bool) {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	return zone, zone.zone.OperationKind != OperationKindDelete
}
func (engine *Engine) QueryZones(matcher func(Zone) bool) []Zone {
	zoneIDs := engine.allZoneIDs()
	sort.Slice(zoneIDs, func(i, j int) bool {
		return zoneIDs[i] < zoneIDs[j]
	})
	var zones []Zone
	for _, zoneID := range zoneIDs {
		zone := engine.Zone(zoneID)
		if matcher(zone) {
			zones = append(zones, zone)
		}
	}
	zoneIDSlicePool.Put(zoneIDs)
	return zones
}
func (engine *Engine) EveryZone() []Zone {
	zoneIDs := engine.allZoneIDs()
	sort.Slice(zoneIDs, func(i, j int) bool {
		return zoneIDs[i] < zoneIDs[j]
	})
	var zones []Zone
	for _, zoneID := range zoneIDs {
		zone := engine.Zone(zoneID)

		zones = append(zones, zone)
	}
	zoneIDSlicePool.Put(zoneIDs)
	return zones
}
func (engine *Engine) Zone(zoneID ZoneID) Zone {
	patchingZone, ok := engine.Patch.Zone[zoneID]
	if ok {
		return Zone{zone: patchingZone}
	}
	currentZone, ok := engine.State.Zone[zoneID]
	if ok {
		return Zone{zone: currentZone}
	}
	return Zone{zone: zoneCore{
		OperationKind: OperationKindDelete,
		engine:        engine,
	}}
}
func (_zone Zone) ID() ZoneID {
	return _zone.zone.ID
}
func (_zone Zone) Path() string {
	return _zone.zone.JSONPath
}
func (_zone Zone) Interactables() []AnyOfItem_Player_ZoneItem {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	var interactables []AnyOfItem_Player_ZoneItem
	for _, anyOfItem_Player_ZoneItemID := range zone.zone.Interactables {
		interactables = append(interactables, zone.zone.engine.anyOfItem_Player_ZoneItem(anyOfItem_Player_ZoneItemID))
	}
	return interactables

}
func (_zone Zone) Items() []ZoneItem {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	var items []ZoneItem
	for _, zoneItemID := range zone.zone.Items {
		items = append(items, zone.zone.engine.ZoneItem(zoneItemID))
	}
	return items

}
func (_zone Zone) Players() []Player {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	var players []Player
	for _, playerID := range zone.zone.Players {
		players = append(players, zone.zone.engine.Player(playerID))
	}
	return players

}
func (_zone Zone) Tags() []string {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	var tags []string
	for _, stringValueID := range zone.zone.Tags {
		tags = append(tags, zone.zone.engine.stringValue(stringValueID).Value)
	}
	return tags

}
func (_zoneItem ZoneItem) Exists() (ZoneItem, bool) {
	zoneItem := _zoneItem.zoneItem.engine.ZoneItem(_zoneItem.zoneItem.ID)
	return zoneItem, zoneItem.zoneItem.OperationKind != OperationKindDelete
}
func (engine *Engine) QueryZoneItems(matcher func(ZoneItem) bool) []ZoneItem {
	zoneItemIDs := engine.allZoneItemIDs()
	sort.Slice(zoneItemIDs, func(i, j int) bool {
		return zoneItemIDs[i] < zoneItemIDs[j]
	})
	var zoneItems []ZoneItem
	for _, zoneItemID := range zoneItemIDs {
		zoneItem := engine.ZoneItem(zoneItemID)
		if matcher(zoneItem) {
			zoneItems = append(zoneItems, zoneItem)
		}
	}
	zoneItemIDSlicePool.Put(zoneItemIDs)
	return zoneItems
}
func (engine *Engine) EveryZoneItem() []ZoneItem {
	zoneItemIDs := engine.allZoneItemIDs()
	sort.Slice(zoneItemIDs, func(i, j int) bool {
		return zoneItemIDs[i] < zoneItemIDs[j]
	})
	var zoneItems []ZoneItem
	for _, zoneItemID := range zoneItemIDs {
		zoneItem := engine.ZoneItem(zoneItemID)
		if zoneItem.zoneItem.HasParent {
			continue
		}
		zoneItems = append(zoneItems, zoneItem)
	}
	zoneItemIDSlicePool.Put(zoneItemIDs)
	return zoneItems
}
func (_zoneItem ZoneItem) ParentKind() (ElementKind, bool) {
	if !_zoneItem.zoneItem.HasParent {
		return "", false
	}
	return _zoneItem.zoneItem.Path[len(_zoneItem.zoneItem.Path)-2].Kind, true
}
func (engine *Engine) ZoneItem(zoneItemID ZoneItemID) ZoneItem {
	patchingZoneItem, ok := engine.Patch.ZoneItem[zoneItemID]
	if ok {
		return ZoneItem{zoneItem: patchingZoneItem}
	}
	currentZoneItem, ok := engine.State.ZoneItem[zoneItemID]
	if ok {
		return ZoneItem{zoneItem: currentZoneItem}
	}
	return ZoneItem{zoneItem: zoneItemCore{
		OperationKind: OperationKindDelete,
		engine:        engine,
	}}
}
func (_zoneItem ZoneItem) ParentZone() Zone {
	zoneItem := _zoneItem.zoneItem.engine.ZoneItem(_zoneItem.zoneItem.ID)
	if !zoneItem.zoneItem.HasParent {
		return Zone{zone: zoneCore{
			OperationKind: OperationKindDelete,
			engine:        zoneItem.zoneItem.engine,
		}}
	}
	parentSeg := zoneItem.zoneItem.Path[len(zoneItem.zoneItem.Path)-2]
	return zoneItem.zoneItem.engine.Zone(ZoneID(parentSeg.ID))
}
func (_zoneItem ZoneItem) ID() ZoneItemID {
	return _zoneItem.zoneItem.ID
}
func (_zoneItem ZoneItem) Path() string {
	return _zoneItem.zoneItem.JSONPath
}
func (_zoneItem ZoneItem) Item() Item {
	zoneItem := _zoneItem.zoneItem.engine.ZoneItem(_zoneItem.zoneItem.ID)

	return zoneItem.zoneItem.engine.Item(zoneItem.zoneItem.Item)
}
func (_zoneItem ZoneItem) Position() Position {
	zoneItem := _zoneItem.zoneItem.engine.ZoneItem(_zoneItem.zoneItem.ID)

	return zoneItem.zoneItem.engine.Position(zoneItem.zoneItem.Position)
}
func (engine *Engine) attackEventTargetRef(attackEventTargetRefID AttackEventTargetRefID) AttackEventTargetRef {
	patchingAttackEventTargetRef, ok := engine.Patch.AttackEventTargetRef[attackEventTargetRefID]
	if ok {
		return AttackEventTargetRef{attackEventTargetRef: patchingAttackEventTargetRef}
	}
	currentAttackEventTargetRef, ok := engine.State.AttackEventTargetRef[attackEventTargetRefID]
	if ok {
		return AttackEventTargetRef{attackEventTargetRef: currentAttackEventTargetRef}
	}
	return AttackEventTargetRef{attackEventTargetRef: attackEventTargetRefCore{
		OperationKind: OperationKindDelete,
		engine:        engine,
	}}
}
func (_attackEventTargetRef AttackEventTargetRef) ID() PlayerID {
	return _attackEventTargetRef.attackEventTargetRef.ReferencedElementID
}
func (engine *Engine) equipmentSetEquipmentRef(equipmentSetEquipmentRefID EquipmentSetEquipmentRefID) EquipmentSetEquipmentRef {
	patchingEquipmentSetEquipmentRef, ok := engine.Patch.EquipmentSetEquipmentRef[equipmentSetEquipmentRefID]
	if ok {
		return EquipmentSetEquipmentRef{equipmentSetEquipmentRef: patchingEquipmentSetEquipmentRef}
	}
	currentEquipmentSetEquipmentRef, ok := engine.State.EquipmentSetEquipmentRef[equipmentSetEquipmentRefID]
	if ok {
		return EquipmentSetEquipmentRef{equipmentSetEquipmentRef: currentEquipmentSetEquipmentRef}
	}
	return EquipmentSetEquipmentRef{equipmentSetEquipmentRef: equipmentSetEquipmentRefCore{
		OperationKind: OperationKindDelete,
		engine:        engine,
	}}
}
func (_equipmentSetEquipmentRef EquipmentSetEquipmentRef) ID() ItemID {
	return _equipmentSetEquipmentRef.equipmentSetEquipmentRef.ReferencedElementID
}
func (engine *Engine) itemBoundToRef(itemBoundToRefID ItemBoundToRefID) ItemBoundToRef {
	patchingItemBoundToRef, ok := engine.Patch.ItemBoundToRef[itemBoundToRefID]
	if ok {
		return ItemBoundToRef{itemBoundToRef: patchingItemBoundToRef}
	}
	currentItemBoundToRef, ok := engine.State.ItemBoundToRef[itemBoundToRefID]
	if ok {
		return ItemBoundToRef{itemBoundToRef: currentItemBoundToRef}
	}
	return ItemBoundToRef{itemBoundToRef: itemBoundToRefCore{
		OperationKind: OperationKindDelete,
		engine:        engine,
	}}
}
func (_itemBoundToRef ItemBoundToRef) ID() PlayerID {
	return _itemBoundToRef.itemBoundToRef.ReferencedElementID
}
func (engine *Engine) playerEquipmentSetRef(playerEquipmentSetRefID PlayerEquipmentSetRefID) PlayerEquipmentSetRef {
	patchingPlayerEquipmentSetRef, ok := engine.Patch.PlayerEquipmentSetRef[playerEquipmentSetRefID]
	if ok {
		return PlayerEquipmentSetRef{playerEquipmentSetRef: patchingPlayerEquipmentSetRef}
	}
	currentPlayerEquipmentSetRef, ok := engine.State.PlayerEquipmentSetRef[playerEquipmentSetRefID]
	if ok {
		return PlayerEquipmentSetRef{playerEquipmentSetRef: currentPlayerEquipmentSetRef}
	}
	return PlayerEquipmentSetRef{playerEquipmentSetRef: playerEquipmentSetRefCore{
		OperationKind: OperationKindDelete,
		engine:        engine,
	}}
}
func (_playerEquipmentSetRef PlayerEquipmentSetRef) ID() EquipmentSetID {
	return _playerEquipmentSetRef.playerEquipmentSetRef.ReferencedElementID
}
func (engine *Engine) playerGuildMemberRef(playerGuildMemberRefID PlayerGuildMemberRefID) PlayerGuildMemberRef {
	patchingPlayerGuildMemberRef, ok := engine.Patch.PlayerGuildMemberRef[playerGuildMemberRefID]
	if ok {
		return PlayerGuildMemberRef{playerGuildMemberRef: patchingPlayerGuildMemberRef}
	}
	currentPlayerGuildMemberRef, ok := engine.State.PlayerGuildMemberRef[playerGuildMemberRefID]
	if ok {
		return PlayerGuildMemberRef{playerGuildMemberRef: currentPlayerGuildMemberRef}
	}
	return PlayerGuildMemberRef{playerGuildMemberRef: playerGuildMemberRefCore{
		OperationKind: OperationKindDelete,
		engine:        engine,
	}}
}
func (_playerGuildMemberRef PlayerGuildMemberRef) ID() PlayerID {
	return _playerGuildMemberRef.playerGuildMemberRef.ReferencedElementID
}
func (engine *Engine) playerTargetRef(playerTargetRefID PlayerTargetRefID) PlayerTargetRef {
	patchingPlayerTargetRef, ok := engine.Patch.PlayerTargetRef[playerTargetRefID]
	if ok {
		return PlayerTargetRef{playerTargetRef: patchingPlayerTargetRef}
	}
	currentPlayerTargetRef, ok := engine.State.PlayerTargetRef[playerTargetRefID]
	if ok {
		return PlayerTargetRef{playerTargetRef: currentPlayerTargetRef}
	}
	return PlayerTargetRef{playerTargetRef: playerTargetRefCore{
		OperationKind: OperationKindDelete,
		engine:        engine,
	}}
}
func (_playerTargetRef PlayerTargetRef) ID() AnyOfPlayer_ZoneItemID {
	return _playerTargetRef.playerTargetRef.ReferencedElementID
}
func (engine *Engine) playerTargetedByRef(playerTargetedByRefID PlayerTargetedByRefID) PlayerTargetedByRef {
	patchingPlayerTargetedByRef, ok := engine.Patch.PlayerTargetedByRef[playerTargetedByRefID]
	if ok {
		return PlayerTargetedByRef{playerTargetedByRef: patchingPlayerTargetedByRef}
	}
	currentPlayerTargetedByRef, ok := engine.State.PlayerTargetedByRef[playerTargetedByRefID]
	if ok {
		return PlayerTargetedByRef{playerTargetedByRef: currentPlayerTargetedByRef}
	}
	return PlayerTargetedByRef{playerTargetedByRef: playerTargetedByRefCore{
		OperationKind: OperationKindDelete,
		engine:        engine,
	}}
}
func (_playerTargetedByRef PlayerTargetedByRef) ID() AnyOfPlayer_ZoneItemID {
	return _playerTargetedByRef.playerTargetedByRef.ReferencedElementID
}
func (engine *Engine) anyOfPlayer_Position(anyOfPlayer_PositionID AnyOfPlayer_PositionID) AnyOfPlayer_Position {
	patchingAnyOfPlayer_Position, ok := engine.Patch.AnyOfPlayer_Position[anyOfPlayer_PositionID]
	if ok {
		return AnyOfPlayer_Position{anyOfPlayer_Position: patchingAnyOfPlayer_Position}
	}
	currentAnyOfPlayer_Position, ok := engine.State.AnyOfPlayer_Position[anyOfPlayer_PositionID]
	if ok {
		return AnyOfPlayer_Position{anyOfPlayer_Position: currentAnyOfPlayer_Position}
	}
	return AnyOfPlayer_Position{anyOfPlayer_Position: anyOfPlayer_PositionCore{
		OperationKind: OperationKindDelete,
		engine:        engine,
	}}
}
func (_anyOfPlayer_Position AnyOfPlayer_Position) ID() AnyOfPlayer_PositionID {
	return _anyOfPlayer_Position.anyOfPlayer_Position.ID
}
func (_anyOfPlayer_Position AnyOfPlayer_Position) Player() Player {
	anyOfPlayer_Position := _anyOfPlayer_Position.anyOfPlayer_Position.engine.anyOfPlayer_Position(_anyOfPlayer_Position.anyOfPlayer_Position.ID)
	return anyOfPlayer_Position.anyOfPlayer_Position.engine.Player(PlayerID(anyOfPlayer_Position.anyOfPlayer_Position.ChildID))
}
func (_anyOfPlayer_Position AnyOfPlayer_Position) Position() Position {
	anyOfPlayer_Position := _anyOfPlayer_Position.anyOfPlayer_Position.engine.anyOfPlayer_Position(_anyOfPlayer_Position.anyOfPlayer_Position.ID)
	return anyOfPlayer_Position.anyOfPlayer_Position.engine.Position(PositionID(anyOfPlayer_Position.anyOfPlayer_Position.ChildID))
}
func (engine *Engine) anyOfPlayer_ZoneItem(anyOfPlayer_ZoneItemID AnyOfPlayer_ZoneItemID) AnyOfPlayer_ZoneItem {
	patchingAnyOfPlayer_ZoneItem, ok := engine.Patch.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItemID]
	if ok {
		return AnyOfPlayer_ZoneItem{anyOfPlayer_ZoneItem: patchingAnyOfPlayer_ZoneItem}
	}
	currentAnyOfPlayer_ZoneItem, ok := engine.State.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItemID]
	if ok {
		return AnyOfPlayer_ZoneItem{anyOfPlayer_ZoneItem: currentAnyOfPlayer_ZoneItem}
	}
	return AnyOfPlayer_ZoneItem{anyOfPlayer_ZoneItem: anyOfPlayer_ZoneItemCore{
		OperationKind: OperationKindDelete,
		engine:        engine,
	}}
}
func (_anyOfPlayer_ZoneItem AnyOfPlayer_ZoneItem) ID() AnyOfPlayer_ZoneItemID {
	return _anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.ID
}
func (_anyOfPlayer_ZoneItem AnyOfPlayer_ZoneItem) Player() Player {
	anyOfPlayer_ZoneItem := _anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.ID)
	return anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.engine.Player(PlayerID(anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.ChildID))
}
func (_anyOfPlayer_ZoneItem AnyOfPlayer_ZoneItem) ZoneItem() ZoneItem {
	anyOfPlayer_ZoneItem := _anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.ID)
	return anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.engine.ZoneItem(ZoneItemID(anyOfPlayer_ZoneItem.anyOfPlayer_ZoneItem.ChildID))
}
func (engine *Engine) anyOfItem_Player_ZoneItem(anyOfItem_Player_ZoneItemID AnyOfItem_Player_ZoneItemID) AnyOfItem_Player_ZoneItem {
	patchingAnyOfItem_Player_ZoneItem, ok := engine.Patch.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItemID]
	if ok {
		return AnyOfItem_Player_ZoneItem{anyOfItem_Player_ZoneItem: patchingAnyOfItem_Player_ZoneItem}
	}
	currentAnyOfItem_Player_ZoneItem, ok := engine.State.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItemID]
	if ok {
		return AnyOfItem_Player_ZoneItem{anyOfItem_Player_ZoneItem: currentAnyOfItem_Player_ZoneItem}
	}
	return AnyOfItem_Player_ZoneItem{anyOfItem_Player_ZoneItem: anyOfItem_Player_ZoneItemCore{
		OperationKind: OperationKindDelete,
		engine:        engine,
	}}
}
func (_anyOfItem_Player_ZoneItem AnyOfItem_Player_ZoneItem) ID() AnyOfItem_Player_ZoneItemID {
	return _anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ID
}
func (_anyOfItem_Player_ZoneItem AnyOfItem_Player_ZoneItem) Item() Item {
	anyOfItem_Player_ZoneItem := _anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ID)
	return anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.Item(ItemID(anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ChildID))
}
func (_anyOfItem_Player_ZoneItem AnyOfItem_Player_ZoneItem) Player() Player {
	anyOfItem_Player_ZoneItem := _anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ID)
	return anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.Player(PlayerID(anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ChildID))
}
func (_anyOfItem_Player_ZoneItem AnyOfItem_Player_ZoneItem) ZoneItem() ZoneItem {
	anyOfItem_Player_ZoneItem := _anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ID)
	return anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.engine.ZoneItem(ZoneItemID(anyOfItem_Player_ZoneItem.anyOfItem_Player_ZoneItem.ChildID))
}
func deduplicateAttackEventIDs(a []AttackEventID, b []AttackEventID) []AttackEventID {
	check := attackEventCheckPool.Get().(map[AttackEventID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := attackEventIDSlicePool.Get().([]AttackEventID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	attackEventCheckPool.Put(check)
	return deduped
}
func deduplicateEquipmentSetIDs(a []EquipmentSetID, b []EquipmentSetID) []EquipmentSetID {
	check := equipmentSetCheckPool.Get().(map[EquipmentSetID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := equipmentSetIDSlicePool.Get().([]EquipmentSetID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	equipmentSetCheckPool.Put(check)
	return deduped
}
func deduplicateGearScoreIDs(a []GearScoreID, b []GearScoreID) []GearScoreID {
	check := gearScoreCheckPool.Get().(map[GearScoreID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := gearScoreIDSlicePool.Get().([]GearScoreID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	gearScoreCheckPool.Put(check)
	return deduped
}
func deduplicateItemIDs(a []ItemID, b []ItemID) []ItemID {
	check := itemCheckPool.Get().(map[ItemID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := itemIDSlicePool.Get().([]ItemID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	itemCheckPool.Put(check)
	return deduped
}
func deduplicatePlayerIDs(a []PlayerID, b []PlayerID) []PlayerID {
	check := playerCheckPool.Get().(map[PlayerID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerIDSlicePool.Get().([]PlayerID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	playerCheckPool.Put(check)
	return deduped
}
func deduplicatePositionIDs(a []PositionID, b []PositionID) []PositionID {
	check := positionCheckPool.Get().(map[PositionID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := positionIDSlicePool.Get().([]PositionID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	positionCheckPool.Put(check)
	return deduped
}
func deduplicateZoneIDs(a []ZoneID, b []ZoneID) []ZoneID {
	check := zoneCheckPool.Get().(map[ZoneID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := zoneIDSlicePool.Get().([]ZoneID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	zoneCheckPool.Put(check)
	return deduped
}
func deduplicateZoneItemIDs(a []ZoneItemID, b []ZoneItemID) []ZoneItemID {
	check := zoneItemCheckPool.Get().(map[ZoneItemID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := zoneItemIDSlicePool.Get().([]ZoneItemID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	zoneItemCheckPool.Put(check)
	return deduped
}
func deduplicateAttackEventTargetRefIDs(a []AttackEventTargetRefID, b []AttackEventTargetRefID) []AttackEventTargetRefID {
	check := attackEventTargetRefCheckPool.Get().(map[AttackEventTargetRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := attackEventTargetRefIDSlicePool.Get().([]AttackEventTargetRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	attackEventTargetRefCheckPool.Put(check)
	return deduped
}
func deduplicateEquipmentSetEquipmentRefIDs(a []EquipmentSetEquipmentRefID, b []EquipmentSetEquipmentRefID) []EquipmentSetEquipmentRefID {
	check := equipmentSetEquipmentRefCheckPool.Get().(map[EquipmentSetEquipmentRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := equipmentSetEquipmentRefIDSlicePool.Get().([]EquipmentSetEquipmentRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	equipmentSetEquipmentRefCheckPool.Put(check)
	return deduped
}
func deduplicateItemBoundToRefIDs(a []ItemBoundToRefID, b []ItemBoundToRefID) []ItemBoundToRefID {
	check := itemBoundToRefCheckPool.Get().(map[ItemBoundToRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := itemBoundToRefIDSlicePool.Get().([]ItemBoundToRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	itemBoundToRefCheckPool.Put(check)
	return deduped
}
func deduplicatePlayerEquipmentSetRefIDs(a []PlayerEquipmentSetRefID, b []PlayerEquipmentSetRefID) []PlayerEquipmentSetRefID {
	check := playerEquipmentSetRefCheckPool.Get().(map[PlayerEquipmentSetRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerEquipmentSetRefIDSlicePool.Get().([]PlayerEquipmentSetRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	playerEquipmentSetRefCheckPool.Put(check)
	return deduped
}
func deduplicatePlayerGuildMemberRefIDs(a []PlayerGuildMemberRefID, b []PlayerGuildMemberRefID) []PlayerGuildMemberRefID {
	check := playerGuildMemberRefCheckPool.Get().(map[PlayerGuildMemberRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerGuildMemberRefIDSlicePool.Get().([]PlayerGuildMemberRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	playerGuildMemberRefCheckPool.Put(check)
	return deduped
}
func deduplicatePlayerTargetRefIDs(a []PlayerTargetRefID, b []PlayerTargetRefID) []PlayerTargetRefID {
	check := playerTargetRefCheckPool.Get().(map[PlayerTargetRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerTargetRefIDSlicePool.Get().([]PlayerTargetRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	playerTargetRefCheckPool.Put(check)
	return deduped
}
func deduplicatePlayerTargetedByRefIDs(a []PlayerTargetedByRefID, b []PlayerTargetedByRefID) []PlayerTargetedByRefID {
	check := playerTargetedByRefCheckPool.Get().(map[PlayerTargetedByRefID]bool)
	for k := range check {
		delete(check, k)
	}
	deduped := playerTargetedByRefIDSlicePool.Get().([]PlayerTargetedByRefID)[:0]
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}
	for val := range check {
		deduped = append(deduped, val)
	}
	playerTargetedByRefCheckPool.Put(check)
	return deduped
}
func (engine Engine) allAttackEventIDs() []AttackEventID {
	stateAttackEventIDs := attackEventIDSlicePool.Get().([]AttackEventID)[:0]
	for attackEventID := range engine.State.AttackEvent {
		stateAttackEventIDs = append(stateAttackEventIDs, attackEventID)
	}
	patchAttackEventIDs := attackEventIDSlicePool.Get().([]AttackEventID)[:0]
	for attackEventID := range engine.Patch.AttackEvent {
		patchAttackEventIDs = append(patchAttackEventIDs, attackEventID)
	}
	dedupedIDs := deduplicateAttackEventIDs(stateAttackEventIDs, patchAttackEventIDs)
	attackEventIDSlicePool.Put(stateAttackEventIDs)
	attackEventIDSlicePool.Put(patchAttackEventIDs)
	return dedupedIDs
}
func (engine Engine) allEquipmentSetIDs() []EquipmentSetID {
	stateEquipmentSetIDs := equipmentSetIDSlicePool.Get().([]EquipmentSetID)[:0]
	for equipmentSetID := range engine.State.EquipmentSet {
		stateEquipmentSetIDs = append(stateEquipmentSetIDs, equipmentSetID)
	}
	patchEquipmentSetIDs := equipmentSetIDSlicePool.Get().([]EquipmentSetID)[:0]
	for equipmentSetID := range engine.Patch.EquipmentSet {
		patchEquipmentSetIDs = append(patchEquipmentSetIDs, equipmentSetID)
	}
	dedupedIDs := deduplicateEquipmentSetIDs(stateEquipmentSetIDs, patchEquipmentSetIDs)
	equipmentSetIDSlicePool.Put(stateEquipmentSetIDs)
	equipmentSetIDSlicePool.Put(patchEquipmentSetIDs)
	return dedupedIDs
}
func (engine Engine) allGearScoreIDs() []GearScoreID {
	stateGearScoreIDs := gearScoreIDSlicePool.Get().([]GearScoreID)[:0]
	for gearScoreID := range engine.State.GearScore {
		stateGearScoreIDs = append(stateGearScoreIDs, gearScoreID)
	}
	patchGearScoreIDs := gearScoreIDSlicePool.Get().([]GearScoreID)[:0]
	for gearScoreID := range engine.Patch.GearScore {
		patchGearScoreIDs = append(patchGearScoreIDs, gearScoreID)
	}
	dedupedIDs := deduplicateGearScoreIDs(stateGearScoreIDs, patchGearScoreIDs)
	gearScoreIDSlicePool.Put(stateGearScoreIDs)
	gearScoreIDSlicePool.Put(patchGearScoreIDs)
	return dedupedIDs
}
func (engine Engine) allItemIDs() []ItemID {
	stateItemIDs := itemIDSlicePool.Get().([]ItemID)[:0]
	for itemID := range engine.State.Item {
		stateItemIDs = append(stateItemIDs, itemID)
	}
	patchItemIDs := itemIDSlicePool.Get().([]ItemID)[:0]
	for itemID := range engine.Patch.Item {
		patchItemIDs = append(patchItemIDs, itemID)
	}
	dedupedIDs := deduplicateItemIDs(stateItemIDs, patchItemIDs)
	itemIDSlicePool.Put(stateItemIDs)
	itemIDSlicePool.Put(patchItemIDs)
	return dedupedIDs
}
func (engine Engine) allPlayerIDs() []PlayerID {
	statePlayerIDs := playerIDSlicePool.Get().([]PlayerID)[:0]
	for playerID := range engine.State.Player {
		statePlayerIDs = append(statePlayerIDs, playerID)
	}
	patchPlayerIDs := playerIDSlicePool.Get().([]PlayerID)[:0]
	for playerID := range engine.Patch.Player {
		patchPlayerIDs = append(patchPlayerIDs, playerID)
	}
	dedupedIDs := deduplicatePlayerIDs(statePlayerIDs, patchPlayerIDs)
	playerIDSlicePool.Put(statePlayerIDs)
	playerIDSlicePool.Put(patchPlayerIDs)
	return dedupedIDs
}
func (engine Engine) allPositionIDs() []PositionID {
	statePositionIDs := positionIDSlicePool.Get().([]PositionID)[:0]
	for positionID := range engine.State.Position {
		statePositionIDs = append(statePositionIDs, positionID)
	}
	patchPositionIDs := positionIDSlicePool.Get().([]PositionID)[:0]
	for positionID := range engine.Patch.Position {
		patchPositionIDs = append(patchPositionIDs, positionID)
	}
	dedupedIDs := deduplicatePositionIDs(statePositionIDs, patchPositionIDs)
	positionIDSlicePool.Put(statePositionIDs)
	positionIDSlicePool.Put(patchPositionIDs)
	return dedupedIDs
}
func (engine Engine) allZoneIDs() []ZoneID {
	stateZoneIDs := zoneIDSlicePool.Get().([]ZoneID)[:0]
	for zoneID := range engine.State.Zone {
		stateZoneIDs = append(stateZoneIDs, zoneID)
	}
	patchZoneIDs := zoneIDSlicePool.Get().([]ZoneID)[:0]
	for zoneID := range engine.Patch.Zone {
		patchZoneIDs = append(patchZoneIDs, zoneID)
	}
	dedupedIDs := deduplicateZoneIDs(stateZoneIDs, patchZoneIDs)
	zoneIDSlicePool.Put(stateZoneIDs)
	zoneIDSlicePool.Put(patchZoneIDs)
	return dedupedIDs
}
func (engine Engine) allZoneItemIDs() []ZoneItemID {
	stateZoneItemIDs := zoneItemIDSlicePool.Get().([]ZoneItemID)[:0]
	for zoneItemID := range engine.State.ZoneItem {
		stateZoneItemIDs = append(stateZoneItemIDs, zoneItemID)
	}
	patchZoneItemIDs := zoneItemIDSlicePool.Get().([]ZoneItemID)[:0]
	for zoneItemID := range engine.Patch.ZoneItem {
		patchZoneItemIDs = append(patchZoneItemIDs, zoneItemID)
	}
	dedupedIDs := deduplicateZoneItemIDs(stateZoneItemIDs, patchZoneItemIDs)
	zoneItemIDSlicePool.Put(stateZoneItemIDs)
	zoneItemIDSlicePool.Put(patchZoneItemIDs)
	return dedupedIDs
}
func (engine Engine) allAttackEventTargetRefIDs() []AttackEventTargetRefID {
	stateAttackEventTargetRefIDs := attackEventTargetRefIDSlicePool.Get().([]AttackEventTargetRefID)[:0]
	for attackEventTargetRefID := range engine.State.AttackEventTargetRef {
		stateAttackEventTargetRefIDs = append(stateAttackEventTargetRefIDs, attackEventTargetRefID)
	}
	patchAttackEventTargetRefIDs := attackEventTargetRefIDSlicePool.Get().([]AttackEventTargetRefID)[:0]
	for attackEventTargetRefID := range engine.Patch.AttackEventTargetRef {
		patchAttackEventTargetRefIDs = append(patchAttackEventTargetRefIDs, attackEventTargetRefID)
	}
	dedupedIDs := deduplicateAttackEventTargetRefIDs(stateAttackEventTargetRefIDs, patchAttackEventTargetRefIDs)
	attackEventTargetRefIDSlicePool.Put(stateAttackEventTargetRefIDs)
	attackEventTargetRefIDSlicePool.Put(patchAttackEventTargetRefIDs)
	return dedupedIDs
}
func (engine Engine) allEquipmentSetEquipmentRefIDs() []EquipmentSetEquipmentRefID {
	stateEquipmentSetEquipmentRefIDs := equipmentSetEquipmentRefIDSlicePool.Get().([]EquipmentSetEquipmentRefID)[:0]
	for equipmentSetEquipmentRefID := range engine.State.EquipmentSetEquipmentRef {
		stateEquipmentSetEquipmentRefIDs = append(stateEquipmentSetEquipmentRefIDs, equipmentSetEquipmentRefID)
	}
	patchEquipmentSetEquipmentRefIDs := equipmentSetEquipmentRefIDSlicePool.Get().([]EquipmentSetEquipmentRefID)[:0]
	for equipmentSetEquipmentRefID := range engine.Patch.EquipmentSetEquipmentRef {
		patchEquipmentSetEquipmentRefIDs = append(patchEquipmentSetEquipmentRefIDs, equipmentSetEquipmentRefID)
	}
	dedupedIDs := deduplicateEquipmentSetEquipmentRefIDs(stateEquipmentSetEquipmentRefIDs, patchEquipmentSetEquipmentRefIDs)
	equipmentSetEquipmentRefIDSlicePool.Put(stateEquipmentSetEquipmentRefIDs)
	equipmentSetEquipmentRefIDSlicePool.Put(patchEquipmentSetEquipmentRefIDs)
	return dedupedIDs
}
func (engine Engine) allItemBoundToRefIDs() []ItemBoundToRefID {
	stateItemBoundToRefIDs := itemBoundToRefIDSlicePool.Get().([]ItemBoundToRefID)[:0]
	for itemBoundToRefID := range engine.State.ItemBoundToRef {
		stateItemBoundToRefIDs = append(stateItemBoundToRefIDs, itemBoundToRefID)
	}
	patchItemBoundToRefIDs := itemBoundToRefIDSlicePool.Get().([]ItemBoundToRefID)[:0]
	for itemBoundToRefID := range engine.Patch.ItemBoundToRef {
		patchItemBoundToRefIDs = append(patchItemBoundToRefIDs, itemBoundToRefID)
	}
	dedupedIDs := deduplicateItemBoundToRefIDs(stateItemBoundToRefIDs, patchItemBoundToRefIDs)
	itemBoundToRefIDSlicePool.Put(stateItemBoundToRefIDs)
	itemBoundToRefIDSlicePool.Put(patchItemBoundToRefIDs)
	return dedupedIDs
}
func (engine Engine) allPlayerEquipmentSetRefIDs() []PlayerEquipmentSetRefID {
	statePlayerEquipmentSetRefIDs := playerEquipmentSetRefIDSlicePool.Get().([]PlayerEquipmentSetRefID)[:0]
	for playerEquipmentSetRefID := range engine.State.PlayerEquipmentSetRef {
		statePlayerEquipmentSetRefIDs = append(statePlayerEquipmentSetRefIDs, playerEquipmentSetRefID)
	}
	patchPlayerEquipmentSetRefIDs := playerEquipmentSetRefIDSlicePool.Get().([]PlayerEquipmentSetRefID)[:0]
	for playerEquipmentSetRefID := range engine.Patch.PlayerEquipmentSetRef {
		patchPlayerEquipmentSetRefIDs = append(patchPlayerEquipmentSetRefIDs, playerEquipmentSetRefID)
	}
	dedupedIDs := deduplicatePlayerEquipmentSetRefIDs(statePlayerEquipmentSetRefIDs, patchPlayerEquipmentSetRefIDs)
	playerEquipmentSetRefIDSlicePool.Put(statePlayerEquipmentSetRefIDs)
	playerEquipmentSetRefIDSlicePool.Put(patchPlayerEquipmentSetRefIDs)
	return dedupedIDs
}
func (engine Engine) allPlayerGuildMemberRefIDs() []PlayerGuildMemberRefID {
	statePlayerGuildMemberRefIDs := playerGuildMemberRefIDSlicePool.Get().([]PlayerGuildMemberRefID)[:0]
	for playerGuildMemberRefID := range engine.State.PlayerGuildMemberRef {
		statePlayerGuildMemberRefIDs = append(statePlayerGuildMemberRefIDs, playerGuildMemberRefID)
	}
	patchPlayerGuildMemberRefIDs := playerGuildMemberRefIDSlicePool.Get().([]PlayerGuildMemberRefID)[:0]
	for playerGuildMemberRefID := range engine.Patch.PlayerGuildMemberRef {
		patchPlayerGuildMemberRefIDs = append(patchPlayerGuildMemberRefIDs, playerGuildMemberRefID)
	}
	dedupedIDs := deduplicatePlayerGuildMemberRefIDs(statePlayerGuildMemberRefIDs, patchPlayerGuildMemberRefIDs)
	playerGuildMemberRefIDSlicePool.Put(statePlayerGuildMemberRefIDs)
	playerGuildMemberRefIDSlicePool.Put(patchPlayerGuildMemberRefIDs)
	return dedupedIDs
}
func (engine Engine) allPlayerTargetRefIDs() []PlayerTargetRefID {
	statePlayerTargetRefIDs := playerTargetRefIDSlicePool.Get().([]PlayerTargetRefID)[:0]
	for playerTargetRefID := range engine.State.PlayerTargetRef {
		statePlayerTargetRefIDs = append(statePlayerTargetRefIDs, playerTargetRefID)
	}
	patchPlayerTargetRefIDs := playerTargetRefIDSlicePool.Get().([]PlayerTargetRefID)[:0]
	for playerTargetRefID := range engine.Patch.PlayerTargetRef {
		patchPlayerTargetRefIDs = append(patchPlayerTargetRefIDs, playerTargetRefID)
	}
	dedupedIDs := deduplicatePlayerTargetRefIDs(statePlayerTargetRefIDs, patchPlayerTargetRefIDs)
	playerTargetRefIDSlicePool.Put(statePlayerTargetRefIDs)
	playerTargetRefIDSlicePool.Put(patchPlayerTargetRefIDs)
	return dedupedIDs
}
func (engine Engine) allPlayerTargetedByRefIDs() []PlayerTargetedByRefID {
	statePlayerTargetedByRefIDs := playerTargetedByRefIDSlicePool.Get().([]PlayerTargetedByRefID)[:0]
	for playerTargetedByRefID := range engine.State.PlayerTargetedByRef {
		statePlayerTargetedByRefIDs = append(statePlayerTargetedByRefIDs, playerTargetedByRefID)
	}
	patchPlayerTargetedByRefIDs := playerTargetedByRefIDSlicePool.Get().([]PlayerTargetedByRefID)[:0]
	for playerTargetedByRefID := range engine.Patch.PlayerTargetedByRef {
		patchPlayerTargetedByRefIDs = append(patchPlayerTargetedByRefIDs, playerTargetedByRefID)
	}
	dedupedIDs := deduplicatePlayerTargetedByRefIDs(statePlayerTargetedByRefIDs, patchPlayerTargetedByRefIDs)
	playerTargetedByRefIDSlicePool.Put(statePlayerTargetedByRefIDs)
	playerTargetedByRefIDSlicePool.Put(patchPlayerTargetedByRefIDs)
	return dedupedIDs
}

type treeFieldIdentifier int

const (
	attackEventIdentifier  treeFieldIdentifier = 1
	equipmentSetIdentifier treeFieldIdentifier = 2
	gearScoreIdentifier    treeFieldIdentifier = 3
	itemIdentifier         treeFieldIdentifier = 4
	playerIdentifier       treeFieldIdentifier = 5
	positionIdentifier     treeFieldIdentifier = 6
	zoneIdentifier         treeFieldIdentifier = 7
	zoneItemIdentifier     treeFieldIdentifier = 8

	attackEvent_targetIdentifier treeFieldIdentifier = 9

	equipmentSet_equipmentIdentifier treeFieldIdentifier = 10
	equipmentSet_nameIdentifier      treeFieldIdentifier = 11

	gearScore_levelIdentifier treeFieldIdentifier = 12
	gearScore_scoreIdentifier treeFieldIdentifier = 13

	item_boundToIdentifier   treeFieldIdentifier = 14
	item_gearScoreIdentifier treeFieldIdentifier = 15
	item_nameIdentifier      treeFieldIdentifier = 16
	item_originIdentifier    treeFieldIdentifier = 17

	player_actionIdentifier        treeFieldIdentifier = 18
	player_equipmentSetsIdentifier treeFieldIdentifier = 19
	player_gearScoreIdentifier     treeFieldIdentifier = 20
	player_guildMembersIdentifier  treeFieldIdentifier = 21
	player_itemsIdentifier         treeFieldIdentifier = 22
	player_positionIdentifier      treeFieldIdentifier = 23
	player_targetIdentifier        treeFieldIdentifier = 24
	player_targetedByIdentifier    treeFieldIdentifier = 25

	position_xIdentifier treeFieldIdentifier = 26
	position_yIdentifier treeFieldIdentifier = 27

	zone_interactablesIdentifier treeFieldIdentifier = 28
	zone_itemsIdentifier         treeFieldIdentifier = 29
	zone_playersIdentifier       treeFieldIdentifier = 30
	zone_tagsIdentifier          treeFieldIdentifier = 31

	zoneItem_itemIdentifier     treeFieldIdentifier = 32
	zoneItem_positionIdentifier treeFieldIdentifier = 33
)

func (t treeFieldIdentifier) toString() string {
	switch t {
	case attackEventIdentifier:
		return "attackEvent"
	case equipmentSetIdentifier:
		return "equipmentSet"
	case gearScoreIdentifier:
		return "gearScore"
	case itemIdentifier:
		return "item"
	case playerIdentifier:
		return "player"
	case positionIdentifier:
		return "position"
	case zoneIdentifier:
		return "zone"
	case zoneItemIdentifier:
		return "zoneItem"

	case attackEvent_targetIdentifier:
		return "target"

	case equipmentSet_equipmentIdentifier:
		return "equipment"
	case equipmentSet_nameIdentifier:
		return "name"

	case gearScore_levelIdentifier:
		return "level"
	case gearScore_scoreIdentifier:
		return "score"

	case item_boundToIdentifier:
		return "boundTo"
	case item_gearScoreIdentifier:
		return "gearScore"
	case item_nameIdentifier:
		return "name"
	case item_originIdentifier:
		return "origin"

	case player_actionIdentifier:
		return "action"
	case player_equipmentSetsIdentifier:
		return "equipmentSets"
	case player_gearScoreIdentifier:
		return "gearScore"
	case player_guildMembersIdentifier:
		return "guildMembers"
	case player_itemsIdentifier:
		return "items"
	case player_positionIdentifier:
		return "position"
	case player_targetIdentifier:
		return "target"
	case player_targetedByIdentifier:
		return "targetedBy"

	case position_xIdentifier:
		return "x"
	case position_yIdentifier:
		return "y"

	case zone_interactablesIdentifier:
		return "interactables"
	case zone_itemsIdentifier:
		return "items"
	case zone_playersIdentifier:
		return "players"
	case zone_tagsIdentifier:
		return "tags"

	case zoneItem_itemIdentifier:
		return "item"
	case zoneItem_positionIdentifier:
		return "position"

	default:
		panic(fmt.Sprintf("no string found for identifier <%d>", t))
	}
}

type segment struct {
	ID         int                 `json:"id"`
	Identifier treeFieldIdentifier `json:"identifier"`
	Kind       ElementKind         `json:"kind"`
	RefID      ComplexID           `json:"refID"`
}
type path []segment

func newPath() path {
	return make(path, 0)
}
func (p path) extendAndCopy(fieldIdentifier treeFieldIdentifier, id int, kind ElementKind, refID ComplexID) path {
	newPath := make(path, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, segment{id, fieldIdentifier, kind, refID})
	return newPath
}
func (p path) toJSONPath() string {
	jsonPath := "$"
	for _, seg := range p {
		jsonPath += "." + seg.Identifier.toString()
		if isSliceFieldIdentifier(seg.Identifier) {
			jsonPath += "[" + strconv.Itoa(seg.ID) + "]"
		}
	}
	return jsonPath
}
func isSliceFieldIdentifier(fieldIdentifier treeFieldIdentifier) bool {
	switch fieldIdentifier {
	case attackEventIdentifier:
		return true
	case equipmentSetIdentifier:
		return true
	case gearScoreIdentifier:
		return true
	case itemIdentifier:
		return true
	case playerIdentifier:
		return true
	case positionIdentifier:
		return true
	case zoneIdentifier:
		return true
	case zoneItemIdentifier:
		return true

	case equipmentSet_equipmentIdentifier:
		return true

	case player_actionIdentifier:
		return true
	case player_equipmentSetsIdentifier:
		return true

	case player_guildMembersIdentifier:
		return true
	case player_itemsIdentifier:
		return true

	case player_targetedByIdentifier:
		return true

	case zone_interactablesIdentifier:
		return true
	case zone_itemsIdentifier:
		return true
	case zone_playersIdentifier:
		return true

	}
	return false
}
func (_ref AttackEventTargetRef) IsSet() (AttackEventTargetRef, bool) {
	ref := _ref.attackEventTargetRef.engine.attackEventTargetRef(_ref.attackEventTargetRef.ID)
	return ref, ref.attackEventTargetRef.ID != AttackEventTargetRefID{}
}
func (_ref AttackEventTargetRef) Unset() {
	ref := _ref.attackEventTargetRef.engine.attackEventTargetRef(_ref.attackEventTargetRef.ID)
	if ref.attackEventTargetRef.OperationKind == OperationKindDelete {
		return
	}
	ref.attackEventTargetRef.engine.deleteAttackEventTargetRef(ref.attackEventTargetRef.ID)
	parent := ref.attackEventTargetRef.engine.AttackEvent(ref.attackEventTargetRef.ParentID).attackEvent
	if parent.OperationKind == OperationKindDelete {
		return
	}
	parent.Target = AttackEventTargetRefID{}
	parent.OperationKind = OperationKindUpdate
	parent.Meta.sign(parent.engine.broadcastingClientID)
	ref.attackEventTargetRef.engine.Patch.AttackEvent[parent.ID] = parent
}
func (_ref AttackEventTargetRef) Get() Player {
	ref := _ref.attackEventTargetRef.engine.attackEventTargetRef(_ref.attackEventTargetRef.ID)
	return ref.attackEventTargetRef.engine.Player(ref.attackEventTargetRef.ReferencedElementID)
}
func (_ref EquipmentSetEquipmentRef) Get() Item {
	ref := _ref.equipmentSetEquipmentRef.engine.equipmentSetEquipmentRef(_ref.equipmentSetEquipmentRef.ID)
	return ref.equipmentSetEquipmentRef.engine.Item(ref.equipmentSetEquipmentRef.ReferencedElementID)
}
func (_ref ItemBoundToRef) IsSet() (ItemBoundToRef, bool) {
	ref := _ref.itemBoundToRef.engine.itemBoundToRef(_ref.itemBoundToRef.ID)
	return ref, ref.itemBoundToRef.ID != ItemBoundToRefID{}
}
func (_ref ItemBoundToRef) Unset() {
	ref := _ref.itemBoundToRef.engine.itemBoundToRef(_ref.itemBoundToRef.ID)
	if ref.itemBoundToRef.OperationKind == OperationKindDelete {
		return
	}
	ref.itemBoundToRef.engine.deleteItemBoundToRef(ref.itemBoundToRef.ID)
	parent := ref.itemBoundToRef.engine.Item(ref.itemBoundToRef.ParentID).item
	if parent.OperationKind == OperationKindDelete {
		return
	}
	parent.BoundTo = ItemBoundToRefID{}
	parent.OperationKind = OperationKindUpdate
	parent.Meta.sign(parent.engine.broadcastingClientID)
	ref.itemBoundToRef.engine.Patch.Item[parent.ID] = parent
}
func (_ref ItemBoundToRef) Get() Player {
	ref := _ref.itemBoundToRef.engine.itemBoundToRef(_ref.itemBoundToRef.ID)
	return ref.itemBoundToRef.engine.Player(ref.itemBoundToRef.ReferencedElementID)
}
func (_ref PlayerEquipmentSetRef) Get() EquipmentSet {
	ref := _ref.playerEquipmentSetRef.engine.playerEquipmentSetRef(_ref.playerEquipmentSetRef.ID)
	return ref.playerEquipmentSetRef.engine.EquipmentSet(ref.playerEquipmentSetRef.ReferencedElementID)
}
func (_ref PlayerGuildMemberRef) Get() Player {
	ref := _ref.playerGuildMemberRef.engine.playerGuildMemberRef(_ref.playerGuildMemberRef.ID)
	return ref.playerGuildMemberRef.engine.Player(ref.playerGuildMemberRef.ReferencedElementID)
}
func (_ref PlayerTargetRef) IsSet() (PlayerTargetRef, bool) {
	ref := _ref.playerTargetRef.engine.playerTargetRef(_ref.playerTargetRef.ID)
	return ref, ref.playerTargetRef.ID != PlayerTargetRefID{}
}
func (_ref PlayerTargetRef) Unset() {
	ref := _ref.playerTargetRef.engine.playerTargetRef(_ref.playerTargetRef.ID)
	if ref.playerTargetRef.OperationKind == OperationKindDelete {
		return
	}
	ref.playerTargetRef.engine.deletePlayerTargetRef(ref.playerTargetRef.ID)
	parent := ref.playerTargetRef.engine.Player(ref.playerTargetRef.ParentID).player
	if parent.OperationKind == OperationKindDelete {
		return
	}
	parent.Target = PlayerTargetRefID{}
	parent.OperationKind = OperationKindUpdate
	parent.Meta.sign(parent.engine.broadcastingClientID)
	ref.playerTargetRef.engine.Patch.Player[parent.ID] = parent
}
func (_ref PlayerTargetRef) Get() anyOfPlayer_ZoneItemRef {
	ref := _ref.playerTargetRef.engine.playerTargetRef(_ref.playerTargetRef.ID)
	anyContainer := ref.playerTargetRef.engine.anyOfPlayer_ZoneItem(ref.playerTargetRef.ReferencedElementID)
	return anyOfPlayer_ZoneItemRef{
		anyOfPlayer_ZoneItem:        anyContainer.anyOfPlayer_ZoneItem,
		anyOfPlayer_ZoneItemWrapper: anyContainer,
	}
}
func (_ref PlayerTargetedByRef) Get() anyOfPlayer_ZoneItemRef {
	ref := _ref.playerTargetedByRef.engine.playerTargetedByRef(_ref.playerTargetedByRef.ID)
	anyContainer := ref.playerTargetedByRef.engine.anyOfPlayer_ZoneItem(ref.playerTargetedByRef.ReferencedElementID)
	return anyOfPlayer_ZoneItemRef{
		anyOfPlayer_ZoneItem:        anyContainer.anyOfPlayer_ZoneItem,
		anyOfPlayer_ZoneItemWrapper: anyContainer,
	}
}
func (engine *Engine) dereferenceAttackEventTargetRefs(playerID PlayerID) {
	allAttackEventTargetRefIDs := engine.allAttackEventTargetRefIDs()
	for _, refID := range allAttackEventTargetRefIDs {
		ref := engine.attackEventTargetRef(refID)

		if ref.attackEventTargetRef.ReferencedElementID == playerID {

			ref.Unset()
		}
	}
	attackEventTargetRefIDSlicePool.Put(allAttackEventTargetRefIDs)
}
func (engine *Engine) dereferenceEquipmentSetEquipmentRefs(itemID ItemID) {
	allEquipmentSetEquipmentRefIDs := engine.allEquipmentSetEquipmentRefIDs()
	for _, refID := range allEquipmentSetEquipmentRefIDs {
		ref := engine.equipmentSetEquipmentRef(refID)

		if ref.equipmentSetEquipmentRef.ReferencedElementID == itemID {
			parent := engine.EquipmentSet(ref.equipmentSetEquipmentRef.ParentID)
			parent.RemoveEquipment(itemID)

		}
	}
	equipmentSetEquipmentRefIDSlicePool.Put(allEquipmentSetEquipmentRefIDs)
}
func (engine *Engine) dereferenceItemBoundToRefs(playerID PlayerID) {
	allItemBoundToRefIDs := engine.allItemBoundToRefIDs()
	for _, refID := range allItemBoundToRefIDs {
		ref := engine.itemBoundToRef(refID)

		if ref.itemBoundToRef.ReferencedElementID == playerID {

			ref.Unset()
		}
	}
	itemBoundToRefIDSlicePool.Put(allItemBoundToRefIDs)
}
func (engine *Engine) dereferencePlayerEquipmentSetRefs(equipmentSetID EquipmentSetID) {
	allPlayerEquipmentSetRefIDs := engine.allPlayerEquipmentSetRefIDs()
	for _, refID := range allPlayerEquipmentSetRefIDs {
		ref := engine.playerEquipmentSetRef(refID)

		if ref.playerEquipmentSetRef.ReferencedElementID == equipmentSetID {
			parent := engine.Player(ref.playerEquipmentSetRef.ParentID)
			parent.RemoveEquipmentSets(equipmentSetID)

		}
	}
	playerEquipmentSetRefIDSlicePool.Put(allPlayerEquipmentSetRefIDs)
}
func (engine *Engine) dereferencePlayerGuildMemberRefs(playerID PlayerID) {
	allPlayerGuildMemberRefIDs := engine.allPlayerGuildMemberRefIDs()
	for _, refID := range allPlayerGuildMemberRefIDs {
		ref := engine.playerGuildMemberRef(refID)

		if ref.playerGuildMemberRef.ReferencedElementID == playerID {
			parent := engine.Player(ref.playerGuildMemberRef.ParentID)
			parent.RemoveGuildMembers(playerID)

		}
	}
	playerGuildMemberRefIDSlicePool.Put(allPlayerGuildMemberRefIDs)
}
func (engine *Engine) dereferencePlayerTargetRefsPlayer(playerID PlayerID) {
	allPlayerTargetRefIDs := engine.allPlayerTargetRefIDs()
	for _, refID := range allPlayerTargetRefIDs {
		ref := engine.playerTargetRef(refID)
		anyContainer := ref.Get()
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind != ElementKindPlayer {
			continue
		}
		if PlayerID(anyContainer.anyOfPlayer_ZoneItem.ChildID) == playerID {

			ref.Unset()
		}
	}
	playerTargetRefIDSlicePool.Put(allPlayerTargetRefIDs)
}
func (engine *Engine) dereferencePlayerTargetRefsZoneItem(zoneItemID ZoneItemID) {
	allPlayerTargetRefIDs := engine.allPlayerTargetRefIDs()
	for _, refID := range allPlayerTargetRefIDs {
		ref := engine.playerTargetRef(refID)
		anyContainer := ref.Get()
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind != ElementKindZoneItem {
			continue
		}
		if ZoneItemID(anyContainer.anyOfPlayer_ZoneItem.ChildID) == zoneItemID {

			ref.Unset()
		}
	}
	playerTargetRefIDSlicePool.Put(allPlayerTargetRefIDs)
}
func (engine *Engine) dereferencePlayerTargetedByRefsPlayer(playerID PlayerID) {
	allPlayerTargetedByRefIDs := engine.allPlayerTargetedByRefIDs()
	for _, refID := range allPlayerTargetedByRefIDs {
		ref := engine.playerTargetedByRef(refID)
		anyContainer := ref.Get()
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind != ElementKindPlayer {
			continue
		}
		if PlayerID(anyContainer.anyOfPlayer_ZoneItem.ChildID) == playerID {
			parent := engine.Player(ref.playerTargetedByRef.ParentID)
			parent.RemoveTargetedByPlayer(playerID)

		}
	}
	playerTargetedByRefIDSlicePool.Put(allPlayerTargetedByRefIDs)
}
func (engine *Engine) dereferencePlayerTargetedByRefsZoneItem(zoneItemID ZoneItemID) {
	allPlayerTargetedByRefIDs := engine.allPlayerTargetedByRefIDs()
	for _, refID := range allPlayerTargetedByRefIDs {
		ref := engine.playerTargetedByRef(refID)
		anyContainer := ref.Get()
		if anyContainer.anyOfPlayer_ZoneItem.ElementKind != ElementKindZoneItem {
			continue
		}
		if ZoneItemID(anyContainer.anyOfPlayer_ZoneItem.ChildID) == zoneItemID {
			parent := engine.Player(ref.playerTargetedByRef.ParentID)
			parent.RemoveTargetedByZoneItem(zoneItemID)

		}
	}
	playerTargetedByRefIDSlicePool.Put(allPlayerTargetedByRefIDs)
}
func (_equipmentSet EquipmentSet) RemoveEquipment(equipmentToRemove ItemID) EquipmentSet {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	if equipmentSet.equipmentSet.OperationKind == OperationKindDelete {
		return equipmentSet
	}
	for i, complexID := range equipmentSet.equipmentSet.Equipment {
		if ItemID(complexID.ChildID) != equipmentToRemove {
			continue
		}
		equipmentSet.equipmentSet.Equipment[i] = equipmentSet.equipmentSet.Equipment[len(equipmentSet.equipmentSet.Equipment)-1]
		equipmentSet.equipmentSet.Equipment[len(equipmentSet.equipmentSet.Equipment)-1] = EquipmentSetEquipmentRefID{}
		equipmentSet.equipmentSet.Equipment = equipmentSet.equipmentSet.Equipment[:len(equipmentSet.equipmentSet.Equipment)-1]
		equipmentSet.equipmentSet.engine.deleteEquipmentSetEquipmentRef(complexID)
		equipmentSet.equipmentSet.OperationKind = OperationKindUpdate
		equipmentSet.equipmentSet.Meta.sign(equipmentSet.equipmentSet.engine.broadcastingClientID)
		equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID] = equipmentSet.equipmentSet
		break
	}
	return equipmentSet
}
func (_player Player) RemoveAction(actionToRemove AttackEventID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	for i, attackEventID := range player.player.Action {
		if attackEventID != actionToRemove {
			continue
		}
		player.player.Action[i] = player.player.Action[len(player.player.Action)-1]
		player.player.Action[len(player.player.Action)-1] = 0
		player.player.Action = player.player.Action[:len(player.player.Action)-1]
		player.player.engine.deleteAttackEvent(attackEventID)
		player.player.OperationKind = OperationKindUpdate
		player.player.Meta.sign(player.player.engine.broadcastingClientID)
		player.player.engine.Patch.Player[player.player.ID] = player.player
		break
	}
	return player
}
func (_player Player) RemoveEquipmentSets(equipmentSetToRemove EquipmentSetID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	for i, complexID := range player.player.EquipmentSets {
		if EquipmentSetID(complexID.ChildID) != equipmentSetToRemove {
			continue
		}
		player.player.EquipmentSets[i] = player.player.EquipmentSets[len(player.player.EquipmentSets)-1]
		player.player.EquipmentSets[len(player.player.EquipmentSets)-1] = PlayerEquipmentSetRefID{}
		player.player.EquipmentSets = player.player.EquipmentSets[:len(player.player.EquipmentSets)-1]
		player.player.engine.deletePlayerEquipmentSetRef(complexID)
		player.player.OperationKind = OperationKindUpdate
		player.player.Meta.sign(player.player.engine.broadcastingClientID)
		player.player.engine.Patch.Player[player.player.ID] = player.player
		break
	}
	return player
}
func (_player Player) RemoveGuildMembers(guildMemberToRemove PlayerID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	for i, complexID := range player.player.GuildMembers {
		if PlayerID(complexID.ChildID) != guildMemberToRemove {
			continue
		}
		player.player.GuildMembers[i] = player.player.GuildMembers[len(player.player.GuildMembers)-1]
		player.player.GuildMembers[len(player.player.GuildMembers)-1] = PlayerGuildMemberRefID{}
		player.player.GuildMembers = player.player.GuildMembers[:len(player.player.GuildMembers)-1]
		player.player.engine.deletePlayerGuildMemberRef(complexID)
		player.player.OperationKind = OperationKindUpdate
		player.player.Meta.sign(player.player.engine.broadcastingClientID)
		player.player.engine.Patch.Player[player.player.ID] = player.player
		break
	}
	return player
}
func (_player Player) RemoveItems(itemToRemove ItemID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	for i, itemID := range player.player.Items {
		if itemID != itemToRemove {
			continue
		}
		player.player.Items[i] = player.player.Items[len(player.player.Items)-1]
		player.player.Items[len(player.player.Items)-1] = 0
		player.player.Items = player.player.Items[:len(player.player.Items)-1]
		player.player.engine.deleteItem(itemID)
		player.player.OperationKind = OperationKindUpdate
		player.player.Meta.sign(player.player.engine.broadcastingClientID)
		player.player.engine.Patch.Player[player.player.ID] = player.player
		break
	}
	return player
}
func (_player Player) RemoveTargetedByPlayer(playerToRemove PlayerID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	for i, complexID := range player.player.TargetedBy {
		if PlayerID(complexID.ChildID) != playerToRemove {
			continue
		}
		player.player.TargetedBy[i] = player.player.TargetedBy[len(player.player.TargetedBy)-1]
		player.player.TargetedBy[len(player.player.TargetedBy)-1] = PlayerTargetedByRefID{}
		player.player.TargetedBy = player.player.TargetedBy[:len(player.player.TargetedBy)-1]
		player.player.engine.deletePlayerTargetedByRef(complexID)
		player.player.OperationKind = OperationKindUpdate
		player.player.Meta.sign(player.player.engine.broadcastingClientID)
		player.player.engine.Patch.Player[player.player.ID] = player.player
		break
	}
	return player
}
func (_player Player) RemoveTargetedByZoneItem(zoneItemToRemove ZoneItemID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	for i, complexID := range player.player.TargetedBy {
		if ZoneItemID(complexID.ChildID) != zoneItemToRemove {
			continue
		}
		player.player.TargetedBy[i] = player.player.TargetedBy[len(player.player.TargetedBy)-1]
		player.player.TargetedBy[len(player.player.TargetedBy)-1] = PlayerTargetedByRefID{}
		player.player.TargetedBy = player.player.TargetedBy[:len(player.player.TargetedBy)-1]
		player.player.engine.deletePlayerTargetedByRef(complexID)
		player.player.OperationKind = OperationKindUpdate
		player.player.Meta.sign(player.player.engine.broadcastingClientID)
		player.player.engine.Patch.Player[player.player.ID] = player.player
		break
	}
	return player
}
func (_zone Zone) RemoveInteractablesItem(itemToRemove ItemID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}
	for i, complexID := range zone.zone.Interactables {
		if ItemID(complexID.ChildID) != itemToRemove {
			continue
		}
		zone.zone.Interactables[i] = zone.zone.Interactables[len(zone.zone.Interactables)-1]
		zone.zone.Interactables[len(zone.zone.Interactables)-1] = AnyOfItem_Player_ZoneItemID{}
		zone.zone.Interactables = zone.zone.Interactables[:len(zone.zone.Interactables)-1]
		zone.zone.engine.deleteAnyOfItem_Player_ZoneItem(complexID, true)
		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.Meta.sign(zone.zone.engine.broadcastingClientID)
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
		break
	}
	return zone
}
func (_zone Zone) RemoveInteractablesPlayer(playerToRemove PlayerID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}
	for i, complexID := range zone.zone.Interactables {
		if PlayerID(complexID.ChildID) != playerToRemove {
			continue
		}
		zone.zone.Interactables[i] = zone.zone.Interactables[len(zone.zone.Interactables)-1]
		zone.zone.Interactables[len(zone.zone.Interactables)-1] = AnyOfItem_Player_ZoneItemID{}
		zone.zone.Interactables = zone.zone.Interactables[:len(zone.zone.Interactables)-1]
		zone.zone.engine.deleteAnyOfItem_Player_ZoneItem(complexID, true)
		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.Meta.sign(zone.zone.engine.broadcastingClientID)
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
		break
	}
	return zone
}
func (_zone Zone) RemoveInteractablesZoneItem(zoneItemToRemove ZoneItemID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}
	for i, complexID := range zone.zone.Interactables {
		if ZoneItemID(complexID.ChildID) != zoneItemToRemove {
			continue
		}
		zone.zone.Interactables[i] = zone.zone.Interactables[len(zone.zone.Interactables)-1]
		zone.zone.Interactables[len(zone.zone.Interactables)-1] = AnyOfItem_Player_ZoneItemID{}
		zone.zone.Interactables = zone.zone.Interactables[:len(zone.zone.Interactables)-1]
		zone.zone.engine.deleteAnyOfItem_Player_ZoneItem(complexID, true)
		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.Meta.sign(zone.zone.engine.broadcastingClientID)
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
		break
	}
	return zone
}
func (_zone Zone) RemoveItems(itemToRemove ZoneItemID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}
	for i, zoneItemID := range zone.zone.Items {
		if zoneItemID != itemToRemove {
			continue
		}
		zone.zone.Items[i] = zone.zone.Items[len(zone.zone.Items)-1]
		zone.zone.Items[len(zone.zone.Items)-1] = 0
		zone.zone.Items = zone.zone.Items[:len(zone.zone.Items)-1]
		zone.zone.engine.deleteZoneItem(zoneItemID)
		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.Meta.sign(zone.zone.engine.broadcastingClientID)
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
		break
	}
	return zone
}
func (_zone Zone) RemovePlayers(playerToRemove PlayerID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}
	for i, playerID := range zone.zone.Players {
		if playerID != playerToRemove {
			continue
		}
		zone.zone.Players[i] = zone.zone.Players[len(zone.zone.Players)-1]
		zone.zone.Players[len(zone.zone.Players)-1] = 0
		zone.zone.Players = zone.zone.Players[:len(zone.zone.Players)-1]
		zone.zone.engine.deletePlayer(playerID)
		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.Meta.sign(zone.zone.engine.broadcastingClientID)
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
		break
	}
	return zone
}
func (_zone Zone) RemoveTags(tagToRemove string) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}
	for i, valID := range zone.zone.Tags {
		if zone.zone.engine.stringValue(valID).Value != tagToRemove {
			continue
		}
		zone.zone.Tags[i] = zone.zone.Tags[len(zone.zone.Tags)-1]
		zone.zone.Tags[len(zone.zone.Tags)-1] = 0
		zone.zone.Tags = zone.zone.Tags[:len(zone.zone.Tags)-1]

		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.Meta.sign(zone.zone.engine.broadcastingClientID)
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
		break
	}
	return zone
}
func (engine *Engine) setBoolValue(id BoolValueID, val bool) {
	boolValue := engine.boolValue(id)
	if boolValue.OperationKind == OperationKindDelete {
		return
	}
	if boolValue.Value == val {
		return
	}
	boolValue.Value = val
	boolValue.OperationKind = OperationKindUpdate
	boolValue.Meta.sign(boolValue.engine.broadcastingClientID)
	engine.Patch.BoolValue[id] = boolValue
}
func (engine *Engine) setFloatValue(id FloatValueID, val float64) {
	floatValue := engine.floatValue(id)
	if floatValue.OperationKind == OperationKindDelete {
		return
	}
	if floatValue.Value == val {
		return
	}
	floatValue.Value = val
	floatValue.OperationKind = OperationKindUpdate
	floatValue.Meta.sign(floatValue.engine.broadcastingClientID)
	engine.Patch.FloatValue[id] = floatValue
}
func (engine *Engine) setIntValue(id IntValueID, val int64) {
	intValue := engine.intValue(id)
	if intValue.OperationKind == OperationKindDelete {
		return
	}
	if intValue.Value == val {
		return
	}
	intValue.Value = val
	intValue.OperationKind = OperationKindUpdate
	intValue.Meta.sign(intValue.engine.broadcastingClientID)
	engine.Patch.IntValue[id] = intValue
}
func (engine *Engine) setStringValue(id StringValueID, val string) {
	stringValue := engine.stringValue(id)
	if stringValue.OperationKind == OperationKindDelete {
		return
	}
	if stringValue.Value == val {
		return
	}
	stringValue.Value = val
	stringValue.OperationKind = OperationKindUpdate
	stringValue.Meta.sign(stringValue.engine.broadcastingClientID)
	engine.Patch.StringValue[id] = stringValue
}
func (_equipmentSet EquipmentSet) SetName(newName string) EquipmentSet {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	if equipmentSet.equipmentSet.OperationKind == OperationKindDelete {
		return equipmentSet
	}
	equipmentSet.equipmentSet.engine.setStringValue(equipmentSet.equipmentSet.Name, newName)
	return equipmentSet
}
func (_gearScore GearScore) SetLevel(newLevel int64) GearScore {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	if gearScore.gearScore.OperationKind == OperationKindDelete {
		return gearScore
	}
	gearScore.gearScore.engine.setIntValue(gearScore.gearScore.Level, newLevel)
	return gearScore
}
func (_gearScore GearScore) SetScore(newScore int64) GearScore {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	if gearScore.gearScore.OperationKind == OperationKindDelete {
		return gearScore
	}
	gearScore.gearScore.engine.setIntValue(gearScore.gearScore.Score, newScore)
	return gearScore
}
func (_item Item) SetName(newName string) Item {
	item := _item.item.engine.Item(_item.item.ID)
	if item.item.OperationKind == OperationKindDelete {
		return item
	}
	item.item.engine.setStringValue(item.item.Name, newName)
	return item
}
func (_position Position) SetX(newX float64) Position {
	position := _position.position.engine.Position(_position.position.ID)
	if position.position.OperationKind == OperationKindDelete {
		return position
	}
	position.position.engine.setFloatValue(position.position.X, newX)
	return position
}
func (_position Position) SetY(newY float64) Position {
	position := _position.position.engine.Position(_position.position.ID)
	if position.position.OperationKind == OperationKindDelete {
		return position
	}
	position.position.engine.setFloatValue(position.position.Y, newY)
	return position
}
func (_attackEvent AttackEvent) SetTarget(playerID PlayerID) AttackEvent {
	attackEvent := _attackEvent.attackEvent.engine.AttackEvent(_attackEvent.attackEvent.ID)
	if attackEvent.attackEvent.OperationKind == OperationKindDelete {
		return attackEvent
	}
	if attackEvent.attackEvent.engine.Player(playerID).player.OperationKind == OperationKindDelete {
		return attackEvent
	}
	if PlayerID(attackEvent.attackEvent.Target.ChildID) == playerID {
		return attackEvent
	}
	if attackEvent.attackEvent.Target != (AttackEventTargetRefID{}) {
		attackEvent.attackEvent.engine.deleteAttackEventTargetRef(attackEvent.attackEvent.Target)
	}

	ref := attackEvent.attackEvent.engine.createAttackEventTargetRef(attackEvent.attackEvent.Path, attackEvent_targetIdentifier, playerID, attackEvent.attackEvent.ID)
	attackEvent.attackEvent.Target = ref.ID
	attackEvent.attackEvent.OperationKind = OperationKindUpdate
	attackEvent.attackEvent.Meta.sign(attackEvent.attackEvent.engine.broadcastingClientID)
	attackEvent.attackEvent.engine.Patch.AttackEvent[attackEvent.attackEvent.ID] = attackEvent.attackEvent
	return attackEvent
}
func (_item Item) SetBoundTo(playerID PlayerID) Item {
	item := _item.item.engine.Item(_item.item.ID)
	if item.item.OperationKind == OperationKindDelete {
		return item
	}
	if item.item.engine.Player(playerID).player.OperationKind == OperationKindDelete {
		return item
	}
	if PlayerID(item.item.BoundTo.ChildID) == playerID {
		return item
	}
	if item.item.BoundTo != (ItemBoundToRefID{}) {
		item.item.engine.deleteItemBoundToRef(item.item.BoundTo)
	}

	ref := item.item.engine.createItemBoundToRef(item.item.Path, item_boundToIdentifier, playerID, item.item.ID)
	item.item.BoundTo = ref.ID
	item.item.OperationKind = OperationKindUpdate
	item.item.Meta.sign(item.item.engine.broadcastingClientID)
	item.item.engine.Patch.Item[item.item.ID] = item.item
	return item
}
func (_player Player) SetTargetPlayer(playerID PlayerID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	if player.player.engine.Player(playerID).player.OperationKind == OperationKindDelete {
		return player
	}
	if PlayerID(player.player.Target.ChildID) == playerID {
		return player
	}
	if player.player.Target != (PlayerTargetRefID{}) {
		player.player.engine.deletePlayerTargetRef(player.player.Target)
	}
	anyContainer := player.player.engine.createAnyOfPlayer_ZoneItem(int(player.player.ID), int(playerID), ElementKindPlayer, player.player.Path, player_targetIdentifier)
	ref := player.player.engine.createPlayerTargetRef(player.player.Path, player_targetIdentifier, anyContainer.anyOfPlayer_ZoneItem.ID, player.player.ID, ElementKindPlayer, int(playerID))
	player.player.Target = ref.ID
	player.player.OperationKind = OperationKindUpdate
	player.player.Meta.sign(player.player.engine.broadcastingClientID)
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return player
}
func (_player Player) SetTargetZoneItem(zoneItemID ZoneItemID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	if player.player.engine.ZoneItem(zoneItemID).zoneItem.OperationKind == OperationKindDelete {
		return player
	}
	if ZoneItemID(player.player.Target.ChildID) == zoneItemID {
		return player
	}
	if player.player.Target != (PlayerTargetRefID{}) {
		player.player.engine.deletePlayerTargetRef(player.player.Target)
	}
	anyContainer := player.player.engine.createAnyOfPlayer_ZoneItem(int(player.player.ID), int(zoneItemID), ElementKindZoneItem, player.player.Path, player_targetIdentifier)
	ref := player.player.engine.createPlayerTargetRef(player.player.Path, player_targetIdentifier, anyContainer.anyOfPlayer_ZoneItem.ID, player.player.ID, ElementKindZoneItem, int(zoneItemID))
	player.player.Target = ref.ID
	player.player.OperationKind = OperationKindUpdate
	player.player.Meta.sign(player.player.engine.broadcastingClientID)
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return player
}

// easyjson:skip
type ComplexID struct {
	Field      treeFieldIdentifier `json:"field"`
	ParentID   int                 `json:"parentID"`
	ChildID    int                 `json:"childID"`
	IsMediator bool                `json:"isMediator"`
}

var (
	complexIDStructCache = make(map[string]ComplexID)
	ComplexIDStringCache = make(map[ComplexID][]byte)
	complexIDZeroString  = []byte("0-0-0-0")
)

func (c ComplexID) MarshalJSON() ([]byte, error) {
	if cachedString, ok := ComplexIDStringCache[c]; ok {
		return cachedString, nil
	}
	var isMediatorBin int
	if c.IsMediator {
		isMediatorBin = 1
	}
	newS := []byte(fmt.Sprintf("\"%d-%d-%d-%d\"", c.Field, c.ParentID, c.ChildID, isMediatorBin))
	ComplexIDStringCache[c] = newS
	return newS, nil
}
func (c *ComplexID) UnmarshalJSON(s []byte) error {
	if bytes.Equal(s, complexIDZeroString) {
		return nil
	}
	asString := string(s)
	if cachedID, ok := complexIDStructCache[asString]; ok {
		c.Field = cachedID.Field
		c.ParentID = cachedID.ParentID
		c.ChildID = cachedID.ChildID
		c.IsMediator = cachedID.IsMediator
		return nil
	}
	idSegments := bytes.Split(s[1:len(s)-1], []byte{'-'})
	ident, _ := strconv.Atoi(string(idSegments[0]))
	c.Field = treeFieldIdentifier(ident)
	c.ParentID, _ = strconv.Atoi(string(idSegments[1]))
	c.ChildID, _ = strconv.Atoi(string(idSegments[2]))
	isMediatorBin, _ := strconv.Atoi(string(idSegments[3]))
	if isMediatorBin == 1 {
		c.IsMediator = true
	}
	complexIDStructCache[asString] = *c
	return nil
}
func (x AttackEventTargetRefID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *AttackEventTargetRefID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = AttackEventTargetRefID(temp)
	return nil
}
func (x EquipmentSetEquipmentRefID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *EquipmentSetEquipmentRefID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = EquipmentSetEquipmentRefID(temp)
	return nil
}
func (x ItemBoundToRefID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *ItemBoundToRefID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = ItemBoundToRefID(temp)
	return nil
}
func (x PlayerEquipmentSetRefID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *PlayerEquipmentSetRefID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = PlayerEquipmentSetRefID(temp)
	return nil
}
func (x PlayerGuildMemberRefID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *PlayerGuildMemberRefID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = PlayerGuildMemberRefID(temp)
	return nil
}
func (x PlayerTargetRefID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *PlayerTargetRefID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = PlayerTargetRefID(temp)
	return nil
}
func (x PlayerTargetedByRefID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *PlayerTargetedByRefID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = PlayerTargetedByRefID(temp)
	return nil
}
func (x AnyOfPlayer_PositionID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *AnyOfPlayer_PositionID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = AnyOfPlayer_PositionID(temp)
	return nil
}
func (x AnyOfPlayer_ZoneItemID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *AnyOfPlayer_ZoneItemID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = AnyOfPlayer_ZoneItemID(temp)
	return nil
}
func (x AnyOfItem_Player_ZoneItemID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *AnyOfItem_Player_ZoneItemID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = AnyOfItem_Player_ZoneItemID(temp)
	return nil
}

type BoolValueID int
type FloatValueID int
type IntValueID int
type StringValueID int
type AttackEventID int
type EquipmentSetID int
type GearScoreID int
type ItemID int
type PlayerID int
type PositionID int
type ZoneID int
type ZoneItemID int
type AttackEventTargetRefID ComplexID
type EquipmentSetEquipmentRefID ComplexID
type ItemBoundToRefID ComplexID
type PlayerEquipmentSetRefID ComplexID
type PlayerGuildMemberRefID ComplexID
type PlayerTargetRefID ComplexID
type PlayerTargetedByRefID ComplexID
type AnyOfPlayer_PositionID ComplexID
type AnyOfPlayer_ZoneItemID ComplexID
type AnyOfItem_Player_ZoneItemID ComplexID
type State struct {
	BoolValue   map[BoolValueID]boolValue     `json:"boolValue"`
	FloatValue  map[FloatValueID]floatValue   `json:"floatValue"`
	IntValue    map[IntValueID]intValue       `json:"intValue"`
	StringValue map[StringValueID]stringValue `json:"stringValue"`

	AttackEvent map[AttackEventID]attackEventCore `json:"attackEvent"`

	EquipmentSet map[EquipmentSetID]equipmentSetCore `json:"equipmentSet"`

	GearScore map[GearScoreID]gearScoreCore `json:"gearScore"`

	Item map[ItemID]itemCore `json:"item"`

	Player map[PlayerID]playerCore `json:"player"`

	Position map[PositionID]positionCore `json:"position"`

	Zone map[ZoneID]zoneCore `json:"zone"`

	ZoneItem map[ZoneItemID]zoneItemCore `json:"zoneItem"`

	AttackEventTargetRef map[AttackEventTargetRefID]attackEventTargetRefCore `json:"attackEventTargetRef"`

	EquipmentSetEquipmentRef map[EquipmentSetEquipmentRefID]equipmentSetEquipmentRefCore `json:"equipmentSetEquipmentRef"`

	ItemBoundToRef map[ItemBoundToRefID]itemBoundToRefCore `json:"itemBoundToRef"`

	PlayerEquipmentSetRef map[PlayerEquipmentSetRefID]playerEquipmentSetRefCore `json:"playerEquipmentSetRef"`

	PlayerGuildMemberRef map[PlayerGuildMemberRefID]playerGuildMemberRefCore `json:"playerGuildMemberRef"`

	PlayerTargetRef map[PlayerTargetRefID]playerTargetRefCore `json:"playerTargetRef"`

	PlayerTargetedByRef map[PlayerTargetedByRefID]playerTargetedByRefCore `json:"playerTargetedByRef"`

	AnyOfPlayer_Position map[AnyOfPlayer_PositionID]anyOfPlayer_PositionCore `json:"anyOfPlayer_Position"`

	AnyOfPlayer_ZoneItem map[AnyOfPlayer_ZoneItemID]anyOfPlayer_ZoneItemCore `json:"anyOfPlayer_ZoneItem"`

	AnyOfItem_Player_ZoneItem map[AnyOfItem_Player_ZoneItemID]anyOfItem_Player_ZoneItemCore `json:"anyOfItem_Player_ZoneItem"`
}

func newState() *State {
	return &State{
		BoolValue:   make(map[BoolValueID]boolValue),
		FloatValue:  make(map[FloatValueID]floatValue),
		IntValue:    make(map[IntValueID]intValue),
		StringValue: make(map[StringValueID]stringValue),

		AttackEvent:  make(map[AttackEventID]attackEventCore),
		EquipmentSet: make(map[EquipmentSetID]equipmentSetCore),
		GearScore:    make(map[GearScoreID]gearScoreCore),
		Item:         make(map[ItemID]itemCore),
		Player:       make(map[PlayerID]playerCore),
		Position:     make(map[PositionID]positionCore),
		Zone:         make(map[ZoneID]zoneCore),
		ZoneItem:     make(map[ZoneItemID]zoneItemCore),

		AttackEventTargetRef:     make(map[AttackEventTargetRefID]attackEventTargetRefCore),
		EquipmentSetEquipmentRef: make(map[EquipmentSetEquipmentRefID]equipmentSetEquipmentRefCore),
		ItemBoundToRef:           make(map[ItemBoundToRefID]itemBoundToRefCore),
		PlayerEquipmentSetRef:    make(map[PlayerEquipmentSetRefID]playerEquipmentSetRefCore),
		PlayerGuildMemberRef:     make(map[PlayerGuildMemberRefID]playerGuildMemberRefCore),
		PlayerTargetRef:          make(map[PlayerTargetRefID]playerTargetRefCore),
		PlayerTargetedByRef:      make(map[PlayerTargetedByRefID]playerTargetedByRefCore),

		AnyOfPlayer_Position:      make(map[AnyOfPlayer_PositionID]anyOfPlayer_PositionCore),
		AnyOfPlayer_ZoneItem:      make(map[AnyOfPlayer_ZoneItemID]anyOfPlayer_ZoneItemCore),
		AnyOfItem_Player_ZoneItem: make(map[AnyOfItem_Player_ZoneItemID]anyOfItem_Player_ZoneItemCore),
	}
}

type metaData struct {
	BroadcastedBy string `json:"broadcastedBy"`
	TouchedByMany bool   `json:"touchedByMany"`
}

func (m *metaData) unsign() {
	m.BroadcastedBy = ""
	m.TouchedByMany = false
}
func (m *metaData) sign(clientID string) {
	if clientID == "" {
		return
	}
	if m.TouchedByMany {
		return
	}
	if m.BroadcastedBy == "" {
		m.BroadcastedBy = clientID
		return
	}
	if m.BroadcastedBy == clientID {
		return
	}
	m.BroadcastedBy = ""
	m.TouchedByMany = true
}

type boolValue struct {
	ID BoolValueID `json:"id"`

	Value bool `json:"value"`

	OperationKind OperationKind `json:"operationKind"`

	JSONPath string   `json:"jsonPath"`
	Path     path     `json:"path"`
	Meta     metaData `json:"meta"`
	engine   *Engine
}
type floatValue struct {
	ID FloatValueID `json:"id"`

	Value float64 `json:"value"`

	OperationKind OperationKind `json:"operationKind"`

	JSONPath string   `json:"jsonPath"`
	Path     path     `json:"path"`
	Meta     metaData `json:"meta"`
	engine   *Engine
}
type intValue struct {
	ID IntValueID `json:"id"`

	Value int64 `json:"value"`

	OperationKind OperationKind `json:"operationKind"`

	JSONPath string   `json:"jsonPath"`
	Path     path     `json:"path"`
	Meta     metaData `json:"meta"`
	engine   *Engine
}
type stringValue struct {
	ID StringValueID `json:"id"`

	Value string `json:"value"`

	OperationKind OperationKind `json:"operationKind"`

	JSONPath string   `json:"jsonPath"`
	Path     path     `json:"path"`
	Meta     metaData `json:"meta"`
	engine   *Engine
}
type attackEventCore struct {
	ID AttackEventID `json:"id"`

	Target AttackEventTargetRefID `json:"target"`

	OperationKind OperationKind `json:"operationKind"`

	HasParent bool `json:"hasParent"`

	JSONPath string   `json:"jsonPath"`
	Path     path     `json:"path"`
	Meta     metaData `json:"meta"`
	engine   *Engine
}
type AttackEvent struct {
	attackEvent attackEventCore
}
type equipmentSetCore struct {
	ID EquipmentSetID `json:"id"`

	Equipment []EquipmentSetEquipmentRefID `json:"equipment"`

	Name StringValueID `json:"name"`

	OperationKind OperationKind `json:"operationKind"`

	HasParent bool `json:"hasParent"`

	JSONPath string   `json:"jsonPath"`
	Path     path     `json:"path"`
	Meta     metaData `json:"meta"`
	engine   *Engine
}
type EquipmentSet struct {
	equipmentSet equipmentSetCore
}
type gearScoreCore struct {
	ID GearScoreID `json:"id"`

	Level IntValueID `json:"level"`

	Score IntValueID `json:"score"`

	OperationKind OperationKind `json:"operationKind"`

	HasParent bool `json:"hasParent"`

	JSONPath string   `json:"jsonPath"`
	Path     path     `json:"path"`
	Meta     metaData `json:"meta"`
	engine   *Engine
}
type GearScore struct {
	gearScore gearScoreCore
}
type itemCore struct {
	ID ItemID `json:"id"`

	BoundTo ItemBoundToRefID `json:"boundTo"`

	GearScore GearScoreID `json:"gearScore"`

	Name StringValueID `json:"name"`

	Origin AnyOfPlayer_PositionID `json:"origin"`

	OperationKind OperationKind `json:"operationKind"`

	HasParent bool `json:"hasParent"`

	JSONPath string   `json:"jsonPath"`
	Path     path     `json:"path"`
	Meta     metaData `json:"meta"`
	engine   *Engine
}
type Item struct {
	item itemCore
}
type playerCore struct {
	ID PlayerID `json:"id"`

	Action []AttackEventID `json:"action"`

	EquipmentSets []PlayerEquipmentSetRefID `json:"equipmentSets"`

	GearScore GearScoreID `json:"gearScore"`

	GuildMembers []PlayerGuildMemberRefID `json:"guildMembers"`

	Items []ItemID `json:"items"`

	Position PositionID `json:"position"`

	Target PlayerTargetRefID `json:"target"`

	TargetedBy []PlayerTargetedByRefID `json:"targetedBy"`

	OperationKind OperationKind `json:"operationKind"`

	HasParent bool `json:"hasParent"`

	JSONPath string   `json:"jsonPath"`
	Path     path     `json:"path"`
	Meta     metaData `json:"meta"`
	engine   *Engine
}
type Player struct {
	player playerCore
}
type positionCore struct {
	ID PositionID `json:"id"`

	X FloatValueID `json:"x"`

	Y FloatValueID `json:"y"`

	OperationKind OperationKind `json:"operationKind"`

	HasParent bool `json:"hasParent"`

	JSONPath string   `json:"jsonPath"`
	Path     path     `json:"path"`
	Meta     metaData `json:"meta"`
	engine   *Engine
}
type Position struct {
	position positionCore
}
type zoneCore struct {
	ID ZoneID `json:"id"`

	Interactables []AnyOfItem_Player_ZoneItemID `json:"interactables"`

	Items []ZoneItemID `json:"items"`

	Players []PlayerID `json:"players"`

	Tags []StringValueID `json:"tags"`

	OperationKind OperationKind `json:"operationKind"`

	HasParent bool `json:"hasParent"`

	JSONPath string   `json:"jsonPath"`
	Path     path     `json:"path"`
	Meta     metaData `json:"meta"`
	engine   *Engine
}
type Zone struct {
	zone zoneCore
}
type zoneItemCore struct {
	ID ZoneItemID `json:"id"`

	Item ItemID `json:"item"`

	Position PositionID `json:"position"`

	OperationKind OperationKind `json:"operationKind"`

	HasParent bool `json:"hasParent"`

	JSONPath string   `json:"jsonPath"`
	Path     path     `json:"path"`
	Meta     metaData `json:"meta"`
	engine   *Engine
}
type ZoneItem struct {
	zoneItem zoneItemCore
}
type attackEventTargetRefCore struct {
	ID AttackEventTargetRefID `json:"id"`

	ParentID AttackEventID `json:"parentID"`

	ReferencedElementID PlayerID `json:"referencedElementID"`

	OperationKind OperationKind `json:"operationKind"`

	Path   path     `json:"path"`
	Meta   metaData `json:"meta"`
	engine *Engine
}
type AttackEventTargetRef struct {
	attackEventTargetRef attackEventTargetRefCore
}
type equipmentSetEquipmentRefCore struct {
	ID EquipmentSetEquipmentRefID `json:"id"`

	ParentID EquipmentSetID `json:"parentID"`

	ReferencedElementID ItemID `json:"referencedElementID"`

	OperationKind OperationKind `json:"operationKind"`

	Path   path     `json:"path"`
	Meta   metaData `json:"meta"`
	engine *Engine
}
type EquipmentSetEquipmentRef struct {
	equipmentSetEquipmentRef equipmentSetEquipmentRefCore
}
type itemBoundToRefCore struct {
	ID ItemBoundToRefID `json:"id"`

	ParentID ItemID `json:"parentID"`

	ReferencedElementID PlayerID `json:"referencedElementID"`

	OperationKind OperationKind `json:"operationKind"`

	Path   path     `json:"path"`
	Meta   metaData `json:"meta"`
	engine *Engine
}
type ItemBoundToRef struct {
	itemBoundToRef itemBoundToRefCore
}
type playerEquipmentSetRefCore struct {
	ID PlayerEquipmentSetRefID `json:"id"`

	ParentID PlayerID `json:"parentID"`

	ReferencedElementID EquipmentSetID `json:"referencedElementID"`

	OperationKind OperationKind `json:"operationKind"`

	Path   path     `json:"path"`
	Meta   metaData `json:"meta"`
	engine *Engine
}
type PlayerEquipmentSetRef struct {
	playerEquipmentSetRef playerEquipmentSetRefCore
}
type playerGuildMemberRefCore struct {
	ID PlayerGuildMemberRefID `json:"id"`

	ParentID PlayerID `json:"parentID"`

	ReferencedElementID PlayerID `json:"referencedElementID"`

	OperationKind OperationKind `json:"operationKind"`

	Path   path     `json:"path"`
	Meta   metaData `json:"meta"`
	engine *Engine
}
type PlayerGuildMemberRef struct {
	playerGuildMemberRef playerGuildMemberRefCore
}
type playerTargetRefCore struct {
	ID PlayerTargetRefID `json:"id"`

	ParentID PlayerID `json:"parentID"`

	ReferencedElementID AnyOfPlayer_ZoneItemID `json:"referencedElementID"`

	OperationKind OperationKind `json:"operationKind"`

	Path   path     `json:"path"`
	Meta   metaData `json:"meta"`
	engine *Engine
}
type PlayerTargetRef struct {
	playerTargetRef playerTargetRefCore
}
type playerTargetedByRefCore struct {
	ID PlayerTargetedByRefID `json:"id"`

	ParentID PlayerID `json:"parentID"`

	ReferencedElementID AnyOfPlayer_ZoneItemID `json:"referencedElementID"`

	OperationKind OperationKind `json:"operationKind"`

	Path   path     `json:"path"`
	Meta   metaData `json:"meta"`
	engine *Engine
}
type PlayerTargetedByRef struct {
	playerTargetedByRef playerTargetedByRefCore
}
type anyOfPlayer_PositionCore struct {
	ID AnyOfPlayer_PositionID `json:"id"`

	ElementKind ElementKind `json:"elementKind"`

	ChildID int `json:"childID"`

	ParentElementPath path `json:"parentElementPath"`

	FieldIdentifier treeFieldIdentifier `json:"fieldIdentifier"`

	OperationKind OperationKind `json:"operationKind"`

	Meta   metaData `json:"meta"`
	engine *Engine
}
type AnyOfPlayer_Position struct {
	anyOfPlayer_Position anyOfPlayer_PositionCore
}
type anyOfPlayer_ZoneItemCore struct {
	ID AnyOfPlayer_ZoneItemID `json:"id"`

	ElementKind ElementKind `json:"elementKind"`

	ChildID int `json:"childID"`

	ParentElementPath path `json:"parentElementPath"`

	FieldIdentifier treeFieldIdentifier `json:"fieldIdentifier"`

	OperationKind OperationKind `json:"operationKind"`

	Meta   metaData `json:"meta"`
	engine *Engine
}
type AnyOfPlayer_ZoneItem struct {
	anyOfPlayer_ZoneItem anyOfPlayer_ZoneItemCore
}
type anyOfItem_Player_ZoneItemCore struct {
	ID AnyOfItem_Player_ZoneItemID `json:"id"`

	ElementKind ElementKind `json:"elementKind"`

	ChildID int `json:"childID"`

	ParentElementPath path `json:"parentElementPath"`

	FieldIdentifier treeFieldIdentifier `json:"fieldIdentifier"`

	OperationKind OperationKind `json:"operationKind"`

	Meta   metaData `json:"meta"`
	engine *Engine
}
type AnyOfItem_Player_ZoneItem struct {
	anyOfItem_Player_ZoneItem anyOfItem_Player_ZoneItemCore
}
type OperationKind string

const (
	OperationKindDelete    OperationKind = "DELETE"
	OperationKindUpdate    OperationKind = "UPDATE"
	OperationKindUnchanged OperationKind = "UNCHANGED"
)

type Engine struct {
	State                *State
	Patch                *State
	Tree                 *Tree
	thisClientID         string
	broadcastingClientID string
	planner              *assemblePlanner
	IDgen                int
}

func newEngine() *Engine {
	return &Engine{
		IDgen:   1,
		Patch:   newState(),
		State:   newState(),
		Tree:    newTree(),
		planner: newAssemblePlanner(),
	}
}
func (engine *Engine) GenerateID() int {
	newID := engine.IDgen
	engine.IDgen = engine.IDgen + 1
	return newID
}
func (engine *Engine) UpdateState() {
	for _, attackEvent := range engine.Patch.AttackEvent {
		engine.deleteAttackEvent(attackEvent.ID)
	}

	for _, boolValue := range engine.Patch.BoolValue {
		if boolValue.OperationKind == OperationKindDelete {
			delete(engine.State.BoolValue, boolValue.ID)
		} else {

			boolValue.OperationKind = OperationKindUnchanged
			boolValue.Meta.unsign()
			engine.State.BoolValue[boolValue.ID] = boolValue
		}
	}
	for _, floatValue := range engine.Patch.FloatValue {
		if floatValue.OperationKind == OperationKindDelete {
			delete(engine.State.FloatValue, floatValue.ID)
		} else {

			floatValue.OperationKind = OperationKindUnchanged
			floatValue.Meta.unsign()
			engine.State.FloatValue[floatValue.ID] = floatValue
		}
	}
	for _, intValue := range engine.Patch.IntValue {
		if intValue.OperationKind == OperationKindDelete {
			delete(engine.State.IntValue, intValue.ID)
		} else {

			intValue.OperationKind = OperationKindUnchanged
			intValue.Meta.unsign()
			engine.State.IntValue[intValue.ID] = intValue
		}
	}
	for _, stringValue := range engine.Patch.StringValue {
		if stringValue.OperationKind == OperationKindDelete {
			delete(engine.State.StringValue, stringValue.ID)
		} else {

			stringValue.OperationKind = OperationKindUnchanged
			stringValue.Meta.unsign()
			engine.State.StringValue[stringValue.ID] = stringValue
		}
	}

	for _, attackEvent := range engine.Patch.AttackEvent {
		if attackEvent.OperationKind == OperationKindDelete {
			delete(engine.State.AttackEvent, attackEvent.ID)
		} else {

			attackEvent.OperationKind = OperationKindUnchanged
			attackEvent.Meta.unsign()
			engine.State.AttackEvent[attackEvent.ID] = attackEvent
		}
	}
	for _, equipmentSet := range engine.Patch.EquipmentSet {
		if equipmentSet.OperationKind == OperationKindDelete {
			delete(engine.State.EquipmentSet, equipmentSet.ID)
		} else {

			equipmentSet.OperationKind = OperationKindUnchanged
			equipmentSet.Meta.unsign()
			engine.State.EquipmentSet[equipmentSet.ID] = equipmentSet
		}
	}
	for _, gearScore := range engine.Patch.GearScore {
		if gearScore.OperationKind == OperationKindDelete {
			delete(engine.State.GearScore, gearScore.ID)
		} else {

			gearScore.OperationKind = OperationKindUnchanged
			gearScore.Meta.unsign()
			engine.State.GearScore[gearScore.ID] = gearScore
		}
	}
	for _, item := range engine.Patch.Item {
		if item.OperationKind == OperationKindDelete {
			delete(engine.State.Item, item.ID)
		} else {

			item.OperationKind = OperationKindUnchanged
			item.Meta.unsign()
			engine.State.Item[item.ID] = item
		}
	}
	for _, player := range engine.Patch.Player {
		if player.OperationKind == OperationKindDelete {
			delete(engine.State.Player, player.ID)
		} else {
			player.Action = player.Action[:0]

			player.OperationKind = OperationKindUnchanged
			player.Meta.unsign()
			engine.State.Player[player.ID] = player
		}
	}
	for _, position := range engine.Patch.Position {
		if position.OperationKind == OperationKindDelete {
			delete(engine.State.Position, position.ID)
		} else {

			position.OperationKind = OperationKindUnchanged
			position.Meta.unsign()
			engine.State.Position[position.ID] = position
		}
	}
	for _, zone := range engine.Patch.Zone {
		if zone.OperationKind == OperationKindDelete {
			delete(engine.State.Zone, zone.ID)
		} else {

			zone.OperationKind = OperationKindUnchanged
			zone.Meta.unsign()
			engine.State.Zone[zone.ID] = zone
		}
	}
	for _, zoneItem := range engine.Patch.ZoneItem {
		if zoneItem.OperationKind == OperationKindDelete {
			delete(engine.State.ZoneItem, zoneItem.ID)
		} else {

			zoneItem.OperationKind = OperationKindUnchanged
			zoneItem.Meta.unsign()
			engine.State.ZoneItem[zoneItem.ID] = zoneItem
		}
	}

	for _, attackEventTargetRef := range engine.Patch.AttackEventTargetRef {
		if attackEventTargetRef.OperationKind == OperationKindDelete {
			delete(engine.State.AttackEventTargetRef, attackEventTargetRef.ID)
		} else {

			attackEventTargetRef.OperationKind = OperationKindUnchanged
			attackEventTargetRef.Meta.unsign()
			engine.State.AttackEventTargetRef[attackEventTargetRef.ID] = attackEventTargetRef
		}
	}
	for _, equipmentSetEquipmentRef := range engine.Patch.EquipmentSetEquipmentRef {
		if equipmentSetEquipmentRef.OperationKind == OperationKindDelete {
			delete(engine.State.EquipmentSetEquipmentRef, equipmentSetEquipmentRef.ID)
		} else {

			equipmentSetEquipmentRef.OperationKind = OperationKindUnchanged
			equipmentSetEquipmentRef.Meta.unsign()
			engine.State.EquipmentSetEquipmentRef[equipmentSetEquipmentRef.ID] = equipmentSetEquipmentRef
		}
	}
	for _, itemBoundToRef := range engine.Patch.ItemBoundToRef {
		if itemBoundToRef.OperationKind == OperationKindDelete {
			delete(engine.State.ItemBoundToRef, itemBoundToRef.ID)
		} else {

			itemBoundToRef.OperationKind = OperationKindUnchanged
			itemBoundToRef.Meta.unsign()
			engine.State.ItemBoundToRef[itemBoundToRef.ID] = itemBoundToRef
		}
	}
	for _, playerEquipmentSetRef := range engine.Patch.PlayerEquipmentSetRef {
		if playerEquipmentSetRef.OperationKind == OperationKindDelete {
			delete(engine.State.PlayerEquipmentSetRef, playerEquipmentSetRef.ID)
		} else {

			playerEquipmentSetRef.OperationKind = OperationKindUnchanged
			playerEquipmentSetRef.Meta.unsign()
			engine.State.PlayerEquipmentSetRef[playerEquipmentSetRef.ID] = playerEquipmentSetRef
		}
	}
	for _, playerGuildMemberRef := range engine.Patch.PlayerGuildMemberRef {
		if playerGuildMemberRef.OperationKind == OperationKindDelete {
			delete(engine.State.PlayerGuildMemberRef, playerGuildMemberRef.ID)
		} else {

			playerGuildMemberRef.OperationKind = OperationKindUnchanged
			playerGuildMemberRef.Meta.unsign()
			engine.State.PlayerGuildMemberRef[playerGuildMemberRef.ID] = playerGuildMemberRef
		}
	}
	for _, playerTargetRef := range engine.Patch.PlayerTargetRef {
		if playerTargetRef.OperationKind == OperationKindDelete {
			delete(engine.State.PlayerTargetRef, playerTargetRef.ID)
		} else {

			playerTargetRef.OperationKind = OperationKindUnchanged
			playerTargetRef.Meta.unsign()
			engine.State.PlayerTargetRef[playerTargetRef.ID] = playerTargetRef
		}
	}
	for _, playerTargetedByRef := range engine.Patch.PlayerTargetedByRef {
		if playerTargetedByRef.OperationKind == OperationKindDelete {
			delete(engine.State.PlayerTargetedByRef, playerTargetedByRef.ID)
		} else {

			playerTargetedByRef.OperationKind = OperationKindUnchanged
			playerTargetedByRef.Meta.unsign()
			engine.State.PlayerTargetedByRef[playerTargetedByRef.ID] = playerTargetedByRef
		}
	}

	for _, anyOfPlayer_Position := range engine.Patch.AnyOfPlayer_Position {
		if anyOfPlayer_Position.OperationKind == OperationKindDelete {
			delete(engine.State.AnyOfPlayer_Position, anyOfPlayer_Position.ID)
		} else {

			anyOfPlayer_Position.OperationKind = OperationKindUnchanged
			anyOfPlayer_Position.Meta.unsign()
			engine.State.AnyOfPlayer_Position[anyOfPlayer_Position.ID] = anyOfPlayer_Position
		}
	}
	for _, anyOfPlayer_ZoneItem := range engine.Patch.AnyOfPlayer_ZoneItem {
		if anyOfPlayer_ZoneItem.OperationKind == OperationKindDelete {
			delete(engine.State.AnyOfPlayer_ZoneItem, anyOfPlayer_ZoneItem.ID)
		} else {

			anyOfPlayer_ZoneItem.OperationKind = OperationKindUnchanged
			anyOfPlayer_ZoneItem.Meta.unsign()
			engine.State.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItem.ID] = anyOfPlayer_ZoneItem
		}
	}
	for _, anyOfItem_Player_ZoneItem := range engine.Patch.AnyOfItem_Player_ZoneItem {
		if anyOfItem_Player_ZoneItem.OperationKind == OperationKindDelete {
			delete(engine.State.AnyOfItem_Player_ZoneItem, anyOfItem_Player_ZoneItem.ID)
		} else {

			anyOfItem_Player_ZoneItem.OperationKind = OperationKindUnchanged
			anyOfItem_Player_ZoneItem.Meta.unsign()
			engine.State.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItem.ID] = anyOfItem_Player_ZoneItem
		}
	}

	for key := range engine.Patch.BoolValue {
		delete(engine.Patch.BoolValue, key)
	}
	for key := range engine.Patch.FloatValue {
		delete(engine.Patch.FloatValue, key)
	}
	for key := range engine.Patch.IntValue {
		delete(engine.Patch.IntValue, key)
	}
	for key := range engine.Patch.StringValue {
		delete(engine.Patch.StringValue, key)
	}

	for key := range engine.Patch.AttackEvent {
		delete(engine.Patch.AttackEvent, key)
	}
	for key := range engine.Patch.EquipmentSet {
		delete(engine.Patch.EquipmentSet, key)
	}
	for key := range engine.Patch.GearScore {
		delete(engine.Patch.GearScore, key)
	}
	for key := range engine.Patch.Item {
		delete(engine.Patch.Item, key)
	}
	for key := range engine.Patch.Player {
		delete(engine.Patch.Player, key)
	}
	for key := range engine.Patch.Position {
		delete(engine.Patch.Position, key)
	}
	for key := range engine.Patch.Zone {
		delete(engine.Patch.Zone, key)
	}
	for key := range engine.Patch.ZoneItem {
		delete(engine.Patch.ZoneItem, key)
	}

	for key := range engine.Patch.AttackEventTargetRef {
		delete(engine.Patch.AttackEventTargetRef, key)
	}
	for key := range engine.Patch.EquipmentSetEquipmentRef {
		delete(engine.Patch.EquipmentSetEquipmentRef, key)
	}
	for key := range engine.Patch.ItemBoundToRef {
		delete(engine.Patch.ItemBoundToRef, key)
	}
	for key := range engine.Patch.PlayerEquipmentSetRef {
		delete(engine.Patch.PlayerEquipmentSetRef, key)
	}
	for key := range engine.Patch.PlayerGuildMemberRef {
		delete(engine.Patch.PlayerGuildMemberRef, key)
	}
	for key := range engine.Patch.PlayerTargetRef {
		delete(engine.Patch.PlayerTargetRef, key)
	}
	for key := range engine.Patch.PlayerTargetedByRef {
		delete(engine.Patch.PlayerTargetedByRef, key)
	}

	for key := range engine.Patch.AnyOfPlayer_Position {
		delete(engine.Patch.AnyOfPlayer_Position, key)
	}
	for key := range engine.Patch.AnyOfPlayer_ZoneItem {
		delete(engine.Patch.AnyOfPlayer_ZoneItem, key)
	}
	for key := range engine.Patch.AnyOfItem_Player_ZoneItem {
		delete(engine.Patch.AnyOfItem_Player_ZoneItem, key)
	}

}

type ReferencedDataStatus string

const (
	ReferencedDataModified  ReferencedDataStatus = "MODIFIED"
	ReferencedDataUnchanged ReferencedDataStatus = "UNCHANGED"
)

type ElementKind string

const (
	ElementKindBoolValue   ElementKind = "bool"
	ElementKindFloatValue  ElementKind = "float64"
	ElementKindIntValue    ElementKind = "int64"
	ElementKindStringValue ElementKind = "string"

	ElementKindAttackEvent  ElementKind = "AttackEvent"
	ElementKindEquipmentSet ElementKind = "EquipmentSet"
	ElementKindGearScore    ElementKind = "GearScore"
	ElementKindItem         ElementKind = "Item"
	ElementKindPlayer       ElementKind = "Player"
	ElementKindPosition     ElementKind = "Position"
	ElementKindZone         ElementKind = "Zone"
	ElementKindZoneItem     ElementKind = "ZoneItem"
)

type Tree struct {
	AttackEvent map[AttackEventID]attackEvent `json:"attackEvent"`

	EquipmentSet map[EquipmentSetID]equipmentSet `json:"equipmentSet"`

	GearScore map[GearScoreID]gearScore `json:"gearScore"`

	Item map[ItemID]item `json:"item"`

	Player map[PlayerID]player `json:"player"`

	Position map[PositionID]position `json:"position"`

	Zone map[ZoneID]zone `json:"zone"`

	ZoneItem map[ZoneItemID]zoneItem `json:"zoneItem"`
}

func newTree() *Tree {
	return &Tree{AttackEvent: make(map[AttackEventID]attackEvent),
		EquipmentSet: make(map[EquipmentSetID]equipmentSet),
		GearScore:    make(map[GearScoreID]gearScore),
		Item:         make(map[ItemID]item),
		Player:       make(map[PlayerID]player),
		Position:     make(map[PositionID]position),
		Zone:         make(map[ZoneID]zone),
		ZoneItem:     make(map[ZoneItemID]zoneItem),
	}
}
func (t *Tree) clear() {
	for key := range t.AttackEvent {
		delete(t.AttackEvent, key)
	}
	for key := range t.EquipmentSet {
		delete(t.EquipmentSet, key)
	}
	for key := range t.GearScore {
		delete(t.GearScore, key)
	}
	for key := range t.Item {
		delete(t.Item, key)
	}
	for key := range t.Player {
		delete(t.Player, key)
	}
	for key := range t.Position {
		delete(t.Position, key)
	}
	for key := range t.Zone {
		delete(t.Zone, key)
	}
	for key := range t.ZoneItem {
		delete(t.ZoneItem, key)
	}

}

type attackEvent struct {
	ID AttackEventID `json:"id"`

	Target *elementReference `json:"target"`

	OperationKind OperationKind `json:"operationKind"`
}
type equipmentSet struct {
	ID EquipmentSetID `json:"id"`

	Equipment map[ItemID]elementReference `json:"equipment"`

	Name *string `json:"name"`

	OperationKind OperationKind `json:"operationKind"`
}
type gearScore struct {
	ID GearScoreID `json:"id"`

	Level *int64 `json:"level"`

	Score *int64 `json:"score"`

	OperationKind OperationKind `json:"operationKind"`
}
type item struct {
	ID ItemID `json:"id"`

	BoundTo *elementReference `json:"boundTo"`

	GearScore *gearScore `json:"gearScore"`

	Name *string `json:"name"`

	Origin interface{} `json:"origin"`

	OperationKind OperationKind `json:"operationKind"`
}
type player struct {
	ID PlayerID `json:"id"`

	Action map[AttackEventID]attackEvent `json:"action"`

	EquipmentSets map[EquipmentSetID]elementReference `json:"equipmentSets"`

	GearScore *gearScore `json:"gearScore"`

	GuildMembers map[PlayerID]elementReference `json:"guildMembers"`

	Items map[ItemID]item `json:"items"`

	Position *position `json:"position"`

	Target *elementReference `json:"target"`

	TargetedBy map[int]elementReference `json:"targetedBy"`

	OperationKind OperationKind `json:"operationKind"`
}
type position struct {
	ID PositionID `json:"id"`

	X *float64 `json:"x"`

	Y *float64 `json:"y"`

	OperationKind OperationKind `json:"operationKind"`
}
type zone struct {
	ID ZoneID `json:"id"`

	Interactables map[int]interface{} `json:"interactables"`

	Items map[ZoneItemID]zoneItem `json:"items"`

	Players map[PlayerID]player `json:"players"`

	Tags []string `json:"tags"`

	OperationKind OperationKind `json:"operationKind"`
}
type zoneItem struct {
	ID ZoneItemID `json:"id"`

	Item *item `json:"item"`

	Position *position `json:"position"`

	OperationKind OperationKind `json:"operationKind"`
}
type elementReference struct {
	OperationKind OperationKind `json:"operationKind"`

	ElementID int `json:"id"`

	ElementKind ElementKind `json:"elementKind"`

	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`

	ElementPath string `json:"elementPath"`
}

var attackEventCheckPool = sync.Pool{New: func() interface{} {
	return make(map[AttackEventID]bool)
}}
var attackEventIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]AttackEventID, 0)
}}
var equipmentSetCheckPool = sync.Pool{New: func() interface{} {
	return make(map[EquipmentSetID]bool)
}}
var equipmentSetIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]EquipmentSetID, 0)
}}
var gearScoreCheckPool = sync.Pool{New: func() interface{} {
	return make(map[GearScoreID]bool)
}}
var gearScoreIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]GearScoreID, 0)
}}
var itemCheckPool = sync.Pool{New: func() interface{} {
	return make(map[ItemID]bool)
}}
var itemIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]ItemID, 0)
}}
var playerCheckPool = sync.Pool{New: func() interface{} {
	return make(map[PlayerID]bool)
}}
var playerIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]PlayerID, 0)
}}
var positionCheckPool = sync.Pool{New: func() interface{} {
	return make(map[PositionID]bool)
}}
var positionIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]PositionID, 0)
}}
var zoneCheckPool = sync.Pool{New: func() interface{} {
	return make(map[ZoneID]bool)
}}
var zoneIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]ZoneID, 0)
}}
var zoneItemCheckPool = sync.Pool{New: func() interface{} {
	return make(map[ZoneItemID]bool)
}}
var zoneItemIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]ZoneItemID, 0)
}}
var attackEventTargetRefCheckPool = sync.Pool{New: func() interface{} {
	return make(map[AttackEventTargetRefID]bool)
}}
var attackEventTargetRefIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]AttackEventTargetRefID, 0)
}}
var equipmentSetEquipmentRefCheckPool = sync.Pool{New: func() interface{} {
	return make(map[EquipmentSetEquipmentRefID]bool)
}}
var equipmentSetEquipmentRefIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]EquipmentSetEquipmentRefID, 0)
}}
var itemBoundToRefCheckPool = sync.Pool{New: func() interface{} {
	return make(map[ItemBoundToRefID]bool)
}}
var itemBoundToRefIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]ItemBoundToRefID, 0)
}}
var playerEquipmentSetRefCheckPool = sync.Pool{New: func() interface{} {
	return make(map[PlayerEquipmentSetRefID]bool)
}}
var playerEquipmentSetRefIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]PlayerEquipmentSetRefID, 0)
}}
var playerGuildMemberRefCheckPool = sync.Pool{New: func() interface{} {
	return make(map[PlayerGuildMemberRefID]bool)
}}
var playerGuildMemberRefIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]PlayerGuildMemberRefID, 0)
}}
var playerTargetRefCheckPool = sync.Pool{New: func() interface{} {
	return make(map[PlayerTargetRefID]bool)
}}
var playerTargetRefIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]PlayerTargetRefID, 0)
}}
var playerTargetedByRefCheckPool = sync.Pool{New: func() interface{} {
	return make(map[PlayerTargetedByRefID]bool)
}}
var playerTargetedByRefIDSlicePool = sync.Pool{New: func() interface{} {
	return make([]PlayerTargetedByRefID, 0)
}}

func (engine *Engine) importPatch(patch *State) {
	for _, boolValue := range patch.BoolValue {
		if boolValue.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		boolValue.Meta.unsign()
		engine.State.BoolValue[boolValue.ID] = boolValue
	}
	for _, floatValue := range patch.FloatValue {
		if floatValue.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		floatValue.Meta.unsign()
		engine.State.FloatValue[floatValue.ID] = floatValue
	}
	for _, intValue := range patch.IntValue {
		if intValue.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		intValue.Meta.unsign()
		engine.State.IntValue[intValue.ID] = intValue
	}
	for _, stringValue := range patch.StringValue {
		if stringValue.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		stringValue.Meta.unsign()
		engine.State.StringValue[stringValue.ID] = stringValue
	}

	for _, attackEvent := range patch.AttackEvent {
		if attackEvent.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		attackEvent.Meta.unsign()
		engine.State.AttackEvent[attackEvent.ID] = attackEvent
	}
	for _, equipmentSet := range patch.EquipmentSet {
		if equipmentSet.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		equipmentSet.Meta.unsign()
		engine.State.EquipmentSet[equipmentSet.ID] = equipmentSet
	}
	for _, gearScore := range patch.GearScore {
		if gearScore.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		gearScore.Meta.unsign()
		engine.State.GearScore[gearScore.ID] = gearScore
	}
	for _, item := range patch.Item {
		if item.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		item.Meta.unsign()
		engine.State.Item[item.ID] = item
	}
	for _, player := range patch.Player {
		if player.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		player.Meta.unsign()
		engine.State.Player[player.ID] = player
	}
	for _, position := range patch.Position {
		if position.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		position.Meta.unsign()
		engine.State.Position[position.ID] = position
	}
	for _, zone := range patch.Zone {
		if zone.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		zone.Meta.unsign()
		engine.State.Zone[zone.ID] = zone
	}
	for _, zoneItem := range patch.ZoneItem {
		if zoneItem.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		zoneItem.Meta.unsign()
		engine.State.ZoneItem[zoneItem.ID] = zoneItem
	}

	for _, attackEventTargetRef := range patch.AttackEventTargetRef {
		if attackEventTargetRef.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		attackEventTargetRef.Meta.unsign()
		engine.State.AttackEventTargetRef[attackEventTargetRef.ID] = attackEventTargetRef
	}
	for _, equipmentSetEquipmentRef := range patch.EquipmentSetEquipmentRef {
		if equipmentSetEquipmentRef.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		equipmentSetEquipmentRef.Meta.unsign()
		engine.State.EquipmentSetEquipmentRef[equipmentSetEquipmentRef.ID] = equipmentSetEquipmentRef
	}
	for _, itemBoundToRef := range patch.ItemBoundToRef {
		if itemBoundToRef.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		itemBoundToRef.Meta.unsign()
		engine.State.ItemBoundToRef[itemBoundToRef.ID] = itemBoundToRef
	}
	for _, playerEquipmentSetRef := range patch.PlayerEquipmentSetRef {
		if playerEquipmentSetRef.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		playerEquipmentSetRef.Meta.unsign()
		engine.State.PlayerEquipmentSetRef[playerEquipmentSetRef.ID] = playerEquipmentSetRef
	}
	for _, playerGuildMemberRef := range patch.PlayerGuildMemberRef {
		if playerGuildMemberRef.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		playerGuildMemberRef.Meta.unsign()
		engine.State.PlayerGuildMemberRef[playerGuildMemberRef.ID] = playerGuildMemberRef
	}
	for _, playerTargetRef := range patch.PlayerTargetRef {
		if playerTargetRef.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		playerTargetRef.Meta.unsign()
		engine.State.PlayerTargetRef[playerTargetRef.ID] = playerTargetRef
	}
	for _, playerTargetedByRef := range patch.PlayerTargetedByRef {
		if playerTargetedByRef.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		playerTargetedByRef.Meta.unsign()
		engine.State.PlayerTargetedByRef[playerTargetedByRef.ID] = playerTargetedByRef
	}

	for _, anyOfPlayer_Position := range patch.AnyOfPlayer_Position {
		if anyOfPlayer_Position.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		anyOfPlayer_Position.Meta.unsign()
		engine.State.AnyOfPlayer_Position[anyOfPlayer_Position.ID] = anyOfPlayer_Position
	}
	for _, anyOfPlayer_ZoneItem := range patch.AnyOfPlayer_ZoneItem {
		if anyOfPlayer_ZoneItem.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		anyOfPlayer_ZoneItem.Meta.unsign()
		engine.State.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItem.ID] = anyOfPlayer_ZoneItem
	}
	for _, anyOfItem_Player_ZoneItem := range patch.AnyOfItem_Player_ZoneItem {
		if anyOfItem_Player_ZoneItem.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		anyOfItem_Player_ZoneItem.Meta.unsign()
		engine.State.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItem.ID] = anyOfItem_Player_ZoneItem
	}

}
