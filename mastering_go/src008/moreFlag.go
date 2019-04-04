package main

import (
	"flag"
	"fmt"
	"strings"
)

type ff struct {
	Names []string
}

func (f *ff) getNames() []string {
	return f.Names
}

func (f *ff) String() string {
	return fmt.Sprint(f.Names)
}

func (f *ff) Set(s string) error {
	if len(f.Names) > 0 {
		return fmt.Errorf("已经有值了")
	}

	names := strings.Split(s, ",")
	for _, item := range names {
		f.Names = append(f.Names, item)
	}

	return nil
}

func main() {
	var f ff
	K := flag.Int("k", 0, "an int")
	O := flag.String("o", "hello", "the name")
	flag.Var(&f, "names", "cmd list")

	flag.Parse()
	fmt.Println("-k:", *K)
	fmt.Println("-o:", *O)

	for i, item := range f.getNames() {
		fmt.Println(i, item)
	}

	fmt.Println("all cmd args:")
	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}
