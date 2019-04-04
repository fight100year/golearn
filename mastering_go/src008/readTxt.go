package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

func lineByline(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("read file failed:%s\n", err)
			break
		}
		fmt.Print(line)
	}

	return nil
}

func wordbyWord(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("read file failed:%s\n", err)
			break
		}

		r := regexp.MustCompile("[^\\s]+")
		words := r.FindAllString(line, -1)
		for i := 0; i < len(words); i++ {
			fmt.Println(words[i])
		}
	}

	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: readTxt <file1> [<file2> ...]\n")
		return
	}

	for _, file := range flag.Args() {
		err := lineByline(file)
		if err != nil {
			fmt.Println(err)
		}

		err = wordbyWord(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
