package state

func (_zone zone) AddPlayer() player {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return player{player: playerCore{OperationKind: OperationKindDelete}}
	}
	player := zone.zone.engine.createPlayer(true)
	zone.zone.Players = append(zone.zone.Players, player.player.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return player
}

func (_zone zone) AddItem() zoneItem {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zoneItem{zoneItem: zoneItemCore{OperationKind: OperationKindDelete}}
	}
	zoneItem := zone.zone.engine.createZoneItem(true)
	zone.zone.Items = append(zone.zone.Items, zoneItem.zoneItem.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return zoneItem
}

func (_zone zone) AddInteractablePlayer() player {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return player{player: playerCore{OperationKind: OperationKindDelete}}
	}
	player := zone.zone.engine.createPlayer(true)
	anyContainer := zone.zone.engine.createAnyOfItem_Player_ZoneItem(false).anyOfItem_Player_ZoneItem
	anyContainer.setPlayer(player.player.ID)
	zone.zone.Interactables = append(zone.zone.Interactables, anyContainer.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return player
}

func (_zone zone) AddInteractableZoneItem() zoneItem {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zoneItem{zoneItem: zoneItemCore{OperationKind: OperationKindDelete}}
	}
	zoneItem := zone.zone.engine.createZoneItem(true)
	anyContainer := zone.zone.engine.createAnyOfItem_Player_ZoneItem(false).anyOfItem_Player_ZoneItem
	anyContainer.setZoneItem(zoneItem.zoneItem.ID)
	zone.zone.Interactables = append(zone.zone.Interactables, anyContainer.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return zoneItem
}

func (_zone zone) AddInteractableItem() item {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return item{item: itemCore{OperationKind: OperationKindDelete}}
	}
	item := zone.zone.engine.createItem(true)
	anyContainer := zone.zone.engine.createAnyOfItem_Player_ZoneItem(false).anyOfItem_Player_ZoneItem
	anyContainer.setItem(item.item.ID)
	zone.zone.Interactables = append(zone.zone.Interactables, anyContainer.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return item
}

func (_zone zone) AddTags(tags ...string) {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return
	}
	zone.zone.Tags = append(zone.zone.Tags, tags...)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
}

func (_player player) AddItem() item {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return item{item: itemCore{OperationKind: OperationKindDelete}}
	}
	item := player.player.engine.createItem(true)
	player.player.Items = append(player.player.Items, item.item.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return item
}

func (_player player) AddGuildMember(playerID PlayerID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return
	}
	if player.player.engine.Player(playerID).player.OperationKind == OperationKindDelete {
		return
	}
	ref := player.player.engine.createPlayerGuildMemberRef(playerID, player.player.ID)
	player.player.GuildMembers = append(player.player.GuildMembers, ref.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
}

func (_player player) AddTargetedByPlayer(playerID PlayerID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return
	}
	if player.player.engine.Player(playerID).player.OperationKind == OperationKindDelete {
		return
	}
	anyContainer := player.player.engine.createAnyOfPlayer_ZoneItem(false).anyOfPlayer_ZoneItem
	anyContainer.setPlayer(playerID)
	ref := player.player.engine.createPlayerTargetedByRef(anyContainer.ID, player.player.ID)
	player.player.TargetedBy = append(player.player.TargetedBy, ref.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
}

func (_player player) AddTargetedByZoneItem(zoneItemID ZoneItemID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return
	}
	if player.player.engine.ZoneItem(zoneItemID).zoneItem.OperationKind == OperationKindDelete {
		return
	}
	anyContainer := player.player.engine.createAnyOfPlayer_ZoneItem(false).anyOfPlayer_ZoneItem
	anyContainer.setZoneItem(zoneItemID)
	ref := player.player.engine.createPlayerTargetedByRef(anyContainer.ID, player.player.ID)
	player.player.TargetedBy = append(player.player.TargetedBy, ref.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
}

func (_player player) AddEquipmentSet(equipmentSetID EquipmentSetID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return
	}
	if player.player.engine.EquipmentSet(equipmentSetID).equipmentSet.OperationKind == OperationKindDelete {
		return
	}
	ref := player.player.engine.createPlayerEquipmentSetRef(equipmentSetID, player.player.ID)
	player.player.EquipmentSets = append(player.player.EquipmentSets, ref.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
}

func (_equipmentSet equipmentSet) AddEquipment(itemID ItemID) {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	if equipmentSet.equipmentSet.OperationKind == OperationKindDelete {
		return
	}
	if equipmentSet.equipmentSet.engine.Item(itemID).item.OperationKind == OperationKindDelete {
		return
	}
	ref := equipmentSet.equipmentSet.engine.createEquipmentSetEquipmentRef(itemID, equipmentSet.equipmentSet.ID)
	equipmentSet.equipmentSet.Equipment = append(equipmentSet.equipmentSet.Equipment, ref.ID)
	equipmentSet.equipmentSet.OperationKind = OperationKindUpdate
	equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID] = equipmentSet.equipmentSet
}
