package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handle(s os.Signal) {
	fmt.Println("received:", s)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				fmt.Println("cautht signal:", sig)
			case syscall.SIGTERM:
				handle(sig)
				os.Exit(0)
			case syscall.SIGUSR2:
				fmt.Println("handing syscall.SIGUSR2")
			default:
				fmt.Println("ignoring:", sig)
			}
		}
	}()

	for {
		fmt.Println(".")
		time.Sleep(20 * time.Second)
	}
}
