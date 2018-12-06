package md5

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)
// 出事话变量,仅用于第一轮
const (
	A uint = 0x67452301
	B uint = 0xefcdab89
	C uint = 0x98badcfe
	D uint = 0x10325476
)

var k = []uint{
	0xd76aa478,0xe8c7b756,0x242070db,0xc1bdceee,
    0xf57c0faf,0x4787c62a,0xa8304613,0xfd469501,0x698098d8,
    0x8b44f7af,0xffff5bb1,0x895cd7be,0x6b901122,0xfd987193,
    0xa679438e,0x49b40821,0xf61e2562,0xc040b340,0x265e5a51,
    0xe9b6c7aa,0xd62f105d,0x02441453,0xd8a1e681,0xe7d3fbc8,
    0x21e1cde6,0xc33707d6,0xf4d50d87,0x455a14ed,0xa9e3e905,
    0xfcefa3f8,0x676f02d9,0x8d2a4c8a,0xfffa3942,0x8771f681,
    0x6d9d6122,0xfde5380c,0xa4beea44,0x4bdecfa9,0xf6bb4b60,
    0xbebfbc70,0x289b7ec6,0xeaa127fa,0xd4ef3085,0x04881d05,
    0xd9d4d039,0xe6db99e5,0x1fa27cf8,0xc4ac5665,0xf4292244,
    0x432aff97,0xab9423a7,0xfc93a039,0x655b59c3,0x8f0ccc92,
    0xffeff47d,0x85845dd1,0x6fa87e4f,0xfe2ce6e0,0xa3014314,
    0x4e0811a1,0xf7537e82,0xbd3af235,0x2ad7d2bb,0xeb86d391,
}

var result = []uint{A, B, C, D}
// 向左位移数
var s = []uint{ 
	7,12,17,22,7,12,17,22,7,12,17,22,7,
	12,17,22,5,9,14,20,5,9,14,20,5,9,14,20,5,9,14,20,
	4,11,16,23,4,11,16,23,4,11,16,23,4,11,16,23,6,10,
	15,21,6,10,15,21,6,10,15,21,6,10,15,21,
}

func Encrypt(clear_text string) string {
	clear_text = hex.EncodeToString([]byte(clear_text))
	clear_text_bin := stuff(clear_text)
	return digest(clear_text_bin)
}

func digest(clear_text_bin string) string {
	for i := 0; i < len(clear_text_bin)/512; i++ {
		trans(clear_text_bin[i*512 : (i+1)*512])
	}
	cipher_text := ""
	for i := 0; i < 4; i++ {
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, uint32(result[i]))
		cipher_text += fmt.Sprintf("%x", buf)
	}
	return cipher_text
}

func div_group(text string) []uint {
	groups := make([]uint, 16)
	for i := 0; i < 15; i++ {
		groups[i] = bin2dec(text[i*32:(i+1)*32])
		fmt.Println("group[", i, "]: ", groups[i])
	}
	return groups
}

func trans(text string) {
	groups := div_group(text)
	var f, g uint
	a := result[0]
	b := result[1]
	c := result[2]
	d := result[3]
	for i:=0; i<64; i++ {
		if i<16 {
			f = F(b,c,d)
			g = i
		} else if i<32 {
			f = G(b,c,d)
			g = (5*i+1)%16
		} else if i<48 {
			f = H(b,c,d);
			g = (3*i+5) %16
		} else {
			f = I(b,c,d)
			g=(7*i)%16
		}
		dtemp := d
		d = c 
		c = b
		b = b + shift(a+f+k[i]+groups[g], s[i])
		a = dtemp
	}
	result[0] = result[0] + a
	result[1] = result[1] + b
	result[2] = result[2] + c
	result[3] = result[3] + d
}

func F(x, y, z uint) uint {
	return (x & y) | ((^x) & z)
}

func G(x, y, z uint) uint {
	return (x & z) | (y & (^z))
}

func H(x, y, z uint) uint {
	return x ^ y ^ z
}

func I(x, y, z uint) uint {
	return y ^ (x | (^z))
}

func shift(x, n uint) uint {
	return uint(x) << uint(n) | uint(x) >> uint((32) - n)
}