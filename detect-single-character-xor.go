package main

import (
	"bufio"
	"fmt"
	"os"
)

func initializeMaxString() func(string) string {
	maxScore := -1
	var maxString string
	return func(str string) string {
		bytes := []byte(str)
		hex := bytesToHex(bytes)
		decodedStr, score := findSingleByteXor(hex)
		if score > maxScore {
			maxScore = score
			maxString = decodedStr
		}
		return maxString
	}
}

func runSingleCharXor() {
	getMaxString := initializeMaxString()
	file, err := os.Open("./4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var maxString string
	for scanner.Scan() {
		maxString = getMaxString(scanner.Text())
	}
	fmt.Println(maxString)
}
