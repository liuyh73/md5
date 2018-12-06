package md5

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

const (
	A = 0x67452301
	B = 0xefcdab89
	C = 0x98badcfe
	D = 0x10325476
)

var result = []int{A, B, C, D}

const (
	S11 = 7
	S12 = 12
	S13 = 17
	S14 = 22

	S21 = 5
	S22 = 9
	S23 = 14
	S24 = 20

	S31 = 4
	S32 = 11
	S33 = 16
	S34 = 23

	S41 = 6
	S42 = 10
	S43 = 15
	S44 = 21
)

func Encrypt(clear_text string) string {
	clear_text = hex.EncodeToString([]byte(clear_text))
	clear_text_bytes := stuff(clear_text)
	fmt.Println(bin2hex(clear_text_bytes))
	return md5(clear_text_bytes)
}

func md5(clear_text_bytes []byte) string {
	for i := 0; i < len(clear_text_bytes)/64; i++ {
		trans(clear_text_bytes[i*64 : (i+1)*64])
	}
	cipher_text := ""
	for i := 0; i < 4; i++ {
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, uint32(result[0]))
		cipher_text += bin2hex(buf)
	}
	return cipher_text
}

func div_group(text []byte) []int {
	groups := make([]int, 16)
	for i := 0; i < 15; i++ {
		groups[i] = int(uint(text[i*4]) | uint(text[i*4+1])<<8 | uint(text[i*4+2])<<16 | uint(text[i*4+3])<<24)
	}
	return groups
}

func trans(text []byte) {
	groups := div_group(text)
	a := result[0]
	b := result[1]
	c := result[2]
	d := result[3]
	/*第一轮*/
	a = FF(a, b, c, d, groups[0], S11, 0xd76aa478)  /* 1 */
	d = FF(d, a, b, c, groups[1], S12, 0xe8c7b756)  /* 2 */
	c = FF(c, d, a, b, groups[2], S13, 0x242070db)  /* 3 */
	b = FF(b, c, d, a, groups[3], S14, 0xc1bdceee)  /* 4 */
	a = FF(a, b, c, d, groups[4], S11, 0xf57c0faf)  /* 5 */
	d = FF(d, a, b, c, groups[5], S12, 0x4787c62a)  /* 6 */
	c = FF(c, d, a, b, groups[6], S13, 0xa8304613)  /* 7 */
	b = FF(b, c, d, a, groups[7], S14, 0xfd469501)  /* 8 */
	a = FF(a, b, c, d, groups[8], S11, 0x698098d8)  /* 9 */
	d = FF(d, a, b, c, groups[9], S12, 0x8b44f7af)  /* 10 */
	c = FF(c, d, a, b, groups[10], S13, 0xffff5bb1) /* 11 */
	b = FF(b, c, d, a, groups[11], S14, 0x895cd7be) /* 12 */
	a = FF(a, b, c, d, groups[12], S11, 0x6b901122) /* 13 */
	d = FF(d, a, b, c, groups[13], S12, 0xfd987193) /* 14 */
	c = FF(c, d, a, b, groups[14], S13, 0xa679438e) /* 15 */
	b = FF(b, c, d, a, groups[15], S14, 0x49b40821) /* 16 */

	/*第二轮*/
	a = GG(a, b, c, d, groups[1], S21, 0xf61e2562)  /* 17 */
	d = GG(d, a, b, c, groups[6], S22, 0xc040b340)  /* 18 */
	c = GG(c, d, a, b, groups[11], S23, 0x265e5a51) /* 19 */
	b = GG(b, c, d, a, groups[0], S24, 0xe9b6c7aa)  /* 20 */
	a = GG(a, b, c, d, groups[5], S21, 0xd62f105d)  /* 21 */
	d = GG(d, a, b, c, groups[10], S22, 0x2441453)  /* 22 */
	c = GG(c, d, a, b, groups[15], S23, 0xd8a1e681) /* 23 */
	b = GG(b, c, d, a, groups[4], S24, 0xe7d3fbc8)  /* 24 */
	a = GG(a, b, c, d, groups[9], S21, 0x21e1cde6)  /* 25 */
	d = GG(d, a, b, c, groups[14], S22, 0xc33707d6) /* 26 */
	c = GG(c, d, a, b, groups[3], S23, 0xf4d50d87)  /* 27 */
	b = GG(b, c, d, a, groups[8], S24, 0x455a14ed)  /* 28 */
	a = GG(a, b, c, d, groups[13], S21, 0xa9e3e905) /* 29 */
	d = GG(d, a, b, c, groups[2], S22, 0xfcefa3f8)  /* 30 */
	c = GG(c, d, a, b, groups[7], S23, 0x676f02d9)  /* 31 */
	b = GG(b, c, d, a, groups[12], S24, 0x8d2a4c8a) /* 32 */

	/*第三轮*/
	a = HH(a, b, c, d, groups[5], S31, 0xfffa3942)  /* 33 */
	d = HH(d, a, b, c, groups[8], S32, 0x8771f681)  /* 34 */
	c = HH(c, d, a, b, groups[11], S33, 0x6d9d6122) /* 35 */
	b = HH(b, c, d, a, groups[14], S34, 0xfde5380c) /* 36 */
	a = HH(a, b, c, d, groups[1], S31, 0xa4beea44)  /* 37 */
	d = HH(d, a, b, c, groups[4], S32, 0x4bdecfa9)  /* 38 */
	c = HH(c, d, a, b, groups[7], S33, 0xf6bb4b60)  /* 39 */
	b = HH(b, c, d, a, groups[10], S34, 0xbebfbc70) /* 40 */
	a = HH(a, b, c, d, groups[13], S31, 0x289b7ec6) /* 41 */
	d = HH(d, a, b, c, groups[0], S32, 0xeaa127fa)  /* 42 */
	c = HH(c, d, a, b, groups[3], S33, 0xd4ef3085)  /* 43 */
	b = HH(b, c, d, a, groups[6], S34, 0x4881d05)   /* 44 */
	a = HH(a, b, c, d, groups[9], S31, 0xd9d4d039)  /* 45 */
	d = HH(d, a, b, c, groups[12], S32, 0xe6db99e5) /* 46 */
	c = HH(c, d, a, b, groups[15], S33, 0x1fa27cf8) /* 47 */
	b = HH(b, c, d, a, groups[2], S34, 0xc4ac5665)  /* 48 */

	/*第四轮*/
	a = II(a, b, c, d, groups[0], S41, 0xf4292244)  /* 49 */
	d = II(d, a, b, c, groups[7], S42, 0x432aff97)  /* 50 */
	c = II(c, d, a, b, groups[14], S43, 0xab9423a7) /* 51 */
	b = II(b, c, d, a, groups[5], S44, 0xfc93a039)  /* 52 */
	a = II(a, b, c, d, groups[12], S41, 0x655b59c3) /* 53 */
	d = II(d, a, b, c, groups[3], S42, 0x8f0ccc92)  /* 54 */
	c = II(c, d, a, b, groups[10], S43, 0xffeff47d) /* 55 */
	b = II(b, c, d, a, groups[1], S44, 0x85845dd1)  /* 56 */
	a = II(a, b, c, d, groups[8], S41, 0x6fa87e4f)  /* 57 */
	d = II(d, a, b, c, groups[15], S42, 0xfe2ce6e0) /* 58 */
	c = II(c, d, a, b, groups[6], S43, 0xa3014314)  /* 59 */
	b = II(b, c, d, a, groups[13], S44, 0x4e0811a1) /* 60 */
	a = II(a, b, c, d, groups[4], S41, 0xf7537e82)  /* 61 */
	d = II(d, a, b, c, groups[11], S42, 0xbd3af235) /* 62 */
	c = II(c, d, a, b, groups[2], S43, 0x2ad7d2bb)  /* 63 */
	b = II(b, c, d, a, groups[9], S44, 0xeb86d391)  /* 64 */
	result[0] = (result[0] + a) & 0xFFFFFFFF
	result[1] = (result[1] + b) & 0xFFFFFFFF
	result[2] = (result[2] + c) & 0xFFFFFFFF
	result[3] = (result[3] + d) & 0xFFFFFFFF
}

func F(x, y, z int) int {
	return (x & y) | ((^x) & z)
}

func G(x, y, z int) int {
	return (x & z) | (y & (^z))
}

func H(x, y, z int) int {
	return x ^ y ^ z
}

func I(x, y, z int) int {
	return y ^ (x | (^z))
}

func FF(a, b, c, d, x, s, ac int) int {
	a += (F(b, c, d) & 0xFFFFFFFF) + x + ac
	a = ((a & 0xFFFFFFFF) << uint(s)) | int(uint(a&0xFFFFFFFF)>>uint(32-s))
	a += b
	return a & 0xFFFFFFFF
}

func GG(a, b, c, d, x, s, ac int) int {
	a += (G(b, c, d) & 0xFFFFFF) + x + ac
	a = ((a & 0xFFFFFFFF) << uint(s)) | int(uint(a&0xFFFFFFFF)>>uint(32-s))
	a += b
	return a & 0xFFFFFFFF
}

func HH(a, b, c, d, x, s, ac int) int {
	a += (H(b, c, d) & 0xFFFFFF) + x + ac
	a = ((a & 0xFFFFFFFF) << uint(s)) | int(uint(a&0xFFFFFFFF)>>uint(32-s))
	a += b
	return a & 0xFFFFFFFF
}
func II(a, b, c, d, x, s, ac int) int {
	a += (I(b, c, d) & 0xFFFFFF) + x + ac
	a = ((a & 0xFFFFFFFF) << uint(s)) | int(uint(a&0xFFFFFFFF)>>uint(32-s))
	a += b
	return a & 0xFFFFFFFF
}
