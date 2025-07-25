package bg

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

type areaHeader struct {
	Signature               Signature
	Version                 Version
	AreaWed                 Resref
	LastSaved               uint32
	AreaFlags               uint32
	AreaNorth               Resref
	AreaNorthFlags          uint32
	AreaEast                Resref
	AreaEastFlags           uint32
	AreaSouth               Resref
	AreaSouthFlags          uint32
	AreaWest                Resref
	AreaWestFlags           uint32
	Areatype                uint16
	Rainprobability         uint16
	SnowProability          uint16
	FogProbability          uint16
	LightningProbability    uint16
	WindSpeed               uint16
	ActorsOffset            uint32
	ActorsCount             uint16
	RegionCount             uint16
	RegionOffset            uint32
	SpawnPointOffset        uint32
	SpawnPointCount         uint32
	EntranceOffset          uint32
	EntranceCount           uint32
	ContainerOffset         uint32
	ContainerCount          uint16
	ItemCount               uint16
	ItemOffset              uint32
	VertexOffset            uint32
	VertexCount             uint16
	AmbientCount            uint16
	AmbientOffset           uint32
	VariableOffset          uint32
	VariableCount           uint16
	TiledObjectFlagCount    uint16
	TiledObjectFlagOffset   uint32
	Script                  Resref
	ExploredSize            uint32
	ExploredOffset          uint32
	DoorsCount              uint32
	DoorsOffset             uint32
	AnimationCount          uint32
	AnimationOffset         uint32
	TiledObjectCount        uint32
	TiledObjectOffset       uint32
	SongEntriesOffset       uint32
	RestInterruptionsOffset uint32
	AutomapOffset           uint32
	AutomapCount            uint32
	ProjectileTrapsOffset   uint32
	ProjectileTrapsCount    uint32
	RestMovieDay            Resref
	RestMovieNight          Resref
	Unknown                 [56]byte
}

type AreaActor struct {
	Name                LongString
	CurrentX            uint16
	CurrentY            uint16
	DestX               uint16
	DestY               uint16
	Flags               uint32
	Type                uint16
	FirstResSlot        byte
	AlignByte           byte
	AnimationType       uint32
	Facing              uint16
	AlignWord           uint16
	ExpirationTime      uint32
	HuntingRange        uint16
	FollowRange         uint16
	TimeOfDayVisible    uint32
	NumberTimesTalkedTo uint32
	Dialog              Resref
	OverrideScript      Resref
	GeneralScript       Resref
	ClassScript         Resref
	RaceScript          Resref
	DefaultScript       Resref
	SpecificScript      Resref
	CreatureData        Resref
	CreatureOffset      uint32
	CreatureSize        uint32
	Unused              [32]uint32
}

type AreaRegion struct {
	Name                    LongString
	Type                    uint16
	BoundingLeft            uint16
	BoundingTop             uint16
	BoundingRight           uint16
	BoundingBottom          uint16
	VertexCount             uint16
	VertexOffset            uint32
	TriggerValue            uint32
	CursorType              uint32
	Destination             Resref
	EntranceName            LongString
	Flags                   uint32
	InformationText         uint32
	TrapDetectionDifficulty uint16
	TrapDisarmingDifficulty uint16
	TrapActivated           uint16
	TrapDetected            uint16
	TrapOriginX             uint16
	TrapOriginY             uint16
	KeyItem                 Resref
	RegionScript            Resref
	TransitionWalkToX       uint16
	TransitionWalkToY       uint16
	Unused                  [15]uint32
}

type AreaSpawnPoint struct {
	Name                LongString
	CoordX              uint16
	CoordY              uint16
	RandomCreatures     [10]Resref
	RandomCreatureCount uint16
	Difficulty          uint16
	SpawnRate           uint16
	Flags               uint16
	LifeSpan            uint32
	HuntingRange        uint32
	FollowRange         uint32
	MaxTypeNum          uint32
	Activated           uint16
	TimeOfDay           uint32
	ProbabilityDay      uint16
	ProbabilityNight    uint16
	Unused              [14]uint32
}

type AreaEntrance struct {
	Name        LongString
	CoordX      uint16
	CoordY      uint16
	Orientation uint16
	Unused      [66]byte
}

type AreaContainer struct {
	Name                    LongString
	CoordX                  uint16
	CoordY                  uint16
	Type                    uint16
	LockDifficulty          uint16
	Flags                   uint32
	TrapDetectionDifficulty uint16
	TrapRemovalDifficulty   uint16
	ContainerTrapped        uint16
	TrapDetected            uint16
	TrapLaunchX             uint16
	TrapLaunchY             uint16
	BoundingTopLeft         uint16
	BoundingTopRight        uint16
	BoundingBottomRight     uint16
	BoundingBottomLeft      uint16
	ItemOffset              uint32
	ItemCount               uint32
	TrapScript              Resref
	VertexOffset            uint32
	VertexCount             uint16
	TriggerRange            uint16
	OwnedBy                 LongString
	KeyType                 Resref
	BreakDifficulty         uint32
	NotPickableString       uint32
	Unused                  [14]uint32
}

type AreaItem struct {
	Resource   Resref
	Expiration uint16
	UsageCount [3]uint16
	Flags      uint32
}

type AreaVertex struct {
	Coordinate uint16
}

type AreaAmbient struct {
	Name            LongString
	CoordinateX     uint16
	CoordinateY     uint16
	Range           uint16
	Alignment1      uint16
	PitchVariance   uint32
	VolumeVariance  uint16
	Volume          uint16
	Sounds          [10]Resref
	SoundCount      uint16
	Alignment2      uint16
	Period          uint32
	PeriodVariance  uint32
	TimeOfDayActive uint32
	Flags           uint32
	Unused          [16]uint32
}

type AreaVariable struct {
	Name       LongString
	Type       uint16
	ResRefType uint16
	DWValue    uint32
	IntValue   int32
	FloatValue float64
	ScriptName LongString
}

type AreaDoor struct {
	Name                    LongString
	DoorID                  Resref
	Flags                   uint32
	OpenDoorVertexOffset    uint32
	OpenDoorVertexCount     uint16
	ClosedDoorVertexCount   uint16
	CloseDoorVertexOffset   uint32
	OpenBoundingLeft        uint16
	OpenBoundingTop         uint16
	OpenBoundingRight       uint16
	OpenBoundingBottom      uint16
	ClosedBoundingLeft      uint16
	ClosedBoundingTop       uint16
	ClosedBoundingRight     uint16
	ClosedBoundingBottom    uint16
	OpenBlockVertexOffset   uint32
	OpenBlockVertexCount    uint16
	ClosedBlockVertexCount  uint16
	ClosedBlockVertexOffset uint32
	HitPoints               uint16
	ArmorClass              uint16
	OpenSound               Resref
	ClosedSound             Resref
	CursorType              uint32
	TrapDetectionDifficulty uint16
	TrapRemovalDifficulty   uint16
	DoorIsTrapped           uint16
	TrapDetected            uint16
	TrapLaunchTargetX       uint16
	TrapLaunchTargetY       uint16
	KeyItem                 Resref
	DoorScript              Resref
	DetectionDifficulty     uint32
	LockDifficulty          uint32
	WalkToX1                uint16
	WalkToY1                uint16
	WalkToX2                uint16
	WalkToY2                uint16
	NotPickableString       uint32
	TriggerName             LongString
	Unused                  [3]uint32
}

type AreaAnimation struct {
	Name             LongString
	CoordX           uint16
	CoordY           uint16
	TimeOfDayVisible uint32
	Animation        Resref
	BamSequence      uint16
	BamFrame         uint16
	Flags            uint32
	Height           int16
	Translucency     uint16
	StartFrameRane   uint16
	Probability      byte
	Period           byte
	Palette          Resref
	Unused           uint32
}

type AreaMapNote struct {
	CoordX uint16
	CoordY uint16
	Note   uint32
	Flags  uint32
	Id     uint32
	Unused [9]uint32
}

type AreaTiledObject struct {
	Name                       LongString
	TileID                     Resref
	Flags                      uint32
	PrimarySearchSquareStart   uint32
	PrimarySearchSquareCount   uint16
	SecondarySearchSquareCount uint16
	SecondarySearcHSquareStart uint32
	Unused                     [12]uint32
}

type AreaProjectileTrap struct {
	Projectile        Resref
	EffectBlockOffset uint32
	EffectBlockSize   uint16
	MissileId         uint16
	DelayCount        uint16
	RepetitionCount   uint16
	CoordX            uint16
	CoordY            uint16
	CoordZ            uint16
	TargetType        byte
	PortraitNum       byte
}

type AreaSong struct {
	DaySong              uint32
	NightSong            uint32
	WinSong              uint32
	BattleSong           uint32
	LoseSong             uint32
	AltMusic0            uint32
	AltMusic1            uint32
	AltMusic2            uint32
	AltMusic3            uint32
	AltMusic4            uint32
	DayAmbient           Resref
	DayAmbientExtended   Resref
	DayAmbientVolume     uint32
	NightAmbient         Resref
	NightAmbientExtended Resref
	NightAmbientVolume   uint32
	Unused               [16]uint32
}

type AreaRestEncounter struct {
	Name                 LongString
	RandomCreatureString [10]uint32
	RandomCreature       [10]Resref
	RandomCreatureNum    uint16
	Difficulty           uint16
	LifeSpan             uint32
	HuntingRange         uint16
	FollowRange          uint16
	MaxTypeNum           uint16
	Activated            uint16
	ProbabilityDay       uint16
	ProbabilityNight     uint16
	Unused               [14]uint32
}

type Area struct {
	areaHeader
	Actors           []AreaActor
	Regions          []AreaRegion
	SpawnPoints      []AreaSpawnPoint
	Entrances        []AreaEntrance
	Containers       []AreaContainer
	Items            []AreaItem
	Vertices         []AreaVertex
	Ambients         []AreaAmbient
	Variables        []AreaVariable
	ExploredBitmask  []byte
	Doors            []AreaDoor
	Animations       []AreaAnimation
	MapNotes         []AreaMapNote
	TiledObjects     []AreaTiledObject
	Traps            []AreaProjectileTrap
	Song             AreaSong
	RestInterruption AreaRestEncounter
}

func OpenArea(r io.ReadSeeker) (*Area, error) {
	area := Area{}

	err := binary.Read(r, binary.LittleEndian, &area.areaHeader)
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
	area.ExploredBitmask = make([]byte, area.ExploredSize)
	r.Seek(int64(area.VariableOffset), io.SeekStart)
	err = binary.Read(r, binary.LittleEndian, &area.ExploredBitmask)
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
	err = binary.Read(r, binary.LittleEndian, &area.Song)
	if err != nil {
		return nil, err
	}
	r.Seek(int64(area.RestInterruptionsOffset), io.SeekStart)
	err = binary.Read(r, binary.LittleEndian, &area.RestInterruption)
	if err != nil {
		return nil, err
	}

	return &area, nil
}

func (are *Area) WriteJson(w io.Writer) error {
	bytes, err := json.MarshalIndent(are, "", "\t")
	if err != nil {
		return err
	}

	_, err = w.Write(bytes)
	return err
}
