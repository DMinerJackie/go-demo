package main

import (
	"github.com/gogo/protobuf/proto"
	"log"
	"reflect"
)

func main() {
	req := &SearchRequest{
		Query:         "keyword=zhangsan",
		PageNumber:    10,
		ResultPerPage: 10,
	}

	data, err := proto.Marshal(req)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}

	log.Printf("marshal result type: %s", reflect.TypeOf(data))

	newReq := &SearchRequest{}
	err = proto.Unmarshal(data, newReq)
	if err != nil {
		log.Fatal("unmarshaling error:", err)
	}

	log.Printf("newReq.Query:%s", newReq.Query)
	log.Printf("newReq.PageNumber:%d", newReq.PageNumber)
	log.Printf("newReq.ResultPerPage:%d", newReq.ResultPerPage)

}
