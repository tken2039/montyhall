package montyhall

import (
	"github.com/tken2039/montyhall/internal/util"
)

type Door struct {
	count   int
	correct int

	beforeSelected int
	afterSelected  int

	opend int
}

func NewDoor(d int) *Door {
	// Add 1 for the following reason
	// util.GenerateRandomNumber(n) -> 0 ~ n-1
	correct := util.GenerateRandomNumber(d) + 1
	door := &Door{
		count:   d,
		correct: correct,
	}

	return door
}

func (d *Door) openDoor() {
	isValid := false
	target := 0
	for !isValid {
		target = util.GenerateRandomNumber(d.count) + 1

		untouchableDoors := []int{d.correct, d.beforeSelected}

		isExist := false
		for _, v := range untouchableDoors {
			if target == v {
				isExist = true
				continue
			}
		}

		if !isExist {
			isValid = true
		}
	}

	d.opend = target
}
