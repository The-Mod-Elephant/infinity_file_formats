package bg

import (
	"encoding/binary"
	"encoding/json"
	"io"
	"maps"
	"reflect"
	"slices"
)

type AreaHeader struct {
	Signature               Signature `json:"signature"`
	Version                 Version   `json:"version"`
	AreaWed                 Resref    `json:"area_wed"`
	LastSaved               uint32    `json:"last_saved"`
	AreaFlags               uint32    `json:"area_flags"`
	AreaNorth               Resref    `json:"area_north"`
	AreaNorthFlags          uint32    `json:"area_north_flags"`
	AreaEast                Resref    `json:"area_east"`
	AreaEastFlags           uint32    `json:"area_east_flags"`
	AreaSouth               Resref    `json:"area_south"`
	AreaSouthFlags          uint32    `json:"area_south_flags"`
	AreaWest                Resref    `json:"area_west"`
	AreaWestFlags           uint32    `json:"area_west_flags"`
	Areatype                uint16    `json:"areatype"`
	Rainprobability         uint16    `json:"rainprobability"`
	SnowProability          uint16    `json:"snow_proability"`
	FogProbability          uint16    `json:"fog_probability"`
	LightningProbability    uint16    `json:"lightning_probability"`
	WindSpeed               uint16    `json:"wind_speed"`
	ActorsOffset            uint32    `json:"actors_offset"`
	ActorsCount             uint16    `json:"actors_count"`
	RegionCount             uint16    `json:"region_count"`
	RegionOffset            uint32    `json:"region_offset"`
	SpawnPointOffset        uint32    `json:"spawn_point_offset"`
	SpawnPointCount         uint32    `json:"spawn_point_count"`
	EntranceOffset          uint32    `json:"entrance_offset"`
	EntranceCount           uint32    `json:"entrance_count"`
	ContainerOffset         uint32    `json:"container_offset"`
	ContainerCount          uint16    `json:"container_count"`
	ItemCount               uint16    `json:"item_count"`
	ItemOffset              uint32    `json:"item_offset"`
	VertexOffset            uint32    `json:"vertex_offset"`
	VertexCount             uint16    `json:"vertex_count"`
	AmbientCount            uint16    `json:"ambient_count"`
	AmbientOffset           uint32    `json:"ambient_offset"`
	VariableOffset          uint32    `json:"variable_offset"`
	VariableCount           uint32    `json:"variable_count"`
	TiledObjectFlagOffset   uint16    `json:"tiled_object_flag_offset"`
	TiledObjectFlagCount    uint16    `json:"tiled_object_flag_count"`
	Script                  Resref    `json:"script"`
	ExploredSize            uint32    `json:"explored_size"`
	ExploredOffset          uint32    `json:"explored_offset"`
	DoorsCount              uint32    `json:"doors_count"`
	DoorsOffset             uint32    `json:"doors_offset"`
	AnimationCount          uint32    `json:"animation_count"`
	AnimationOffset         uint32    `json:"animation_offset"`
	TiledObjectCount        uint32    `json:"tiled_object_count"`
	TiledObjectOffset       uint32    `json:"tiled_object_offset"`
	SongEntriesOffset       uint32    `json:"song_entries_offset"`
	RestInterruptionsOffset uint32    `json:"rest_interruptions_offset"`
	AutomapOffset           uint32    `json:"automap_offset"`
	AutomapCount            uint32    `json:"automap_count"`
	ProjectileTrapsOffset   uint32    `json:"projectile_traps_offset"`
	ProjectileTrapsCount    uint32    `json:"projectile_traps_count"`
	RestMovieDay            Resref    `json:"rest_movie_day"`
	RestMovieNight          Resref    `json:"rest_movie_night"`
	Unknown                 [56]byte  `json:"unknown"`
}

type AreaActor struct {
	Name                LongString `json:"name"`
	CurrentX            uint16     `json:"current_x"`
	CurrentY            uint16     `json:"current_y"`
	DestX               uint16     `json:"dest_x"`
	DestY               uint16     `json:"dest_y"`
	Flags               uint32     `json:"flags"`
	Type                uint16     `json:"type"`
	FirstResSlot        byte       `json:"first_res_slot"`
	AlignByte           byte       `json:"align_byte"`
	AnimationType       uint32     `json:"animation_type"`
	Facing              uint16     `json:"facing"`
	AlignWord           uint16     `json:"align_word"`
	ExpirationTime      uint32     `json:"expiration_time"`
	HuntingRange        uint16     `json:"hunting_range"`
	FollowRange         uint16     `json:"follow_range"`
	TimeOfDayVisible    uint32     `json:"time_of_day_visible"`
	NumberTimesTalkedTo uint32     `json:"number_times_talked_to"`
	Dialog              Resref     `json:"dialog"`
	OverrideScript      Resref     `json:"override_script"`
	GeneralScript       Resref     `json:"general_script"`
	ClassScript         Resref     `json:"class_script"`
	RaceScript          Resref     `json:"race_script"`
	DefaultScript       Resref     `json:"default_script"`
	SpecificScript      Resref     `json:"specific_script"`
	CreatureData        Resref     `json:"creature_data"`
	CreatureOffset      uint32     `json:"creature_offset"`
	CreatureSize        uint32     `json:"creature_size"`
	Unused              [32]uint32 `json:"unused"`
}

type AreaRegion struct {
	Name                    LongString `json:"name"`
	Type                    uint16     `json:"type"`
	BoundingLeft            uint16     `json:"bounding_left"`
	BoundingTop             uint16     `json:"bounding_top"`
	BoundingRight           uint16     `json:"bounding_right"`
	BoundingBottom          uint16     `json:"bounding_bottom"`
	VertexCount             uint16     `json:"vertex_count"`
	VertexOffset            uint32     `json:"vertex_offset"`
	TriggerValue            uint32     `json:"trigger_value"`
	CursorType              uint32     `json:"cursor_type"`
	Destination             Resref     `json:"destination"`
	EntranceName            LongString `json:"entrance_name"`
	Flags                   uint32     `json:"flags"`
	InformationText         uint32     `json:"information_text"`
	TrapDetectionDifficulty uint16     `json:"trap_detection_difficulty"`
	TrapDisarmingDifficulty uint16     `json:"trap_disarming_difficulty"`
	TrapActivated           uint16     `json:"trap_activated"`
	TrapDetected            uint16     `json:"trap_detected"`
	TrapOriginX             uint16     `json:"trap_origin_x"`
	TrapOriginY             uint16     `json:"trap_origin_y"`
	KeyItem                 Resref     `json:"key_item"`
	RegionScript            Resref     `json:"region_script"`
	TransitionWalkToX       uint16     `json:"transition_walk_to_x"`
	TransitionWalkToY       uint16     `json:"transition_walk_to_y"`
	Unused                  [15]uint32 `json:"unused"`
}

type AreaSpawnPoint struct {
	Name                LongString `json:"name"`
	CoordX              uint16     `json:"coord_x"`
	CoordY              uint16     `json:"coord_y"`
	RandomCreatures     [10]Resref `json:"random_creatures"`
	RandomCreatureCount uint16     `json:"random_creature_count"`
	Difficulty          uint16     `json:"difficulty"`
	SpawnRate           uint16     `json:"spawn_rate"`
	Flags               uint16     `json:"flags"`
	LifeSpan            uint32     `json:"life_span"`
	HuntingRange        uint32     `json:"hunting_range"`
	FollowRange         uint32     `json:"follow_range"`
	MaxTypeNum          uint32     `json:"max_type_num"`
	Activated           uint16     `json:"activated"`
	TimeOfDay           uint32     `json:"time_of_day"`
	ProbabilityDay      uint16     `json:"probability_day"`
	ProbabilityNight    uint16     `json:"probability_night"`
	Unused              [14]uint32 `json:"unused"`
}

type AreaEntrance struct {
	Name        LongString `json:"name"`
	CoordX      uint16     `json:"coord_x"`
	CoordY      uint16     `json:"coord_y"`
	Orientation uint16     `json:"orientation"`
	Unused      [66]byte   `json:"unused"`
}

type AreaContainer struct {
	Name                    LongString `json:"name"`
	CoordX                  uint16     `json:"coord_x"`
	CoordY                  uint16     `json:"coord_y"`
	Type                    uint16     `json:"type"`
	LockDifficulty          uint16     `json:"lock_difficulty"`
	Flags                   uint32     `json:"flags"`
	TrapDetectionDifficulty uint16     `json:"trap_detection_difficulty"`
	TrapRemovalDifficulty   uint16     `json:"trap_removal_difficulty"`
	ContainerTrapped        uint16     `json:"container_trapped"`
	TrapDetected            uint16     `json:"trap_detected"`
	TrapLaunchX             uint16     `json:"trap_launch_x"`
	TrapLaunchY             uint16     `json:"trap_launch_y"`
	BoundingTopLeft         uint16     `json:"bounding_top_left"`
	BoundingTopRight        uint16     `json:"bounding_top_right"`
	BoundingBottomRight     uint16     `json:"bounding_bottom_right"`
	BoundingBottomLeft      uint16     `json:"bounding_bottom_left"`
	ItemOffset              uint32     `json:"item_offset"`
	ItemCount               uint32     `json:"item_count"`
	TrapScript              Resref     `json:"trap_script"`
	VertexOffset            uint32     `json:"vertex_offset"`
	VertexCount             uint16     `json:"vertex_count"`
	TriggerRange            uint16     `json:"trigger_range"`
	OwnedBy                 LongString `json:"owned_by"`
	KeyType                 Resref     `json:"key_type"`
	BreakDifficulty         uint32     `json:"break_difficulty"`
	NotPickableString       uint32     `json:"not_pickable_string"`
	Unused                  [14]uint32 `json:"unused"`
}

type AreaItem struct {
	Resource   Resref    `json:"resource"`
	Expiration uint16    `json:"expiration"`
	UsageCount [3]uint16 `json:"usage_count"`
	Flags      uint32    `json:"flags"`
}

type AreaVertex struct {
	Coordinate uint16 `json:"coordinate"`
}

type AreaAmbient struct {
	Name            LongString `json:"name"`
	CoordinateX     uint16     `json:"coordinate_x"`
	CoordinateY     uint16     `json:"coordinate_y"`
	Range           uint16     `json:"range"`
	Alignment1      uint16     `json:"alignment1"`
	PitchVariance   uint32     `json:"pitch_variance"`
	VolumeVariance  uint16     `json:"volume_variance"`
	Volume          uint16     `json:"volume"`
	Sounds          [10]Resref `json:"sounds"`
	SoundCount      uint16     `json:"sound_count"`
	Alignment2      uint16     `json:"alignment2"`
	Period          uint32     `json:"period"`
	PeriodVariance  uint32     `json:"period_variance"`
	TimeOfDayActive uint32     `json:"time_of_day_active"`
	Flags           uint32     `json:"flags"`
	Unused          [16]uint32 `json:"unused"`
}

type AreaVariable struct {
	Name       LongString `json:"name"`
	Type       uint16     `json:"type"`
	ResRefType uint16     `json:"res_ref_type"`
	DWValue    uint32     `json:"dw_value"`
	IntValue   int32      `json:"int_value"`
	FloatValue float64    `json:"float_value"`
	ScriptName LongString `json:"script_name"`
}

type AreaDoor struct {
	Name                    LongString `json:"name"`
	DoorID                  Resref     `json:"door_id"`
	Flags                   uint32     `json:"flags"`
	OpenDoorVertexOffset    uint32     `json:"open_door_vertex_offset"`
	OpenDoorVertexCount     uint16     `json:"open_door_vertex_count"`
	ClosedDoorVertexCount   uint16     `json:"closed_door_vertex_count"`
	CloseDoorVertexOffset   uint32     `json:"close_door_vertex_offset"`
	OpenBoundingLeft        uint16     `json:"open_bounding_left"`
	OpenBoundingTop         uint16     `json:"open_bounding_top"`
	OpenBoundingRight       uint16     `json:"open_bounding_right"`
	OpenBoundingBottom      uint16     `json:"open_bounding_bottom"`
	ClosedBoundingLeft      uint16     `json:"closed_bounding_left"`
	ClosedBoundingTop       uint16     `json:"closed_bounding_top"`
	ClosedBoundingRight     uint16     `json:"closed_bounding_right"`
	ClosedBoundingBottom    uint16     `json:"closed_bounding_bottom"`
	OpenBlockVertexOffset   uint32     `json:"open_block_vertex_offset"`
	OpenBlockVertexCount    uint16     `json:"open_block_vertex_count"`
	ClosedBlockVertexCount  uint16     `json:"closed_block_vertex_count"`
	ClosedBlockVertexOffset uint32     `json:"closed_block_vertex_offset"`
	HitPoints               uint16     `json:"hit_points"`
	ArmorClass              uint16     `json:"armor_class"`
	OpenSound               Resref     `json:"open_sound"`
	ClosedSound             Resref     `json:"closed_sound"`
	CursorType              uint32     `json:"cursor_type"`
	TrapDetectionDifficulty uint16     `json:"trap_detection_difficulty"`
	TrapRemovalDifficulty   uint16     `json:"trap_removal_difficulty"`
	DoorIsTrapped           uint16     `json:"door_is_trapped"`
	TrapDetected            uint16     `json:"trap_detected"`
	TrapLaunchTargetX       uint16     `json:"trap_launch_target_x"`
	TrapLaunchTargetY       uint16     `json:"trap_launch_target_y"`
	KeyItem                 Resref     `json:"key_item"`
	DoorScript              Resref     `json:"door_script"`
	DetectionDifficulty     uint32     `json:"detection_difficulty"`
	LockDifficulty          uint32     `json:"lock_difficulty"`
	WalkToX1                uint16     `json:"walk_to_x1"`
	WalkToY1                uint16     `json:"walk_to_y1"`
	WalkToX2                uint16     `json:"walk_to_x2"`
	WalkToY2                uint16     `json:"walk_to_y2"`
	NotPickableString       uint32     `json:"not_pickable_string"`
	TriggerName             LongString `json:"trigger_name"`
	Unused                  [3]uint32  `json:"unused"`
}

type AreaAnimation struct {
	Name             LongString `json:"name"`
	CoordX           uint16     `json:"coord_x"`
	CoordY           uint16     `json:"coord_y"`
	TimeOfDayVisible uint32     `json:"time_of_day_visible"`
	Animation        Resref     `json:"animation"`
	BamSequence      uint16     `json:"bam_sequence"`
	BamFrame         uint16     `json:"bam_frame"`
	Flags            uint32     `json:"flags"`
	Height           int16      `json:"height"`
	Translucency     uint16     `json:"translucency"`
	StartFrameRane   uint16     `json:"start_frame_rane"`
	Probability      byte       `json:"probability"`
	Period           byte       `json:"period"`
	Palette          Resref     `json:"palette"`
	Unused           uint32     `json:"unused"`
}

type AreaMapNote struct {
	CoordX uint16    `json:"coord_x"`
	CoordY uint16    `json:"coord_y"`
	Note   uint32    `json:"note"`
	Flags  uint32    `json:"flags"`
	Id     uint32    `json:"id"`
	Unused [9]uint32 `json:"unused"`
}

type AreaTiledObject struct {
	Name                       LongString `json:"name"`
	TileID                     Resref     `json:"tile_id"`
	Flags                      uint32     `json:"flags"`
	PrimarySearchSquareStart   uint32     `json:"primary_search_square_start"`
	PrimarySearchSquareCount   uint16     `json:"primary_search_square_count"`
	SecondarySearchSquareCount uint16     `json:"secondary_search_square_count"`
	SecondarySearcHSquareStart uint32     `json:"secondary_searc_h_square_start"`
	Unused                     [12]uint32 `json:"unused"`
}

type AreaProjectileTrap struct {
	Projectile        Resref `json:"projectile"`
	EffectBlockOffset uint32 `json:"effect_block_offset"`
	EffectBlockSize   uint16 `json:"effect_block_size"`
	MissileId         uint16 `json:"missile_id"`
	DelayCount        uint16 `json:"delay_count"`
	RepetitionCount   uint16 `json:"repetition_count"`
	CoordX            uint16 `json:"coord_x"`
	CoordY            uint16 `json:"coord_y"`
	CoordZ            uint16 `json:"coord_z"`
	TargetType        byte   `json:"target_type"`
	PortraitNum       byte   `json:"portrait_num"`
}

type AreaSong struct {
	DaySong              uint32     `json:"day_song"`
	NightSong            uint32     `json:"night_song"`
	WinSong              uint32     `json:"win_song"`
	BattleSong           uint32     `json:"battle_song"`
	LoseSong             uint32     `json:"lose_song"`
	AltMusic0            uint32     `json:"alt_music0"`
	AltMusic1            uint32     `json:"alt_music1"`
	AltMusic2            uint32     `json:"alt_music2"`
	AltMusic3            uint32     `json:"alt_music3"`
	AltMusic4            uint32     `json:"alt_music4"`
	DayAmbient           Resref     `json:"day_ambient"`
	DayAmbientExtended   Resref     `json:"day_ambient_extended"`
	DayAmbientVolume     uint32     `json:"day_ambient_volume"`
	NightAmbient         Resref     `json:"night_ambient"`
	NightAmbientExtended Resref     `json:"night_ambient_extended"`
	NightAmbientVolume   uint32     `json:"night_ambient_volume"`
	Unused               [16]uint32 `json:"unused"`
}

type AreaRestEncounter struct {
	Name                 LongString `json:"name"`
	RandomCreatureString [10]strref `json:"random_creature_string"`
	RandomCreature       [10]Resref `json:"random_creature"`
	RandomCreatureNum    uint16     `json:"random_creature_num"`
	Difficulty           uint16     `json:"difficulty"`
	LifeSpan             uint32     `json:"life_span"`
	HuntingRange         uint16     `json:"hunting_range"`
	FollowRange          uint16     `json:"follow_range"`
	MaxTypeNum           uint16     `json:"max_type_num"`
	Activated            uint16     `json:"activated"`
	ProbabilityDay       uint16     `json:"probability_day"`
	ProbabilityNight     uint16     `json:"probability_night"`
	Unused               [56]uint8  `json:"unused"`
}

type Area struct {
	AreaHeader
	Actors            []AreaActor          `json:"actors"`
	Regions           []AreaRegion         `json:"regions"`
	SpawnPoints       []AreaSpawnPoint     `json:"spawn_points"`
	Entrances         []AreaEntrance       `json:"entrances"`
	Containers        []AreaContainer      `json:"containers"`
	Items             []AreaItem           `json:"items"`
	Vertices          []AreaVertex         `json:"vertices"`
	Ambients          []AreaAmbient        `json:"ambients"`
	Variables         []AreaVariable       `json:"variables"`
	ExploredBitmasks  []byte               `json:"explored_bitmasks"`
	Doors             []AreaDoor           `json:"doors"`
	Animations        []AreaAnimation      `json:"animations"`
	MapNotes          []AreaMapNote        `json:"map_notes"`
	TiledObjects      []AreaTiledObject    `json:"tiled_objects"`
	Traps             []AreaProjectileTrap `json:"traps"`
	Songs             AreaSong             `json:"songs"`
	RestInterruptions AreaRestEncounter    `json:"rest_interruptions"`
	Filename          string               `json:"-"`
}

func (a *Area) Equal(other *Area) bool {
	if !reflect.DeepEqual(a.AreaHeader, other.AreaHeader) {
		return false
	}
	if !slices.Equal(a.Actors, other.Actors) {
		return false
	}
	if !slices.Equal(a.Regions, other.Regions) {
		return false
	}
	if !slices.Equal(a.SpawnPoints, other.SpawnPoints) {
		return false
	}
	if !slices.Equal(a.Entrances, other.Entrances) {
		return false
	}
	if !slices.Equal(a.Containers, other.Containers) {
		return false
	}
	if !slices.Equal(a.Items, other.Items) {
		return false
	}
	if !slices.Equal(a.Vertices, other.Vertices) {
		return false
	}
	if !slices.Equal(a.Ambients, other.Ambients) {
		return false
	}
	if !slices.Equal(a.Variables, other.Variables) {
		return false
	}
	if !slices.Equal(a.ExploredBitmasks, other.ExploredBitmasks) {
		return false
	}
	if !slices.Equal(a.Doors, other.Doors) {
		return false
	}
	if !slices.Equal(a.Animations, other.Animations) {
		return false
	}
	if !slices.Equal(a.MapNotes, other.MapNotes) {
		return false
	}
	if !slices.Equal(a.TiledObjects, other.TiledObjects) {
		return false
	}
	if !slices.Equal(a.Traps, other.Traps) {
		return false
	}
	if !reflect.DeepEqual(a.Songs, other.Songs) {
		return false
	}
	if !reflect.DeepEqual(a.RestInterruptions, other.RestInterruptions) {
		return false
	}
	return true
}

func OpenArea(r io.ReadSeeker) (*Area, error) {
	area := Area{}

	err := binary.Read(r, binary.LittleEndian, &area.AreaHeader)
	if err != nil {
		return nil, err
	}
	area.Actors = make([]AreaActor, area.ActorsCount)
	r.Seek(int64(area.ActorsOffset), io.SeekStart)
	err = binary.Read(r, binary.LittleEndian, &area.Actors)
	if err != nil {
		return nil, err
	}
	area.Regions = make([]AreaRegion, area.RegionCount)
	r.Seek(int64(area.RegionOffset), io.SeekStart)
	err = binary.Read(r, binary.LittleEndian, &area.Regions)
	if err != nil {
		return nil, err
	}
	area.SpawnPoints = make([]AreaSpawnPoint, area.SpawnPointCount)
	r.Seek(int64(area.SpawnPointOffset), io.SeekStart)
	err = binary.Read(r, binary.LittleEndian, &area.SpawnPoints)
	if err != nil {
		return nil, err
	}
	area.Entrances = make([]AreaEntrance, area.EntranceCount)
	r.Seek(int64(area.EntranceOffset), io.SeekStart)
	err = binary.Read(r, binary.LittleEndian, &area.Entrances)
	if err != nil {
		return nil, err
	}
	area.Containers = make([]AreaContainer, area.ContainerCount)
	r.Seek(int64(area.ContainerOffset), io.SeekStart)
	err = binary.Read(r, binary.LittleEndian, &area.Containers)
	if err != nil {
		return nil, err
	}
	area.Items = make([]AreaItem, area.ItemCount)
	r.Seek(int64(area.ItemOffset), io.SeekStart)
	err = binary.Read(r, binary.LittleEndian, &area.Items)
	if err != nil {
		return nil, err
	}
	area.Vertices = make([]AreaVertex, area.VertexCount)
	r.Seek(int64(area.VertexOffset), io.SeekStart)
	err = binary.Read(r, binary.LittleEndian, &area.Vertices)
	if err != nil {
		return nil, err
	}
	area.Ambients = make([]AreaAmbient, area.AmbientCount)
	r.Seek(int64(area.AmbientOffset), io.SeekStart)
	err = binary.Read(r, binary.LittleEndian, &area.Ambients)
	if err != nil {
		return nil, err
	}
	area.Variables = make([]AreaVariable, area.VariableCount)
	r.Seek(int64(area.VariableOffset), io.SeekStart)
	err = binary.Read(r, binary.LittleEndian, &area.Variables)
	if err != nil {
		return nil, err
	}
	area.ExploredBitmasks = make([]byte, area.ExploredSize)
	r.Seek(int64(area.VariableOffset), io.SeekStart)
	err = binary.Read(r, binary.LittleEndian, &area.ExploredBitmasks)
	if err != nil {
		return nil, err
	}
	area.Doors = make([]AreaDoor, area.DoorsCount)
	r.Seek(int64(area.DoorsOffset), io.SeekStart)
	err = binary.Read(r, binary.LittleEndian, &area.Doors)
	if err != nil {
		return nil, err
	}
	area.Animations = make([]AreaAnimation, area.AnimationCount)
	r.Seek(int64(area.AnimationOffset), io.SeekStart)
	err = binary.Read(r, binary.LittleEndian, &area.Animations)
	if err != nil {
		return nil, err
	}
	area.MapNotes = make([]AreaMapNote, area.AutomapCount)
	r.Seek(int64(area.AutomapOffset), io.SeekStart)
	err = binary.Read(r, binary.LittleEndian, &area.MapNotes)
	if err != nil {
		return nil, err
	}
	area.TiledObjects = make([]AreaTiledObject, area.TiledObjectCount)
	r.Seek(int64(area.TiledObjectOffset), io.SeekStart)
	err = binary.Read(r, binary.LittleEndian, &area.TiledObjects)
	if err != nil {
		return nil, err
	}
	area.Traps = make([]AreaProjectileTrap, area.ProjectileTrapsCount)
	r.Seek(int64(area.ProjectileTrapsOffset), io.SeekStart)
	err = binary.Read(r, binary.LittleEndian, &area.Traps)
	if err != nil {
		return nil, err
	}
	r.Seek(int64(area.SongEntriesOffset), io.SeekStart)
	err = binary.Read(r, binary.LittleEndian, &area.Songs)
	if err != nil {
		return nil, err
	}
	r.Seek(int64(area.RestInterruptionsOffset), io.SeekStart)
	err = binary.Read(r, binary.LittleEndian, &area.RestInterruptions)
	if err != nil {
		return nil, err
	}

	return &area, nil
}

func (area *Area) WriteJson(w io.Writer) error {
	bytes, err := json.MarshalIndent(area, "", "\t")
	if err != nil {
		return err
	}

	_, err = w.Write(bytes)
	return err
}

func (are *Area) Write(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, are.AreaHeader); err != nil {
		return err
	}
	order := map[uint32]func() error{
		are.ActorsOffset:                  func() error { return binary.Write(w, binary.LittleEndian, are.Actors) },
		are.RegionOffset:                  func() error { return binary.Write(w, binary.LittleEndian, are.Regions) },
		are.SpawnPointOffset:              func() error { return binary.Write(w, binary.LittleEndian, are.SpawnPoints) },
		are.EntranceOffset:                func() error { return binary.Write(w, binary.LittleEndian, are.Entrances) },
		are.ContainerOffset:               func() error { return binary.Write(w, binary.LittleEndian, are.Containers) },
		are.ItemOffset:                    func() error { return binary.Write(w, binary.LittleEndian, are.Items) },
		are.VertexOffset:                  func() error { return binary.Write(w, binary.LittleEndian, are.Vertices) },
		are.AmbientOffset:                 func() error { return binary.Write(w, binary.LittleEndian, are.Ambients) },
		are.VariableOffset:                func() error { return binary.Write(w, binary.LittleEndian, are.Variables) },
		uint32(are.TiledObjectFlagOffset): func() error { return binary.Write(w, binary.LittleEndian, are.TiledObjects) },
		are.ExploredOffset:                func() error { return binary.Write(w, binary.LittleEndian, are.ExploredBitmasks) },
		are.DoorsOffset:                   func() error { return binary.Write(w, binary.LittleEndian, are.Doors) },
		are.AnimationOffset:               func() error { return binary.Write(w, binary.LittleEndian, are.Animations) },
		are.TiledObjectOffset:             func() error { return binary.Write(w, binary.LittleEndian, are.TiledObjects) },
		are.SongEntriesOffset:             func() error { return binary.Write(w, binary.LittleEndian, are.Songs) },
		are.RestInterruptionsOffset:       func() error { return binary.Write(w, binary.LittleEndian, are.RestInterruptions) },
		are.AutomapOffset:                 func() error { return binary.Write(w, binary.LittleEndian, are.MapNotes) },
		are.ProjectileTrapsOffset:         func() error { return binary.Write(w, binary.LittleEndian, are.Traps) },
	}
	for _, key := range slices.Sorted(maps.Keys(order)) {
		if err := order[key](); err != nil {
			return err
		}
	}

	return nil
}
