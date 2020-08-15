package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 打开一个zip文件供阅读。
	r, err := zip.OpenReader("testdata/readme.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close() // 遍历存档中的文件，// 打印他们的一些内容。
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
}
