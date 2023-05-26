package main

import printChannel "test-project/Practice"

var myChannel chan string      // Exported channel variable
var anotherChannel chan string // Exported channel variable

func main() {
	myChannel = make(chan string)
	//time.Sleep(time.Second)
	anotherChannel = make(chan string)

	go func() {
		myChannel <- "mychanal"
	}()

	go func() {
		anotherChannel <- "anotherchanal"
	}()

	printChannel.PrintChannel(myChannel, anotherChannel)

}
