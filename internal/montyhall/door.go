package montyhall

import (
	"fmt"

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

func (d *Door) openDoor() error {
	isValid := false
	target := 0

	if d.count < 3 {
		return fmt.Errorf("the number of doors must be at least 3. given door count: %v", d.count)
	}

	untouchableDoors := []int{d.correct, d.beforeSelected}
	for !isValid {
		target = util.GenerateRandomNumber(d.count) + 1

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

	return nil
}
