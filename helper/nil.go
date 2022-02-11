package helper

import (
	blog "belajar-go/log"
	"fmt"
	"reflect"
)

func generateNewMap(v ...string) map[int]string {
	if len(v) == 0 {
		return nil
	}
	res := make(map[int]string)
	for i := 0; i < len(v); i++ {
		res[i] = v[i]
	}
	return res
}

func GoNil() {
	blog.Banner("Nil")

	// g := generateNewMap("hi", "earth", "from", "mars")
	g := generateNewMap()

	if g == nil {
		fmt.Println("Kosong")
	} else {
		fmt.Println(g)
	}

	gtype := reflect.TypeOf(g)
	fmt.Println("Return type:", gtype)
}
