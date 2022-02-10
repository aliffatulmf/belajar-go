package helper

import (
	blog "belajar-go/log"
	"fmt"
)

type Struct struct {
	Name, color  string
	Moon, Number int
}

func (s Struct) planetName() string {
	return s.Name
}

func earthColor(s Struct) {
	fmt.Printf("%s is %s\n", s.Name, s.color)
}

func GoStruct() {
	blog.Banner("Struct")

	fmt.Println("Type 1")
	s1 := Struct{
		Name:   "Earth",
		color:  "blue",
		Moon:   1,
		Number: 3,
	}

	fmt.Println("Name:", s1.Name)
	fmt.Println("Moon:", s1.Moon)
	fmt.Println("Number:", s1.Number)

	fmt.Println("Type 2")
	var s2 Struct
	s2 = Struct{
		Name:   "Mars",
		color:  "red",
		Moon:   2,
		Number: 4,
	}

	fmt.Println("Name:", s2.Name)
	fmt.Println("Moon:", s2.Moon)
	fmt.Println("Number:", s2.Number)

	fmt.Println("Type 3")
	fmt.Println(s2.planetName())

	fmt.Println("Type 4")
	earthColor(s1)
	earthColor(s2)
}

type mataramStruct struct {
	is string
}

func mataramScopeEx(mataram mataramStruct) {
	mataram.is = "blue"
	fmt.Println("Mataram is", mataram.is)
}

func StructScope() {
	blog.Banner("Struct Scope")
	mataram := mataramStruct{
		is: "red",
	}

	fmt.Println("Mataram is", mataram.is)
	mataramScopeEx(mataram)
	fmt.Println("Mataram is", mataram.is)

	fmt.Println("-----------")
	mataramScopeIn := func() {
		mataram.is = "blue" // overwrite mataram is
	}

	mataramScopeIn()
	fmt.Println("Mataram is", mataram.is)

	mataramScopeEx(mataram)
	fmt.Println("HOOH")
}
