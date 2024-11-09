package main

import (
	"fmt"
	"os"
	"syscall"
)

func aa() {
}

func main() {
	cwd, _ := os.Getwd()
	procAttr := syscall.ProcAttr{
		Dir:   cwd,
		Files: nil, //[]uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
		Sys:   &syscall.SysProcAttr{},
	}
	var forever chan (struct{})
	go func() {
		pid, h, err := syscall.StartProcess("/usr/bin/bash", []string{}, &procAttr)
		if err != nil {
			panic(err)
		}
		fmt.Printf("pid1: %v, pid2: %v\n", os.Getpid(), pid)
		fileLock := syscall.Flock_t{}
		err = syscall.FcntlFlock(h, syscall.F_GETLK, &fileLock)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Flock= %+v", fileLock)
	}()

	<-forever
	// defer syscall.Close(handler)
}
