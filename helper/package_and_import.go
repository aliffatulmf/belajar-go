package helper

import (
	blog "belajar-go/log"
	"fmt"
)

func CanImport() {
	blog.Banner("Import")
	fmt.Println("Work")
}

func cannotImport() {
	fmt.Println("Fail")
}
