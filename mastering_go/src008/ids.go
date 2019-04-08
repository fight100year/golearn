package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	fmt.Println("user id:", os.Getuid())

	u, _ := user.Current()
	fmt.Print("group ids: ")
	gid, _ := u.GroupIds()
	for _, i := range gid {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
