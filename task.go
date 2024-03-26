package main

type Task struct {
	title string
	desc  string
}

func NewTask(options ...func(*Task) *Task) Task {
	t := Task{}
	for _, opt := range options {
		opt(&t)
	}
	return t
}

func WithTitle(title string) func(*Task) *Task {
	return func(t *Task) *Task {
		t.title = title
		return t
	}
}

func WithDesc(desc string) func(*Task) *Task {
	return func(t *Task) *Task {
		t.desc = desc
		return t
	}
}
