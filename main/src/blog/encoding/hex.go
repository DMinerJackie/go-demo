package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

func main() {
	src := []byte("48656c6c6f20476f7068657221")

	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", dst[:n])

	// DecodeString
	const s = "48656c6c6f20476f7068657221"
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", decoded)

	// Encode
	src1 := []byte("Hello Gopher!")
	dst1 := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst1, src1)

	fmt.Printf("%s\n", dst1)

	data := []byte("jackie")
	hash := md5.New()
	bytes := hash.Sum(data)
	fmt.Printf("%x\n", bytes)

	// EncodeString
	src2 := []byte("Hello")
	encodedStr := hex.EncodeToString(src2)
	val, _ := strconv.ParseInt(encodedStr, 16, 64)
	fmt.Printf("%s %d\n", encodedStr, val)
}
