package state

type assemblePlanner struct {
	updatedPaths          []path
	updatedReferencePaths map[ComplexID]path
	updatedElementPaths   map[int]path
	includedElements      map[int]bool // used to determine ReferencedDataStatus during assembling
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
	// we want to find all nodes which have updated and collect their paths.
	// later we will loop over the paths we have collected, and "walk" them (assembleBranch)
	// in order to assemble the tree from top to bottom (leaf nodes to root nodes)
	for _, boolValue := range patch.BoolValue {
		ap.updatedElementPaths[int(boolValue.ID)] = boolValue.path
	}
	for _, floatValue := range patch.FloatValue {
		ap.updatedElementPaths[int(floatValue.ID)] = floatValue.path
	}
	for _, intValue := range patch.IntValue {
		ap.updatedElementPaths[int(intValue.ID)] = intValue.path
	}
	for _, stringValue := range patch.StringValue {
		ap.updatedElementPaths[int(stringValue.ID)] = stringValue.path
	}
	for _, attackEvent := range patch.AttackEvent {
		ap.updatedElementPaths[int(attackEvent.ID)] = attackEvent.path
	}
	for _, equipmentSet := range patch.EquipmentSet {
		ap.updatedElementPaths[int(equipmentSet.ID)] = equipmentSet.path
	}
	for _, gearScore := range patch.GearScore {
		ap.updatedElementPaths[int(gearScore.ID)] = gearScore.path
	}
	for _, item := range patch.Item {
		ap.updatedElementPaths[int(item.ID)] = item.path
	}
	for _, player := range patch.Player {
		ap.updatedElementPaths[int(player.ID)] = player.path
	}
	for _, position := range patch.Position {
		ap.updatedElementPaths[int(position.ID)] = position.path
	}
	for _, zone := range patch.Zone {
		ap.updatedElementPaths[int(zone.ID)] = zone.path
	}
	for _, zoneItem := range patch.ZoneItem {
		ap.updatedElementPaths[int(zoneItem.ID)] = zoneItem.path
	}
	for _, attackEventTargetRef := range patch.AttackEventTargetRef {
		ap.updatedReferencePaths[ComplexID(attackEventTargetRef.ID)] = attackEventTargetRef.path
	}
	for _, equipmentSetEquipmentRef := range patch.EquipmentSetEquipmentRef {
		ap.updatedReferencePaths[ComplexID(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
	}
	for _, itemBoundToRef := range patch.ItemBoundToRef {
		ap.updatedReferencePaths[ComplexID(itemBoundToRef.ID)] = itemBoundToRef.path
	}
	for _, playerEquipmentSetRef := range patch.PlayerEquipmentSetRef {
		ap.updatedReferencePaths[ComplexID(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
	}
	for _, playerGuildMemberRef := range patch.PlayerGuildMemberRef {
		ap.updatedReferencePaths[ComplexID(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
	}
	for _, playerTargetRef := range patch.PlayerTargetRef {
		ap.updatedReferencePaths[ComplexID(playerTargetRef.ID)] = playerTargetRef.path
	}
	for _, playerTargetedByRef := range patch.PlayerTargetedByRef {
		ap.updatedReferencePaths[ComplexID(playerTargetedByRef.ID)] = playerTargetedByRef.path
	}

	previousLen := 0
	// we'd be pretty much done collecting the required paths, but we also want to
	// build all paths ending with a reference which references an updated element.
	// (and want to populate includedElements map of course)
	// this needs to happen recursively, consider this example (-> = reference):
	// A -> B -> C -> ^D
	// Since "D" has updated (^), but no other element, we'd only include "C".
	// However, now that "C" is considered updated by reference, we also want
	// to include "B". This is why recursiveness is required.
	for {
		// here we populate out includedElements with all newly collected paths segments
		// so we can check if any of these  elements are referenced by any reference
		// in the loop below
		for _, p := range ap.updatedElementPaths {
			for _, seg := range p {
				ap.includedElements[seg.id] = true
			}
		}
		// add all elements of the updated reference paths to the includedElements
		for _, p := range ap.updatedReferencePaths {
			for _, seg := range p {
				if seg.refID != (ComplexID{}) {
					// ommitting ref segments as the actual element is already included by
					// the previous segment.
				} else {
					ap.includedElements[seg.id] = true
				}
			}
		}

		// we check if any new elements are involved, which could
		// mean that new paths containing references need to be looked at
		if previousLen == len(ap.includedElements) {
			break
		}

		previousLen = len(ap.includedElements)

		for _, attackEventTargetRef := range patch.AttackEventTargetRef {
			if _, ok := ap.includedElements[int(attackEventTargetRef.ReferencedElementID)]; ok {
				ap.updatedReferencePaths[ComplexID(attackEventTargetRef.ID)] = attackEventTargetRef.path
			}
		}
		// again, events can't ever be present in state, but well keep this for consistency
		for _, attackEventTargetRef := range state.AttackEventTargetRef {
			if _, ok := ap.updatedReferencePaths[ComplexID(attackEventTargetRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(attackEventTargetRef.ReferencedElementID)]; ok {
				ap.updatedReferencePaths[ComplexID(attackEventTargetRef.ID)] = attackEventTargetRef.path
			}
		}

		for _, equipmentSetEquipmentRef := range patch.EquipmentSetEquipmentRef {
			// if the reference references an element that has updated its path is collected
			// so that all segments can later be added to includedElements
			if _, ok := ap.includedElements[int(equipmentSetEquipmentRef.ReferencedElementID)]; ok {
				ap.updatedReferencePaths[ComplexID(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
			}
		}
		// we also loop over all references in state because a reference which may not have updated
		// itself may still reference an element which has updated
		for _, equipmentSetEquipmentRef := range state.EquipmentSetEquipmentRef {
			if _, ok := ap.updatedReferencePaths[ComplexID(equipmentSetEquipmentRef.ID)]; ok {
				// we don't need to do the check if the reference is already included
				continue
			}
			if _, ok := ap.includedElements[int(equipmentSetEquipmentRef.ReferencedElementID)]; ok {
				ap.updatedReferencePaths[ComplexID(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
			}
		}

		for _, itemBoundToRef := range patch.ItemBoundToRef {
			if _, ok := ap.includedElements[int(itemBoundToRef.ReferencedElementID)]; ok {
				ap.updatedReferencePaths[ComplexID(itemBoundToRef.ID)] = itemBoundToRef.path
			}
		}
		for _, itemBoundToRef := range state.ItemBoundToRef {
			if _, ok := ap.updatedReferencePaths[ComplexID(itemBoundToRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(itemBoundToRef.ReferencedElementID)]; ok {
				ap.updatedReferencePaths[ComplexID(itemBoundToRef.ID)] = itemBoundToRef.path
			}
		}

		for _, playerEquipmentSetRef := range patch.PlayerEquipmentSetRef {
			if _, ok := ap.includedElements[int(playerEquipmentSetRef.ReferencedElementID)]; ok {
				ap.updatedReferencePaths[ComplexID(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
			}
		}
		for _, playerEquipmentSetRef := range state.PlayerEquipmentSetRef {
			if _, ok := ap.updatedReferencePaths[ComplexID(playerEquipmentSetRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(playerEquipmentSetRef.ReferencedElementID)]; ok {
				ap.updatedReferencePaths[ComplexID(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
			}
		}

		for _, playerGuildMemberRef := range patch.PlayerGuildMemberRef {
			if _, ok := ap.includedElements[int(playerGuildMemberRef.ReferencedElementID)]; ok {
				ap.updatedReferencePaths[ComplexID(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
			}
		}
		for _, playerGuildMemberRef := range state.PlayerGuildMemberRef {
			if _, ok := ap.updatedReferencePaths[ComplexID(playerGuildMemberRef.ID)]; ok {
				continue
			}
			if _, ok := ap.includedElements[int(playerGuildMemberRef.ReferencedElementID)]; ok {
				ap.updatedReferencePaths[ComplexID(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
			}
		}

		for _, playerTargetRef := range patch.PlayerTargetRef {
			// if the reference exists in the patch, the anyContainer HAS to exist in patch as well
			// as both are always created and destroyed in unison
			anyContainer := patch.AnyOfPlayer_ZoneItem[playerTargetRef.ReferencedElementID]
			if _, ok := ap.includedElements[int(anyContainer.ChildID)]; ok {
				ap.updatedReferencePaths[ComplexID(playerTargetRef.ID)] = playerTargetRef.path
			}
		}
		for _, playerTargetRef := range state.PlayerTargetRef {
			if _, ok := ap.updatedReferencePaths[ComplexID(playerTargetRef.ID)]; ok {
				continue
			}
			// if the reference exists in the state, the anyContainer HAS to exist in state as well
			// as both are always created and destroyed in unison
			anyContainer := state.AnyOfPlayer_ZoneItem[playerTargetRef.ReferencedElementID]
			if _, ok := ap.includedElements[int(anyContainer.ChildID)]; ok {
				ap.updatedReferencePaths[ComplexID(playerTargetRef.ID)] = playerTargetRef.path
			}
		}

		for _, playerTargetedByRef := range patch.PlayerTargetedByRef {
			anyContainer := patch.AnyOfPlayer_ZoneItem[playerTargetedByRef.ReferencedElementID]
			if _, ok := ap.includedElements[int(anyContainer.ChildID)]; ok {
				ap.updatedReferencePaths[ComplexID(playerTargetedByRef.ID)] = playerTargetedByRef.path
			}
		}
		for _, playerTargetedByRef := range state.PlayerTargetedByRef {
			if _, ok := ap.updatedReferencePaths[ComplexID(playerTargetedByRef.ID)]; ok {
				continue
			}
			anyContainer := state.AnyOfPlayer_ZoneItem[playerTargetedByRef.ReferencedElementID]
			if _, ok := ap.includedElements[int(anyContainer.ChildID)]; ok {
				ap.updatedReferencePaths[ComplexID(playerTargetedByRef.ID)] = playerTargetedByRef.path
			}
		}
	}

	// merge paths into one slice, for convencience (they are recycled anyway)
	for _, p := range ap.updatedElementPaths {
		ap.updatedPaths = append(ap.updatedPaths, p)
	}
	for _, p := range ap.updatedReferencePaths {
		ap.updatedPaths = append(ap.updatedPaths, p)
	}
}

func (ap *assemblePlanner) fill(state *State) {
	for _, boolValue := range state.BoolValue {
		ap.updatedElementPaths[int(boolValue.ID)] = boolValue.path
	}
	for _, floatValue := range state.FloatValue {
		ap.updatedElementPaths[int(floatValue.ID)] = floatValue.path
	}
	for _, intValue := range state.IntValue {
		ap.updatedElementPaths[int(intValue.ID)] = intValue.path
	}
	for _, stringValue := range state.StringValue {
		ap.updatedElementPaths[int(stringValue.ID)] = stringValue.path
	}
	for _, attackEvent := range state.AttackEvent {
		ap.updatedElementPaths[int(attackEvent.ID)] = attackEvent.path
	}
	for _, equipmentSet := range state.EquipmentSet {
		ap.updatedElementPaths[int(equipmentSet.ID)] = equipmentSet.path
	}
	for _, gearScore := range state.GearScore {
		ap.updatedElementPaths[int(gearScore.ID)] = gearScore.path
	}
	for _, item := range state.Item {
		ap.updatedElementPaths[int(item.ID)] = item.path
	}
	for _, player := range state.Player {
		ap.updatedElementPaths[int(player.ID)] = player.path
	}
	for _, position := range state.Position {
		ap.updatedElementPaths[int(position.ID)] = position.path
	}
	for _, zone := range state.Zone {
		ap.updatedElementPaths[int(zone.ID)] = zone.path
	}
	for _, zoneItem := range state.ZoneItem {
		ap.updatedElementPaths[int(zoneItem.ID)] = zoneItem.path
	}
	for _, attackEventTargetRef := range state.AttackEventTargetRef {
		ap.updatedReferencePaths[ComplexID(attackEventTargetRef.ID)] = attackEventTargetRef.path
	}
	for _, equipmentSetEquipmentRef := range state.EquipmentSetEquipmentRef {
		ap.updatedReferencePaths[ComplexID(equipmentSetEquipmentRef.ID)] = equipmentSetEquipmentRef.path
	}
	for _, itemBoundToRef := range state.ItemBoundToRef {
		ap.updatedReferencePaths[ComplexID(itemBoundToRef.ID)] = itemBoundToRef.path
	}
	for _, playerEquipmentSetRef := range state.PlayerEquipmentSetRef {
		ap.updatedReferencePaths[ComplexID(playerEquipmentSetRef.ID)] = playerEquipmentSetRef.path
	}
	for _, playerGuildMemberRef := range state.PlayerGuildMemberRef {
		ap.updatedReferencePaths[ComplexID(playerGuildMemberRef.ID)] = playerGuildMemberRef.path
	}
	for _, playerTargetRef := range state.PlayerTargetRef {
		ap.updatedReferencePaths[ComplexID(playerTargetRef.ID)] = playerTargetRef.path
	}
	for _, playerTargetedByRef := range state.PlayerTargetedByRef {
		ap.updatedReferencePaths[ComplexID(playerTargetedByRef.ID)] = playerTargetedByRef.path
	}

	// merge paths into one slice, for convencience (they are recycled anyway)
	for _, p := range ap.updatedElementPaths {
		ap.updatedPaths = append(ap.updatedPaths, p)
	}
	for _, p := range ap.updatedReferencePaths {
		ap.updatedPaths = append(ap.updatedPaths, p)
	}
}
