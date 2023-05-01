package main

import (
	"fmt"
	"strings"
	"test-project/bill"
)

func main() {
	//str := "hey hello world"
	//fmt.Print(replace(str, "hey", "wow"))
	myBill := bill.SetBill()
	newBill := bill.Bill{
		Name:  "wowow",
		Items: []string{"low", "USD"},
	}
	fmt.Println(myBill.GetBill(), newBill.GetBill())
}

func replace(fullStr string, replaceableStr string, newStr string) string {
	fmt.Println(getInit("Forkan alam"))
	return strings.ReplaceAll(fullStr, replaceableStr, newStr)
}

func getInit(n string) interface{} {
	if n == "" {
		return false
	}
	s := strings.ToUpper(n)
	sepArr := strings.Split(s, " ")
	fst := sepArr[0][0:1]
	snd := sepArr[1][0:1]
	if len(sepArr) > 2 {
		snd = sepArr[2][0:1]
	}
	return fst + snd
}
