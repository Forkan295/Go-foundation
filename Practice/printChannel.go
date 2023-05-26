package printChannel

import (
	"fmt"
)

func PrintChannel(myChannel chan string, anotherChannel chan string) {
	select {
	case printmyChannel := <-myChannel:
		fmt.Println(printmyChannel)
	case printanotherChannel := <-anotherChannel:
		fmt.Println(printanotherChannel)
	}
}
