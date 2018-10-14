package main

import (
	"fmt"
	"strconv"
	"strings"
)

func xorBytes(orig1, orig2 []byte) []byte {
	bytes1 := bytesToHex(orig1)
	bytes2 := bytesToHex(orig2)
	bytes3 := make([]byte, len(bytes1))
	for i, byte1 := range bytes1 {
		byte2 := bytes2[i]
		bytes3[i] = byte1 ^ byte2
	}
	return bytes3
}

func printHex(hex []byte) {
	chars := make([]string, len(hex))
	for i, v := range hex {
		if v <= 9 {
			chars[i] = strconv.Itoa(int(v))
		} else {
			chars[i] = string(97 + v - 10)
		}
	}
	fmt.Println(strings.Join(chars, ""))
}

func runHexXor() {
	orig1 := []byte("1c0111001f010100061a024b53535009181c")
	orig2 := []byte("686974207468652062756c6c277320657965")
	xordvalue := xorBytes(orig1, orig2)
	printHex(xordvalue)
}
