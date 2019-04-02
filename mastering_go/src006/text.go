package main

import (
	"fmt"
	"os"
	"text/template"
)

// Entry 数据
type Entry struct {
	Number int
	Square int
}

func main() {
	arg := os.Args
	if len(arg) != 2 {
		fmt.Println("no template file")
		return
	}

	file := arg[1]
	DATA := [][]int{{-1, 1}, {-2, 4}, {-3, 9}, {-4, 16}}

	var e []Entry
	for _, i := range DATA {
		if len(i) == 2 {
			temp := Entry{Number: i[0], Square: i[1]}
			e = append(e, temp)
		}
	}

	t := template.Must(template.ParseGlob(file))
	t.Execute(os.Stdout, e)
}
