package bg

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

type stoItemTypes uint32

const (
	BooksMisc stoItemTypes = iota
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

var stoItemTypeName = map[stoItemTypes]string{
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

func (it stoItemTypes) String() string {
	return stoItemTypeName[it]
}

type stoHeader struct {
	Signature, Version           [4]byte
	Type                         uint32
	Name                         uint32
	StoreFlags                   [4]uint8
	SellPriceMarkup              uint32
	BuyPriceMarkup               uint32
	DepreciationRate             uint32
	PercentageChanceStealFailure uint16
	Capacity                     uint16
	Unknown                      [8]uint8
	OffsetToItemsPurchased       uint32
	CountOfItemsPurchased        uint32
	OffsetToItemsForSale         uint32
	CountOfItemsForSale          uint32
	Lore                         uint32
	IdPrice                      uint32
	RumoursTavern                Resref
	OffsetToDrinks               uint32
	CountOfDrinks                uint32
	RumoursTemple                Resref
	RoomFlags                    [4]uint8
	PriceOfAPeasantRoom          uint32
	PriceOfAMerchantRoom         uint32
	PriceOfANobleRoom            uint32
	PriceOfARoyalRoom            uint32
	OffsetToCures                uint32
	CountOfCures                 uint32
	Unused                       [36]uint8
}

type stoItems struct {
	FileName           Resref
	ItemExpirationTime uint16
	Charges1           uint16
	Charges2           uint16
	Charges3           uint16
	Flags              uint32
	Amount             uint32
	InfiniteSupplyFlag uint32
}

type stoDrinks struct {
	RumourResource    Resref
	Name              uint32
	Price             uint32
	AlcoholicStrength uint32
}

type stoCures struct {
	FileNameOfSpell Resref
	SpellPrice      uint32
}

type STO struct {
	Header             stoHeader
	Items              []stoItems
	Drinks             []stoDrinks
	Cures              []stoCures
	ItemsPurchasedHere []stoItemTypes
	Filename           string
}

func (sto *STO) Write(w io.Writer) error {
	err := binary.Write(w, binary.LittleEndian, sto.Header)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, sto.Items)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, sto.Drinks)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, sto.Cures)
	if err != nil {
		return err
	}
	return nil
}

func OpenSTO(r io.ReadSeeker) (*STO, error) {
	sto := STO{}

	err := binary.Read(r, binary.LittleEndian, &sto.Header)
	if err != nil {
		return nil, err
	}

	sto.Items, err = parseArray[stoItems](r, sto.Header.CountOfItemsForSale, sto.Header.OffsetToItemsForSale)
	if err != nil {
		return nil, err
	}

	sto.Drinks, err = parseArray[stoDrinks](r, sto.Header.CountOfDrinks, sto.Header.OffsetToDrinks)
	if err != nil {
		return nil, err
	}

	sto.Cures, err = parseArray[stoCures](r, sto.Header.CountOfCures, sto.Header.OffsetToCures)
	if err != nil {
		return nil, err
	}

	sto.ItemsPurchasedHere, err = parseArray[stoItemTypes](r, sto.Header.CountOfItemsPurchased, sto.Header.OffsetToItemsPurchased)
	if err != nil {
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
