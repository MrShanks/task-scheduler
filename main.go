package main

import (
	"fmt"
	"sync"
	"time"
)

var scheduledTasks map[Task]bool
var rwm sync.RWMutex

type Scheduler struct {
	wg    sync.WaitGroup
	tasks chan (Task)
}

func (s *Scheduler) Run() {
	for {
		select {
		case task, ok := <-s.tasks:
			if !ok {
				return
			}
			rwm.Lock()
			scheduledTasks[task] = true
			rwm.Unlock()
		default:
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func Execute() {
	for {
		for key := range scheduledTasks {
			if len(scheduledTasks) == 0 {
				return
			}
			if time.Now().Format("Jan 02, 2006 3:04 PM") == key.execTime {
				fmt.Printf("%s: %s \n", key.title, key.execTime)
				delete(scheduledTasks, key)
			}
		}
		fmt.Println("Next run in 1 min")
		time.Sleep(1 * time.Minute)

	}
}

func (s *Scheduler) AddTask(task Task) {
	defer s.wg.Done()
	s.tasks <- task
}

func main() {
	s := Scheduler{tasks: make(chan Task)}
	scheduledTasks = make(map[Task]bool)

	go func() {
		defer close(s.tasks)
		s.wg.Add(1)
		go s.AddTask(
			NewTask(
				WithTitle(fmt.Sprintf("test-%d", 1)),
				WithExecTime(time.Now().Add(1*time.Minute).Format("Jan 02, 2006 3:04 PM")),
			),
		)
		s.wg.Add(1)
		go s.AddTask(
			NewTask(
				WithTitle(fmt.Sprintf("test-%d", 2)),
				WithExecTime(time.Now().Add(2*time.Minute).Format("Jan 02, 2006 3:04 PM")),
			),
		)
		s.wg.Add(1)
		go s.AddTask(
			NewTask(
				WithTitle(fmt.Sprintf("test-%d", 3)),
				WithExecTime(time.Now().Add(5*time.Minute).Format("Jan 02, 2006 3:04 PM")),
			),
		)
		s.wg.Wait()
	}()

	go s.Run()
	Execute()

}
