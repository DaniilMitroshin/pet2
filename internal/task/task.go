package task

import (
	"fmt"
	"strings"
)

type task struct {
	title       string
	description string
	setTags     map[string]struct{}
	set_lists   map[string]struct{}
}

func (t *task) String() string {
	tags := make([]string, 0, len(t.setTags))
	for k := range t.setTags {
		tags = append(tags, k)
	}

	lists := make([]string, 0, len(t.set_lists))
	for k := range t.set_lists {
		tags = append(lists, k)
	}

	return fmt.Sprintf("Заголовок: %s \n Описание: %s \n Теги: %s \n Списки: %s \n", t.title, t.description, strings.Join(tags, ", "), strings.Join(lists, ", "))
}
