package main

import (
	"fmt"
	"reflect"
)

func changeStruct(in interface{}, values map[string]interface{}) error {
	if in == nil {
		return fmt.Errorf("in is null")
	}

	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("in is not struct")
	}

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)

		valueInterFace, ok := values[typeField.Name]
		if !ok {
			return fmt.Errorf("field %v not found in values", typeField.Name)
		}
		value := reflect.ValueOf(valueInterFace)
		val.Field(i).Set(value)
	}

	return nil
}

func main() {
	type person struct {
		Name string
		Age  int
	}
	type data struct {
		Person person
		Count  int
	}
	a := data{
		Person: person{Name: "andrey", Age: 17},
		Count:  7,
	}

	values := make(map[string]interface{})
	values["Person"] = person{Name: "miha", Age: 27}
	values["Count"] = 14

	fmt.Println(changeStruct(&a, values))
	fmt.Println(a)
}
