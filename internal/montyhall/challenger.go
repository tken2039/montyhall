package montyhall

import (
	"errors"

	"github.com/tken2039/montyhall/internal/util"
)

type Challenger struct {
	willChange bool
}

func NewChallenger(wc bool) *Challenger {
	return &Challenger{wc}
}

func (c *Challenger) selectDoor(door int, untouchableDoors ...int) (int, error) {
	isValid := false
	target := 0

	if door <= len(untouchableDoors) {
		return -1, errors.New("selectable door is not exists")
	}

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

	return target, nil
}
