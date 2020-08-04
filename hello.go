package main

import (
	"fmt"
	"net/http"
)
func handler(writer http.ResponseWriter,request *http.Request){
	fmt.Printf(writer,"Hello World, %s!",request.URL.Path[1:])
}
func main(){
	http.HaddleFunc("/",handler)
	http.ListenAndServer(":8080",nil)
}