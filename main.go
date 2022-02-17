package main

import (
	"belajar-go/helper"
	blog "belajar-go/log"
	"fmt"
	"math"
)

func variable() {
	blog.Banner("variable")
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
	blog.Banner("constanta")
	const id int8 = 1
	const firstName string = "Alif"
	const _ string = "Fatul" // unused

	fmt.Println(id, firstName)
}

func conversion() {
	blog.Banner("conversion")
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
	blog.Banner("Type Declaration / Alias")
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
	blog.Banner("Math")
	x := 8
	y := 4
	z := 4

	op := Operator{
		c: 8,
		b: [3]int{x, y, z},
	}

	op.Balok()
	op.Kubus()

	blog.Banner("Operator")
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
	blog.Banner("Array & Slice")

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
	blog.Banner("Map")
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
	blog.Banner("If Expression")

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
	blog.Banner("Switch Expression")

	switchlib1("planet", "earth")
	switchlib1("moon", "io")
	switchlib1("loop", "saturn")
	switchlib1("loopbreak", "loop")

	switchlib2("europa")
	switchlib2("io")

	switchlib3(200)
}

func breakcontinue() {
	blog.Banner("Break & Continue")

	for i := 0; i < 15; i++ {
		if i%2 == 0 {
			continue
		}

		if i == 13 {
			break
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func function() {
	blog.Banner("Function")
	fmt.Println("This is function in golang")
}

func functionParameter(a string, b int, c float64) {
	blog.Banner("Function Parameter")
	fmt.Println(a, b, c)
}

func functionReturnValue(a, b int) int {
	blog.Banner("Function Return Value")
	return a * b
}

func functionReturnMultipleValue(a, b int) (int, int) {
	blog.Banner("Function Return Multiple Value 1")
	return a + b, b - a
}

func functionVariadic(a ...string) string {
	blog.Banner("Function Variadic")
	var res string
	for _, v := range a {
		res += fmt.Sprintf("%s ", v)
	}

	return res
}

func functionValue(name string) string {
	blog.Banner("Function Value")
	return fmt.Sprintf("Earth is %s", name)
}

type Exec func(uint) uint

func functionAsParameter(number uint, exec Exec) string {
	blog.Banner("Function as Parameter")
	n := exec(number)

	if n == 0 {
		return fmt.Sprintf("%d prime", number)
	} else {
		return fmt.Sprintf("%d not prime", number)
	}
}

func exec(i uint) uint {
	var res uint

	switch {
	case i == 1:
		res = 1
	case i%2 > 0:
		res = i % 2
	case i%3 > 0:
		res = i % 3
	default:
		res = 1
	}

	return res
}

// factorial
func recursiveFunction(a uint) uint {
	if a == 1 {
		return 1
	}

	return a * recursiveFunction(a-1)
}

type AnonymousHelper func(string)

func anonymousFunctionHelper(s string, f AnonymousHelper) {
	f(s)
}
func anonymousFunction() {
	blog.Banner("Anonymous Function")
	a := func(s string) {
		fmt.Println("The", s)
	}

	anonymousFunctionHelper("Sun", a)
	anonymousFunctionHelper("Moon", func(s string) {
		fmt.Println(s, "light")
	})
}

func closure() {
	blog.Banner("Closure")
	n := 0

	increment := func() {
		n := 0
		n++
		fmt.Println(n)
	}
	increment()
	fmt.Println(n)
	increment()
	fmt.Println(n)
}

func exception() {
	message := recover()
	fmt.Println("Exit:", message)
}

func error(b bool) {
	if b {
		panic("throw panic")
	}
}

func deferPanicRecover() {
	blog.Banner("Defer, Panic, Recover")

	defer exception()
	fmt.Println("Step 1")

	error(true)

	fmt.Println("Step 2")
}

func functionExec() {
	// function
	function()

	// parameter
	functionParameter("Alif", 21, 8.0)

	// return single value
	v := functionReturnValue(8, 4)
	fmt.Println(v)

	// ruturn multiple value
	m1, m2 := functionReturnMultipleValue(8, 10)
	fmt.Println(m1, m2)

	// variadic
	ss := []string{"hello", "my", "name", "is", "al"}
	r := functionVariadic("hello", "my", "name", "is", "alif")
	fmt.Println(r)
	r2 := functionVariadic(ss...)
	fmt.Println(r2)

	// function value
	ff := functionValue
	fmt.Println(ff("the third planet from the sun"))

	// function as parameter
	fp := functionAsParameter(9, exec)
	fmt.Println(fp)

	// recursive function
	blog.Banner("Recursive Function")
	rf := recursiveFunction(5)
	fmt.Println("5!:", rf)

	// anonymous function
	anonymousFunction()
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
	breakcontinue()
	functionExec()
	closure()
	deferPanicRecover()
	helper.CanImport()
	helper.GoStruct()
	helper.StructScope()
	helper.GoStructFunction()
	helper.GoInterface()
	helper.EmtpyInterface()
	helper.GoNil()
	helper.ErrorInterface()
	helper.TypeAssertions()
	helper.Pointer()
}
