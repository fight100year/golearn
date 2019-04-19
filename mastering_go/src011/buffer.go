package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func createBuffer(buf *[]byte, cnt int) {
	if cnt == 0 {
		return
	}

	for i := 0; i < cnt; i++ {
		if len(*buf) > cnt {
			return
		}
		*buf = append(*buf, byte(random(0, 100)))
	}
}

func create(dst string, b, f int) error {
	_, err := os.Stat(dst)
	if err == nil {
		return fmt.Errorf("file %s already exists. ", dst)
	}

	file, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := make([]byte, 0)
	for {
		createBuffer(&buf, b)
		buf = buf[:b]
		if _, err := file.Write(buf); err != nil {
			return err
		}

		if f < 0 {
			break
		}
		f = f - len(buf)
	}

	return err
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	file := "output.file"
	bufferSize, _ := strconv.Atoi(os.Args[1])
	fileSize, _ := strconv.Atoi(os.Args[2])

	err := create(file, bufferSize, fileSize)
	if err != nil {
		fmt.Println(err)
	}

	err = os.Remove(file)
	if err != nil {
		fmt.Println(err)
	}
}
