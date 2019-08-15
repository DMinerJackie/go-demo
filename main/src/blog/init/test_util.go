/**
配合blog目录下的init用于验证init方法的执行时机
具体参见：https://zhuanlan.zhihu.com/p/34211611
*/
package init

import "fmt"

func init() {
	fmt.Println("init test_util")
}

func Hello() {
	fmt.Println("hello func")
}
