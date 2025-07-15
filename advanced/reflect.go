package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	name   string
	age    int
	salary float64
}

func main() {
	x := 36
	v := reflect.ValueOf(x)
	t := v.Type()

	fmt.Println(v)
	fmt.Println(v.IsZero())
	fmt.Println(t)
	fmt.Println(t.Kind())

	emp := Employee{
		name:   "John",
		age:    30,
		salary: 500000.0,
	}

	empReflectValue := reflect.ValueOf(emp)

	for i := range empReflectValue.NumField() {
		fmt.Printf("Field %d: with value: %v\n", i, empReflectValue.Field(i))
	}
}
