package bg

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
)

const (
	NUM_ATTACK_TYPES = 6
)

type ItmHeader struct {
	Signature              Signature `json:"signature"`
	Version                Version   `json:"version"`
	GenericName            uint32    `json:"generic_name"`
	IdentifiedName         uint32    `json:"identified_name"`
	UsedUpItemID           Resref    `json:"used_up_item_id"`
	ItemFlags              uint32    `json:"item_flags"`
	ItemType               uint16    `json:"item_type"`
	NotUsableBy            uint32    `json:"not_usable_by"`
	AnimationType          uint16    `json:"animation_type"`
	MinLevelRequired       uint16    `json:"min_level_required"`
	MinStrRequired         uint16    `json:"min_str_required"`
	MinStrBonusRequired    uint8     `json:"min_str_bonus_required"`
	NotUsableBy2a          uint8     `json:"not_usable_by2a"`
	MinIntRequired         uint8     `json:"min_int_required"`
	NotUsableBy2b          uint8     `json:"not_usable_by2b"`
	MinDexRequired         uint8     `json:"min_dex_required"`
	NotUsableBy2c          uint8     `json:"not_usable_by2c"`
	MinWisRequired         uint8     `json:"min_wis_required"`
	NotUsableBy2d          uint8     `json:"not_usable_by2d"`
	MinConRequired         uint8     `json:"min_con_required"`
	ProficiencyType        uint8     `json:"proficiency_type"`
	MinChrRequired         uint16    `json:"min_chr_required"`
	BaseValue              uint32    `json:"base_value"`
	MaxStackable           uint16    `json:"max_stackable"`
	ItemIcon               Resref    `json:"item_icon"`
	LoreValue              uint16    `json:"lore_value"`
	GroundIcon             Resref    `json:"ground_icon"`
	BaseWeight             uint32    `json:"base_weight"`
	GenericDescription     uint32    `json:"generic_description"`
	IdentifiedDescription  uint32    `json:"identified_description"`
	DescriptionPicture     Resref    `json:"description_picture"`
	Attributes             uint32    `json:"attributes"`
	AbilityOffset          uint32    `json:"ability_offset"`
	AbilityCount           uint16    `json:"ability_count"`
	EffectsOffset          uint32    `json:"effects_offset"`
	EquippedStartingEffect uint16    `json:"equipped_starting_effect"`
	EquippedEffectCount    uint16    `json:"equiped_effect_count"`
}

type itmAbility struct {
	Type                 uint16                   `json:"type"`
	QuickSlotType        uint8                    `json:"quick_slot_type"`
	LargeDamageDice      uint8                    `json:"large_damage_dice"`
	QuickSlotIcon        Resref                   `json:"quick_slot_icon"`
	ActionType           uint8                    `json:"action_type"`
	ActionCount          uint8                    `json:"action_count"`
	Range                uint16                   `json:"range"`
	LauncherType         uint8                    `json:"launcher_type"`
	LargeDamageDiceCount uint8                    `json:"large_damage_dice_count"`
	SpeedFactor          uint8                    `json:"speed_factor"`
	LargeDamageDiceBonus uint8                    `json:"large_damage_dice_bonus"`
	Thac0Bonus           int16                    `json:"thac0_bonus"`
	DamageDice           uint8                    `json:"damage_dice"`
	School               uint8                    `json:"school"`
	DamageDiceCount      uint8                    `json:"damage_dice_count"`
	SecondaryType        uint8                    `json:"secondary_type"`
	DamageDiceBonus      uint16                   `json:"damage_dice_bonus"`
	DamageType           uint16                   `json:"damage_type"`
	EffectCount          uint16                   `json:"effect_count"`
	StartingEffect       uint16                   `json:"starting_effect"`
	MaxUsageCount        uint16                   `json:"max_usage_count"`
	UsageFlags           uint16                   `json:"usage_flags"`
	AbilityFlags         uint32                   `json:"ability_flags"`
	MissileType          uint16                   `json:"missile_type"`
	AttackProbability    [NUM_ATTACK_TYPES]uint16 `json:"attack_probability"`
}

type ItmEffect struct {
	EffectID         uint16 `json:"effect_id"`
	TargetType       uint8  `json:"target_type"`
	SpellLevel       uint8  `json:"spell_level"`
	EffectAmount     int32  `json:"effect_amount"`
	Flags            uint32 `json:"flags"`
	DurationType     uint16 `json:"duration_type"`
	Duration         uint32 `json:"duration"`
	ProbabilityUpper uint8  `json:"probability_upper"`
	ProbabilityLower uint8  `json:"probability_lower"`
	Res              Resref `json:"res"`
	NumDice          uint32 `json:"num_dice"`
	DiceSize         uint32 `json:"dice_size"`
	SavingThrow      uint32 `json:"saving_throw"`
	SaveMod          int32  `json:"save_mod"`
	Special          uint32 `json:"special"`
}

type ITM struct {
	ItmHeader
	Abilities []itmAbility `json:"abilities"`
	Effects   []ItmEffect  `json:"effects"`
	Filename  string       `json:"-"`
}

func (itm *ITM) Write(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, itm.ItmHeader); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, itm.Abilities); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, itm.Effects); err != nil {
		return err
	}
	return nil
}

func (itm *ITM) Tp2Block(baseNum int) (string, []int) {
	stringIds := []int{
		int(itm.GenericName),
		int(itm.IdentifiedName),
		int(itm.GenericDescription),
		int(itm.IdentifiedDescription),
	}
	out := "COPY_EXISTING ~" + itm.Filename + "~ ~override~\n"
	if int(itm.GenericName) >= 0 {
		out += fmt.Sprintf("\tSAY 0x0008 #%d\n", baseNum+int(itm.GenericName))
	}
	if int(itm.IdentifiedName) >= 0 {
		out += fmt.Sprintf("\tSAY 0x000C #%d\n", baseNum+int(itm.IdentifiedName))
	}
	if int(itm.GenericDescription) >= 0 {
		out += fmt.Sprintf("\tSAY 0x0050 #%d\n", baseNum+int(itm.GenericDescription))
	}
	if int(itm.IdentifiedDescription) >= 0 {
		out += fmt.Sprintf("\tSAY 0x0054 #%d\n", baseNum+int(itm.IdentifiedDescription))
	}

	if int(itm.GenericName) >= 0 {
		out += fmt.Sprintf("STRING_SET #%d = @%d\n", baseNum+int(itm.GenericName), baseNum+int(itm.GenericName))
	}
	if int(itm.IdentifiedName) >= 0 {
		out += fmt.Sprintf("STRING_SET #%d = @%d\n", baseNum+int(itm.IdentifiedName), baseNum+int(itm.IdentifiedName))
	}
	if int(itm.GenericDescription) >= 0 {
		out += fmt.Sprintf("STRING_SET #%d = @%d\n", baseNum+int(itm.GenericDescription), baseNum+int(itm.GenericDescription))
	}
	if int(itm.IdentifiedDescription) >= 0 {
		out += fmt.Sprintf("STRING_SET #%d = @%d\n", baseNum+int(itm.IdentifiedDescription), baseNum+int(itm.IdentifiedDescription))
	}
	return out, stringIds

}

func (itm *ITM) Strings() map[string]int {
	names := map[string]int{}
	if int(itm.GenericName) >= 0 {
		names["genericName"] = int(itm.GenericName)
	}
	if int(itm.IdentifiedName) >= 0 {
		names["identifiedName"] = int(itm.IdentifiedName)
	}
	if int(itm.GenericDescription) >= 0 {
		names["genericDescription"] = int(itm.GenericDescription)
	}
	if int(itm.IdentifiedDescription) >= 0 {
		names["identifiedDescription"] = int(itm.IdentifiedDescription)
	}
	return names
}

func OpenITM(r io.ReadSeeker) (*ITM, error) {
	itm := &ITM{}

	err := binary.Read(r, binary.LittleEndian, &itm.ItmHeader)
	if err != nil {
		return nil, err
	}

	itm.Abilities = make([]itmAbility, itm.AbilityCount)
	_, err = r.Seek(int64(itm.AbilityOffset), io.SeekStart)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.LittleEndian, &itm.Abilities)
	if err != nil {
		return nil, err
	}
	effectsCount := 0
	for _, ability := range itm.Abilities {
		effectsCount += int(ability.EffectCount)
	}
	effectsCount += int(itm.EquippedEffectCount)
	itm.Effects = make([]ItmEffect, effectsCount)
	r.Seek(int64(itm.EffectsOffset), io.SeekStart)
	binary.Read(r, binary.LittleEndian, &itm.Effects)

	return itm, nil
}

func (itm *ITM) WriteJson(w io.Writer) error {
	bytes, err := json.MarshalIndent(itm, "", "\t")
	if err != nil {
		return err
	}

	_, err = w.Write(bytes)
	return err
}
