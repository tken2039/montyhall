package montyhall

import (
	"github.com/tken2039/montyhall/internal/util"
)

type Scenario struct {
	// game options
	door       *Door
	challenger *Challenger
}

func NewScenario(v *Verification) *Scenario {
	scenario := &Scenario{}

	scenario.door = NewDoor(v.doorCount)
	scenario.challenger = NewChallenger(v.willChange)

	return scenario
}

type doors struct {
	doorCount int
	correct   int
}

func NewDoors(doorCount int) *doors {
	d := &doors{}

	d.doorCount = doorCount
	d.correct = util.GenerateRandomNumber(doorCount) + 1

	return d
}

func (sc *Scenario) playMontyHall() (bool, error) {
	// player selects one door

	if err := sc.selectDoor(); err != nil {
		return false, err
	}

	// monty opens one door other than the door selected by player
	if err := sc.openDoor(); err != nil {
		return false, err
	}

	// player can re-select the door
	if sc.challenger.willChange {
		// re-select
		if err := sc.reSelectDoor(); err != nil {
			return false, err
		}
	} else {
		// unchange from before-selected
		sc.door.afterSelected = sc.door.beforeSelected
	}

	// The result of opening the door was...
	isHit := isHit(sc.door.correct, sc.door.afterSelected)

	return isHit, nil
}

func (sc *Scenario) selectDoor() error {
	selected, err := sc.challenger.selectDoor(sc.door.count)
	if err != nil {
		return err
	}

	sc.door.beforeSelected = selected

	return nil
}

func (sc *Scenario) reSelectDoor() error {
	selected, err := sc.challenger.selectDoor(sc.door.count, sc.door.beforeSelected, sc.door.opend)
	if err != nil {
		return err
	}

	sc.door.afterSelected = selected

	return nil
}

func (sc *Scenario) openDoor() error {
	err := sc.door.openDoor()
	if err != nil {
		return err
	}

	return nil
}

func isHit(correct int, selected int) bool {
	return correct == selected
}
