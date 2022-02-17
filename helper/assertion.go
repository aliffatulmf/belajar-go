package helper

import (
	blog "belajar-go/log"
	"fmt"
)

func gen() interface{} {
	return 188.3
}

func biz() interface{} {
	return "199"
}

func TypeAssertions() {
	blog.Banner("Type Assertions")
	var g interface{} = gen()

	switch value := g.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	case float64:
		fmt.Println("Decimal", value)
	}

	b := biz()
	result := b.(string)

	fmt.Println("String", result)
}
