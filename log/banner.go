package log

import (
	"fmt"
	"strings"
)

func Banner(name string) {
	var buff strings.Builder
	for i := 0; i <= 10; i++ {
		buff.WriteString("=")
	}

	fmt.Println(buff.String(), strings.ToTitle(name), buff.String())
}
