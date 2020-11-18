package variable

import "fmt"

func Exec()  {
	var a int
	b:="10"
	var c []int
	var d = []int{1,2,3,4}
	f := []string{"a","b","c"}
	fmt.Println(a,b,c,d,f)
	f = append(f,"abcd")
	fmt.Println("new array f",f)

}
