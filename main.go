package main

import (
	"fmt"
	"goTraining/constant"
	"goTraining/variable"
	"os"
)

func main()  {
	choice := os.Args[1]
	fmt.Println(choice)

	switch choice {
	case "v": variable.Exec()
	case "c": constant.Exec()

	}
}
