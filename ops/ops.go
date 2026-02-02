package ops

import "fmt"

func GetValue(arg int) int {
	fmt.Println("ARG ", arg)
	return 2 + arg
}
