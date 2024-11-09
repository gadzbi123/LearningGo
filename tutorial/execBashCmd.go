package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "ls")
	go func() {
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Failed to exec bash: %v\n", err)
		}
	}()
	t1 := time.NewTicker(2 * time.Second)
	defer t1.Stop()
	<-t1.C
	fmt.Println("time has passed")
	cancel()
	fmt.Println("Running forever")
	var forever chan (struct{})
	<-forever
}
