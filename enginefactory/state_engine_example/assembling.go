package state

type assembleConfig struct {
	forceInclude bool
}

func (se *Engine) assembleGearScore(gearScoreID GearScoreID, check *recursionCheck, config assembleConfig) (GearScore, bool) {
	if check != nil {
		if alreadyExists := check.gearScore[gearScoreID]; alreadyExists {
			return GearScore{}, false
		} else {
			check.gearScore[gearScoreID] = true
		}
	}

	gearScoreData, hasUpdated := se.Patch.GearScore[gearScoreID]
	if !hasUpdated && !config.forceInclude {
		return GearScore{}, false
	}

	var gearScore GearScore

	gearScore.ID = gearScoreData.ID
	gearScore.OperationKind_ = gearScoreData.OperationKind_
	gearScore.Level = gearScoreData.Level
	gearScore.Score = gearScoreData.Score
	return gearScore, true
}

func (se *Engine) assemblePosition(positionID PositionID, check *recursionCheck, config assembleConfig) (Position, bool) {
	if check != nil {
		if alreadyExists := check.position[positionID]; alreadyExists {
			return Position{}, false
		} else {
			check.position[positionID] = true
		}
	}

	positionData, hasUpdated := se.Patch.Position[positionID]
	if !hasUpdated && !config.forceInclude {
		return Position{}, false
	}

	var position Position

	position.ID = positionData.ID
	position.OperationKind_ = positionData.OperationKind_
	position.X = positionData.X
	position.Y = positionData.Y
	return position, true
}

func (se *Engine) assembleEquipmentSet(equipmentSetID EquipmentSetID, check *recursionCheck, config assembleConfig) (EquipmentSet, bool) {
	if check != nil {
		if alreadyExists := check.equipmentSet[equipmentSetID]; alreadyExists {
			return EquipmentSet{}, false
		} else {
			check.equipmentSet[equipmentSetID] = true
		}
	}

	equipmentSetData, hasUpdated := se.Patch.EquipmentSet[equipmentSetID]
	if !hasUpdated {
		equipmentSetData = se.State.EquipmentSet[equipmentSetID]
	}

	var equipmentSet EquipmentSet

	for _, refID := range mergeEquipmentSetEquipmentRefIDs(se.State.EquipmentSet[equipmentSetID].Equipment, se.Patch.EquipmentSet[equipmentSetID].Equipment) {
		if ref, refHasUpdated := se.assembleEquipmentSetEquipmentRef(refID, check, config); refHasUpdated {
			hasUpdated = true
			equipmentSet.Equipment = append(equipmentSet.Equipment, ref)
		}
	}

	equipmentSet.ID = equipmentSetData.ID
	equipmentSet.OperationKind_ = equipmentSetData.OperationKind_
	equipmentSet.Name = equipmentSetData.Name
	return equipmentSet, hasUpdated || config.forceInclude
}

func (se *Engine) assembleItem(itemID ItemID, check *recursionCheck, config assembleConfig) (Item, bool) {
	if check != nil {
		if alreadyExists := check.item[itemID]; alreadyExists {
			return Item{}, false
		} else {
			check.item[itemID] = true
		}
	}

	itemData, hasUpdated := se.Patch.Item[itemID]
	if !hasUpdated {
		itemData = se.State.Item[itemID]
	}

	var item Item

	if refs, refHasUpdated := se.assembleItemBoundToRef(itemID, check, config); refHasUpdated {
		item.BoundTo = refs
		hasUpdated = true
	}
	if treeGearScore, gearScoreHasUpdated := se.assembleGearScore(itemData.GearScore, check, config); gearScoreHasUpdated {
		hasUpdated = true
		item.GearScore = &treeGearScore
	}

	item.ID = itemData.ID
	item.OperationKind_ = itemData.OperationKind_
	item.Name = itemData.Name
	return item, hasUpdated || config.forceInclude
}

func (se *Engine) assembleZoneItem(zoneItemID ZoneItemID, check *recursionCheck, config assembleConfig) (ZoneItem, bool) {
	if check != nil {
		if alreadyExists := check.zoneItem[zoneItemID]; alreadyExists {
			return ZoneItem{}, false
		} else {
			check.zoneItem[zoneItemID] = true
		}
	}

	zoneItemData, hasUpdated := se.Patch.ZoneItem[zoneItemID]
	if !hasUpdated {
		zoneItemData = se.State.ZoneItem[zoneItemID]
	}

	var zoneItem ZoneItem

	if treeItem, itemHasUpdated := se.assembleItem(zoneItemData.Item, check, config); itemHasUpdated {
		hasUpdated = true
		zoneItem.Item = &treeItem
	}
	if treePosition, positionHasUpdated := se.assemblePosition(zoneItemData.Position, check, config); positionHasUpdated {
		hasUpdated = true
		zoneItem.Position = &treePosition
	}

	zoneItem.ID = zoneItemData.ID
	zoneItem.OperationKind_ = zoneItemData.OperationKind_
	return zoneItem, hasUpdated || config.forceInclude

}

func (se *Engine) assemblePlayer(playerID PlayerID, check *recursionCheck, config assembleConfig) (Player, bool) {
	if check != nil {
		if alreadyExists := check.player[playerID]; alreadyExists {
			return Player{}, false
		} else {
			check.player[playerID] = true
		}
	}

	playerData, hasUpdated := se.Patch.Player[playerID]
	if !hasUpdated {
		playerData = se.State.Player[playerID]
	}

	var player Player

	for _, refID := range mergePlayerEquipmentSetRefIDs(se.State.Player[playerID].EquipmentSets, se.Patch.Player[playerID].EquipmentSets) {
		if ref, refHasUpdated := se.assemblePlayerEquipmentSetRef(refID, check, config); refHasUpdated {
			hasUpdated = true
			player.EquipmentSets = append(player.EquipmentSets, ref)
		}
	}
	if treeGearScore, gearScoreHasUpdated := se.assembleGearScore(playerData.GearScore, check, config); gearScoreHasUpdated {
		hasUpdated = true
		player.GearScore = &treeGearScore
	}
	for _, refID := range mergePlayerGuildMemberRefIDs(se.State.Player[playerID].GuildMembers, se.Patch.Player[playerID].GuildMembers) {
		if ref, refHasUpdated := se.assemblePlayerGuildMemberRef(refID, check, config); refHasUpdated {
			hasUpdated = true
			player.GuildMembers = append(player.GuildMembers, ref)
		}
	}
	for _, itemID := range mergeItemIDs(se.State.Player[playerData.ID].Items, se.Patch.Player[playerData.ID].Items) {
		if treeItem, itemHasUpdated := se.assembleItem(itemID, check, config); itemHasUpdated {
			hasUpdated = true
			player.Items = append(player.Items, treeItem)
		}
	}
	if treePosition, positionHasUpdated := se.assemblePosition(playerData.Position, check, config); positionHasUpdated {
		hasUpdated = true
		player.Position = &treePosition
	}

	player.ID = playerData.ID
	player.OperationKind_ = playerData.OperationKind_
	return player, hasUpdated || config.forceInclude
}

func (se *Engine) assembleZone(zoneID ZoneID, check *recursionCheck, config assembleConfig) (Zone, bool) {
	if check != nil {
		if alreadyExists := check.zone[zoneID]; alreadyExists {
			return Zone{}, false
		} else {
			check.zone[zoneID] = true
		}
	}

	zoneData, hasUpdated := se.Patch.Zone[zoneID]
	if !hasUpdated {
		zoneData = se.State.Zone[zoneID]
	}

	var zone Zone

	for _, zoneItemID := range mergeZoneItemIDs(se.State.Zone[zoneData.ID].Items, se.Patch.Zone[zoneData.ID].Items) {
		if treeZoneItem, zoneItemHasUpdated := se.assembleZoneItem(zoneItemID, check, config); zoneItemHasUpdated {
			hasUpdated = true
			zone.Items = append(zone.Items, treeZoneItem)
		}
	}
	for _, playerID := range mergePlayerIDs(se.State.Zone[zoneData.ID].Players, se.Patch.Zone[zoneData.ID].Players) {
		if treePlayer, playerHasUpdated := se.assemblePlayer(playerID, check, config); playerHasUpdated {
			hasUpdated = true
			zone.Players = append(zone.Players, treePlayer)
		}
	}

	zone.ID = zoneData.ID
	zone.OperationKind_ = zoneData.OperationKind_
	zone.Tags = zoneData.Tags
	return zone, hasUpdated || config.forceInclude
}

func (se *Engine) assembleItemBoundToRef(itemID ItemID, check *recursionCheck, config assembleConfig) (*ElementReference, bool) {
	stateItem := se.State.Item[itemID]
	patchItem, itemIsInPatch := se.Patch.Item[itemID]

	// ref not set at all
	if stateItem.BoundTo == 0 && (!itemIsInPatch || patchItem.BoundTo == 0) {
		return nil, false
	}

	// force include
	if config.forceInclude {
		ref := se.itemBoundToRef(patchItem.BoundTo)
		referencedElement := se.Player(ref.itemBoundToRef.ReferencedElementID).player
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, hasUpdatedDownstream := se.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		return &ElementReference{ref.itemBoundToRef.OperationKind_, int(referencedElement.ID), ElementKindPlayer, referencedDataStatus}, true
	}

	// ref was definitely created
	if stateItem.BoundTo == 0 && (itemIsInPatch && patchItem.BoundTo != 0) {
		config.forceInclude = true
		ref := se.itemBoundToRef(patchItem.BoundTo)
		referencedElement := se.Player(ref.itemBoundToRef.ReferencedElementID).player
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, hasUpdatedDownstream := se.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		return &ElementReference{OperationKindUpdate, int(referencedElement.ID), ElementKindPlayer, referencedDataStatus}, true
	}

	// ref was definitely removed
	if stateItem.BoundTo != 0 && (itemIsInPatch && patchItem.BoundTo == 0) {
		ref := se.itemBoundToRef(stateItem.BoundTo)
		referencedElement := se.Player(ref.itemBoundToRef.ReferencedElementID).player
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, hasUpdatedDownstream := se.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		return &ElementReference{OperationKindDelete, int(referencedElement.ID), ElementKindPlayer, referencedDataStatus}, true
	}

	// immediate replacement of refs
	if stateItem.BoundTo != 0 && (itemIsInPatch && patchItem.BoundTo != 0) {
		if stateItem.BoundTo != patchItem.BoundTo {
			ref := se.itemBoundToRef(patchItem.BoundTo)
			referencedElement := se.Player(ref.itemBoundToRef.ReferencedElementID).player
			if check == nil {
				check = newRecursionCheck()
			}
			referencedDataStatus := ReferencedDataUnchanged
			if _, hasUpdatedDownstream := se.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
				referencedDataStatus = ReferencedDataModified
			}
			return &ElementReference{OperationKindUpdate, int(referencedElement.ID), ElementKindPlayer, referencedDataStatus}, true
		}
	}

	// OperationKindUpdate element got updated
	if stateItem.BoundTo != 0 {
		ref := se.itemBoundToRef(stateItem.BoundTo)
		if check == nil {
			check = newRecursionCheck()
		}
		if _, hasUpdatedDownstream := se.assemblePlayer(ref.ID(), check, config); hasUpdatedDownstream {
			return &ElementReference{OperationKindUnchanged, int(ref.ID()), ElementKindPlayer, ReferencedDataModified}, true
		}
	}

	return nil, false
}

func (se *Engine) assemblePlayerGuildMemberRef(refID PlayerGuildMemberRefID, check *recursionCheck, config assembleConfig) (ElementReference, bool) {
	if config.forceInclude {
		ref := se.playerGuildMemberRef(refID).playerGuildMemberRef
		referencedElement := se.Player(ref.ReferencedElementID).player
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, hasUpdatedDownstream := se.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		return ElementReference{ref.OperationKind_, int(ref.ReferencedElementID), ElementKindPlayer, referencedDataStatus}, true
	}

	if patchRef, hasUpdated := se.Patch.PlayerGuildMemberRef[refID]; hasUpdated || config.forceInclude {
		referencedElement := se.Player(patchRef.ReferencedElementID).player
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, hasUpdatedDownstream := se.assemblePlayer(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		if patchRef.OperationKind_ == OperationKindUpdate {
			config.forceInclude = true
		}
		return ElementReference{patchRef.OperationKind_, int(patchRef.ReferencedElementID), ElementKindPlayer, referencedDataStatus}, true
	}

	ref := se.playerGuildMemberRef(refID).playerGuildMemberRef
	if check == nil {
		check = newRecursionCheck()
	}
	if _, hasUpdatedDownstream := se.assemblePlayer(ref.ReferencedElementID, check, config); hasUpdatedDownstream {
		return ElementReference{OperationKindUnchanged, int(ref.ReferencedElementID), ElementKindPlayer, ReferencedDataModified}, true
	}

	return ElementReference{}, false
}

func (se *Engine) assemblePlayerEquipmentSetRef(refID PlayerEquipmentSetRefID, check *recursionCheck, config assembleConfig) (ElementReference, bool) {
	if config.forceInclude {
		ref := se.playerEquipmentSetRef(refID).playerEquipmentSetRef
		referencedElement := se.EquipmentSet(ref.ReferencedElementID).equipmentSet
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, hasUpdatedDownstream := se.assembleEquipmentSet(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		return ElementReference{ref.OperationKind_, int(ref.ReferencedElementID), ElementKindEquipmentSet, referencedDataStatus}, true
	}

	if patchRef, hasUpdated := se.Patch.PlayerEquipmentSetRef[refID]; hasUpdated {
		referencedElement := se.EquipmentSet(patchRef.ReferencedElementID).equipmentSet
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, hasUpdatedDownstream := se.assembleEquipmentSet(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		if patchRef.OperationKind_ == OperationKindUpdate {
			config.forceInclude = true
		}
		return ElementReference{patchRef.OperationKind_, int(patchRef.ReferencedElementID), ElementKindEquipmentSet, referencedDataStatus}, true
	}

	ref := se.playerEquipmentSetRef(refID).playerEquipmentSetRef
	if check == nil {
		check = newRecursionCheck()
	}
	if _, hasUpdatedDownstream := se.assembleEquipmentSet(ref.ReferencedElementID, check, config); hasUpdatedDownstream {
		return ElementReference{OperationKindUnchanged, int(ref.ReferencedElementID), ElementKindEquipmentSet, ReferencedDataModified}, true
	}

	return ElementReference{}, false
}

func (se *Engine) assembleEquipmentSetEquipmentRef(refID EquipmentSetEquipmentRefID, check *recursionCheck, config assembleConfig) (ElementReference, bool) {
	if config.forceInclude {
		ref := se.equipmentSetEquipmentRef(refID).equipmentSetEquipmentRef
		referencedElement := se.Item(ref.ReferencedElementID).item
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, hasUpdatedDownstream := se.assembleItem(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		return ElementReference{ref.OperationKind_, int(ref.ReferencedElementID), ElementKindItem, referencedDataStatus}, true
	}

	if patchRef, hasUpdated := se.Patch.EquipmentSetEquipmentRef[refID]; hasUpdated {
		referencedElement := se.Item(patchRef.ReferencedElementID).item
		if check == nil {
			check = newRecursionCheck()
		}
		referencedDataStatus := ReferencedDataUnchanged
		if _, hasUpdatedDownstream := se.assembleItem(referencedElement.ID, check, config); hasUpdatedDownstream {
			referencedDataStatus = ReferencedDataModified
		}
		if patchRef.OperationKind_ == OperationKindUpdate {
			config.forceInclude = true
		}
		return ElementReference{patchRef.OperationKind_, int(patchRef.ReferencedElementID), ElementKindItem, referencedDataStatus}, true
	}

	ref := se.equipmentSetEquipmentRef(refID).equipmentSetEquipmentRef
	if check == nil {
		check = newRecursionCheck()
	}
	if _, hasUpdatedDownstream := se.assembleItem(ref.ReferencedElementID, check, config); hasUpdatedDownstream {
		return ElementReference{OperationKindUnchanged, int(ref.ReferencedElementID), ElementKindItem, ReferencedDataModified}, true
	}

	return ElementReference{}, false
}

func (se *Engine) assembleTree() Tree {

	config := assembleConfig{
		forceInclude: false,
	}

	for _, equipmentSetData := range se.Patch.EquipmentSet {
		equipmentSet, hasUpdated := se.assembleEquipmentSet(equipmentSetData.ID, nil, config)
		if hasUpdated {
			se.Tree.EquipmentSet[equipmentSetData.ID] = equipmentSet
		}
	}
	for _, gearScoreData := range se.Patch.GearScore {
		if !gearScoreData.HasParent_ {
			gearScore, hasUpdated := se.assembleGearScore(gearScoreData.ID, nil, config)
			if hasUpdated {
				se.Tree.GearScore[gearScoreData.ID] = gearScore
			}
		}
	}
	for _, itemData := range se.Patch.Item {
		if !itemData.HasParent_ {
			item, hasUpdated := se.assembleItem(itemData.ID, nil, config)
			if hasUpdated {
				se.Tree.Item[itemData.ID] = item
			}
		}
	}
	for _, playerData := range se.Patch.Player {
		if !playerData.HasParent_ {
			player, hasUpdated := se.assemblePlayer(playerData.ID, nil, config)
			if hasUpdated {
				se.Tree.Player[playerData.ID] = player
			}
		}
	}
	for _, positionData := range se.Patch.Position {
		if !positionData.HasParent_ {
			position, hasUpdated := se.assemblePosition(positionData.ID, nil, config)
			if hasUpdated {
				se.Tree.Position[positionData.ID] = position
			}
		}
	}
	for _, zoneData := range se.Patch.Zone {
		zone, hasUpdated := se.assembleZone(zoneData.ID, nil, config)
		if hasUpdated {
			se.Tree.Zone[zoneData.ID] = zone
		}
	}
	for _, zoneItemData := range se.Patch.ZoneItem {
		if !zoneItemData.HasParent_ {
			zoneItem, hasUpdated := se.assembleZoneItem(zoneItemData.ID, nil, config)
			if hasUpdated {
				se.Tree.ZoneItem[zoneItemData.ID] = zoneItem
			}
		}
	}

	for _, equipmentSetData := range se.State.EquipmentSet {
		if _, ok := se.Tree.EquipmentSet[equipmentSetData.ID]; !ok {
			equipmentSet, hasUpdated := se.assembleEquipmentSet(equipmentSetData.ID, nil, config)
			if hasUpdated {
				se.Tree.EquipmentSet[equipmentSetData.ID] = equipmentSet
			}
		}
	}
	for _, gearScoreData := range se.State.GearScore {
		if !gearScoreData.HasParent_ {
			if _, ok := se.Tree.GearScore[gearScoreData.ID]; !ok {
				gearScore, hasUpdated := se.assembleGearScore(gearScoreData.ID, nil, config)
				if hasUpdated {
					se.Tree.GearScore[gearScoreData.ID] = gearScore
				}
			}
		}
	}
	for _, itemData := range se.State.Item {
		if !itemData.HasParent_ {
			if _, ok := se.Tree.Item[itemData.ID]; !ok {
				item, hasUpdated := se.assembleItem(itemData.ID, nil, config)
				if hasUpdated {
					se.Tree.Item[itemData.ID] = item
				}
			}
		}
	}
	for _, playerData := range se.State.Player {
		if !playerData.HasParent_ {
			if _, ok := se.Tree.Player[playerData.ID]; !ok {
				player, hasUpdated := se.assemblePlayer(playerData.ID, nil, config)
				if hasUpdated {
					se.Tree.Player[playerData.ID] = player
				}
			}
		}
	}
	for _, positionData := range se.State.Position {
		if !positionData.HasParent_ {
			if _, ok := se.Tree.Position[positionData.ID]; !ok {
				position, hasUpdated := se.assemblePosition(positionData.ID, nil, config)
				if hasUpdated {
					se.Tree.Position[positionData.ID] = position
				}
			}
		}
	}
	for _, zoneData := range se.State.Zone {
		if _, ok := se.Tree.Zone[zoneData.ID]; !ok {
			zone, hasUpdated := se.assembleZone(zoneData.ID, nil, config)
			if hasUpdated {
				se.Tree.Zone[zoneData.ID] = zone
			}
		}
	}
	for _, zoneItemData := range se.State.ZoneItem {
		if !zoneItemData.HasParent_ {
			if _, ok := se.Tree.ZoneItem[zoneItemData.ID]; !ok {
				zoneItem, hasUpdated := se.assembleZoneItem(zoneItemData.ID, nil, config)
				if hasUpdated {
					se.Tree.ZoneItem[zoneItemData.ID] = zoneItem
				}
			}
		}
	}

	return se.Tree
}
