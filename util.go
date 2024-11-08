package rollmelette

import "encoding/hex"

func Bin2Hex(bin []byte) string {
	hx := hex.EncodeToString(bin)
	return "0x" + string(hx)
}

func PadBytes(bin []byte, size int) []byte {
	tmp := make([]byte, size)
	l := len(bin)
	copy(tmp[(size-l):], bin)
	return tmp
}
