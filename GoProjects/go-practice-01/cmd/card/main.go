package main

import "fmt"

func main() {
	var name string = "NotA65"
	var group string = "Main-03"
	const course = "Бэкенд-разработка на языке Go"
	complete := 1
	interest := true

	fmt.Println("Участник:", name)
	fmt.Println("Группа:", group)
	fmt.Printf("Курс: %s\n", course)
	fmt.Println("Завершено работ:", complete)
	fmt.Println("Интерес к бэкенду:", interest)
}
