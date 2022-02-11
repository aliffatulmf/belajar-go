package helper

import (
	blog "belajar-go/log"
	"fmt"
	"reflect"
)

func typeOf(v interface{}) interface{} {
	return reflect.TypeOf(v)
}

func EmtpyInterface() {
	blog.Banner("Empty Interface")

	fmt.Println(typeOf(1))
	fmt.Println(typeOf("string"))
	fmt.Println(typeOf(false))
	fmt.Println(typeOf(make([]string, 5)))
	fmt.Println(typeOf(map[string]string{"test": "test"}))
}
