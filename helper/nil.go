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

type Message struct {
	message string
}

type MessageInterface interface {
	Get() string
	Print()
}

func New(message string) MessageInterface {
	return Message{message}
}

func (m Message) Get() string {
	return m.message
}

func (m Message) Print() {
	fmt.Println(m.message)
}

func Cetak(s string) (string, MessageInterface) {
	if s == "" {
		return "", New("ERROR: Cetak Kosong")
	}

	return s, nil
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

	msg, err := Cetak("CETAK")

	if err != nil {
		err.Print()
	} else {
		fmt.Println(msg)
	}
}
