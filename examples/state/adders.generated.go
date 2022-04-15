package state

func (_zone Zone) AddPlayer() Player {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return Player{player: playerCore{OperationKind: OperationKindDelete, engine: zone.zone.engine}}
	}
	player := zone.zone.engine.createPlayer(zone.zone.Path, zone_playersIdentifier)
	zone.zone.Players = append(zone.zone.Players, player.player.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.Meta.sign(zone.zone.engine.BroadcastingClientID)
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return player
}

func (_zone Zone) AddItem() ZoneItem {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return ZoneItem{zoneItem: zoneItemCore{OperationKind: OperationKindDelete, engine: zone.zone.engine}}
	}
	zoneItem := zone.zone.engine.createZoneItem(zone.zone.Path, zone_itemsIdentifier)
	zone.zone.Items = append(zone.zone.Items, zoneItem.zoneItem.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.Meta.sign(zone.zone.engine.BroadcastingClientID)
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return zoneItem
}

func (_zone Zone) AddInteractablePlayer() Player {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return Player{player: playerCore{OperationKind: OperationKindDelete, engine: zone.zone.engine}}
	}
	player := zone.zone.engine.createPlayer(zone.zone.Path, zone_interactablesIdentifier)
	anyContainer := zone.zone.engine.createAnyOfItem_Player_ZoneItem(int(zone.zone.ID), int(player.player.ID), ElementKindPlayer, zone.zone.Path, zone_interactablesIdentifier).anyOfItem_Player_ZoneItem
	zone.zone.Interactables = append(zone.zone.Interactables, anyContainer.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.Meta.sign(zone.zone.engine.BroadcastingClientID)
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return player
}

func (_zone Zone) AddInteractableZoneItem() ZoneItem {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return ZoneItem{zoneItem: zoneItemCore{OperationKind: OperationKindDelete, engine: zone.zone.engine}}
	}
	zoneItem := zone.zone.engine.createZoneItem(zone.zone.Path, zone_interactablesIdentifier)
	anyContainer := zone.zone.engine.createAnyOfItem_Player_ZoneItem(int(zone.zone.ID), int(zoneItem.zoneItem.ID), ElementKindZoneItem, zone.zone.Path, zone_interactablesIdentifier).anyOfItem_Player_ZoneItem
	zone.zone.Interactables = append(zone.zone.Interactables, anyContainer.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.Meta.sign(zone.zone.engine.BroadcastingClientID)
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return zoneItem
}

func (_zone Zone) AddInteractableItem() Item {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return Item{item: itemCore{OperationKind: OperationKindDelete, engine: zone.zone.engine}}
	}
	item := zone.zone.engine.createItem(zone.zone.Path, zone_interactablesIdentifier)
	anyContainer := zone.zone.engine.createAnyOfItem_Player_ZoneItem(int(zone.zone.ID), int(item.item.ID), ElementKindItem, zone.zone.Path, zone_interactablesIdentifier).anyOfItem_Player_ZoneItem
	zone.zone.Interactables = append(zone.zone.Interactables, anyContainer.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.Meta.sign(zone.zone.engine.BroadcastingClientID)
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return item
}

func (_zone Zone) AddTag(tag string) {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return
	}
	tagValue := zone.zone.engine.createStringValue(zone.zone.Path, zone_tagsIdentifier, tag)
	zone.zone.Tags = append(zone.zone.Tags, tagValue.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.Meta.sign(zone.zone.engine.BroadcastingClientID)
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
}

func (_player Player) AddItem() Item {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return Item{item: itemCore{OperationKind: OperationKindDelete, engine: player.player.engine}}
	}
	item := player.player.engine.createItem(player.player.Path, player_itemsIdentifier)
	player.player.Items = append(player.player.Items, item.item.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.Meta.sign(player.player.engine.BroadcastingClientID)
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return item
}

func (_player Player) AddAction() AttackEvent {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return AttackEvent{attackEvent: attackEventCore{OperationKind: OperationKindDelete, engine: player.player.engine}}
	}
	attackEvent := player.player.engine.createAttackEvent(player.player.Path, player_actionIdentifier)
	player.player.Action = append(player.player.Action, attackEvent.attackEvent.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.Meta.sign(player.player.engine.BroadcastingClientID)
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return attackEvent
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
	player.player.Meta.sign(player.player.engine.BroadcastingClientID)
	player.player.engine.Patch.Player[player.player.ID] = player.player
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
	player.player.Meta.sign(player.player.engine.BroadcastingClientID)
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
	player.player.Meta.sign(player.player.engine.BroadcastingClientID)
	player.player.engine.Patch.Player[player.player.ID] = player.player
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
	player.player.Meta.sign(player.player.engine.BroadcastingClientID)
	player.player.engine.Patch.Player[player.player.ID] = player.player
}

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
	equipmentSet.equipmentSet.Meta.sign(equipmentSet.equipmentSet.engine.BroadcastingClientID)
	equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID] = equipmentSet.equipmentSet
}