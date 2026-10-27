package main

import "fmt"

// statusLabel превращает код статуса в понятную человеку подпись.
func statusLabel(status string) string {
	// TODO: обработайте new, in_progress, done через switch.
	// TODO: предусмотрите ветку default для незнакомого статуса.
	switch status {
	case "new":
		return "Новая"
	case "in_progress":
		return "В процессе"
	case "done":
		return "Выполнена"
	}
	return "Неизвестный результат"
}

func findStatus(statuses []string, target string) int {
	for i, status := range statuses {
		if status == target {
			return i
		}
	}
	return -1
}

func completionPercent(statuses []string) float64 {
	if len(statuses) == 0 {
		return 0
	}
	counts := make(map[string]int)
	for _, status := range statuses {
		counts[status]++
	}
	return 100 * float64(counts["done"]) / float64(len(statuses))
}

// countStatus считает количество элементов с заданным статусом.
func countStatus(statuses []string, target string) int {
	// TODO: пройдите по срезу через for range.
	// TODO: сравните каждый элемент с target и верните счётчик.
	count := 0
	for _, status := range statuses {
		if status == target {
			count++
		}
	}
	return count
}

func main() {
	statuses := []string{"breaking", "in_progress", "done", "done"}
	statuses = append(statuses, "new")

	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))

	counts := make(map[string]int)
	for _, status := range statuses {
		counts[status]++
	}
	fmt.Printf("Стаитистика: new=%d, in_progress=%d, done=%d\n", counts["new"], counts["in_progress"], counts["done"])
	// TODO: выведите new, in_progress, done в фиксированном порядке.
	target := "done"
	fmt.Printf("Первое вхождение %s находится под индексом %d\n", target, findStatus(statuses, target))
	fmt.Printf("Доля выполненных задач составляет %.1f\n", completionPercent(statuses))
}
