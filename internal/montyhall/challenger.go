package montyhall

import (
	"github.com/tken2039/montyhall/internal/util"
)

type Challenger struct {
	willChange bool
}

func NewChallenger(wc bool) *Challenger {
	return &Challenger{wc}
}

func (c *Challenger) selectDoor(door int, untouchableDoors ...int) int {
	isValid := false
	target := 0
	for !isValid {
		target = util.GenerateRandomNumber(door) + 1

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

	return target
}
