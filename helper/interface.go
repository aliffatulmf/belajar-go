package helper

import (
	blog "belajar-go/log"
	"fmt"
)

type planet struct {
	moon     int
	position int
}
type iplanet interface {
	hasMoon() bool
	getPosition() int
}

func (p planet) hasMoon() bool {
	if p.moon > 0 {
		return true
	}
	return false
}

func (p planet) getPosition() int {
	return p.position
}

func GoInterface() {
	blog.Banner("Interface")
	var p iplanet = planet{
		moon:     2,
		position: 4,
	}

	moon := p.hasMoon()
	position := p.getPosition()

	fmt.Println(moon)
	fmt.Println(position)
}
