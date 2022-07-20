package collections

import (
	"fmt"
	"testing"

	"github.com/NethermindEth/juno/pkg/common"
)

func TestBS(t *testing.T) {
	tests := []struct {
		hex      string
		expected []bool
	}{
		{"0x11", []bool{true, false, false, false, true}},
		{"0x11", []bool{false, false, true, false, false, false, true}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestBSCreate%d", i+1), func(t *testing.T) {
			bs := NewBitSet(len(test.expected), common.FromHex(test.hex))
			for i := 0; i < bs.Len(); i++ {
				if bs.Get(i) != test.expected[i] {
					t.Errorf("bs.Get(%d) == %v, want %v", i, bs.Get(i), test.expected[i])
				}
			}

			bs.Set(0)
			if !bs.Get(0) {
				t.Errorf("bs.Get(0) == %v, want %v", bs.Get(0), true)
			}

			bs.Clear(0)
			if bs.Get(0) {
				t.Errorf("bs.Get(0) == %v, want %v", bs.Get(0), false)
			}

			bs.Clear(0)
			if bs.Get(0) {
				t.Errorf("bs.Get(0) == %v, want %v", bs.Get(0), false)
			}
		})
	}
}