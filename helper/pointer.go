package helper

import (
	blog "belajar-go/log"
	"fmt"
)

type Address struct {
	City, Province, Country string
}

type Interface interface {
	tptr(city, province, country string)
}

func (a *Address) tptr(city, province, country string) {
	// a = Address{city, province, country}
	a.City = city
	a.Province = province
	a.Country = country
}

func paramPtr(a *Address) {
	a.City = "Banda Aceh"
	a.Province = "Banda"
	a.Country = "Indonesia"
}

func paramNonPtr(a Address) {
	a.City = "Jakarta Pusat"
	a.Province = "DKI Jakarta"
	a.Country = "Indonesia"
}

func Pointer() {
	blog.Banner("Pointer")

	var a int = 10
	var b *int = &a

	fmt.Println(*b)

	var addr1 Address = Address{"Semarang", "Jawa Tengah", "Indonesia"}
	var addr2 *Address = &addr1
	// addr2.City = "Wonogiri"

	addr2 = &Address{"Wonogiri", "Jawa Tengah", "Indonesia"}
	addr2.City = "Grobogan"

	fmt.Println(addr1)
	fmt.Println(addr2)

	paramPtr(addr2)
	paramNonPtr(addr1)

	fmt.Println(addr2)
	fmt.Println(addr1)

	var o Interface = &Address{"Solo", "Jawa Tengah", "Indonesia"}
	o.tptr("Boyolali", "Jawa Tengah", "Indonesia")

	fmt.Println(o)
}
