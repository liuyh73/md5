package md5

import "fmt"

func hexadecimalToBinary(data byte) string {
	switch data {
	case '0':
		return "0000"
	case '1':
		return "0001"
	case '2':
		return "0010"
	case '3':
		return "0011"
	case '4':
		return "0100"
	case '5':
		return "0101"
	case '6':
		return "0110"
	case '7':
		return "0111"
	case '8':
		return "1000"
	case '9':
		return "1001"
	case 'a':
		return "1010"
	case 'b':
		return "1011"
	case 'c':
		return "1100"
	case 'd':
		return "1101"
	case 'e':
		return "1110"
	case 'f':
		return "1111"
	}
	return ""
}

func binText(text string) string {
	binText := ""
	for i := 0; i < len(text); i++ {
		binText += hexadecimalToBinary(text[i])
	}
	return binText
}

// 小端，低位在低字节
func bin2dec(binstr string) uint32 {
	num := make([]uint32, 4)
	for i := 0; i < 32; i++ {
		num[i/8] = num[i/8] << 1
		num[i/8] += uint32(binstr[i] - '0')
	}
	n := num[3]
	for i := 2; i >= 0; i-- {
		n = n << 8
		n += num[i]
	}
	return n
}

// 小端，低位在低字节
func dec2hex(num uint32) string {
	buf := make([]byte, 4)
	for i := 0; i < 4; i++ {
		buf[i] = byte(num >> uint(8*i))
	}
	return fmt.Sprintf("%x", buf)
}

// func hex2bin(str string) []byte {
// 	if len(str)%2 == 1 {
// 		str = string('0') + str
// 	}
// 	strLen := len(str)
// 	binBytes := make([]byte, 0)
// 	for i := 0; i < strLen; i += 2 {
// 		bt, _ := strconv.ParseInt(str[i:i+2], 16, 32)
// 		fmt.Println(bt)
// 		binBytes = append(binBytes, byte(bt))
// 	}
// 	return binBytes
// }

// func bin2hex(bts []byte) string {
// 	return fmt.Sprintf("%x", bts)
// }
