package printChannel

import (
	"fmt"
	"strconv"
)

func main() {
	res := callBack(func(i int, i2 int) int {
		return i + i2
	}, "wow")

	fmt.Println(res)
}

func callBack(getSum func(i int, s int) int, name string) string {
	f := getSum(3, 4)
	return name + " " + strconv.Itoa(f)
}
