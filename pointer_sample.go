package main

import (
	"fmt"
	"reflect"
)

func main() {
	foo := &[]bool{true}[0]
	bar := &[]bool{true}[0]

	fmt.Println(foo)
	fmt.Println(bar)

	fmt.Println(*foo)
	fmt.Println(*bar)

	if foo == bar {
		fmt.Println("foo == bar true")
	} else {
		fmt.Println("foo == bar false")
	}

	if *foo == *bar {
		fmt.Println("*foo == *bar true")
	} else {
		fmt.Println("*foo == *bar false")
	}

	if reflect.DeepEqual(foo, bar) {
		fmt.Println("reflect.DeepEqual(foo,bar) true")
	} else {
		fmt.Println("reflect.DeepEqual(foo,bar) true")
	}
}
