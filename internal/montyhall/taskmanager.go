package montyhall

const maxTask int = 100000

func splitTask(taskCount int) []int {
	quotient := int(taskCount / maxTask)
	surplus := int(taskCount % maxTask)

	task := []int{}
	for i := 0; i < quotient; i++ {
		task = append(task, maxTask)
	}

	if surplus != 0 {
		task = append(task, surplus)
	}

	return task
}
