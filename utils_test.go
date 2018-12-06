package md5

import (
	"testing"
)

func TestHex2bin(t *testing.T) {
	str := "123"
	if bin2hex(hex2bin(str)) == "0123" {
		t.Log("hex2bin 测试通过")
	} else {
		t.Error("hex2bin 测试失败")
	}
}
