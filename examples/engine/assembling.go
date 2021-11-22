package state

import "fmt"

func (engine *Engine) assembleGearScorePath(element *gearScore, p path, pIndex int, includedElements map[int]bool) {

	gearScoreData, ok := engine.Patch.GearScore[element.ID]
	if !ok {
		gearScoreData = engine.State.GearScore[element.ID]
	}

	element.OperationKind = gearScoreData.OperationKind
	element.Level = gearScoreData.Level
	element.Score = gearScoreData.Score

	_ = gearScoreData
}

func (engine *Engine) assemblePositionPath(element *position, p path, pIndex int, includedElements map[int]bool) {

	positionData, ok := engine.Patch.Position[element.ID]
	if !ok {
		positionData = engine.State.Position[element.ID]
	}

	element.OperationKind = positionData.OperationKind
	element.X = positionData.X
	element.Y = positionData.Y

	_ = positionData
}

func (engine *Engine) assembleEquipmentSetPath(element *equipmentSet, p path, pIndex int, includedElements map[int]bool) {

	equipmentSetData, ok := engine.Patch.EquipmentSet[element.ID]
	if !ok {
		equipmentSetData = engine.State.EquipmentSet[element.ID]
	}

	element.OperationKind = equipmentSetData.OperationKind
	element.Name = equipmentSetData.Name

	if pIndex+1 == len(p) {
		return
	}

	switch p[pIndex+1].identifier {
	case equipmentSet_equipmentIdentifier:
		ref := engine.equipmentSetEquipmentRef(EquipmentSetEquipmentRefID(p[pIndex+1].refID)).equipmentSetEquipmentRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.Item(ref.ReferencedElementID).item
		treeRef := elementReference{
			OperationKind:        ref.OperationKind,
			ElementID:            int(ref.ReferencedElementID),
			ElementKind:          ElementKindItem,
			ReferencedDataStatus: referencedDataStatus,
			ElementPath:          referencedElement.Path,
		}
		if element.Equipment == nil {
			element.Equipment = make(map[ItemID]elementReference)
		}
		element.Equipment[referencedElement.ID] = treeRef
	}

	_ = equipmentSetData
}

func (engine *Engine) assembleItemPath(element *item, p path, pIndex int, includedElements map[int]bool) {

	itemData, ok := engine.Patch.Item[element.ID]
	if !ok {
		itemData = engine.State.Item[element.ID]
	}

	element.OperationKind = itemData.OperationKind
	element.Name = itemData.Name

	if pIndex+1 == len(p) {
		return
	}

	switch p[pIndex+1].identifier {
	case item_boundToIdentifier:
		ref := engine.itemBoundToRef(ItemBoundToRefID(p[pIndex+1].id)).itemBoundToRef
		if element.BoundTo != nil && ref.OperationKind == OperationKindDelete {
			break
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.Player(ref.ReferencedElementID).player
		treeRef := elementReference{
			OperationKind:        ref.OperationKind,
			ElementID:            int(ref.ReferencedElementID),
			ElementKind:          ElementKindPlayer,
			ReferencedDataStatus: referencedDataStatus,
			ElementPath:          referencedElement.Path,
		}
		element.BoundTo = &treeRef
	case item_gearScoreIdentifier:
		child := element.GearScore
		if child == nil {
			child = &gearScore{ID: itemData.GearScore}
		}
		engine.assembleGearScorePath(child, p, pIndex+1, includedElements)
		element.GearScore = child
	case item_originIdentifier:
		switch p[pIndex+1].kind {
		case ElementKindPlayer:
			child, ok := element.Origin.(*player)
			if !ok || child == nil {
				child = &player{ID: PlayerID(p[pIndex+1].id)}
			}
			engine.assemblePlayerPath(child, p, pIndex+1, includedElements)
			element.Origin = child
		case ElementKindPosition:
			child, ok := element.Origin.(*position)
			if !ok || child == nil {
				child = &position{ID: PositionID(p[pIndex+1].id)}
			}
			engine.assemblePositionPath(child, p, pIndex+1, includedElements)
			element.Origin = child
		}
	}

	_ = itemData
}

func (engine *Engine) assembleZoneItemPath(element *zoneItem, p path, pIndex int, includedElements map[int]bool) {

	zoneItemData, ok := engine.Patch.ZoneItem[element.ID]
	if !ok {
		zoneItemData = engine.State.ZoneItem[element.ID]
	}

	element.OperationKind = zoneItemData.OperationKind

	if pIndex+1 == len(p) {
		return
	}

	switch p[pIndex+1].identifier {
	case zoneItem_itemIdentifier:
		child := element.Item
		if child == nil {
			child = &item{ID: zoneItemData.Item}
		}
		engine.assembleItemPath(child, p, pIndex+1, includedElements)
		element.Item = child
	case zoneItem_positionIdentifier:
		child := element.Position
		if child == nil {
			child = &position{ID: zoneItemData.Position}
		}
		engine.assemblePositionPath(child, p, pIndex+1, includedElements)
		element.Position = child
	}

	_ = zoneItemData
}

func (engine *Engine) assemblePlayerPath(element *player, p path, pIndex int, includedElements map[int]bool) {

	playerData, ok := engine.Patch.Player[element.ID]
	if !ok {
		playerData = engine.State.Player[element.ID]
	}

	element.OperationKind = playerData.OperationKind

	if pIndex+1 == len(p) {
		return
	}

	switch p[pIndex+1].identifier {
	case player_equipmentSetsIdentifier:
		ref := engine.playerEquipmentSetRef(PlayerEquipmentSetRefID(p[pIndex+1].refID)).playerEquipmentSetRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.EquipmentSet(ref.ReferencedElementID).equipmentSet
		treeRef := elementReference{
			OperationKind:        ref.OperationKind,
			ElementID:            int(ref.ReferencedElementID),
			ElementKind:          ElementKindEquipmentSet,
			ReferencedDataStatus: referencedDataStatus,
			ElementPath:          referencedElement.Path,
		}
		if element.EquipmentSets == nil {
			element.EquipmentSets = make(map[EquipmentSetID]elementReference)
		}
		element.EquipmentSets[referencedElement.ID] = treeRef
	case player_gearScoreIdentifier:
		child := element.GearScore
		if child == nil {
			child = &gearScore{ID: playerData.GearScore}
		}
		engine.assembleGearScorePath(child, p, pIndex+1, includedElements)
		element.GearScore = child
	case player_guildMembersIdentifier:
		ref := engine.playerGuildMemberRef(PlayerGuildMemberRefID(p[pIndex+1].refID)).playerGuildMemberRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[int(ref.ReferencedElementID)]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		referencedElement := engine.Player(ref.ReferencedElementID).player
		treeRef := elementReference{
			OperationKind:        ref.OperationKind,
			ElementID:            int(ref.ReferencedElementID),
			ElementKind:          ElementKindPlayer,
			ReferencedDataStatus: referencedDataStatus,
			ElementPath:          referencedElement.Path,
		}
		if element.GuildMembers == nil {
			element.GuildMembers = make(map[PlayerID]elementReference)
		}
		element.GuildMembers[referencedElement.ID] = treeRef
	case player_itemsIdentifier:
		if element.Items == nil {
			element.Items = make(map[ItemID]item)
		}
		child, ok := element.Items[ItemID(p[pIndex+1].id)]
		if !ok {
			child = item{ID: ItemID(p[pIndex+1].id)}
		}
		engine.assembleItemPath(&child, p, pIndex+1, includedElements)
		element.Items[child.ID] = child
	case player_positionIdentifier:
		child := element.Position
		if child == nil {
			child = &position{ID: playerData.Position}
		}
		engine.assemblePositionPath(child, p, pIndex+1, includedElements)
		element.Position = child
	case player_targetIdentifier:
		ref := engine.playerTargetRef(PlayerTargetRefID(p[pIndex+1].refID)).playerTargetRef
		if element.Target != nil && ref.OperationKind == OperationKindDelete {
			break
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[p[pIndex+1].id]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		// TODO sometimes player first, sometimes zoneItem
		fmt.Println(p[pIndex+1].kind, ref, playerData.Target)
		switch p[pIndex+1].kind {
		case ElementKindPlayer:
			referencedElement := engine.Player(PlayerID(p[pIndex+1].id)).player
			treeRef := elementReference{
				OperationKind:        ref.OperationKind,
				ElementID:            p[pIndex+1].id,
				ElementKind:          ElementKindPlayer,
				ReferencedDataStatus: referencedDataStatus,
				ElementPath:          referencedElement.Path,
			}
			element.Target = &treeRef
		case ElementKindZoneItem:
			referencedElement := engine.ZoneItem(ZoneItemID(p[pIndex+1].id)).zoneItem
			treeRef := elementReference{
				OperationKind:        ref.OperationKind,
				ElementID:            p[pIndex+1].id,
				ElementKind:          ElementKindZoneItem,
				ReferencedDataStatus: referencedDataStatus,
				ElementPath:          referencedElement.Path,
			}
			element.Target = &treeRef
		}
	case player_targetedByIdentifier:
		if element.TargetedBy == nil {
			element.TargetedBy = make(map[int]elementReference)
		}
		ref := engine.playerTargetedByRef(PlayerTargetedByRefID(p[pIndex+1].refID)).playerTargetedByRef
		referencedDataStatus := ReferencedDataUnchanged
		if _, ok := includedElements[p[pIndex+1].id]; ok {
			referencedDataStatus = ReferencedDataModified
		}
		switch p[pIndex+1].kind {
		case ElementKindPlayer:
			referencedElement := engine.Player(PlayerID(p[pIndex+1].id)).player
			treeRef := elementReference{
				OperationKind:        ref.OperationKind,
				ElementID:            p[pIndex+1].id,
				ElementKind:          ElementKindPlayer,
				ReferencedDataStatus: referencedDataStatus,
				ElementPath:          referencedElement.Path,
			}
			element.TargetedBy[p[pIndex+1].id] = treeRef
		case ElementKindZoneItem:
			referencedElement := engine.ZoneItem(ZoneItemID(p[pIndex+1].id)).zoneItem
			treeRef := elementReference{
				OperationKind:        ref.OperationKind,
				ElementID:            p[pIndex+1].id,
				ElementKind:          ElementKindPlayer,
				ReferencedDataStatus: referencedDataStatus,
				ElementPath:          referencedElement.Path,
			}
			element.TargetedBy[p[pIndex+1].id] = treeRef
		}
	}

	_ = playerData
}

func (engine *Engine) assembleZonePath(element *zone, p path, pIndex int, includedElements map[int]bool) {

	zoneData, ok := engine.Patch.Zone[element.ID]
	if !ok {
		zoneData = engine.State.Zone[element.ID]
	}

	element.OperationKind = zoneData.OperationKind
	element.Tags = zoneData.Tags[:]

	if pIndex+1 == len(p) {
		return
	}

	switch p[pIndex+1].identifier {
	case zone_interactablesIdentifier:
		if element.Interactables == nil {
			element.Interactables = make(map[int]interface{})
		}
		switch p[pIndex+1].kind {
		case ElementKindItem:
			child, ok := element.Interactables[p[pIndex+1].id].(item)
			if !ok {
				child = item{ID: ItemID(p[pIndex+1].id)}
			}
			engine.assembleItemPath(&child, p, pIndex+1, includedElements)
			element.Interactables[p[pIndex+1].id] = child
		case ElementKindPlayer:
			child, ok := element.Interactables[p[pIndex+1].id].(player)
			if !ok {
				child = player{ID: PlayerID(p[pIndex+1].id)}
			}
			engine.assemblePlayerPath(&child, p, pIndex+1, includedElements)
			element.Interactables[p[pIndex+1].id] = child
		case ElementKindZoneItem:
			child, ok := element.Interactables[p[pIndex+1].id].(zoneItem)
			if !ok {
				child = zoneItem{ID: ZoneItemID(p[pIndex+1].id)}
			}
			engine.assembleZoneItemPath(&child, p, pIndex+1, includedElements)
			element.Interactables[p[pIndex+1].id] = child
		}
	case zone_itemsIdentifier:
		if element.Items == nil {
			element.Items = make(map[ZoneItemID]zoneItem)
		}
		child, ok := element.Items[ZoneItemID(p[pIndex+1].id)]
		if !ok {
			child = zoneItem{ID: ZoneItemID(p[pIndex+1].id)}
		}
		engine.assembleZoneItemPath(&child, p, pIndex+1, includedElements)
		element.Items[child.ID] = child
	case zone_playersIdentifier:
		if element.Players == nil {
			element.Players = make(map[PlayerID]player)
		}
		child, ok := element.Players[PlayerID(p[pIndex+1].id)]
		if !ok {
			child = player{ID: PlayerID(p[pIndex+1].id)}
		}
		engine.assemblePlayerPath(&child, p, pIndex+1, includedElements)
		element.Players[child.ID] = child
	}

	_ = zoneData
}

// 1. get all basic elements and references out of patch, put their paths in updatedReferencePaths
// 2. go through all paths and put ids in includedElements map, save len(includedElements)
// 3. get all references out of STATE, check if they reference element in includedElements, if TRUE put reference path into updatedByReferecePaths
// 4. if len(updatedByReferecePaths) != 0: go through all updatedByReferecePaths and put ids in includedElements map, ELSE continue with 6.
// 5. back to step 3.

// TODO PROBLEM?? if a reference is Set and then Unset and Set again, does that mess with pathuilding, as referenceID might be 0 in patch or state
// TODO what happens if you call SetPlayer? will the path be built with a player-update or a position-delete??

func (engine *Engine) assembleUpdateTree() Tree {

	for key := range engine.Tree.EquipmentSet {
		delete(engine.Tree.EquipmentSet, key)
	}
	for key := range engine.Tree.GearScore {
		delete(engine.Tree.GearScore, key)
	}
	for key := range engine.Tree.Item {
		delete(engine.Tree.Item, key)
	}
	for key := range engine.Tree.Player {
		delete(engine.Tree.Player, key)
	}
	for key := range engine.Tree.Position {
		delete(engine.Tree.Position, key)
	}
	for key := range engine.Tree.Zone {
		delete(engine.Tree.Zone, key)
	}
	for key := range engine.Tree.ZoneItem {
		delete(engine.Tree.ZoneItem, key)
	}

	updatedReferencePaths := make(map[int]path)
	updatedElementPaths := make(map[int]path)
	// TODO possibly big performance boost
	// updatedElements := make(map[int]bool)

	for _, equipmentSet := range engine.Patch.EquipmentSet {
		updatedElementPaths[int(equipmentSet.ID)] = equipmentSet.path
	}
	for _, gearScore := range engine.Patch.GearScore {
		updatedElementPaths[int(gearScore.ID)] = gearScore.path
	}
	for _, item := range engine.Patch.Item {
		updatedElementPaths[int(item.ID)] = item.path
	}
	for _, player := range engine.Patch.Player {
		updatedElementPaths[int(player.ID)] = player.path
	}
	for _, position := range engine.Patch.Position {
		updatedElementPaths[int(position.ID)] = position.path
	}
	for _, zone := range engine.Patch.Zone {
		updatedElementPaths[int(zone.ID)] = zone.path
	}
	for _, zoneItem := range engine.Patch.ZoneItem {
		updatedElementPaths[int(zoneItem.ID)] = zoneItem.path
	}
	for _, equipmentSetEquipmentRef := range engine.Patch.EquipmentSetEquipmentRef {
		updatedReferencePaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
	}
	for _, itemBoundToRef := range engine.Patch.ItemBoundToRef {
		updatedReferencePaths[int(itemBoundToRef.ID)] = itemBoundToRef.path
	}
	for _, playerEquipmentSetRef := range engine.Patch.PlayerEquipmentSetRef {
		updatedReferencePaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
	}
	for _, playerGuildMemberRef := range engine.Patch.PlayerGuildMemberRef {
		updatedReferencePaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
	}
	for _, playerTargetRef := range engine.Patch.PlayerTargetRef {
		updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
	}
	for _, playerTargetedByRef := range engine.Patch.PlayerTargetedByRef {
		updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
	}

	includedElements := make(map[int]bool)

	prevousLength := 0
	for {
		for _, p := range updatedElementPaths {
			for _, seg := range p {
				includedElements[seg.id] = true
			}
		}
		for _, p := range updatedReferencePaths {
			for _, seg := range p {
				if seg.refID != 0 {
					includedElements[seg.refID] = true
				} else {
					includedElements[seg.id] = true
				}
			}
		}

		if prevousLength == len(includedElements) {
			break
		}

		prevousLength = len(includedElements)

		for _, equipmentSetEquipmentRef := range engine.Patch.EquipmentSetEquipmentRef {
			if _, ok := includedElements[int(equipmentSetEquipmentRef.ReferencedElementID)]; ok {
				updatedReferencePaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
			}
		}
		for _, equipmentSetEquipmentRef := range engine.State.EquipmentSetEquipmentRef {
			if _, ok := updatedReferencePaths[int(equipmentSetEquipmentRef.ID)]; ok {
				continue
			}
			if _, ok := includedElements[int(equipmentSetEquipmentRef.ReferencedElementID)]; ok {
				updatedReferencePaths[int(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
			}
		}

		for _, itemBoundToRef := range engine.Patch.ItemBoundToRef {
			if _, ok := includedElements[int(itemBoundToRef.ReferencedElementID)]; ok {
				updatedReferencePaths[int(itemBoundToRef.ID)] = itemBoundToRef.path
			}
		}
		for _, itemBoundToRef := range engine.State.ItemBoundToRef {
			if _, ok := updatedReferencePaths[int(itemBoundToRef.ID)]; ok {
				continue
			}
			if _, ok := includedElements[int(itemBoundToRef.ReferencedElementID)]; ok {
				updatedReferencePaths[int(itemBoundToRef.ID)] = itemBoundToRef.path
			}
		}

		for _, playerEquipmentSetRef := range engine.Patch.PlayerEquipmentSetRef {
			if _, ok := includedElements[int(playerEquipmentSetRef.ReferencedElementID)]; ok {
				updatedReferencePaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
			}
		}
		for _, playerEquipmentSetRef := range engine.State.PlayerEquipmentSetRef {
			if _, ok := updatedReferencePaths[int(playerEquipmentSetRef.ID)]; ok {
				continue
			}
			if _, ok := includedElements[int(playerEquipmentSetRef.ReferencedElementID)]; ok {
				updatedReferencePaths[int(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
			}
		}

		for _, playerGuildMemberRef := range engine.Patch.PlayerGuildMemberRef {
			if _, ok := includedElements[int(playerGuildMemberRef.ReferencedElementID)]; ok {
				updatedReferencePaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
			}
		}
		for _, playerGuildMemberRef := range engine.State.PlayerGuildMemberRef {
			if _, ok := updatedReferencePaths[int(playerGuildMemberRef.ID)]; ok {
				continue
			}
			if _, ok := includedElements[int(playerGuildMemberRef.ReferencedElementID)]; ok {
				updatedReferencePaths[int(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
			}
		}

		for _, playerTargetRef := range engine.Patch.PlayerTargetRef {
			anyContainer := engine.anyOfPlayer_ZoneItem(playerTargetRef.ReferencedElementID)
			switch anyContainer.anyOfPlayer_ZoneItem.ElementKind {
			case ElementKindPlayer:
				if _, ok := includedElements[int(anyContainer.anyOfPlayer_ZoneItem.Player)]; ok {
					updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
				}
			case ElementKindZoneItem:
				if _, ok := includedElements[int(anyContainer.anyOfPlayer_ZoneItem.ZoneItem)]; ok {
					updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
				}
			}
		}
		for _, playerTargetRef := range engine.State.PlayerTargetRef {
			if _, ok := updatedReferencePaths[int(playerTargetRef.ID)]; ok {
				continue
			}
			anyContainer := engine.anyOfPlayer_ZoneItem(playerTargetRef.ReferencedElementID)
			switch anyContainer.anyOfPlayer_ZoneItem.ElementKind {
			case ElementKindPlayer:
				if _, ok := includedElements[int(anyContainer.anyOfPlayer_ZoneItem.Player)]; ok {
					updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
				}
			case ElementKindZoneItem:
				if _, ok := includedElements[int(anyContainer.anyOfPlayer_ZoneItem.Player)]; ok {
					updatedReferencePaths[int(playerTargetRef.ID)] = playerTargetRef.path
				}
			}
		}

		for _, playerTargetedByRef := range engine.Patch.PlayerTargetedByRef {
			anyContainer := engine.anyOfPlayer_ZoneItem(playerTargetedByRef.ReferencedElementID)
			switch anyContainer.anyOfPlayer_ZoneItem.ElementKind {
			case ElementKindPlayer:
				if _, ok := includedElements[int(anyContainer.anyOfPlayer_ZoneItem.Player)]; ok {
					updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
				}
			case ElementKindZoneItem:
				if _, ok := includedElements[int(anyContainer.anyOfPlayer_ZoneItem.Player)]; ok {
					updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
				}
			}
		}
		for _, playerTargetedByRef := range engine.State.PlayerTargetedByRef {
			if _, ok := updatedReferencePaths[int(playerTargetedByRef.ID)]; ok {
				continue
			}
			anyContainer := engine.anyOfPlayer_ZoneItem(playerTargetedByRef.ReferencedElementID)
			switch anyContainer.anyOfPlayer_ZoneItem.ElementKind {
			case ElementKindPlayer:
				if _, ok := includedElements[int(anyContainer.anyOfPlayer_ZoneItem.Player)]; ok {
					updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
				}
			case ElementKindZoneItem:
				if _, ok := includedElements[int(anyContainer.anyOfPlayer_ZoneItem.Player)]; ok {
					updatedReferencePaths[int(playerTargetedByRef.ID)] = playerTargetedByRef.path
				}
			}
		}

	}

	updatedPaths := make(map[int]path)
	for id, path := range updatedElementPaths {
		updatedPaths[id] = path
	}
	for id, path := range updatedReferencePaths {
		updatedPaths[id] = path
	}

	fmt.Println(updatedPaths)

	for _, elementPath := range updatedPaths {
		switch elementPath[0].identifier {
		case equipmentSetIdentifier:
			child, ok := engine.Tree.EquipmentSet[EquipmentSetID(elementPath[0].id)]
			if !ok {
				child = equipmentSet{ID: EquipmentSetID(elementPath[0].id)}
			}
			engine.assembleEquipmentSetPath(&child, elementPath, 0, includedElements)
			engine.Tree.EquipmentSet[EquipmentSetID(elementPath[0].id)] = child
		case gearScoreIdentifier:
			child, ok := engine.Tree.GearScore[GearScoreID(elementPath[0].id)]
			if !ok {
				child = gearScore{ID: GearScoreID(elementPath[0].id)}
			}
			engine.assembleGearScorePath(&child, elementPath, 0, includedElements)
			engine.Tree.GearScore[GearScoreID(elementPath[0].id)] = child
		case itemIdentifier:
			child, ok := engine.Tree.Item[ItemID(elementPath[0].id)]
			if !ok {
				child = item{ID: ItemID(elementPath[0].id)}
			}
			engine.assembleItemPath(&child, elementPath, 0, includedElements)
			engine.Tree.Item[ItemID(elementPath[0].id)] = child
		case playerIdentifier:
			child, ok := engine.Tree.Player[PlayerID(elementPath[0].id)]
			if !ok {
				child = player{ID: PlayerID(elementPath[0].id)}
			}
			engine.assemblePlayerPath(&child, elementPath, 0, includedElements)
			engine.Tree.Player[PlayerID(elementPath[0].id)] = child
		case positionIdentifier:
			child, ok := engine.Tree.Position[PositionID(elementPath[0].id)]
			if !ok {
				child = position{ID: PositionID(elementPath[0].id)}
			}
			engine.assemblePositionPath(&child, elementPath, 0, includedElements)
			engine.Tree.Position[PositionID(elementPath[0].id)] = child
		case zoneIdentifier:
			child, ok := engine.Tree.Zone[ZoneID(elementPath[0].id)]
			if !ok {
				child = zone{ID: ZoneID(elementPath[0].id)}
			}
			engine.assembleZonePath(&child, elementPath, 0, includedElements)
			engine.Tree.Zone[ZoneID(elementPath[0].id)] = child
		case zoneItemIdentifier:
			child, ok := engine.Tree.ZoneItem[ZoneItemID(elementPath[0].id)]
			if !ok {
				child = zoneItem{ID: ZoneItemID(elementPath[0].id)}
			}
			engine.assembleZoneItemPath(&child, elementPath, 0, includedElements)
			engine.Tree.ZoneItem[ZoneItemID(elementPath[0].id)] = child
		}
	}

	return engine.Tree
}
