package main

import (
	"fmt"
	"time"
)

func main(){
	go count("sheep")
	count("starts")
}

func count(s string){
	for i:=0;true;i++{
		fmt.Println(s)
		time.Sleep(time.Millisecond * 500)
	}
}