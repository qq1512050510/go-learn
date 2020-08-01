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

}