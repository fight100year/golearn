package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("usage: column 20 <file1> [<file2> ...]")
		os.Exit(1)
	}

	tmp, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("第一个参数是整数:", tmp)
		return
	}

	if tmp < 0 {
		fmt.Println("第一个参数是正整数:", tmp)
		os.Exit(1)
	}

	for _, filename := range args[2:] {
		fmt.Println("\t\t", filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("open file failed:", filename)
			continue
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println("read file failed:", err)
			}

			data := strings.Fields(line) // 按空格分割字符串
			if len(data) >= tmp {
				fmt.Println(data[tmp-1])
			}
		}
	}
}
