package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}
	err := cmd.Start()
	if err != nil {
		fmt.Println("start:", err)
		return
	}

	err = cmd.Wait()
	fmt.Println("state:%v\n", err)
	wpid := cmd.Process.Pid

	var r syscall.PtraceRegs
	err = syscall.PtraceGetRegs(wpid, &r)
	if err != nil {
		fmt.Println("regs:", err)
		return
	}
	fmt.Printf("regs:%#v\n", r)
	fmt.Printf("R15=%d, Gs=%d\n", r.R15, r.Gs)

	time.Sleep(2 * time.Second)
}
