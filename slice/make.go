package main

import "fmt"

func main(){
  a := make([]int, 5)
  fmt.Println("a: len:%d, cap:%d, %v\n", len(a), cap(a), a)

  b := make([]int, 5, 10)
  fmt.Println("b: len:%d, cap:%d, %v\n", len(b), cap(b), b)

}
