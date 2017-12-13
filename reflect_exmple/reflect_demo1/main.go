package main

import (
	"fmt"
	"reflect"
)

func MemberInfo(inf interface{}) {
	ref_inf := reflect.TypeOf(inf)

	for {
		if reflect.Struct == ref_inf.Kind() {
			fmt.Printf("\n%-8v %v 个字段:\n", ref_inf, ref_inf.NumField())

			for i := 0; i < ref_inf.NumField(); i++ {
				fmt.Println(ref_inf.Field(i).Name)
			}
		}

		fmt.Printf("\n%-8v %v 个方法:\n", ref_inf, ref_inf.NumMethod())

		for i := 0; i < ref_inf.NumMethod(); i++ {
			fmt.Println(ref_inf.Method(i).Name)
		}

		if reflect.Ptr == ref_inf.Kind() {
			ref_inf = ref_inf.Elem()
		} else {
			break
		}
	}
}

type Company struct {
	CompanyName string
	team        []*Team
	string      //匿名类型
}

type Team struct {
	hum []*Human
}

func (team *Team) RunObject() {
}

func (team *Team) CostTime() {
}

func (team *Team) CostMoney() {
}

type Human struct {
}

func (hum *Human) GetIDCard() {
}

func (hum *Human) GetName() {
}

func (hum *Human) GetWight() {
}

func (hum *Human) GetHight() {
}

func main() {
	MemberInfo(&Company{})
	MemberInfo(&Team{})
	MemberInfo(&Human{})
}
