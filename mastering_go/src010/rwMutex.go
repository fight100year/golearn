package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type secret struct {
	rwm sync.RWMutex
	m   sync.Mutex
	pwd string
}

var password = secret{pwd: "123"}

func change(c *secret, pwd string) {
	c.rwm.Lock()

	fmt.Println("change")
	time.Sleep(5 * time.Second)
	c.pwd = pwd

	c.rwm.Unlock()
}

func show(c *secret) string {
	c.rwm.RLock()

	// fmt.Println("show")
	time.Sleep(1 * time.Second)
	defer c.rwm.RUnlock()

	return c.pwd
}

func showWithLock(c *secret) string {
	c.rwm.Lock()

	// fmt.Println("show with lock")
	time.Sleep(1 * time.Second)
	defer c.rwm.Unlock()

	return c.pwd
}

func main() {
	var showFunc = func(c *secret) string {
		return ""
	}

	if len(os.Args) != 2 {
		fmt.Println("using sync.RWMutex")
		showFunc = show
	} else {
		fmt.Println("using sync.Mutex")
		showFunc = showWithLock
	}

	var w sync.WaitGroup
	fmt.Println("password:", showFunc(&password))

	for i := 0; i < 15; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			fmt.Println("go pass:", showFunc(&password))
		}()
	}

	w.Add(1)
	go func() {
		defer w.Done()
		change(&password, "456")
	}()

	w.Wait()
	fmt.Println("pass:", showFunc(&password))
}
