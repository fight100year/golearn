package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("epoch time:", time.Now().Unix)
	t := time.Now()
	fmt.Println(t, t.Format(time.RFC3339))
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	time.Sleep(time.Second)
	t1 := time.Now()
	fmt.Println("time pass: ", t1.Sub(t))

	f := t.Format("01 January 2000")
	fmt.Println(f)
	l, _ := time.LoadLocation("Europe/Paris")
	lt := t.In(l)
	fmt.Println("paris:", lt)
}
