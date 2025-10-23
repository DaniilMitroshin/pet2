package task

import (
	"bufio"
	"fmt"
	"os"
)

var storage Storage

func add() {
	scanner := bufio.NewScanner(os.Stdin)
	var name string
	fmt.Println("Введите название задачи:")
	for scanner.Scan() {
		name = scanner.Text()
		if len(name) < 3 {
			fmt.Println("Название должно состоять из 3 или более символов, попробуйте еще раз:")
			continue
		}
	}
	fmt.Println("Введите описание задачи:")
	scanner.Scan()
	description := scanner.Text()
	storage.add(NewTask(name, description, false))
	fmt.Println("Задача добавлена под номером", storage.getTempIndex())
}

//номер задачи нумеруется с 1, индекс+1

func list() {
	fmt.Println("Список задач:")
	for i, t := range storage.tasks {
		fmt.Print(i+1, "Задача: ", t)
	}
}
