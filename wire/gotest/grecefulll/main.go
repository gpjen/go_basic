package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type job struct {
	wait chan struct{}
	f    func() error
}

func NewJob(f func() error) *job {
	return &job{
		wait: make(chan struct{}, 1),
		f:    f,
	}
}

func (j *job) Execute() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	go func() {
		defer func() {
			j.wait <- struct{}{}
		}()
		j.f()
	}()

	select {
	case <-j.wait:
		{
			return
		}
	case <-exit:
		{
			fmt.Println("KILLING SIGNAL RECIVE...")
			<-j.wait
			fmt.Println("SHUTDOWN DONE...")
		}
	}
}

func doSomething() {
	fmt.Println("PROCESSING SOMETHING....")

	time.Sleep(5 * time.Second)
	fmt.Println("PROCESSING DONE....")
}

func main() {
	fmt.Println("---- Gracefull Shutdown -----")

	job := NewJob(func() error {
		doSomething()
		return nil
	})

	job.Execute()

}
