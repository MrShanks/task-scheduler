package main

import (
	"fmt"
	"time"
)

func main() {
	s := Scheduler{tasks: make(chan Task)}

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
			*NewTaskFromFile("deploy"),
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
	s.Execute()

}
