package main

import "fmt"

func hexToUnicode(hex []byte) []byte {
	unicode := make([]byte, len(hex)/2)
	for i := range unicode {
		v1 := hex[i*2]
		v2 := hex[i*2+1]
		unicode[i] = v1<<4 + v2
	}
	return unicode
}

func hexToString(hex []byte) string {
	unicode := hexToUnicode(hex)
	str := string(unicode)
	return str
}

func makeOfN(values []byte, length int) []byte {
	bytes := make([]byte, length)
	for i := 0; i < length; {
		for j := 0; j < len(values) && i < length; j++ {
			bytes[i] = values[j]
			i++
		}
	}
	return bytes
}

func scoreString(str string) (sum int) {
	for _, v := range str {
		char := byte(v)
		if char >= 97 && char <= 122 {
			sum += 2
		} else if char >= 65 && char <= 91 {
			sum += 1
		}
	}
	return
}

func findSingleByteXor(hex []byte) (string, int) {
	maxScore := -1
	var bestGuessString string
	for char := 0; char < 256; char++ {
		bytes := []byte{
			byte(char >> 4),
			byte(char % 16),
		}
		xorByteArray := makeOfN(bytes, len(hex))
		xorredHex := xorBytes(hex, xorByteArray)
		xorredString := hexToString(xorredHex)
		xorScore := scoreString(xorredString)
		if xorScore > maxScore {
			bestGuessString = xorredString
			maxScore = xorScore
		}
	}
	return bestGuessString, maxScore
}

func runSingleByteXorCipher() {
	hexBytes := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	hex := bytesToHex(hexBytes)
	str, _ := findSingleByteXor(hex)
	fmt.Println(str)
}
