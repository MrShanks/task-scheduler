package main

import "testing"

func TestNewTask(t *testing.T) {
	t.Run("New Task is empty", func(t *testing.T) {
		want := &Task{}
		got := NewTask()

		if got != want {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})
}
