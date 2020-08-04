package main

//import "fmt"
import (
	"fmt"
	) 
func main(){ 
	var a string = "Test"
	var b bool
	var c int
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("hello,go")
	var s string
	fmt.Printf("%v %v %q\n",a,b,s)
	var intVal int 

	//intVal :=1 // 这时候会产生编译错误

	intVal,intVal1 := 1,2 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
	fmt.Printf("%v %v",intVal,intVal1)
	f := "Runoob" // var f string = "Runoob"

	var m,n int = 1,2

	fmt.Println(f)
	fmt.Printf("%v和%v\n",m,n)
	m,n = n,m
	fmt.Printf("%v和%v\n",m,n)

	println("Hello World!")

	println("提交测试")
	const(
		aa = iota
		bb
		cc
		ee
		dd = "ha"
		ff = 100
		gg
		hh = iota
		ii
	)
	fmt.Println(aa,bb,cc,dd,ee,ff,gg,hh,ii)

	/*
	select case
	*/
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received",i1,"from c1\n")
	case c2<- i2:
		fmt.Printf("sent",i2,"to c2\n")
	case i3, ok:=(<-c3):// same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received",i3," from c3\n")
		}else{
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

}