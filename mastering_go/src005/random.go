package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	MIN := 0
	MAX := 100
	TOTAL := 100
	SEED := time.Now().Unix()

	arg := os.Args
	switch len(arg) {
	case 2:
		fmt.Println("usage: go run random.go MIN MAX TOTAL SEED")
		MIN, _ = strconv.Atoi(arg[1])
		MAX = MIN + 100
	case 3:
		fmt.Println("usage: go run random.go MIN MAX TOTAL SEED")
		MIN, _ = strconv.Atoi(arg[1])
		MAX, _ = strconv.Atoi(arg[2])
	case 4:
		fmt.Println("usage: go run random.go MIN MAX TOTAL SEED")
		MIN, _ = strconv.Atoi(arg[1])
		MAX, _ = strconv.Atoi(arg[2])
		TOTAL, _ = strconv.Atoi(arg[3])
	case 5:
		fmt.Println("usage: go run random.go MIN MAX TOTAL SEED")
		MIN, _ = strconv.Atoi(arg[1])
		MAX, _ = strconv.Atoi(arg[2])
		TOTAL, _ = strconv.Atoi(arg[3])
		SEED, _ = strconv.ParseInt(arg[4], 10, 64)
	default:
		fmt.Println("using default values")
	}

	rand.Seed(SEED)
	for i := 0; i < TOTAL; i++ {
		r := random(MIN, MAX)
		fmt.Print(r, " ")
	}
	fmt.Println()
}
