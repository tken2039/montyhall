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
	doorCount  int
	correct    int
	doorStatus []bool
}

func NewDoors(doorCount int) *doors {
	d := &doors{}

	d.doorCount = doorCount
	d.correct = util.GenerateRandomNumber(doorCount) + 1

	return d
}

func (sc *Scenario) playMontyHall() bool {
	// player selects one door
	sc.selectDoor()

	// monty opens one door other than the door selected by player
	sc.openDoor()

	// player can re-select the door
	if sc.challenger.willChange {
		// re-select
		sc.reSelectDoor()
	} else {
		// unchange from before-selected
		sc.door.afterSelected = sc.door.beforeSelected
	}

	// The result of opening the door was...
	isHit := isHit(sc.door.correct, sc.door.afterSelected)

	return isHit
}

func (sc *Scenario) selectDoor() {
	selected := sc.challenger.selectDoor(sc.door.count)

	sc.door.beforeSelected = selected
}

func (sc *Scenario) reSelectDoor() {
	selected := sc.challenger.selectDoor(sc.door.count, sc.door.beforeSelected, sc.door.opend)

	sc.door.afterSelected = selected
}

func (sc *Scenario) openDoor() {
	sc.door.openDoor()
}

func isHit(correct int, selected int) bool {
	return correct == selected
}
