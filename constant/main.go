package constant

import (
	"errors"
	"fmt"
)

const Pi = 3.14

type MenuType string
const (
	PLUS  MenuType = "PLUS"
	MINUS          = "MINUS"
)



func Exec()  {
	fmt.Println(Pi)
	// Menutype check
	data := "test"
	fmt.Println(MenuType("PLUS")==PLUS)
	switch MenuType(data) {
	case PLUS, MINUS:
		fmt.Println("pass")
	default:
		fmt.Println(errors.New("error can't parse type"))
	}
}
