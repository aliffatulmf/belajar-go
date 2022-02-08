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
}

// Operator end

func main() {
	fmt.Println("Let's GO!")

	variable()
	constant()
	conversion()
	alias()
	operator()
}
