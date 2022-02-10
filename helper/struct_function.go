package helper

import (
	blog "belajar-go/log"
	"fmt"
)

type Meme struct {
	name, text string
}

func basic(m Meme) {
	fmt.Println(m.name, m.text)
}

func (m Meme) structFunction() {
	fmt.Println("Struct Function")
	fmt.Println(m.name, m.text)
}

func GoStructFunction() {
	blog.Banner("Struct Function")
	m := Meme{
		name: "Ariel",
		text: "kalian luar biasa",
	}

	basic(m)
	m.structFunction()
}
