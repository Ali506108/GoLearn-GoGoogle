package main

import (
	"fmt"
	"reflect"
)

type T struct {
	F1 int
	F2 float64
	F3 string
}

func main() {
	data := T{
		1,
		3.14156,
		"London",
	}

	fmt.Println("A : ", data)

	r := reflect.ValueOf(&data).Elem()
	fmt.Println("String value :", r.String())

	typeOfdata := r.Type()

	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		toData := typeOfdata.Field(i).Name
		fmt.Printf("Field %s  (%v) = %v\n", toData, f.Type(), f.Interface())

		k := reflect.TypeOf(r.Field(i).Interface()).Kind()

		switch k {
		case reflect.Int:
			r.Field(i).SetInt(-43)
		case reflect.String:
			r.Field(i).SetString("Boston")
		}
	}

	fmt.Println("A : ", data)

}
