package md5

func stuff(clear_text string) []byte {
	length := len(clear_text)
	if len(clear_text)%128 != 112 {
		clear_text += "90"
	}
	for len(clear_text)%128 != 112 {
		clear_text += "00"
	}
	clear_text += int64(length)
	return clear_text
}
