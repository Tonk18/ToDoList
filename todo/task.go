package todo

import "time"

type Task struct {
	Title       string
	Discription string
	IsDone      bool

	CreatedAt time.Time
	DoneAt    *time.Time
}

func NewTask(title string, discription string) Task {
	return Task{
		Title:       title,
		Discription: discription,
		IsDone:      false,
		CreatedAt:   time.Now(),
		DoneAt:      nil,
	}
}

func (t *Task) Done() {
	doneTime := time.Now()
	t.IsDone = true
	t.DoneAt = &doneTime
}

func (t *Task) Undone() {
	t.IsDone = false
	t.DoneAt = nil

}
