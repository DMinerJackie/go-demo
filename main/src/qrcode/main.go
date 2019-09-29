package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"image/color"
	"log"
	"net/http"
	"time"
)

func Qrcode(w http.ResponseWriter, req *http.Request) {
	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(500)
			return
		}
	}()
	q, err := qrcode.New(fmt.Sprintf("https://www.cnblogs.com/bigdatazj?t=%d", time.Now().Unix()), qrcode.Medium)
	if err != nil {
		return
	}
	q.BackgroundColor = color.RGBA64{200, 250, 190, 255}
	//color.RGBA{200,90,190,100}
	q.ForegroundColor = color.RGBA{11, 230, 204, 255}
	png, err := q.PNG(128)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(png)))
	w.Write(png)
}

func QR() ([]byte, error) {
	return qrcode.Encode("https://www.baidu.com", qrcode.Medium, 256)
}

func main() {
	http.HandleFunc("/qrcode", Qrcode)
	log.Fatal(http.ListenAndServe(":8008", nil))
}
