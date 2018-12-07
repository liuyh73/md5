package md5

import (
	"math/rand"
	"testing"
)

func TestDec2hex(t *testing.T) {
	for i := 1; i <= 10; i++ {
		hexStr := ""
		dec := rand.Uint32()
		deci := dec
		for dec != 0 {
			str := ""
			if dec%16 < 10 {
				str = string((dec % 16) + '0')
			} else {
				str = string((dec%16 - 10) + 'a')
			}
			dec >>= 4
			if dec%16 < 10 {
				str = string((dec%16)+'0') + str
			} else {
				str = string((dec%16-10)+'a') + str
			}
			dec >>= 4
			hexStr = hexStr + str
		}
		if hexStr == dec2hex(uint32(deci)) {
			t.Logf("The %d test passes.\n", i)
		} else {
			t.Errorf("The %d test fails.\n", i)
		}
	}
}
