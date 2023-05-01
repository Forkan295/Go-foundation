package bill

import (
	"fmt"
	"strings"
)

type Bill struct {
	Name  string
	Items []string
}

func SetBill() Bill {
	return Bill{
		Name:  "hell",
		Items: []string{"sdf", "asdfs"},
	}
}

// change or format a struct object property value
func (b Bill) GetBill() Bill {
	b.Name = strings.ToUpper(b.Name)
	return b
}

func (b Bill) Format() string {
	var fs string
	for _, v := range b.Items {
		fs += fmt.Sprintf("%v has %v \n", b.Name, v)
	}
	return fs
}
