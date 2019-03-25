package main

/*
# include <stdio.h>
void callC() {
  printf("in c code");
}
*/
import "C"
import "fmt"

func main() {
	C.callC()
	fmt.Println("done")

}
