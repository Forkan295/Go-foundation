package main

import (
	"fmt"
)

func main() {
	myChannel := make(chan string)
	//time.Sleep(time.Second)
	anotherChannel := make(chan string)

	go func() {
		myChannel <- "mychanal"
	}()

	go func() {
		anotherChannel <- "anotherchanal"
	}()

	select {
	case printmyChannel := <-myChannel:
		fmt.Println(printmyChannel)
	case printanotherChannel := <-anotherChannel:
		fmt.Println(printanotherChannel)
	}
}
