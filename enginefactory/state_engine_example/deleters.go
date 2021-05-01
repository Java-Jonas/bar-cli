package state

func (engine *Engine) DeletePlayer(playerID PlayerID) {
	player := engine.Player(playerID).player
	if player.HasParent_ {
		return
	}
	engine.deletePlayer(playerID)
}
func (engine *Engine) deletePlayer(playerID PlayerID) {
	player := engine.Player(playerID).player
	engine.dereferenceItemBoundToRefs(playerID)
	engine.dereferencePlayerGuildMemberRefs(playerID)
	engine.dereferencePlayerTargetPlayerRefs(playerID)
	engine.dereferencePlayerTargetedByPlayerRefs(playerID)
	engine.deleteGearScore(player.GearScore)
	for _, guildMember := range player.GuildMembers {
		engine.deletePlayerGuildMemberRef(guildMember)
	}
	for _, itemID := range player.Items {
		engine.deleteItem(itemID)
	}
	engine.deletePosition(player.Position)
	engine.deletePlayerTargetRef(player.Target)
	for _, targetedBy := range player.TargetedBy {
		engine.deletePlayerTargetedByRef(targetedBy)
	}
	player.OperationKind_ = OperationKindDelete
	engine.Patch.Player[player.ID] = player
}

func (engine *Engine) DeleteGearScore(gearScoreID GearScoreID) {
	gearScore := engine.GearScore(gearScoreID).gearScore
	if gearScore.HasParent_ {
		return
	}
	engine.deleteGearScore(gearScoreID)
}
func (engine *Engine) deleteGearScore(gearScoreID GearScoreID) {
	gearScore := engine.GearScore(gearScoreID).gearScore
	gearScore.OperationKind_ = OperationKindDelete
	engine.Patch.GearScore[gearScore.ID] = gearScore
}

func (engine *Engine) DeletePosition(positionID PositionID) {
	position := engine.Position(positionID).position
	if position.HasParent_ {
		return
	}
	engine.deletePosition(positionID)
}
func (engine *Engine) deletePosition(positionID PositionID) {
	position := engine.Position(positionID).position
	position.OperationKind_ = OperationKindDelete
	engine.Patch.Position[position.ID] = position
}

func (engine *Engine) DeleteItem(itemID ItemID) {
	item := engine.Item(itemID).item
	if item.HasParent_ {
		return
	}
	engine.deleteItem(itemID)
}
func (engine *Engine) deleteItem(itemID ItemID) {
	item := engine.Item(itemID).item
	engine.deleteItemBoundToRef(item.BoundTo)
	engine.deleteGearScore(item.GearScore)
	engine.deleteAnyOfPlayerZone(item.Origin, true)
	item.OperationKind_ = OperationKindDelete
	engine.Patch.Item[item.ID] = item
}

func (engine *Engine) DeleteZoneItem(zoneItemID ZoneItemID) {
	zoneItem := engine.ZoneItem(zoneItemID).zoneItem
	if zoneItem.HasParent_ {
		return
	}
	engine.deleteZoneItem(zoneItemID)
}
func (engine *Engine) deleteZoneItem(zoneItemID ZoneItemID) {
	zoneItem := engine.ZoneItem(zoneItemID).zoneItem
	engine.dereferencePlayerTargetZoneItemRefs(zoneItemID)
	engine.dereferencePlayerTargetedByZoneItemRefs(zoneItemID)
	engine.deleteItem(zoneItem.Item)
	engine.deletePosition(zoneItem.Position)
	zoneItem.OperationKind_ = OperationKindDelete
	engine.Patch.ZoneItem[zoneItem.ID] = zoneItem
}

func (engine *Engine) DeleteZone(zoneID ZoneID) {
	engine.deleteZone(zoneID)
}
func (engine *Engine) deleteZone(zoneID ZoneID) {
	zone := engine.Zone(zoneID).zone
	for _, zoneItemID := range zone.Items {
		engine.deleteZoneItem(zoneItemID)
	}
	for _, playerID := range zone.Players {
		engine.deletePlayer(playerID)
	}
	zone.OperationKind_ = OperationKindDelete
	engine.Patch.Zone[zone.ID] = zone
}

func (engine *Engine) DeleteEquipmentSet(equipmentSetID EquipmentSetID) {
	engine.deleteEquipmentSet(equipmentSetID)
}
func (engine *Engine) deleteEquipmentSet(equipmentSetID EquipmentSetID) {
	equipmentSet := engine.EquipmentSet(equipmentSetID).equipmentSet
	engine.dereferencePlayerEquipmentSetRefs(equipmentSetID)
	for _, equipmentSet := range equipmentSet.Equipment {
		engine.deleteEquipmentSetEquipmentRef(equipmentSet)
	}
	equipmentSet.OperationKind_ = OperationKindDelete
	engine.Patch.EquipmentSet[equipmentSet.ID] = equipmentSet
}

func (engine *Engine) deletePlayerGuildMemberRef(playerGuildMemberRefID PlayerGuildMemberRefID) {
	playerGuildMemberRef := engine.playerGuildMemberRef(playerGuildMemberRefID).playerGuildMemberRef
	playerGuildMemberRef.OperationKind_ = OperationKindDelete
	engine.Patch.PlayerGuildMemberRef[playerGuildMemberRef.ID] = playerGuildMemberRef
}

func (engine *Engine) deletePlayerEquipmentSetRef(playerEquipmentSetRefID PlayerEquipmentSetRefID) {
	playerEquipmentSetRef := engine.playerEquipmentSetRef(playerEquipmentSetRefID).playerEquipmentSetRef
	playerEquipmentSetRef.OperationKind_ = OperationKindDelete
	engine.Patch.PlayerEquipmentSetRef[playerEquipmentSetRef.ID] = playerEquipmentSetRef
}

func (engine *Engine) deleteItemBoundToRef(itemBoundToRefID ItemBoundToRefID) {
	itemBoundToRef := engine.itemBoundToRef(itemBoundToRefID).itemBoundToRef
	itemBoundToRef.OperationKind_ = OperationKindDelete
	engine.Patch.ItemBoundToRef[itemBoundToRef.ID] = itemBoundToRef
}

func (engine *Engine) deleteEquipmentSetEquipmentRef(equipmentSetEquipmentRefID EquipmentSetEquipmentRefID) {
	equipmentSetEquipmentRef := engine.equipmentSetEquipmentRef(equipmentSetEquipmentRefID).equipmentSetEquipmentRef
	equipmentSetEquipmentRef.OperationKind_ = OperationKindDelete
	engine.Patch.EquipmentSetEquipmentRef[equipmentSetEquipmentRef.ID] = equipmentSetEquipmentRef
}

func (engine *Engine) deletePlayerTargetRef(playerTargetRefID PlayerTargetRefID) {
	playerTargetRef := engine.playerTargetRef(playerTargetRefID).playerTargetRef
	engine.deleteAnyOfPlayerZoneItem(playerTargetRef.ReferencedElementID, false)
	playerTargetRef.OperationKind_ = OperationKindDelete
	engine.Patch.PlayerTargetRef[playerTargetRef.ID] = playerTargetRef
}

func (engine *Engine) deletePlayerTargetedByRef(playerTargetedByRefID PlayerTargetedByRefID) {
	playerTargetedByRef := engine.playerTargetedByRef(playerTargetedByRefID).playerTargetedByRef
	engine.deleteAnyOfPlayerZoneItem(playerTargetedByRef.ReferencedElementID, false)
	playerTargetedByRef.OperationKind_ = OperationKindDelete
	engine.Patch.PlayerTargetedByRef[playerTargetedByRef.ID] = playerTargetedByRef
}

func (engine *Engine) deleteAnyOfPlayerZoneItem(anyOfPlayerZoneItemID AnyOfPlayerZoneItemID, deleteChild bool) {
	anyOfPlayerZoneItem := engine.anyOfPlayerZoneItem(anyOfPlayerZoneItemID).anyOfPlayerZoneItem
	if deleteChild {
		anyOfPlayerZoneItem.deleteChild()
	}
	anyOfPlayerZoneItem.OperationKind_ = OperationKindDelete
	engine.Patch.AnyOfPlayerZoneItem[anyOfPlayerZoneItem.ID] = anyOfPlayerZoneItem
}

func (engine *Engine) deleteAnyOfPlayerZone(anyOfPlayerZoneID AnyOfPlayerZoneID, deleteChild bool) {
	anyOfPlayerZone := engine.anyOfPlayerZone(anyOfPlayerZoneID).anyOfPlayerZone
	if deleteChild {
		anyOfPlayerZone.deleteChild()
	}
	anyOfPlayerZone.OperationKind_ = OperationKindDelete
	engine.Patch.AnyOfPlayerZone[anyOfPlayerZone.ID] = anyOfPlayerZone
}

func (engine *Engine) deleteAnyOfItemPlayerZoneItem(anyOfItemPlayerZoneItemID AnyOfItemPlayerZoneItemID, deleteChild bool) {
	anyOfItemPlayerZoneItem := engine.anyOfItemPlayerZoneItem(anyOfItemPlayerZoneItemID).anyOfItemPlayerZoneItem
	if deleteChild {
		anyOfItemPlayerZoneItem.deleteChild()
	}
	anyOfItemPlayerZoneItem.OperationKind_ = OperationKindDelete
	engine.Patch.AnyOfItemPlayerZoneItem[anyOfItemPlayerZoneItem.ID] = anyOfItemPlayerZoneItem
}
