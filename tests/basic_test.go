package main

import (
	"testing"
	"strings"
	"fmt"
)

type Base struct {
	a int
	af func()
}
func (b *Base)method(){
	fmt.Printf("this is Base\n")
}
type ObjectBase struct {
	Base
	b int
	ob func()
	af func()
}
func (b *ObjectBase)method(){
	fmt.Printf("this is ObjectBase\n")
}
func TestTrim(t *testing.T) {
	//test compose
	objectBase := ObjectBase{}
	objectBase.af = func() {
		fmt.Printf("this is af\n")
	}
	objectBase.af()
	objectBase.method()
	objectBase.Base.method()
	fmt.Printf("%#v \n", objectBase)

	ch := make(chan int)
	go func() {
		fmt.Printf("%#v\n", <-ch)
	}()
	ch <- 1

	close(ch)
	select {
	case v, ok := <-ch:
		fmt.Printf("closed ch is:%#v  %#v\n", v, ok)
	}
	//select {
	//case ch<-2://panic
	//	fmt.Printf("closed ch send is:%#v  %#v\n")
	//}
	//ok := ch<-2 //panic



	t.Run("Trim", func(t *testing.T){
		if strings.Trim(" ,.sasdf#$,. ", " ,.#") != "sasdf#$" {
			fmt.Println("assert Trim fail")
			t.Fail()
		}
	})
}

