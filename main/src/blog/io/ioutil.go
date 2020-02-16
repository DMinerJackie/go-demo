package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	IOUtil("/Users/jackie/Downloads/ipip_threat_full_2019-11-30.txt.gz")
}

func IOUtil(file string) {
	f, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	bytes, _ := ioutil.ReadAll(f)
	fmt.Println(string(bytes))
	if content, err := ioutil.ReadFile(file); err == nil {
		result := strings.Replace(string(content), "\n", "\n", 1)
		fmt.Println(result)
	}
}
