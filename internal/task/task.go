package task

import (
	"fmt"
)

type Task struct {
	title       string
	description string
	isDone      bool
	//setTags     map[string]struct{}
	//set_lists   map[string]struct{}
}

func NewTask(title, description string, isdone bool) Task {
	return Task{
		title:       title,
		description: description,
		isDone:      isdone,
	}
}

func (t *Task) String() string {
	/*
		tags := make([]string, 0, len(t.setTags))
		for k := range t.setTags {
			tags = append(tags, k)
		}

		lists := make([]string, 0, len(t.set_lists))
		for k := range t.set_lists {
			tags = append(lists, k)
		}
	*/

	var doneText string
	if t.isDone {
		doneText = "[DONE]"
	} else {
		doneText = "[TODO]"
	}
	return fmt.Sprintf("%s \nЗаголовок: %s \nОписание: %s \n", doneText, t.title, t.description)

	//return fmt.Sprintf("%s \nЗаголовок: %s \nОписание: %s \nТеги: %s \nСписки: %s \n",doneText, t.title, t.description, strings.Join(tags, ", "), strings.Join(lists, ", "))

}
