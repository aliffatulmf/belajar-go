package main

import (
	"fmt"
	"math"
	"strings"
)

func Banner(name string) {
	var buff strings.Builder
	for i := 0; i <= 10; i++ {
		buff.WriteString("=")
	}

	fmt.Println(buff.String(), strings.ToTitle(name), buff.String())
}

func variable() {
	Banner("variable")
	var number int8
	var text string
	var char byte

	number = 10
	text = "Golang"
	char = 'G'

	fmt.Println("Number", number)
	fmt.Println("Text", text)
	fmt.Println("Character", char, "=>", string(char))

	// multiple variable
	var (
		benar bool
		salah string
	)
	benar = true
	salah = "benarkan"

	fmt.Println("Benar", benar)
	fmt.Println("Salah", salah)

	// simple
	username := "aliffatul"

	fmt.Println("username", username)
}

func constant() {
	Banner("constanta")
	const id int8 = 1
	const firstName string = "Alif"
	const _ string = "Fatul" // unused

	fmt.Println(id, firstName)
}

func conversion() {
	Banner("conversion")
	printMinimal := func(s ...interface{}) {
		fmt.Printf("Before %v(%T), After %v(%T)\n", s[0], s[0], s[1], s[1])
	}

	c := func(s string) []rune {
		var result []rune
		for _, v := range s {
			result = append(result, v)
		}
		return result
	}

	var i8 int8 = 8
	var i8tou8 uint8 = uint8(i8)

	var i8tof32 float32 = float32(i8)
	var text string = "Golang"
	var strtorune []rune = c(text)

	printMinimal(i8, i8tou8)
	printMinimal(i8, i8tof32)
	printMinimal(text, strtorune)
}

func alias() {
	Banner("Type Declaration / Alias")
	type huruf string
	type angka int

	var s huruf = "string"
	var i angka = 1024

	fmt.Println(s, i)
}

// Operator start
type Operator struct {
	c float64
	b [3]int
}

type IOperator interface {
	Kubus()
	Balok()
}

func (o *Operator) Kubus() {
	fmt.Println("Kubus:", math.Pow(o.c, 3))
}
func (o *Operator) Balok() {
	fmt.Println("Balok:", o.b[0]*o.b[1]*o.b[2])
}

func operator() {
	Banner("Math")
	x := 8
	y := 4
	z := 4

	op := Operator{
		c: 8,
		b: [3]int{x, y, z},
	}

	op.Balok()
	op.Kubus()

	Banner("Operator")
	fmt.Println("A == A", "A" == "A")                          // equal
	fmt.Println("A != A", "A" != "A")                          // no equal
	fmt.Println("4 < 5", 4 < 5)                                // lower than
	fmt.Println("4 < 6(true) && 7 > 8(false)", 4 < 6 && 7 > 8) // and
	fmt.Println("4 < 6(true) || 7 > 8(false)", 4 < 6 || 7 > 8) // or

	a := 0
	a += 8
	a *= 2
	a /= 4
	a %= 2
	a -= -2

	fmt.Println(a)
}

// Operator end

func slowRemoveIndex(array []string, index int) []string {
	a, b := array[:index], array[index+1:]
	return append(a, b...)
}

func tryArrayAndSlice() {
	Banner("Array & Slice")

	// https://randomwordgenerator.com/
	a := [4]string{"texture", "transmission", "sow", "mosaic"}
	b := a[0:3]
	bb := append(b, "meadow")
	bb[len(bb)-1] = "last"
	c := a[3:]
	d := append(c, "connection")
	d[0] = "float"

	fmt.Printf("a %v, LEN: %d, CAP: %d\n", a, len(a), cap(a))
	fmt.Printf("b %v, LEN: %d, CAP: %d\n", b, len(b), cap(b))
	fmt.Printf("bb %v, LEN: %d, CAP: %d\n", bb, len(bb), cap(bb))
	fmt.Printf("c %v, LEN: %d, CAP: %d\n", c, len(c), cap(c))
	fmt.Printf("d %v, LEN: %d, CAP: %d\n", d, len(d), cap(d))

	ss := slowRemoveIndex(bb, 2)
	fmt.Println(ss, len(ss), cap(ss))
}

func maps() {
	Banner("Map")
	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	m2 := make(map[int]string)
	m2[1] = "a"
	m2[2] = "b"
	m2[3] = "c"
	m2[4] = "d"
	m2[5] = "e"

	fmt.Println(m)
	for _, v := range m {
		fmt.Printf("%v ", v)
	}
	fmt.Println() // new line
	fmt.Println(m2)
	delete(m2, 1)
	fmt.Println(m2)

	for _, v := range m2 {
		fmt.Printf("%v ", v)
	}
	fmt.Println() // new line
}

func ifexp() {
	Banner("If Expression")

	earth := "planet"

	if earth != "planet" {
		fmt.Println("Earth is moon")
	} else {
		fmt.Println("Earth is planet")
	}

	star := "red"

	if star == "hot" {
		fmt.Println("Star is hot")
	} else if star == "cold" {
		fmt.Println("Star is cold")
	} else {
		fmt.Println("Not star")
	}

	red := "mars"
	mars := "planet"

	if "mars" == red {
		if mars == "planet" {
			fmt.Println("Mars is planet")
		} else {
			fmt.Println("Mars not a planet")
		}
	} else {
		fmt.Println("Mars is not red")
	}

	word := "SunIsBiggerThanEarth" // SunIsBigger 11

	if 11 > len(word) {
		fmt.Println("Not true")
	} else if len(word) > 11 {
		fmt.Println(word[:11])
	}
}

func switchlib1(id, name string) {
	switch id {
	case "planet":
		fmt.Println(name, "is planet")
	case "moon":
		fmt.Println(name, "is moon")
	case "loop":
		{
			for x := 0; x < 10; x++ {
				fmt.Println("Write planet", name)
			}
		}
	case "loopbreak":
		{
			for x := 0; x < 10; x++ {
				if x == 5 {
					break
				}
				fmt.Println("Write planet", name)
			}
		}
	default:
		fmt.Println("none")
	}

}

func switchlib2(name string) {
	switch {
	case name == "proxima centauri":
		fmt.Println("Star")
	case name == "europa":
		fmt.Println("Jupyters moon")
	}
}

func switchlib3(n uint) {
	switch sum := n + 80; sum > 100 {
	case true:
		fmt.Println("sum > 100")
	case false:
		fmt.Println("sum < 100")
	}
}

func switchexp() {
	Banner("Switch Expression")

	switchlib1("planet", "earth")
	switchlib1("moon", "io")
	switchlib1("loop", "saturn")
	switchlib1("loopbreak", "loop")

	switchlib2("europa")
	switchlib2("io")

	switchlib3(200)
}

func main() {
	fmt.Println("Let's GO!")

	variable()
	constant()
	conversion()
	alias()
	operator()
	tryArrayAndSlice()
	maps()
	ifexp()
	switchexp()
}
