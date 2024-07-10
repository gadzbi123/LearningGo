package main

import (
	"fmt"
	"reflect"
	"sync"
)

func main() {
	my_mutex := sync.Mutex{}
	// (&my_mutex).state
	mut := reflect.TypeOf(my_mutex)
	fmt.Println("Mut has", mut.NumField(), "fields:")
	for i := 0; i < mut.NumField(); i++ {
		fmt.Println("Field", mut.Field(i).Name)
	}
	mutPtr := reflect.TypeOf(&sync.Mutex{})
	fmt.Println("*Mut has", mutPtr.NumMethod(), "methods:")
	for i := 0; i < mutPtr.NumMethod(); i++ {
		fmt.Println(mutPtr.Method(i).Name)
	}
}
