package main

import (
	"fmt"
	"reflect"
)

type inf interface {
	Method1()
	Method2()
}

type structfunc struct {
	sayhi func()
}

func (sf structfunc) Method1() {
}

func (sf structfunc) Method2() {
}

func main() {
	struct1 := reflect.TypeOf(structfunc{})
	struct2 := reflect.TypeOf(new(inf)).Elem()

	Test(struct1) //实例是正确的
	Test(struct2) //如果是接口类型，是非法的
}

func Test(t reflect.Type) {
	if t.NumMethod() > 0 {
		fmt.Printf("\n--- %s ---\n", t)
		fmt.Println(t.Method(0).Type)
		fmt.Println(t.Method(0).Func.String())
	} else {
		fmt.Println("None Method")
	}
}
