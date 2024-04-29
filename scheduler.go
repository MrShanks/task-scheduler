package main

import (
	"fmt"
	"sync"
	"time"
)

type Scheduler struct {
	wg             sync.WaitGroup
	tasks          chan (Task)
	scheduledTasks map[Task]bool
	rwmutex        sync.RWMutex
}

func (s *Scheduler) Run() {
	s.scheduledTasks = make(map[Task]bool)
	for {
		select {
		case task, ok := <-s.tasks:
			if !ok {
				return
			}
			s.rwmutex.Lock()
			s.scheduledTasks[task] = true
			s.rwmutex.Unlock()
		default:
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func (s *Scheduler) AddTask(task Task) {
	defer s.wg.Done()
	s.tasks <- task
}

func (s *Scheduler) Execute() {
	for {
		for task := range s.scheduledTasks {
			if len(s.scheduledTasks) == 0 {
				return
			}
			if time.Now().Format("Jan 02, 2006 3:04 PM") == task.ExecTime {
				fmt.Printf("%s: %s \n", task.Title, task.ExecTime)
				delete(s.scheduledTasks, task)
			}
		}
		fmt.Println("Next run in 1 min")
		time.Sleep(1 * time.Minute)

	}
}
