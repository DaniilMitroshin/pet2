package task

type Storage struct {
	tasks     []Task
	tempIndex int
}

func (s *Storage) getTempIndex() int {
	return s.tempIndex
}

func (s *Storage) add(t Task) {
	s.tasks = append(s.tasks, t)
	s.tempIndex++
}
