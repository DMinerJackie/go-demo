package cocurrency

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestBigFile(t *testing.T) {
	//ReadFile("/Users/jackie/Downloads/ipip_threat_full_2019-11-30.txt", print)

	//ReadBigFile("/Users/jackie/Downloads/ipip_threat_full_2019-11-30.txt")
	ReadBigFile("/Users/jackie/Library/Containers/com.tencent.WeWorkMac/Data/Library/Application Support/WXWork/Data/1688850815531185/Cache/File/2019-12/ipip_threat_full5_4.dat")
}

func ReadFile(filePath string, fu func(b []byte)) (err error) {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	var line []byte
	for {
		line, _, err = buf.ReadLine()
		//lineStr := strings.TrimSpace(string(line))
		fu(line)
		//fmt.Println(lineStr)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		//return nil
	}
}

type in interface {
	PrintContent([]byte)
}

type handle func([]byte)

func (h handle) PrintContent(b []byte) {
	fmt.Println(string(b))
}

func print(b []byte) {
	fmt.Println(string(b))
}

func ReadBigFile(fileName string) error {
	f, err := os.Open(fileName)

	//b, err := ioutil.ReadAll(f)
	//fmt.Println(binary.LittleEndian.Uint32(b[10:14]))
	//var ipBlackList []uint32
	//startIndex, offset := uint32(14), uint32(4)
	//ipNum := binary.LittleEndian.Uint32(b[10:14])
	//for i := uint32(1); i <= ipNum; i++ {
	//	endIndex := startIndex + offset
	//	ip := binary.LittleEndian.Uint32(b[startIndex:endIndex])
	//	ipBlackList = append(ipBlackList, ip)
	//	startIndex = endIndex
	//}
	if err != nil {
		fmt.Println("can't opened this file")
		return err
	}
	defer f.Close()
	s := make([]byte, 4096)
	for {
		switch nr, err := f.Read(s[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // EOF
			return nil
		case nr > 0:
			fmt.Println(string(s[0:10]))
			//fmt.Println(string(s[10:11]))
			//var a []byte = []byte{0, 1, 2, 3}
			fmt.Println(len(s))
			fmt.Println(binary.LittleEndian.Uint32(s[10:14]))
			fmt.Println(binary.LittleEndian.Uint32(s[14:17]))
			fmt.Println(binary.LittleEndian.Uint32(s[18:22]))

			fmt.Println(StringIpToInt("1.0.134.40"))

			x := int32(16811560)
			//x := int32(16777394)

			bytesBuffer := bytes.NewBuffer([]byte{})
			binary.Write(bytesBuffer, binary.LittleEndian, x)
			//fmt.Println(strconv.ParseUint(string(s[11:12]), 16, 32))
			fmt.Println(string(s[10:14]))
			//fmt.Println(string(s[11:18]))
		}
	}
	return nil
}

func StringIpToInt(ipstring string) int {
	ipSegs := strings.Split(ipstring, ".")
	var ipInt int = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}
