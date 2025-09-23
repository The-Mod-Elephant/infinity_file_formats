package bg

import (
	"encoding/binary"
	"encoding/json"
	"io"
	"maps"
	"reflect"
	"slices"
)

type CreHeader struct {
	Signature                    Signature   `json:"signature"`
	Version                      Version     `json:"version"`
	LongName                     strref      `json:"long_name"`
	ShortName                    strref      `json:"short_name"`
	Flags                        uint32      `json:"flags"`
	XPValue                      uint32      `json:"xp_value"`
	XP                           uint32      `json:"xp"`
	Gold                         uint32      `json:"gold"`
	PermanentStatusFlags         uint32      `json:"permanent_status_flags"`
	HitPoints                    uint16      `json:"hit_points"`
	MaxHitPointsBase             uint16      `json:"max_hit_points_base"`
	AnimationType                uint32      `json:"animation_type"`
	MetalColor                   uint8       `json:"metal_color"`
	MinorColor                   uint8       `json:"minor_color"`
	MajorColor                   uint8       `json:"major_color"`
	SkinColor                    uint8       `json:"skin_color"`
	LeatherColor                 uint8       `json:"leather_color"`
	ArmorColor                   uint8       `json:"armor_color"`
	HairColor                    uint8       `json:"hair_color"`
	EffStructureVersion          uint8       `json:"eff_structure_version"`
	SmallPortrait                Resref      `json:"small_portrait"`
	LargePortrait                Resref      `json:"large_portrait"`
	Reputation                   int8        `json:"reputation"`
	HideInShadowsBase            uint8       `json:"hide_in_shadows_base"`
	ArmorClass                   int16       `json:"armor_class"`
	ArmorClassBase               int16       `json:"armor_class_base"`
	ArmorClassCurshingAdjustment int16       `json:"armor_class_curshing_adjustment"`
	ArmorClassMissileAdjustment  int16       `json:"armor_class_missile_adjustment"`
	ArmorClassPiercingAdjustment int16       `json:"armor_class_piercing_adjustment"`
	ArmorClassSlashingAdjustment int16       `json:"armor_class_slashing_adjustment"`
	Thac0                        int8        `json:"thac0"`
	NumberOfAttacksBase          uint8       `json:"number_of_attacks_base"`
	SaveVsDeathBase              uint8       `json:"save_vs_death_base"`
	SaveVsWandsBase              uint8       `json:"save_vs_wands_base"`
	SaveVsPolyBase               uint8       `json:"save_vs_poly_base"`
	SaveVsBreathBase             uint8       `json:"save_vs_breath_base"`
	SaveVsSpellBase              uint8       `json:"save_vs_spell_base"`
	ResistFireBase               int8        `json:"resist_fire_base"`
	ResistColdBase               int8        `json:"resist_cold_base"`
	ResistElectricityBase        int8        `json:"resist_electricity_base"`
	ResistAcidBase               int8        `json:"resist_acid_base"`
	ResistMagicBase              int8        `json:"resist_magic_base"`
	ResistMagicFireBase          int8        `json:"resist_magic_fire_base"`
	ResistMagicColdBase          int8        `json:"resist_magic_cold_base"`
	ResistSlashingBase           int8        `json:"resist_slashing_base"`
	ResistCrushingBase           int8        `json:"resist_crushing_base"`
	ResistPiercingBase           int8        `json:"resist_piercing_base"`
	ResistMissileBase            int8        `json:"resist_missile_base"`
	DetectIllusionBase           uint8       `json:"detect_illusion_base"`
	SetTrapsBase                 uint8       `json:"set_traps_base"`
	LoreBase                     uint8       `json:"lore_base"`
	LockPickingBase              uint8       `json:"lock_picking_base"`
	MoveSilentlyBase             uint8       `json:"move_silently_base"`
	FindTrapsBase                uint8       `json:"find_traps_base"`
	PickPocketBase               uint8       `json:"pick_pocket_base"`
	Fatigue                      uint8       `json:"fatigue"`
	Intoxication                 uint8       `json:"intoxication"`
	LuckBase                     int8        `json:"luck_base"`
	ProficiencyLargeSwords       uint8       `json:"proficiency_large_swords"`
	ProficiencySmallSwords       uint8       `json:"proficiency_small_swords"`
	ProficiencyBows              uint8       `json:"proficiency_bows"`
	ProficiencySpears            uint8       `json:"proficiency_spears"`
	ProficiencyBlunt             uint8       `json:"proficiency_blunt"`
	ProficiencySpiked            uint8       `json:"proficiency_spiked"`
	ProficiencyAxes              uint8       `json:"proficiency_axes"`
	ProficiencyMissiles          uint8       `json:"proficiency_missiles"`
	UnusedProficiencies          [7]uint8    `json:"unused_proficiencies"`
	NightmareMode                uint8       `json:"nightmare_mode"`
	Translucency                 uint8       `json:"translucency"`
	ReputationLossIfKilled       uint8       `json:"reputation_loss_if_killed"`
	ReputationLossIfJoinsParty   uint8       `json:"reputation_loss_if_joins_party"`
	ReputationLossIfLeavesParty  uint8       `json:"reputation_loss_if_leaves_party"`
	UndeadLevel                  uint8       `json:"undead_level"`
	TrackingBase                 uint8       `json:"tracking_base"`
	TrackingTarget               LongString  `json:"tracking_target"`
	Strrefs                      [100]strref `json:"strrefs"`
	LevelFirstClass              uint8       `json:"level_first_class"`
	LevelSecondClass             uint8       `json:"level_second_class"`
	LevelThirdClass              uint8       `json:"level_third_class"`
	Sex                          uint8       `json:"sex"`
	Strength                     uint8       `json:"strength"`
	StrengthBonus                uint8       `json:"strength_bonus"`
	Intelligence                 uint8       `json:"intelligence"`
	Wisdom                       uint8       `json:"wisdom"`
	Dexterity                    uint8       `json:"dexterity"`
	Constitution                 uint8       `json:"constitution"`
	Charisma                     uint8       `json:"charisma"`
	Morale                       uint8       `json:"morale"`
	MoraleBreak                  uint8       `json:"morale_break"`
	RacialEnemy                  uint8       `json:"racial_enemy"`
	MoraleRecoveryTime           uint16      `json:"morale_recovery_time"`
	Kit                          uint32      `json:"kit"`
	OverrideScript               Resref      `json:"override_script"`
	ClassScript                  Resref      `json:"class_script"`
	RaceScript                   Resref      `json:"race_script"`
	GeneralScript                Resref      `json:"general_script"`
	DefaultScript                Resref      `json:"default_script"`
	EnemyAlly                    uint8       `json:"enemy_ally"`
	General                      uint8       `json:"general"`
	Race                         uint8       `json:"race"`
	Class                        uint8       `json:"class"`
	Specific                     uint8       `json:"specific"`
	Gender                       uint8       `json:"gender"`
	ObjectReferences             [5]uint8    `json:"object_references"`
	Alignment                    uint8       `json:"alignment"`
	GlobalActorEnumeration       uint16      `json:"global_actor_enumeration"`
	LocalActorEnumeration        uint16      `json:"local_actor_enumeration"`
	DeathVariable                LongString  `json:"death_variable"`
	KnownSpellListOffset         uint32      `json:"known_spell_list_offset"`
	KnownSpellListCount          uint32      `json:"known_spell_list_count"`
	MemorizationLevelListOffset  uint32      `json:"memorization_level_list_offset"`
	MemorizationLevelListCount   uint32      `json:"memorization_level_list_count"`
	MemorizationSpellListOffset  uint32      `json:"memorization_spell_list_offset"`
	MemorizationSpellListCount   uint32      `json:"memorization_spell_list_count"`
	EquipmentListOffset          uint32      `json:"equipment_list_offset"`
	ItemListOffset               uint32      `json:"item_list_offset"`
	ItemListCount                uint32      `json:"item_list_count"`
	EffectListOffset             uint32      `json:"effect_list_offset"`
	EffectListCount              uint32      `json:"effect_list_count"`
	Dialog                       Resref      `json:"dialog"`
}

type CreKnownSpell struct {
	SpellName  Resref `json:"spell_name"`
	SpellLevel uint16 `json:"spell_level"`
	SpellType  uint16 `json:"spell_type"`
}

type CreMemorizedSpellLevel struct {
	SpellLevel             uint16 `json:"spell_level"`
	BaseCount              uint16 `json:"base_count"`
	Count                  uint16 `json:"count"`
	MagicType              uint16 `json:"magic_type"`
	MemorizedStartingSpell uint32 `json:"memorized_starting_spell"`
	MemorizedCount         uint32 `json:"memorized_count"`
}

type CreMemorizedSpell struct {
	SpellName Resref `json:"spell_name"`
	Memorised uint32 `json:"memorised"`
}

type DynamicFlagsJson struct {
	Identified  uint8 `json:"identified"`
	Unstealable uint8 `json:"unstealable"`
	Stolen      uint8 `json:"stolen"`
	Undroppable uint8 `json:"undroppable"`
}

type CreItemHeader struct {
	ItemName               Resref `json:"item_name"`
	ItemExpirationTimeHour uint8  `json:"item_expiration_time_hour"`
	ItemExpirationTime     uint8  `json:"item_expiration_time"`
	Quantity1              uint16 `json:"quantity_1"`
	Quantity2              uint16 `json:"quantity_2"`
	Quantity3              uint16 `json:"quantity_3"`
}

type creItemJson struct {
	CreItemHeader
	DynamicFlagsJson
}

type CreItem struct {
	CreItemHeader
	Flags uint32 `json:"flags"`
}

func (d *CreItem) MarshalJSON() ([]byte, error) {
	out := creItemJson{
		CreItemHeader: d.CreItemHeader,
	}
	out.Identified = uint8((d.Flags >> 0) & 1)
	out.Unstealable = uint8((d.Flags >> 1) & 1)
	out.Stolen = uint8((d.Flags >> 2) & 1)
	out.Undroppable = uint8((d.Flags >> 3) & 1)

	return json.Marshal(&out)
}

func (d *CreItem) UnmarshalJSON(b []byte) error {
	in := creItemJson{}
	if err := json.Unmarshal(b, &in); err != nil {
		return err
	}
	d.CreItemHeader = in.CreItemHeader
	d.Flags = uint32((in.Undroppable << 3) | (in.Stolen << 2) | (in.Unstealable << 1) | in.Identified)

	return nil
}

type CreEquipment struct {
	Helmet                uint16 `json:"helmet"`
	Armor                 uint16 `json:"armor"`
	Shield                uint16 `json:"shield"`
	Gloves                uint16 `json:"gloves"`
	LeftRing              uint16 `json:"left_ring"`
	RightRing             uint16 `json:"right_ring"`
	Amulet                uint16 `json:"amulet"`
	Belt                  uint16 `json:"belt"`
	Boots                 uint16 `json:"boots"`
	Weapon1               uint16 `json:"weapon_1"`
	Weapon2               uint16 `json:"weapon_2"`
	Weapon3               uint16 `json:"weapon_3"`
	Weapon4               uint16 `json:"weapon_4"`
	Quiver1               uint16 `json:"quiver_1"`
	Quiver2               uint16 `json:"quiver_2"`
	Quiver3               uint16 `json:"quiver_3"`
	Quiver4               uint16 `json:"quiver_4"`
	Cloak                 uint16 `json:"cloak"`
	QuickItem1            uint16 `json:"quick_item_1"`
	QuickItem2            uint16 `json:"quick_item_2"`
	QuickItem3            uint16 `json:"quick_item_3"`
	InventoryItem1        uint16 `json:"inventory_item_1"`
	InventoryItem2        uint16 `json:"inventory_item_2"`
	InventoryItem3        uint16 `json:"inventory_item_3"`
	InventoryItem4        uint16 `json:"inventory_item_4"`
	InventoryItem5        uint16 `json:"inventory_item_5"`
	InventoryItem6        uint16 `json:"inventory_item_6"`
	InventoryItem7        uint16 `json:"inventory_item_7"`
	InventoryItem8        uint16 `json:"inventory_item_8"`
	InventoryItem9        uint16 `json:"inventory_item_9"`
	InventoryItem10       uint16 `json:"inventory_item_10"`
	InventoryItem11       uint16 `json:"inventory_item_11"`
	InventoryItem12       uint16 `json:"inventory_item_12"`
	InventoryItem13       uint16 `json:"inventory_item_13"`
	InventoryItem14       uint16 `json:"inventory_item_14"`
	InventoryItem15       uint16 `json:"inventory_item_15"`
	InventoryItem16       uint16 `json:"inventory_item_16"`
	MagicWeapon           uint16 `json:"magic_weapon"`
	SelectedWeapon        uint16 `json:"selected_weapon"`
	SelectedWeaponAbility uint16 `json:"selected_weapon_ability"`
}

type CRE struct {
	CreHeader
	KnownSpells          []CreKnownSpell          `json:"known_spells"`
	MemorizedSpellLevels []CreMemorizedSpellLevel `json:"memorized_spell_levels"`
	MemorizedSpells      []CreMemorizedSpell      `json:"memorized_spells"`
	Effects              []ItmEffect              `json:"effects"`
	Effectsv2            []EffEffect              `json:"effectsv2"`
	Items                []CreItem                `json:"item_table"`
	ItemSlots            CreEquipment             `json:"item_slots"`
	Filename             string                   `json:"-"`
}

func (cre *CRE) Equal(other *CRE) bool {
	if !reflect.DeepEqual(cre.CreHeader, other.CreHeader) {
		return false
	}
	if !slices.Equal(cre.KnownSpells, other.KnownSpells) {
		return false
	}
	if !slices.Equal(cre.MemorizedSpells, other.MemorizedSpells) {
		return false
	}
	if !slices.Equal(cre.Effects, other.Effects) {
		return false
	}
	if !slices.Equal(cre.Effectsv2, other.Effectsv2) {
		return false
	}
	if !slices.Equal(cre.Items, other.Items) {
		return false
	}
	if !reflect.DeepEqual(cre.ItemSlots, other.ItemSlots) {
		return false
	}
	return true
}

func OpenCre(r io.ReadSeeker) (*CRE, error) {
	cre := &CRE{}

	if err := binary.Read(r, binary.LittleEndian, &cre.CreHeader); err != nil {
		return nil, err
	}

	cre.KnownSpells = make([]CreKnownSpell, cre.KnownSpellListCount)
	if _, err := r.Seek(int64(cre.KnownSpellListOffset), io.SeekStart); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &cre.KnownSpells); err != nil {
		return nil, err
	}

	cre.MemorizedSpellLevels = make([]CreMemorizedSpellLevel, cre.MemorizationLevelListCount)
	_, err := r.Seek(int64(cre.MemorizationLevelListOffset), io.SeekStart)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.LittleEndian, &cre.MemorizedSpellLevels)
	if err != nil {
		return nil, err
	}
	cre.MemorizedSpells = make([]CreMemorizedSpell, cre.MemorizationSpellListCount)
	_, err = r.Seek(int64(cre.MemorizationSpellListOffset), io.SeekStart)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.LittleEndian, &cre.MemorizedSpells)
	if err != nil {
		return nil, err
	}
	cre.Items = make([]CreItem, cre.ItemListCount)
	_, err = r.Seek(int64(cre.ItemListOffset), io.SeekStart)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.LittleEndian, &cre.Items)
	if err != nil {
		return nil, err
	}
	_, err = r.Seek(int64(cre.EquipmentListOffset), io.SeekStart)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.LittleEndian, &cre.ItemSlots)
	if err != nil {
		return nil, err
	}
	switch cre.EffStructureVersion {
	case 0:
		cre.Effects = make([]ItmEffect, cre.EffectListCount)
		_, err = r.Seek(int64(cre.EffectListOffset), io.SeekStart)
		if err != nil {
			return nil, err
		}
		err = binary.Read(r, binary.LittleEndian, &cre.Effects)
		if err != nil {
			return nil, err
		}
	case 1:
		cre.Effectsv2 = make([]EffEffect, cre.EffectListCount)
		_, err = r.Seek(int64(cre.EffectListOffset), io.SeekStart)
		if err != nil {
			return nil, err
		}
		err = binary.Read(r, binary.LittleEndian, &cre.Effectsv2)
		if err != nil {
			return nil, err
		}

	}
	return cre, nil
}

func (cre *CRE) Write(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, cre.CreHeader); err != nil {
		return err
	}
	order := map[uint32]func() error{
		cre.KnownSpellListOffset:        func() error { return binary.Write(w, binary.LittleEndian, cre.KnownSpells) },
		cre.MemorizationLevelListOffset: func() error { return binary.Write(w, binary.LittleEndian, cre.MemorizedSpellLevels) },
		cre.MemorizationSpellListOffset: func() error { return binary.Write(w, binary.LittleEndian, cre.MemorizedSpells) },
		cre.EquipmentListOffset:         func() error { return binary.Write(w, binary.LittleEndian, cre.ItemSlots) },
		cre.ItemListOffset:              func() error { return binary.Write(w, binary.LittleEndian, cre.Items) },
		cre.EffectListOffset: func() error {
			switch cre.EffStructureVersion {
			case 0:
				if err := binary.Write(w, binary.LittleEndian, cre.Effects); err != nil {
					return err
				}
			case 1:
				if err := binary.Write(w, binary.LittleEndian, cre.Effectsv2); err != nil {
					return err
				}
			}
			return nil
		},
	}
	for _, key := range slices.Sorted(maps.Keys(order)) {
		if err := order[key](); err != nil {
			return err
		}
	}

	return nil
}

func (cre *CRE) WriteJson(w io.Writer) error {
	bytes, err := json.MarshalIndent(cre, "", "\t")
	if err != nil {
		return err
	}

	_, err = w.Write(bytes)
	return err
}
