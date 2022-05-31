package montyhall

import (
	"fmt"
	"sync"
)

type Verification struct {
	doorCount  int
	tryCount   int
	willChange bool

	result *Result

	detailMode bool
}

type Result struct {
	hitCount int
	mutex    *sync.Mutex
}

func NewValification(doorCount int, tryCount int, willChange bool, detailMode bool) *Verification {
	v := &Verification{}

	v.doorCount = doorCount
	v.tryCount = tryCount
	v.willChange = willChange
	v.detailMode = detailMode

	v.result = &Result{
		mutex: &sync.Mutex{},
	}

	return v
}

func (v *Verification) Start() error {
	v.startVerification()

	return nil
}

func (v *Verification) startVerification() {
	v.showPreferences()

	tasks := splitTask(v.tryCount)
	wg := &sync.WaitGroup{}
	wg.Add(len(tasks))

	if v.detailMode {
		fmt.Printf("\n- Work Detail -\n")
		fmt.Println("Start verification.")
		fmt.Printf("Number of workers: %v\n", len(tasks))
	}

	for i, t := range tasks {
		tryCount := t
		workerNum := i
		go func() {
			defer wg.Done()

			v.verify(tryCount, workerNum+1)
		}()
	}

	wg.Wait()

	if v.detailMode {
		fmt.Println("---------------")
	}

	v.showResults()
}

func (v *Verification) verify(tryCount int, workerNum int) {
	scenario := NewScenario(v)
	var hitCount int

	for i := 0; i < tryCount-1; i++ {
		isHit := scenario.playMontyHall()

		if isHit {
			hitCount += 1
		}
	}

	if v.detailMode {
		fmt.Printf("HitRate (worker[%v]): %v / %v\n", workerNum, hitCount, tryCount)
	}

	v.addHitCount(hitCount)
}

func (v *Verification) addHitCount(hitCount int) {
	v.result.mutex.Lock()
	defer v.result.mutex.Unlock()

	v.result.hitCount += hitCount
}

func (v *Verification) showPreferences() {
	fmt.Println("\n- Preferences -")

	fmt.Printf("Number of doors: %v\n", v.doorCount)
	fmt.Printf("Number of trials: %v\n", v.tryCount)

	doorChange := "no"
	if v.willChange {
		doorChange = "yes"
	}
	fmt.Printf("Changes after Monty Hall opened the door: %v\n", doorChange)

	fmt.Println("---------------")
}

func (v *Verification) showResults() {
	fmt.Println("\n- Result -")

	fmt.Printf("hit count: %v\n", v.result.hitCount)

	hitRate := float64(v.result.hitCount) / float64(v.tryCount)
	fmt.Printf("hit rate: %.2f %%\n", hitRate*100)

	fmt.Println("----------")
}
