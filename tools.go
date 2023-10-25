package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

func StrToBinary(s string) (res string) {
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res
}

func ToString(s string) (res string) {
	arr := make([]string, 8)
	j := 0
	for i, c := range s {
		if i%8 == 0 && i != 0 {
			j++
		}
		arr[j] += string(c)
	}

	for _, a := range arr {
		tmp, err := strconv.ParseUint(a, 2, 0)
		if err != nil {
			log.Fatal(err)
		}
		res += string(tmp)
	}

	return res
}

func Hex(s string) string {
	return hex.EncodeToString([]byte(s))
}

func HexToBinary(s string) string {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	return StrToBinary(string(decoded))
}

func Base64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64ToBinary(s string) string {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	return StrToBinary(string(decoded))
}
