package main

import "fmt"

func main() {
	var min int
	fmt.Print("Введите количество минут: ")
	_, err := fmt.Scan(&min)
	if err != nil {
		fmt.Print("Нужно ввести целое число")
		return
	}
	if min <= 0 {
		fmt.Print("Нужно ввести положительное число")
		return
	}
	hour := min / 60
	last_min := min % 60

	fmt.Println("Полных часов:", hour)
	fmt.Println("Оставшихся минут:", last_min)
}
