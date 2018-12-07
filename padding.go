package md5

func padding(clear_text string) string {
	clear_text_bin := binText(clear_text)
	length := int64(len(clear_text_bin))
	if len(clear_text_bin)%512 != 448 {
		clear_text_bin += "1"
	}
	for len(clear_text_bin)%512 != 448 {
		clear_text_bin += "0"
	}
	lenStr := ""
	for i := 0; i < 64; i++ {
		lenStr = string((length&1)+'0') + lenStr
		length = length >> 1
		if (i+1)%8 == 0 {
			clear_text_bin += lenStr
			lenStr = ""
		}
	}
	return clear_text_bin
}
