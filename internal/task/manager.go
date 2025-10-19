package task

import (
	"bufio"
	"fmt"
	"os"
)

//type TaskManager interface{}
//Commands := make(map[string] func())

func add() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите название задачи:")
	for scanner.Scan() {
		//line:= scanner.Text()

	}
}
