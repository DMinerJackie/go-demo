package main

import (
	"context"
	"fmt"
)

type MyContext struct {
	context.Context
	KeyValue map[string]bool
}

type FilterFunc func(*MyContext)

type FilterFuncChain []FilterFunc

type CombinedFunc struct {
	CF    FilterFuncChain
	MyCtx *MyContext
}

//type FilterFunc struct {
//}

func main() {
	//ctx := context.TODO()
	//newCtx := context.WithValue(ctx, "key", "value")
	myContext := MyContext{Context: context.TODO(), KeyValue: map[string]bool{"key": false}}

	cf := CombinedFilter(&myContext, F1, F2, F3)
	DoFilter(cf)
}

func DoFilter(cf *CombinedFunc) {
	for _, f := range cf.CF {
		f(cf.MyCtx)
		continued := cf.MyCtx.KeyValue["key"]
		fmt.Println("result:", continued)
		if !continued {
			fmt.Println("stopped")
			return
		}
	}
}

func CombinedFilter(ctx *MyContext, ff ...FilterFunc) *CombinedFunc {
	return &CombinedFunc{
		CF:    ff,
		MyCtx: ctx,
	}
}

func F1(ctx *MyContext) {
	ctx.KeyValue["key"] = true
	fmt.Println(ctx.KeyValue["key"])
	//return ctx.KeyValue["key"]
}

func F2(ctx *MyContext) {
	ctx.KeyValue["key"] = false
	fmt.Println(ctx.KeyValue["key"])
	//return ctx.KeyValue["key"]
}

func F3(ctx *MyContext) {
	ctx.KeyValue["key"] = false
	fmt.Println(ctx.KeyValue["key"])
	//return ctx.KeyValue["key"]
}
