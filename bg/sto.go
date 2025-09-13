package bg

import (
	"encoding/binary"
	"encoding/json"
	"io"
	"reflect"
	"slices"
)

type StoItemTypes uint32

const (
	BooksMisc StoItemTypes = iota
	AmuletsAndNecklaces
	Armor
	BeltsAndGirdles
	Boots
	Arrows
	BracersAndGauntlets
	Headgear
	Keys
	Potions
	Rings
	Scrolls
	Shields
	Food
	Bullets
	Bows
	Daggers
	Maces
	Slings
	SmallSwords
	LargeSwords
	Hammers
	MorningStars
	Flails
	Darts
	Axes
	Quarterstaff
	Crossbow
	HandToHandWeapons
	Spears
	Halberds
	CrossbowBolts
	CloaksAndRobes
	GoldPieces
	Gems
	Wands
	ContainersBrokenArmor
	BooksBrokenShieldsBracelets
	FamiliarsBrokenSwordsEarrings
	TattoosPST
	LensesPST
	BucklersTeeth
	Candles
	Unknown1
	Clubs
	Unknown2
	Unknown3
	LargeShieldsIWD
	Unknown4
	MediumShieldsIWD
	Notes
	Unknown5
	Unknown6
	SmallShields
	Unknown7
	TelescopesIWD
	DrinksIWD
	GreatSwordsIWD
	Container
	FurPelt
	LeatherArmor
	StuddedLeatherArmor
	ChainMail
	SplintMail
	HalfPlate
	FullPlate
	HideArmor
	Robe
	Unknown8
	BastardSword
	Scarf
	FoodIWD2
	Hat
	Gauntlet
)

var stoItemTypeName = map[StoItemTypes]string{
	BooksMisc:                     "BooksMisc",
	AmuletsAndNecklaces:           "AmuletsAndNecklaces",
	Armor:                         "Armor",
	BeltsAndGirdles:               "BeltsAndGirdles",
	Boots:                         "Boots",
	Arrows:                        "Arrows",
	BracersAndGauntlets:           "BracersAndGauntlets",
	Headgear:                      "Headgear",
	Keys:                          "Keys",
	Potions:                       "Potions",
	Rings:                         "Rings",
	Scrolls:                       "Scrolls",
	Shields:                       "Shields",
	Food:                          "Food",
	Bullets:                       "Bullets",
	Bows:                          "Bows",
	Daggers:                       "Daggers",
	Maces:                         "Maces",
	Slings:                        "Slings",
	SmallSwords:                   "SmallSwords",
	LargeSwords:                   "LargeSwords",
	Hammers:                       "Hammers",
	MorningStars:                  "MorningStars",
	Flails:                        "Flails",
	Darts:                         "Darts",
	Axes:                          "Axes",
	Quarterstaff:                  "Quarterstaff",
	Crossbow:                      "Crossbow",
	HandToHandWeapons:             "HandToHandWeapons",
	Spears:                        "Spears",
	Halberds:                      "Halberds",
	CrossbowBolts:                 "CrossbowBolts",
	CloaksAndRobes:                "CloaksAndRobes",
	GoldPieces:                    "GoldPieces",
	Gems:                          "Gems",
	Wands:                         "Wands",
	ContainersBrokenArmor:         "ContainersBrokenArmor",
	BooksBrokenShieldsBracelets:   "BooksBrokenShieldsBracelets",
	FamiliarsBrokenSwordsEarrings: "FamiliarsBrokenSwordsEarrings",
	TattoosPST:                    "TattoosPST",
	LensesPST:                     "LensesPST",
	BucklersTeeth:                 "BucklersTeeth",
	Candles:                       "Candles",
	Unknown1:                      "Unknown1",
	Clubs:                         "Clubs",
	Unknown2:                      "Unknown2",
	Unknown3:                      "Unknown3",
	LargeShieldsIWD:               "LargeShieldsIWD",
	Unknown4:                      "Unknown4",
	MediumShieldsIWD:              "MediumShieldsIWD",
	Notes:                         "Notes",
	Unknown5:                      "Unknown5",
	Unknown6:                      "Unknown6",
	SmallShields:                  "SmallShields",
	Unknown7:                      "Unknown7",
	TelescopesIWD:                 "TelescopesIWD",
	DrinksIWD:                     "DrinksIWD",
	GreatSwordsIWD:                "GreatSwordsIWD",
	Container:                     "Container",
	FurPelt:                       "FurPelt",
	LeatherArmor:                  "LeatherArmor",
	StuddedLeatherArmor:           "StuddedLeatherArmor",
	ChainMail:                     "ChainMail",
	SplintMail:                    "SplintMail",
	HalfPlate:                     "HalfPlate",
	FullPlate:                     "FullPlate",
	HideArmor:                     "HideArmor",
	Robe:                          "Robe",
	Unknown8:                      "Unknown8",
	BastardSword:                  "BastardSword",
	Scarf:                         "Scarf",
	FoodIWD2:                      "FoodIWD2",
	Hat:                           "Hat",
	Gauntlet:                      "Gauntlet",
}

func (it StoItemTypes) String() string {
	return stoItemTypeName[it]
}

type StoHeader struct {
	Signature                    Signature `json:"signature"`
	Version                      Version   `json:"version"`
	StoreType                    uint32    `json:"store_type"`
	Name                         uint32    `json:"name"`
	StoreFlags                   [4]uint8  `json:"store_flags"`
	SellPriceMarkup              uint32    `json:"sell_price_markup"`
	BuyPriceMarkup               uint32    `json:"buy_price_markup"`
	DepreciationRate             uint32    `json:"depreciation_rate"`
	PercentageChanceStealFailure uint16    `json:"percentage_chance_steal_failure"`
	Capacity                     uint16    `json:"capacity"`
	Unknown                      [8]uint8  `json:"unknown"`
	OffsetToItemsPurchased       uint32    `json:"offset_to_items_purchased"`
	CountOfItemsPurchased        uint32    `json:"count_of_items_purchased"`
	OffsetToItemsForSale         uint32    `json:"offset_to_items_for_sale"`
	CountOfItemsForSale          uint32    `json:"count_of_items_for_sale"`
	Lore                         uint32    `json:"lore"`
	IdPrice                      uint32    `json:"id_price"`
	RumoursTavern                Resref    `json:"rumours_tavern"`
	OffsetToDrinks               uint32    `json:"offset_to_drinks"`
	CountOfDrinks                uint32    `json:"count_of_drinks"`
	RumoursTemple                Resref    `json:"rumours_temple"`
	RoomFlags                    [4]uint8  `json:"room_flags"`
	PriceOfAPeasantRoom          uint32    `json:"price_of_a_peasant_room"`
	PriceOfAMerchantRoom         uint32    `json:"price_of_a_merchant_room"`
	PriceOfANobleRoom            uint32    `json:"price_of_a_noble_room"`
	PriceOfARoyalRoom            uint32    `json:"price_of_a_royal_room"`
	OffsetToCures                uint32    `json:"offset_to_cures"`
	CountOfCures                 uint32    `json:"count_of_cures"`
	Unused                       [36]uint8 `json:"unused"`
}

type StoItems struct {
	FileNameOfItem     Resref `json:"filename_of_item"`
	ItemExpirationTime uint16 `json:"item_expiration_time"`
	Charges1           uint16 `json:"charges1"`
	Charges2           uint16 `json:"charges2"`
	Charges3           uint16 `json:"charges3"`
	Flags              uint32 `json:"flags"`
	Amount             uint32 `json:"amount"`
	InfiniteSupplyFlag uint32 `json:"infinite_supply_flag"`
}

type StoDrinks struct {
	RumourResource    Resref `json:"rumour_resource"`
	Name              uint32 `json:"name"`
	Price             uint32 `json:"price"`
	AlcoholicStrength uint32 `json:"alcoholic_strength"`
}

type StoCures struct {
	FileNameOfSpell Resref `json:"file_name_of_spell"`
	SpellPrice      uint32 `json:"spell_price"`
}

type STO struct {
	StoHeader
	Items              []StoItems     `json:"items_for_sale"`
	Drinks             []StoDrinks    `json:"drinks_for_sale"`
	Cures              []StoCures     `json:"cures_for_sale"`
	ItemsPurchasedHere []StoItemTypes `json:"items_purchased_here"`
	Filename           string         `json:"-"`
}

func (sto *STO) Equal(other *STO) bool {
	if !reflect.DeepEqual(sto.StoHeader, other.StoHeader) {
		return false
	}
	if !reflect.DeepEqual(sto.Cures, other.Cures) {
		return false
	}
	if !reflect.DeepEqual(sto.Items, other.Items) {
		return false
	}
	if !reflect.DeepEqual(sto.Drinks, other.Drinks) {
		return false
	}
	if !slices.Equal(sto.ItemsPurchasedHere, other.ItemsPurchasedHere) {
		return false
	}
	return true
}

func (sto *STO) Write(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, sto.StoHeader); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, sto.Items); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, sto.Drinks); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, sto.Cures); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, sto.ItemsPurchasedHere); err != nil {
		return err
	}
	return nil
}

func OpenSTO(r io.ReadSeeker) (*STO, error) {
	sto := STO{}

	if err := binary.Read(r, binary.LittleEndian, &sto.StoHeader); err != nil {
		return nil, err
	}
	sto.Items = make([]StoItems, sto.CountOfItemsForSale)
	if err := parseArray(r, sto.OffsetToItemsForSale, sto.Items); err != nil {
		return nil, err
	}
	sto.Drinks = make([]StoDrinks, sto.CountOfDrinks)
	if err := parseArray(r, sto.OffsetToDrinks, sto.Drinks); err != nil {
		return nil, err
	}
	sto.Cures = make([]StoCures, sto.CountOfCures)
	if err := parseArray(r, sto.OffsetToCures, sto.Cures); err != nil {
		return nil, err
	}
	sto.ItemsPurchasedHere = make([]StoItemTypes, sto.CountOfItemsPurchased)
	if err := parseArray(r, sto.OffsetToItemsPurchased, sto.ItemsPurchasedHere); err != nil {
		return nil, err
	}
	return &sto, nil
}

func (sto *STO) WriteJson(w io.Writer) error {
	bytes, err := json.MarshalIndent(sto, "", "\t")
	if err != nil {
		return err
	}

	_, err = w.Write(bytes)
	return err
}
