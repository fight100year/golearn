package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	pid, _, _ := syscall.Syscall(39, 0, 0, 0)
	fmt.Println("my pid is:", pid)

	uid, _, _ := syscall.Syscall(24, 0, 0, 0)
	fmt.Println("user id is:", uid)

	msg := []byte{'H', 'e', 'l', 'l', 'o', '\n'}
	fd := 1
	syscall.Write(fd, msg)

	fmt.Println("using syscall.Exec()")
	command := "/bin/cat"
	env := os.Environ()
	syscall.Exec(command, []string{"ls", "-a", "-x"}, env)
}
