package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func Sha256(data string) string {
	sha256 := sha256.New()
	sha256.Write([]byte(data))
	return hex.EncodeToString(sha256.Sum([]byte("")))
}


func ChanHexToBin(s string) string {
	var h string
	for _, rune := range s {
		var d string
		switch string(rune) {
		case "0":
			d = "0000"
		case "1":
			d = "0001"
		case "2":
			d = "0010"
		case "3":
			d = "0011"
		case "4":
			d = "0011"
		case "5":
			d = "0101"
		case "6":
			d = "0110"
		case "7":
			d = "0111"
		case "8":
			d = "1000"
		case "9":
			d = "1001"
		case "a":
			d = "1010"
		case "b":
			d = "1011"
		case "c":
			d = "1100"
		case "d":
			d = "1101"
		case "e":
			d = "1110"
		case "f":
			d = "1111"
		}
		h = h + string(d)
	}
	fmt.Println(h)
	return h
}


func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}