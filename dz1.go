package main

import (
	"fmt"
	"strings"
	"github.com/beevik/ntp"
)

func main() {
	if time, err := ntp.Time("pool.ntp.org"); err == nil {
		fmt.Println(time)
	}

	fmt.Println(unpack("a4bc2d5e"))
	fmt.Println(unpack("abcd"))
	fmt.Println(unpack("45"))
	fmt.Println(unpack(`qwe\4\5`))
	fmt.Println(unpack(`qwe\45`))
	fmt.Println(unpack(`qwe\\5`))
}

func unpack(input string) string {
	var (
		prefix = false
	    repeat, out string
	)
	for _, v := range input {
		if v >= '0' && v <= '9' && !prefix && repeat == "" {
			continue
		}
		if v == '\\' && !prefix {
			out = out + repeat
			repeat = ""
			prefix = true
			continue
		} else if v >= '0' && v <= '9' && !prefix && repeat != "" {
			out = out + strings.Repeat(repeat, int(v)-'0')
			repeat = ""
		} else {
			out = out + repeat
			repeat = string(v)
		}
		prefix = false

	}
	return out + repeat
}
