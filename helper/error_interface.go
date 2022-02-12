package helper

import (
	blog "belajar-go/log"
	"errors"
	"fmt"
)

func earthIs(s string) (string, error) {
	if s != "planet" {
		return "", errors.New("Error: Earth is planet")
	}

	return "planet", nil
}

func ErrorInterface() {
	blog.Banner("Error Interface")
	res, err := earthIs("planet")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Earth is", res)
	}

}
