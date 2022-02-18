package helper

import (
	blog "belajar-go/log"
	"fmt"

	go_greeting "github.com/aliffatulmf/go-greeting-mod"
)

func GoModules() {
	blog.Banner("Go Modules")
	fmt.Println("Go Greeting Modules", go_greeting.Alo())
}
