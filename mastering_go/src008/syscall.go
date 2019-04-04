package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

var maxSyscalls = 0

const syscallFile = "SYSCALLS"

func main() {
	var calls []string
	f, err := os.Open(syscallFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, " ", "", -1)
		line = strings.Replace(line, "SYS_", "", -1)
		temp := strings.ToLower(strings.Split(line, "=")[0])
		calls = append(calls, temp)
		maxSyscalls++
	}

	counter := make([]int, maxSyscalls)
	var regs syscall.PtraceRegs
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}

	err = cmd.Start()
	err = cmd.Wait()
	if err != nil {
		fmt.Println("wait:", err)
	}

	pid := cmd.Process.Pid
	fmt.Println("process id:", pid)

	before := true
	cnt := 0
	for {
		if before {
			err := syscall.PtraceGetRegs(pid, &regs)
			if err != nil {
				break
			}

			if regs.Orig_rax > uint64(maxSyscalls) {
				fmt.Println("unknown:", regs.Orig_rax)
				return
			}

			calls[regs.Orig_rax]++
			cnt++
		}

		err = syscall.PtraceSyscall(pid, 0)
		if err != nil {
			fmt.Println("ptracesyscall:", err)
			return
		}

		_, err = syscall.Wait4(pid, nil, 0, nil)
		if err != nil {
			fmt.Println("wait4:", err)
			return
		}

		before = !before
	}

	for i, x := range calls {
		if x != 0 {
			fmt.Println(calls[i], "->", x)
		}
	}

	fmt.Println("total system calls:", cnt)
}
