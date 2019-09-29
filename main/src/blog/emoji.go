package main

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	//fmt.Println(math.Log2(8))
	ch1 := make(chan string)
	go func() {
		time.Sleep(time.Second)
		ch1 <- "hello"
	}()

	select {
	case <-ch1:
		fmt.Println("hello cha1")

	default:
		fmt.Println("end")
	}

	//time.Sleep(time.Second * 2)

	score := 2*math.Log(4) +
		float64(3) +
		float64(4) +
		float64(4) + rand.Float64()
	fmt.Println(score)
	fmt.Println(math.Log(4))

	//c := make(chan string)
	////c <- "joe"
	//timeout := time.After(5 * time.Second)
	//go func() {
	//	for {
	//		select {
	//		case s := <-c:
	//			fmt.Println(s)
	//		case <-timeout:
	//			fmt.Println("You talk too much.")
	//			return
	//		}
	//	}
	//}()

	//ch := make(chan string)
	//
	//go func() {
	//	select {
	//	case <-ch:
	//		fmt.Println("done")
	//	default:
	//		fmt.Println("default")
	//	}
	//}()
	//
	//time.Sleep(time.Second * 3)
	//
	//close(ch)
	//time.Sleep(time.Second * 2)

	////var c coder = &Gopher{"Go"}
	////c.code()
	////c.debug()
	//
	//map1 := map[string]int32{"阿斯顿发jjj":11, "hello": 12}
	////delete(map1, "hello")
	//fmt.Println(map1)
	//var sqlString string
	//for key, value := range map1 {
	//	//sqlString = sqlString + key + "=" + strconv.Itoa(int(value)) + ", "
	//	sqlString = sqlString + fmt.Sprintf("%s=%d, ", key, value)
	//}
	//
	//lastIndex := strings.LastIndex(sqlString, ",")
	//prefix := []byte(sqlString)[0:lastIndex]
	//fmt.Println(string(prefix))
	//
	//
	//hello := "hello"
	//world := "world"
	//for i := 0; i < 3; i++ {
	//	fmt.Println(fmt.Sprintf("%s=%s, ", hello, world))
	//}
	//
	//var m1 map[int64]int64
	//m2 := map[int64]int64{1:1,2:2}
	//m1 = m2
	//fmt.Println(m1)

	//dbgInfs := []DebugInfo{
	//	DebugInfo{"debug", `File: "test.txt" Not Found`, "Cynhard"},
	//	DebugInfo{"", "Logic error", "Gopher"},
	//}
	//
	//if data, err := json.Marshal(dbgInfs); err == nil {
	//	fmt.Printf("%s\n", data)
	//}

	//data := `[{"level":"debug","message":"File Not Found","author":"Cynhard"},` +
	//	`{"level":""}]`
	//
	//var dbgInfos []DebugInfo
	//json.Unmarshal([]byte(data), &dbgInfos)
	//
	//fmt.Println(dbgInfos)

	fmt.Println(FilterEmoji("1😉aa"))
	fmt.Println(FilterEmoji("hh"))

	UnicodeEmojiDecode(UnicodeEmojiCode("哈哈哈🍻"))
	UnicodeEmojiDecode(UnicodeEmojiCode("✌🤩🌸🌸🌺🌺"))

	stopCh := make(chan int)
	go func() {
		time.Sleep(time.Second * 10)
		stopCh <- 1
	}()

	for {
		select {
		case <-stopCh:
			fmt.Println("stop")
			return
		default:

		}
	}

	time.Sleep(time.Second * 2)

}

type DebugInfo struct {
	Level  string `json:"level,omitempty"` // level 解码为 Level
	Msg    string `json:"message"`         // message 解码为 Msg
	Author string `json:"-"`               // 忽略Author
}

func (dbgInfo DebugInfo) String() string {
	return fmt.Sprintf("{Level: %s, Msg: %s}", dbgInfo.Level, dbgInfo.Msg)
}

func FilterEmoji(content string) string {
	new_content := ""
	for _, value := range content {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			new_content += string(value)
		}
	}
	return new_content
}

//表情转换
func UnicodeEmojiCode(s string) string {
	ret := ""
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if len(string(rs[i])) == 4 {
			u := `[\u` + strconv.FormatInt(int64(rs[i]), 16) + `]`
			ret += u
		} else {
			ret += string(rs[i])
		}
	}
	return ret
}

//表情解码
func UnicodeEmojiDecode(s string) string {
	//emoji表情的数据表达式
	re := regexp.MustCompile("\\[[\\\\u0-9a-zA-Z]+\\]")
	//提取emoji数据表达式
	reg := regexp.MustCompile("\\[\\\\u|]")
	src := re.FindAllString(s, -1)
	for i := 0; i < len(src); i++ {
		e := reg.ReplaceAllString(src[i], "")
		p, err := strconv.ParseInt(e, 16, 32)
		if err == nil {
			s = strings.Replace(s, src[i], string(rune(p)), -1)
		}
	}
	return s
}
