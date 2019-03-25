package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2

	m2 := map[string]int{
		"k1": 1,
		"k2": 2,
	}

	m = nil
	m["k1"] = 1
	fmt.Println("m:", m)
	fmt.Println("m2:", m2)

	delete(m2, "k1")
	delete(m2, "k1")
	delete(m2, "k1")
	delete(m2, "k1")
	m2["k2"] = 3
	m2["k3"] = 4
	fmt.Println("m2:", m2)

	_, ok := m2["k1"]
	if ok {
		fmt.Println("exists")
	} else {
		fmt.Println("not exists")
	}

}
