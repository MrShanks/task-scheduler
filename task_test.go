package main

import "testing"

func TestNewTask(t *testing.T) {
	t.Run("New Task is empty", func(t *testing.T) {
		want := Task{}
		got := NewTask()

		if got != want {
			t.Errorf("Expected %#v, got %#v", want, got)
		}
	})

	t.Run("New Task WithTitle has title test", func(t *testing.T) {
		want := Task{Title: "test"}
		got := NewTask(WithTitle("test"))

		if got.Title != want.Title {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})

	t.Run("New Task WithDesc has desc test", func(t *testing.T) {
		want := Task{Desc: "test"}
		got := NewTask(WithDesc("test"))

		if got.Desc != want.Desc {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})
}
