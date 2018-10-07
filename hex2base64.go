package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const hex0 = byte('0')
const hex9 = byte('9')
const hexa = byte('a')

func bytesToHex(bytes []byte) []byte {
	hexBytes := make([]byte, len(bytes))
	for i, v := range bytes {
		if v <= hex9 {
			hexBytes[i] = v - hex0
		} else {
			hexBytes[i] = v - hexa + 10
		}
	}
	return hexBytes
}

func getDivSize(size, length int) int {
	return int(math.Ceil(float64(size) / float64(length)))
}

func bitToPosition(bit bool, pos uint8) byte {
	if !bit {
		return 0
	}
	return byte(1 << pos)
}

func getBit(bits []bool, n int) bool {
	if n < len(bits) {
		return bits[n]
	}
	return false
}

func hexToBits(hex []byte) []bool {
	bits := make([]bool, len(hex)*4)
	for i, v := range hex {
		bits[i*4] = v&8 > 0
		bits[i*4+1] = v&4 > 0
		bits[i*4+2] = v&2 > 0
		bits[i*4+3] = v&1 > 0
	}
	return bits
}

func bitsToBase64(bits []bool) []byte {
	limit := getDivSize(len(bits), 6)
	base64 := make([]byte, limit)
	for i := 0; i < limit; i++ {
		sum := byte(0)
		for j := 0; j < 6; j++ {
			temp := bitToPosition(getBit(bits, i*6+j), uint8(5-j))
			sum += temp
		}
		base64[i] = sum
	}
	return base64
}

func printBase64(base64 []byte) {
	chars := make([]string, len(base64))
	for i, v := range base64 {
		if v <= 25 {
			chars[i] = string(65 + v)
		} else if v <= 51 {
			chars[i] = string(97 + v - 26)
		} else if v <= 61 {
			chars[i] = strconv.Itoa(int(v - 52))
		} else if v == 62 {
			chars[i] = "+"
		} else {
			chars[i] = "/"
		}
	}
	fmt.Println(strings.Join(chars, ""))
}

func run() {
	orig := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	bytes := []byte(orig)
	hex := bytesToHex(bytes)
	bits := hexToBits(hex)
	base64 := bitsToBase64(bits)
	printBase64(base64)
}

func main() {
	run()
}
