package bg

import (
	"encoding/json"
	"strings"
)

type LongString struct {
	Value [32]byte
}

func (r *LongString) String() string {
	str := strings.Split(string(r.Value[0:]), "\x00")[0]
	return str
}

func (r *LongString) MarshalJSON() ([]byte, error) {
	return []byte("\"" + r.String() + "\""), nil
}

type Resref struct {
	Name [8]byte
}

func NewResref(name string) Resref {
	r := Resref{}
	copy(r.Name[:], []byte(name))
	return r
}

func (r *Resref) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *Resref) Valid() bool {
	return r.String() != ""
}

func (r Resref) String() string {
	str := strings.Split(string(r.Name[0:]), "\x00")[0]
	return str
}
