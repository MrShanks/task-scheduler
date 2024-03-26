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
		want := Task{title: "test"}
		got := NewTask(WithTitle("test"))

		if got.title != want.title {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})

	t.Run("New Task WithDesc has desc test", func(t *testing.T) {
		want := Task{desc: "test"}
		got := NewTask(WithDesc("test"))

		if got.desc != want.desc {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})
}
