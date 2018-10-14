package main

func stringToHex(str string) []byte {
	hex := make([]byte, len(str)*2)
	for i, v := range str {
		char := byte(v)
		hex[i*2] = char >> 4
		hex[i*2+1] = char % 16
	}
	return hex
}

func encodeHex(hex, key []byte) []byte {
	xorKey := makeOfN(key, len(hex))
	xorred := xorBytes(hex, xorKey)
	return xorred
}

func encodeRepeatingKey(str, keyStr string) []byte {
	hex := stringToHex(str)
	key := stringToHex(keyStr)
	return encodeHex(hex, key)
}

func runImplementRepeatingKeyXor() {
	printHex(encodeRepeatingKey("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal", "ICE"))
}
