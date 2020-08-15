package main

import (
	"encoding/binary"

	"fmt"

	"os"
)

type stlHead struct {
	Name [14]byte

	FaceNum uint32
}

func main() {
	var a []byte = []byte{0, 1, 2, 3}
	fmt.Println(a)
	fmt.Println(binary.BigEndian.Uint32(a))
	fmt.Println(binary.LittleEndian.Uint32(a))

	file, err := os.Open("/Users/jackie/Library/Containers/com.tencent.WeWorkMac/Data/Library/Application Support/WXWork/Data/1688850815531185/Cache/File/2019-12/ipip_threat_full5_4.dat")

	if err != nil {

		fmt.Print(err)

		return

	}

	head := new(stlHead)

	if err := binary.Read(file, binary.LittleEndian, head); err != nil {

		fmt.Print(err)

		return

	}

	fmt.Printf("name: %s\r\n", head.Name)

	fmt.Println("FaceNum: ", head.FaceNum)

}
