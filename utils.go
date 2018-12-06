package md5

import (
	"fmt"
	"strconv"
)

func hex2bin(str string) []byte {
	if len(str)%2 == 1 {
		str = string('0') + str
	}
	strLen := len(str)
	binBytes := make([]byte, 0)
	for i := 0; i < strLen; i += 2 {
		bt, _ := strconv.ParseInt(str[i:i+2], 16, 32)
		fmt.Println(bt)
		binBytes = append(binBytes, byte(bt))
	}
	return binBytes
}

func bin2hex(bts []byte) string {
	return fmt.Sprintf("%x", bts)
}
