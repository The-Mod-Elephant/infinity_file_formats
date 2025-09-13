package bg

import (
	"encoding/binary"
	"encoding/json"
	"io"
	"strings"
)

type BG interface {
	Write(w io.Writer) error
	WriteJson(w io.Writer) error
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type LongString [32]byte

func (l *LongString) String() string {
	return string(l[:])
}

func (l *LongString) MarshalJSON() ([]byte, error) {
	return []byte(l.String()), nil
}

func (l *LongString) UnmarshalJSON(b []byte) error {
	for i := range min(len(b)-2, 32) {
		l[i] = b[i+1]
	}
	return nil
}

type Signature [4]byte

func (s Signature) String() string {
	return string(s[:])
}

func (s *Signature) UnmarshalJSON(b []byte) error {
	if len(b) > 2 {
		for i := range min(len(b)-2, 4) {
			s[i] = b[i+1]
		}
	}
	return nil
}

func (s *Signature) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *Signature) Valid() bool {
	return len(s) != 0
}

type Version [4]byte

func (v Version) String() string {
	return string(v[0:])
}

func (v *Version) UnmarshalJSON(b []byte) error {
	if len(b) > 2 {
		for i := range min(len(b)-2, 4) {
			v[i] = b[i+1]
		}
	}
	return nil
}

func (v *Version) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

func (v *Version) Valid() bool {
	return len(v) != 0
}

type Resref [8]byte

func NewResref(name string) Resref {
	return Resref([]byte(name))
}

func (r *Resref) UnmarshalJSON(b []byte) error {
	if len(b) > 2 {
		for i := range min(len(b)-2, 8) {
			r[i] = b[i+1]
		}
	}
	return nil
}

func (r *Resref) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *Resref) Valid() bool {
	return len(r) != 0
}

func (r Resref) String() string {
	return strings.Split(string(r[0:]), "\x00")[0]
}

type strref uint32

func parseArray(r io.ReadSeeker, start uint32, out any) error {
	if _, err := r.Seek(int64(start), io.SeekStart); err != nil {
		return err
	}
	return binary.Read(r, binary.LittleEndian, out)
}
