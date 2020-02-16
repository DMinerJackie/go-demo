package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {

	s1 := int64(time.Since(time.Unix(int64(1575268765), 0)).Hours()) / 2160
	fmt.Println(s1)
	reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	rgx := regexp.MustCompile(reg)
	s := []string{"18505921256", "13489594009", "12759029321"}
	for _, v := range s {
		fmt.Println(rgx.MatchString(v))
	}

	//virtualPhoneReg := "^1([7][0-9]|[71]|[62]|[65]|[67])"
	//rgx := regexp.MustCompile(virtualPhoneReg)
	//if userDetailRes.BindTel && rgx.MatchString(userDetailRes.HideTel) {
	//	hitVirtualPhone = true
	//}
}
