package bg

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

type SplHeader struct {
	Signature             Signature `json:"signature"`
	Version               Version   `json:"version"`
	GenericName           uint32    `json:"generic_name"`
	IdentifiedName        uint32    `json:"identified_name"`
	UsedUpItemID          Resref    `json:"used_up_item_id"`
	ItemFlags             uint32    `json:"item_flags"`
	ItemType              uint16    `json:"item_type"`
	NotUsableBy           uint32    `json:"not_usable_by"`
	AnimationType         [2]uint8  `json:"animation_type"`
	MinLevelRequired      uint8     `json:"min_level_required"`
	School                uint8     `json:"school"`
	MinStrRequired        uint8     `json:"min_str_required"`
	SecondaryType         uint8     `json:"secondary_type"`
	MinStrBonusRequired   uint8     `json:"min_str_bonus_required"`
	NotUsableBy2a         uint8     `json:"not_usable_by2a"`
	MinIntRequired        uint8     `json:"min_int_required"`
	NotUsableBy2b         uint8     `json:"not_usable_by2b"`
	MinDexRequired        uint8     `json:"min_dex_required"`
	NotUsableBy2c         uint8     `json:"not_usable_by2c"`
	MinWisRequired        uint8     `json:"min_wis_required"`
	NotUsableBy2d         uint8     `json:"not_usable_by2d"`
	MinConRequired        uint16    `json:"min_con_required"`
	MinChrRequired        uint16    `json:"min_chr_required"`
	SpellLevel            uint32    `json:"spell_level"`
	MaxStackable          uint16    `json:"max_stackable"`
	ItemIcon              Resref    `json:"item_icon"`
	LoreValue             uint16    `json:"lore_value"`
	GroundIcon            Resref    `json:"ground_icon"`
	BaseWeight            uint32    `json:"base_weight"`
	GenericDescription    uint32    `json:"generic_description"`
	IdentifiedDescription uint32    `json:"identified_description"`
	DescriptionPicture    Resref    `json:"description_picture"`
	Attributes            uint32    `json:"attributes"`
	AbilityOffset         uint32    `json:"ability_offset"`
	AbilityCount          uint16    `json:"ability_count"`
	EffectsOffset         uint32    `json:"effects_offset"`
	CastingStartingEffect uint16    `json:"casting_starting_effect"`
	CastingEffectCount    uint16    `json:"casting_effect_count"`
}

type SplAbility struct {
	Type            uint16 `json:"type"`
	QuickSlotType   uint16 `json:"quick_slot_type"`
	QuickSlotIcon   Resref `json:"quick_slot_icon"`
	ActionType      uint8  `json:"action_type"`
	ActionCount     uint8  `json:"action_count"`
	Range           uint16 `json:"range"`
	MinCasterLevel  uint16 `json:"min_caster_level"`
	SpeedFactor     uint16 `json:"speed_factor"`
	TimesPerDay     uint16 `json:"times_per_day"`
	DamageDice      uint16 `json:"damage_dice"`
	DamageDiceCount uint16 `json:"damage_dice_count"`
	DamageDiceBonus uint16 `json:"damage_dice_bonus"`
	DamageType      uint16 `json:"damage_type"`
	EffectCount     uint16 `json:"effect_count"`
	StartingEffect  uint16 `json:"starting_effect"`
	MaxUsageCount   uint16 `json:"max_usage_count"`
	UsageFlags      uint16 `json:"usage_flags"`
	MissileType     uint16 `json:"missile_type"`
}

type SPL struct {
	SplHeader
	Abilities []SplAbility `json:"abilities"`
	Effects   []ItmEffect  `json:"effects"`
	Filename  string       `json:"-"`
}

func (spl *SPL) Write(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, spl.SplHeader); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, spl.Abilities); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, spl.Effects); err != nil {
		return err
	}
	return nil
}

func OpenSPL(r io.ReadSeeker) (*SPL, error) {
	spl := SPL{}

	err := binary.Read(r, binary.LittleEndian, &spl.SplHeader)
	if err != nil {
		return nil, err
	}

	spl.Abilities = make([]SplAbility, spl.AbilityCount)
	if err := parseArray(r, spl.AbilityOffset, spl.Abilities); err != nil {
		return nil, err
	}

	effectsCount := 0
	for _, ability := range spl.Abilities {
		effectsCount += int(ability.EffectCount)
	}
	effectsCount += int(spl.CastingEffectCount)
	spl.Effects = make([]ItmEffect, effectsCount)
	r.Seek(int64(spl.EffectsOffset), io.SeekStart)
	binary.Read(r, binary.LittleEndian, &spl.Effects)

	return &spl, nil
}

func (spl *SPL) WriteJson(w io.Writer) error {
	bytes, err := json.MarshalIndent(spl, "", "\t")
	if err != nil {
		return err
	}

	_, err = w.Write(bytes)
	return err
}
