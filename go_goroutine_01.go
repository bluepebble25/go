package main

import (
	"fmt"
	"time" //java Thread.sleep(5000)
)

func x(){
	for i:=0; i<10; i++ {
		fmt.Print("x")
	}
}

func y(){
	for i:=0; i<10; i++ {
		fmt.Print("y")
	}
}
func z(){
	fmt.Println("z")
}
func k(){
	fmt.Println("k")
}
func main() {
	defer z()
	defer k()
	go x()
	go y()
	time.Sleep(time.Second * 2)
	fmt.Println("end main()")
}
