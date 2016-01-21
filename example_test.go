package sdnv_test

import (
	"fmt"
	"github.com/bantl23/sdnv"
)

func Example_marshal() {
	s := sdnv.NewSdnv(10)
	data := s.Marshal()
	fmt.Println("data", data)
}

func Example_unmarshal() {
	s := sdnv.NewSdnv(0)
	data := []byte{0x95, 0x3c}
	err := s.Unmarshal(data)
	if err == nil {
		fmt.Println("sdnv", s)
	} else {
		fmt.Println("error", err)
	}
}
