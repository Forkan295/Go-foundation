package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("go routing working 1")
	go fmt.Println("go routing working 2")
	go fmt.Println("go routing working 3")
	time.Sleep(time.Second)
	go fmt.Println("go routing working 4")
	go fmt.Println("go routing working 5")
	time.Sleep(time.Second * 3)
	go fmt.Println("go routing working 6")
	time.Sleep(time.Second * 3)
	fmt.Println("go routing working 7")
}
