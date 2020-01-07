package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	var (
		ctx        context.Context
		cancelFunc context.CancelFunc
		cmd        *exec.Cmd
		c          chan string
	)

	c = make(chan string)

	ctx, cancelFunc = context.WithCancel(context.TODO())

	go func() {
		var (
			output []byte
			err    error
		)

		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 5; echo 5")
		if output, err = cmd.CombinedOutput(); err != nil {
			fmt.Println("err: ", err)
		}

		fmt.Println(string(output))

		a := "close"

		c <- a
	}()

	time.Sleep(time.Second)

	cancelFunc()

	<-c
}
