package main

import (
	"fmt"
	"reflect"
)

type User struct {
	UserName    string
	description string
	age         int
}

type Record struct {
	Field  string
	Field2 float64
	Field3 User
}

func createUser(username string, desc string, age int) User {
	user := User{
		UserName:    username,
		description: desc,
		age:         age,
	}

	return user
}

func main() {
	// утечка пмяти и висячие указатели и сложнасть отладки
	name := "Alex"
	desc := "I'm a system architecture"
	age := 20
	Data := Record{
		"Alex_Uk",
		23.42,
		createUser(name, desc, age),
	}
	fmt.Println(Data)

	ref := reflect.ValueOf(Data)
	tType := ref.Type()

	fmt.Println("String value is : ", ref.String())
	fmt.Printf("the value is %s ", tType)

	for i := 0; i < ref.NumField(); i++ {
		fmt.Printf("\t%s", tType.Field(i).Name)
		fmt.Printf("\twitch type : %s ", ref.Field(i).Type())
		fmt.Printf("\tand value_%v_\n", ref.Field(i).Interface())

		sim := reflect.TypeOf(ref.Field(i).Interface()).Kind()
		fmt.Println("Similar")
		if sim == reflect.Struct {
			fmt.Println(ref.Field(i).Type())
		}
	}

}
