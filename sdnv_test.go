package sdnv_test

import (
	"gitlab.com/bantl23/sdnv"
	"testing"
)

func TestMarshal(t *testing.T) {
	values := []sdnv.Sdnv{{0xabc}, {0x1234}, {0x4234}, {0x7f}}
	answers := [][]byte{{0x95, 0x3c}, {0xa4, 0x34}, {0x81, 0x84, 0x34}, {0x7f}}
	for i := 0; i < len(values); i++ {
		encoded, err := values[i].Marshal()
		if err != nil {
			t.Error(err)
		} else {
			if equal(encoded, answers[i]) == false {
				t.Error("expected encoded value", encoded, "==", answers)
			}
		}
	}

	val := sdnv.Sdnv{-1}
	_, err := val.Marshal()
	if err == nil {
		t.Error("expected marshal error but received none")
	}
}

func TestUnmarshal(t *testing.T) {
	values := []sdnv.Sdnv{{0xabc}, {0x1234}, {0x4234}, {0x7f}}
	for i := 0; i < len(values); i++ {
		encoded, err := values[i].Marshal()
		if err != nil {
			t.Error(err)
		} else {
			ans := sdnv.Sdnv{0}
			err := ans.Unmarshal(encoded)
			if err != nil {
				t.Error(err)
			} else {
				if ans.Value != values[i].Value {
					t.Error("expected unmarshal value", ans.Value, "==", values[i].Value)
				}
			}
		}
	}

	val := []byte{0xff, 0xff, 0xff}
	ans := sdnv.Sdnv{0}
	err := ans.Unmarshal(val)
	if err == nil {
		t.Error("Expected error but got none")
	}
}
