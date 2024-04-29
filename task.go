package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Task struct {
	Title    string `yaml:"title"`
	Desc     string `yaml:"desc"`
	ExecTime string `yaml:"execTime"`
}

func NewTaskFromFile(taskName string) *Task {

	data, err := os.ReadFile(fmt.Sprintf("%s.yaml", taskName))
	if err != nil {
		panic(err)
	}

	var task Task
	if err := yaml.Unmarshal(data, &task); err != nil {
		panic(err)
	}
	return &task
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
		t.Title = title
		return t
	}
}

func WithDesc(desc string) func(*Task) *Task {
	return func(t *Task) *Task {
		t.Desc = desc
		return t
	}
}

func WithExecTime(execTime string) func(*Task) *Task {
	return func(t *Task) *Task {
		t.ExecTime = execTime
		return t
	}
}
