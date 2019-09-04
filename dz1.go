package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"github.com/vjeantet/jodaTime"
	"strings"
	"time"
)

func main() {
	fmt.Println("First exercise:")
	if response, err := ntp.Query("pool.ntp.org"); err == nil {
		now := time.Now()
		fmt.Printf("Current time: %s\n", jodaTime.Format("dd.MM.YYYY HH:mm:ss:SSS", now))
		fmt.Printf("Corrected time: %s\n", jodaTime.Format("dd.MM.YYYY HH:mm:ss:SSS", now.Add(response.ClockOffset)))
	} else {
		fmt.Println("Failure")
	}

	fmt.Println("\n\nSecond exercise:")
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
