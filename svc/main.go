package main

import (
	"fmt"

	"github.com/liuyh73/md5"
)

func main() {
	fmt.Println(md5.Encrypt("The quick brown fox jumps over the lazy dog")) //9e107d9d372bb6826bd81d3542a419d6
	fmt.Println(md5.Encrypt("The quick brown fox jumps over the lazy cog")) // 1055d3e698d289f2af8663725127bd4b
	fmt.Println(md5.Encrypt(""))                                            // d41d8cd98f00b204e9800998ecf8427e
}
