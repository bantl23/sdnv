package sdnv

import (
	"errors"
)

type Sdnv struct {
	Value uint64
}

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

func (s *Sdnv) Unmarshal(data []byte) error {
	s.Value = uint64(0)
	for i := 0; i < len(data); i++ {
		s.Value = s.Value << 7
		s.Value = s.Value + uint64(data[i]&0x7f)
		if (data[i] >> 7) == 0 {
			break
		} else if i == (len(data) - 1) {
			return errors.New("Reached end of input without seeing end of SDNV")
		}
	}
	return nil
}
