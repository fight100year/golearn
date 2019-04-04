package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(s os.Signal) {
	fmt.Println("caught a signal:", s)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGIO)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				fmt.Println("cautht signal:", sig)
			case syscall.SIGIO:
				handleSignal(sig)
				return
			}
		}
	}()

	for {
		fmt.Println(".")
		time.Sleep(20 * time.Second)
	}
}
