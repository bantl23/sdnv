package sdnv_test

import (
	"github.com/bantl23/sdnv"
	"testing"
)

func TestNewSdnv(t *testing.T) {
	for i := uint64(0); i < 10; i++ {
		s := sdnv.NewSdnv(i)
		if i != s.Value {
			t.Error("expected", i, "==", s.Value)
		}
	}
}

func TestMarshal(t *testing.T) {
	values := []sdnv.Sdnv{{0xabc, 0}, {0x1234, 0}, {0x4234, 0}, {0x7f, 0}}
	answers := [][]byte{{0x95, 0x3c}, {0xa4, 0x34}, {0x81, 0x84, 0x34}, {0x7f}}
	for i := 0; i < len(values); i++ {
		encoded := values[i].Marshal()
		if equal(encoded, answers[i]) == false {
			t.Error("expected encoded value", encoded, "==", answers)
		}
	}
}

func TestUnmarshal(t *testing.T) {
	values := []sdnv.Sdnv{{0xabc, 0}, {0x1234, 0}, {0x4234, 0}, {0x7f, 0}}
	for i := 0; i < len(values); i++ {
		encoded := values[i].Marshal()
		ans := sdnv.Sdnv{0, 0}
		err := ans.Unmarshal(encoded)
		if err != nil {
			t.Error(err)
		} else {
			if ans.Value != values[i].Value {
				t.Error("expected unmarshal value", ans.Value, "==", values[i].Value)
			}
		}
	}

	val := []byte{0xff, 0xff, 0xff}
	ans := sdnv.Sdnv{0, 0}
	err := ans.Unmarshal(val)
	if err == nil {
		t.Error("Expected error but got none")
	}
}
