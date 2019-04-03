package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 3.0
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))
	fmt.Println("value:", reflect.ValueOf(x).String())

	fmt.Println("kind is float64:", reflect.ValueOf(x).Kind() == reflect.Float64)
	fmt.Println("value:", reflect.ValueOf(x).Float())

	var i uint8 = 'x'
	fmt.Println("kind is unit8:", reflect.ValueOf(i).Kind() == reflect.Uint8)
	i = uint8(reflect.ValueOf(i).Uint())

	fmt.Println(reflect.ValueOf(x).Interface())

	v := reflect.ValueOf(&x)
	fmt.Println("settability of v", v.CanSet())
	fmt.Println("settability of x", v.Elem().CanSet())
	v.Elem().SetFloat(4.0)
	// v.SetFloat(4.0)
	fmt.Println(x)
	reflect.ValueOf(&x).Elem().SetFloat(5.0)
	fmt.Println(x)

	type T struct {
		I int
		S string
	}

	t := T{1, "hello"}
	s := reflect.ValueOf(&t).Elem()
	st := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, st.Field(i).Name, f.Type(), f)
		if i == 0 {
			s.Field(0).SetInt(2)
		} else if i == 1 {
			f.SetString("world")
		}
	}

	fmt.Println("t:", t)
}
