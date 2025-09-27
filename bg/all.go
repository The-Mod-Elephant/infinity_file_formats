package bg

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"unicode/utf8"
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

func parseArray(r io.ReadSeeker, start uint32, data any) error {
	if _, err := r.Seek(int64(start), io.SeekStart); err != nil {
		return err
	}
	return binary.Read(r, binary.LittleEndian, data)
}

func asciiBytesToString(b []byte) string {
	out := ""
	for _, c := range b {
		out += string(c)
	}
	return out
}

type LongString [32]byte

func (l *LongString) UnmarshalJSON(b []byte) error {
	var decoded string
	if err := json.Unmarshal(b, &decoded); err != nil {
		return fmt.Errorf("%v, error for byte slice of %v", err, b)
	}
	for i := 0; len(decoded) > 0 && i < 32; i++ {
		asciiRune, size := utf8.DecodeRune([]byte(decoded))
		l[i] = byte(asciiRune)
		decoded = decoded[size:]
	}
	return nil
}

func (l *LongString) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}

func (l *LongString) Valid() bool {
	return len(l) != 0
}

func (l *LongString) String() string {
	return asciiBytesToString(l[:])
}

type Signature [4]byte

func (s *Signature) UnmarshalJSON(b []byte) error {
	var decoded string
	if err := json.Unmarshal(b, &decoded); err != nil {
		return err
	}
	*s = Signature([]byte(decoded))
	return nil
}

func (s *Signature) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *Signature) Valid() bool {
	return len(s) != 0
}

func (s Signature) String() string {
	return asciiBytesToString(s[:])
}

type Version [4]byte

func (v *Version) UnmarshalJSON(b []byte) error {
	var decoded string
	if err := json.Unmarshal(b, &decoded); err != nil {
		return err
	}
	*v = Version([]byte(decoded))
	return nil
}

func (v *Version) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

func (v *Version) Valid() bool {
	return len(v) != 0
}

func (v Version) String() string {
	return asciiBytesToString(v[:])
}

type Resref [8]byte

func (r *Resref) UnmarshalJSON(b []byte) error {
	var decoded string
	if err := json.Unmarshal(b, &decoded); err != nil {
		return fmt.Errorf("%v, error for byte slice of %v", err, b)
	}
	for i := 0; len(decoded) > 0 && i < 8; i++ {
		asciiRune, size := utf8.DecodeRune([]byte(decoded))
		r[i] = byte(asciiRune)
		decoded = decoded[size:]
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
	return asciiBytesToString(r[:])
}

type strref uint32
