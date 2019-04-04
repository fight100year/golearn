package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// D 是否打印目录
var D = false

// F 是否打印常规文件
var F = false

func walk(path string, info os.FileInfo, err error) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	mode := fileInfo.Mode()
	if mode.IsRegular() && F {
		fmt.Println("+", path)
		return nil
	}

	if mode.IsDir() && D {
		fmt.Println("*", path)
		return nil
	}

	fmt.Println(path)
	return nil
}

func main() {
	dd := flag.Bool("d", false, "目录")
	ff := flag.Bool("f", false, "常规文件")
	flag.Parse()
	flags := flag.Args()

	path := "."
	if len(flags) == 1 {
		path = flags[0]
	}

	D = *dd
	F = *ff
	err := filepath.Walk(path, walk)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
