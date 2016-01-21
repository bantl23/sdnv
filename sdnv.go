package sdnv

import (
	"errors"
	"reflect"
)

// Sdnv struct containing the value and
// encoded byte length information
type Sdnv struct {
	Value  uint64
	EncLen uint64
}

// NewSdnv creates and initializes
// a new sdnv struct
func NewSdnv(val uint64) *Sdnv {
	s := new(Sdnv)
	s.Value = val
	return s
}

// Marshal returns and sdnv encoded byte array
func (s Sdnv) Marshal() []byte {
	data := []byte{}
	flag := byte(0)
	done := false
	for done == false {
		newbits := byte(s.Value & 0x7f)
		s.Value = s.Value >> 7
		newbyte := byte(newbits + flag)
		data = append([]byte{newbyte}, data...)
		if flag == 0 {
			flag = 0x80
		}
		if s.Value == 0 {
			done = true
		}
	}
	return data
}

// Unmarshal unencodes a byte array into an
// sdnv structure
func (s *Sdnv) Unmarshal(data []byte) error {
	s.Value = uint64(0)
	s.EncLen = 0
	length := int(reflect.TypeOf(s.Value).Size())
	if len(data) < length {
		length = len(data)
	}
	for i := 0; i < length; i++ {
		s.Value = s.Value << 7
		s.Value = s.Value + uint64(data[i]&0x7f)
		if (data[i] >> 7) == 0 {
			s.EncLen += 1
			break
		} else if i == (length - 1) {
			return errors.New("Reached end of input without seeing end of SDNV")
		}
	}
	return nil
}
