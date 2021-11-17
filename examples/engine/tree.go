package state

type ReferencedDataStatus string

const (
	ReferencedDataModified  ReferencedDataStatus = "MODIFIED"
	ReferencedDataUnchanged ReferencedDataStatus = "UNCHANGED"
)

type ElementKind string

const (
	ElementKindEquipmentSet ElementKind = "EquipmentSet"
	ElementKindGearScore    ElementKind = "GearScore"
	ElementKindItem         ElementKind = "Item"
	ElementKindPlayer       ElementKind = "Player"
	ElementKindPosition     ElementKind = "Position"
	ElementKindZone         ElementKind = "Zone"
	ElementKindZoneItem     ElementKind = "ZoneItem"
)

type Tree struct {
	EquipmentSet map[EquipmentSetID]equipmentSet `json:"equipmentSet"`
	GearScore    map[GearScoreID]gearScore       `json:"gearScore"`
	Item         map[ItemID]item                 `json:"item"`
	Player       map[PlayerID]player             `json:"player"`
	Position     map[PositionID]position         `json:"position"`
	Zone         map[ZoneID]zone                 `json:"zone"`
	ZoneItem     map[ZoneItemID]zoneItem         `json:"zoneItem"`
}

func newTree() Tree {
	return Tree{
		EquipmentSet: make(map[EquipmentSetID]equipmentSet),
		GearScore:    make(map[GearScoreID]gearScore),
		Item:         make(map[ItemID]item),
		Player:       make(map[PlayerID]player),
		Position:     make(map[PositionID]position),
		Zone:         make(map[ZoneID]zone),
		ZoneItem:     make(map[ZoneItemID]zoneItem),
	}
}

type zoneItem struct {
	ID            ZoneItemID    `json:"id"`
	Item          *item         `json:"item"`
	Position      *position     `json:"position"`
	OperationKind OperationKind `json:"operationKind"`
}
type zoneItemReference struct {
	OperationKind        OperationKind        `json:"operationKind"`
	ElementID            ZoneItemID           `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
}

type item struct {
	ID            ItemID           `json:"id"`
	BoundTo       *playerReference `json:"boundTo"`
	GearScore     *gearScore       `json:"gearScore"`
	Name          string           `json:"name"`
	Origin        interface{}      `json:"origin"`
	OperationKind OperationKind    `json:"operationKind"`
}
type itemReference struct {
	OperationKind        OperationKind        `json:"operationKind"`
	ElementID            ItemID               `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
}

type equipmentSet struct {
	ID            EquipmentSetID           `json:"id"`
	Equipment     map[ItemID]itemReference `json:"equipment"`
	Name          string                   `json:"name"`
	OperationKind OperationKind            `json:"operationKind"`
}
type equipmentSetReference struct {
	OperationKind        OperationKind        `json:"operationKind"`
	ElementID            EquipmentSetID       `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
}

type position struct {
	ID            PositionID    `json:"id"`
	X             float64       `json:"x"`
	Y             float64       `json:"y"`
	OperationKind OperationKind `json:"operationKind"`
}
type positionReference struct {
	OperationKind        OperationKind        `json:"operationKind"`
	ElementID            PositionID           `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
}

type gearScore struct {
	ID            GearScoreID   `json:"id"`
	Level         int           `json:"level"`
	Score         int           `json:"score"`
	OperationKind OperationKind `json:"operationKind"`
}
type gearScoreReference struct {
	OperationKind        OperationKind        `json:"operationKind"`
	ElementID            GearScoreID          `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
}

type player struct {
	ID            PlayerID                                 `json:"id"`
	EquipmentSets map[EquipmentSetID]equipmentSetReference `json:"equipmentSets"`
	GearScore     *gearScore                               `json:"gearScore"`
	GuildMembers  map[PlayerID]playerReference             `json:"guildMembers"`
	Items         map[ItemID]item                          `json:"items"`
	Position      *position                                `json:"position"`
	Target        *anyOfPlayer_ZoneItemReference           `json:"target"`
	TargetedBy    map[int]anyOfPlayer_ZoneItemReference    `json:"targetedBy"`
	OperationKind OperationKind                            `json:"operationKind"`
}
type playerReference struct {
	OperationKind        OperationKind        `json:"operationKind"`
	ElementID            PlayerID             `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
}

type zone struct {
	ID            ZoneID                  `json:"id"`
	Interactables map[int]interface{}     `json:"interactables"`
	Items         map[ZoneItemID]zoneItem `json:"items"`
	Players       map[PlayerID]player     `json:"players"`
	Tags          []string                `json:"tags"`
	OperationKind OperationKind           `json:"operationKind"`
}
type zoneReference struct {
	OperationKind        OperationKind        `json:"operationKind"`
	ElementID            ZoneID               `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
}

type anyOfPlayer_ZoneItemReference struct {
	OperationKind        OperationKind        `json:"operationKind"`
	ElementID            int                  `json:"id"`
	ElementKind          ElementKind          `json:"elementKind"`
	ReferencedDataStatus ReferencedDataStatus `json:"referencedDataStatus"`
	ElementPath          string               `json:"elementPath"`
}

type recursionCheck struct {
	equipmentSet map[EquipmentSetID]bool
	gearScore    map[GearScoreID]bool
	item         map[ItemID]bool
	player       map[PlayerID]bool
	position     map[PositionID]bool
	zone         map[ZoneID]bool
	zoneItem     map[ZoneItemID]bool
}

func newRecursionCheck() *recursionCheck {
	return &recursionCheck{
		equipmentSet: make(map[EquipmentSetID]bool),
		gearScore:    make(map[GearScoreID]bool),
		item:         make(map[ItemID]bool),
		player:       make(map[PlayerID]bool),
		position:     make(map[PositionID]bool),
		zone:         make(map[ZoneID]bool),
		zoneItem:     make(map[ZoneItemID]bool),
	}
}

type assembleCache struct {
	equipmentSet map[EquipmentSetID]equipmentSetCacheElement
	gearScore    map[GearScoreID]gearScoreCacheElement
	item         map[ItemID]itemCacheElement
	player       map[PlayerID]playerCacheElement
	position     map[PositionID]positionCacheElement
	zone         map[ZoneID]zoneCacheElement
	zoneItem     map[ZoneItemID]zoneItemCacheElement
}

func newAssembleCache() assembleCache {
	return assembleCache{
		equipmentSet: make(map[EquipmentSetID]equipmentSetCacheElement),
		gearScore:    make(map[GearScoreID]gearScoreCacheElement),
		item:         make(map[ItemID]itemCacheElement),
		player:       make(map[PlayerID]playerCacheElement),
		position:     make(map[PositionID]positionCacheElement),
		zone:         make(map[ZoneID]zoneCacheElement),
		zoneItem:     make(map[ZoneItemID]zoneItemCacheElement),
	}
}

type equipmentSetCacheElement struct {
	hasUpdated   bool
	equipmentSet equipmentSet
}
type gearScoreCacheElement struct {
	hasUpdated bool
	gearScore  gearScore
}
type itemCacheElement struct {
	hasUpdated bool
	item       item
}
type playerCacheElement struct {
	hasUpdated bool
	player     player
}
type positionCacheElement struct {
	hasUpdated bool
	position   position
}
type zoneCacheElement struct {
	hasUpdated bool
	zone       zone
}
type zoneItemCacheElement struct {
	hasUpdated bool
	zoneItem   zoneItem
}
