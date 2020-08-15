package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	d := int8(32)
	e := int8(35)
	fmt.Printf("%c, %c", d, e)

	fmt.Println("=========Contains=========")
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("foo")))
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("bar")))
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("")))
	fmt.Println(bytes.Contains([]byte(""), []byte("")))

	fmt.Println("=========Count=========")
	fmt.Println(bytes.Count([]byte("cheese"), []byte("e")))
	fmt.Println(bytes.Count([]byte("five"), []byte(""))) // before & after each rune

	fmt.Println("=========Equal=========")
	fmt.Println(bytes.Equal([]byte("Go"), []byte("go")))
	fmt.Println(bytes.EqualFold([]byte("Go"), []byte("go")))

	fmt.Println("=========Fields=========")
	fmt.Printf("Fields are: %q\n", bytes.Fields([]byte("  foo bar  baz   ")))

	fmt.Println("=========FieldsFunc=========")
	f := func(c rune) bool { return !unicode.IsLetter(c) && !unicode.IsNumber(c) }
	fmt.Printf("Fields are: %q\n", bytes.FieldsFunc([]byte("  foo1;bar2,baz3..."), f))

	fmt.Println("=========HasPrefix=========")
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("Go")))
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("C")))
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("")))

	fmt.Println("=========SuffixPrefix=========")
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("go")))
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("O")))
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("Ami")))
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("")))

	fmt.Println("=========Index=========")
	fmt.Println(bytes.Index([]byte("chicken"), []byte("ken")))
	fmt.Println(bytes.Index([]byte("chicken"), []byte("dmr")))

	fmt.Println("=========IndexAny=========")
	fmt.Println(bytes.IndexAny([]byte("chicken"), "aeiouy"))
	fmt.Println(bytes.IndexAny([]byte("crwth"), "aeiouy"))

	fmt.Println("=========IndexRune=========")
	fmt.Println(bytes.IndexRune([]byte("chicken"), 'k'))
	fmt.Println(bytes.IndexRune([]byte("chicken"), 'd'))

	fmt.Println("=========Joins=========")
	s := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	fmt.Printf("%s", bytes.Join(s, []byte(", ")))

	fmt.Println("=========LastIndex=========")
	fmt.Println(bytes.Index([]byte("go gopher"), []byte("go")))
	fmt.Println(bytes.LastIndex([]byte("go gopher"), []byte("go")))
	fmt.Println(bytes.LastIndex([]byte("go gopher"), []byte("rodent")))

	fmt.Println("=========Map=========")
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Printf("%s", bytes.Map(rot13, []byte("'Twas brillig and the slithy gopher...")))

	fmt.Println("=========Repeat=========")
	fmt.Printf("ba%s", bytes.Repeat([]byte("na"), 2))

	fmt.Println("=========Replace=========")
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("k"), []byte("ky"), 2))
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("oink"), []byte("moo"), -1))

	fmt.Println("=========Split=========")
	fmt.Printf("%q\n", bytes.Split([]byte("a,b,c"), []byte(",")))
	fmt.Printf("%q\n", bytes.Split([]byte("a man a plan a canal panama"), []byte("a ")))
	fmt.Printf("%q\n", bytes.Split([]byte(" xyz "), []byte("")))
	fmt.Printf("%q\n", bytes.Split([]byte(""), []byte("Bernardo O'Higgins")))

	fmt.Println("=========Title=========")
	fmt.Printf("%s\n", bytes.Title([]byte("her royal highness")))

	fmt.Println("=========ToTitle=========")
	fmt.Printf("%s\n", bytes.ToTitle([]byte("loud noises")))
	fmt.Printf("%s\n", bytes.ToTitle([]byte("хлеб")))

	fmt.Println("=========Buffer=========")
	var b bytes.Buffer // 缓冲区不需要初始化。
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "world!")
	b.WriteTo(os.Stdout)

	fmt.Println("=========ToTitle=========")
	buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	fmt.Println()
	dec := base64.NewDecoder(base64.StdEncoding, buf)
	io.Copy(os.Stdout, dec)
}
