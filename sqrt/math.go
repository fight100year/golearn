package main

import (
	"fmt"
)

func Sqrt(x float64) float64{
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z * z - x) / (2 * z)
		fmt.Println(i, "th :", z)
	}

  return z
}

func test(x int) {
  switch x{
  case 1:
  case 2:
    fmt.Println(x, " < 5")
  case 10:
    fmt.Println(x, " > 5")
  default:
    fmt.Println("unknown")
  }
}

func main(){
	fmt.Println(Sqrt(10))
  test(2)
  test(1)
  test(10)
  test(11)
}
