package main

import "fmt"

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

func (t Task) Summary() string {
	return fmt.Sprintf("#%d %s [%s], приоритет %d", t.ID, t.Title, t.Status, t.Priority)
}

func (t *Task) Complete() {
	t.Status = "done"
}

func (t *Task) SetPriority(priority int) bool {
	if priority < 1 || priority > 3 {
		return false
	}
	t.Priority = priority
	return true
}

func findTaskByID(tasks []Task, id int) *Task {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	return nil
}

func tasksByStatus(tasks []Task, status string) []Task {
	sort_tasks := []Task{}
	for _, task := range tasks {
		if task.Status == status {
			sort_tasks = append(sort_tasks, task)
		}
	}
	return sort_tasks
}

func main() {
	tasks := []Task{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3},
	}

	task2 := findTaskByID(tasks, 2)
	task2.Complete()
	task2.SetPriority(3)

	counts := make(map[string]int)
	for _, task := range tasks {
		counts[task.Status]++
	}

	for _, task := range tasks {
		fmt.Println(task.Summary())
	}
	fmt.Println("Задач завершено:", counts["done"])
	if findTaskByID(tasks, 99) == nil {
		fmt.Println("Задача не найдена")
	}

	sort_tasks := tasksByStatus(tasks, "done")
	for _, task := range sort_tasks {
		fmt.Println(task.Summary())
	}
}
