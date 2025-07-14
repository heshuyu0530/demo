package model

type Task struct {
    Id       int
	Content string
}

var nextId int = 0

// NewTask creates a new task

func NewTask(content string) *Task {
	nextId++
	return &Task{
		Id: nextId,
		Content: content,
	}
}