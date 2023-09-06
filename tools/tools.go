package tools

import "fmt"

func ConditionalPrint(cond bool, params ...any) {
	if cond {
		fmt.Println(params...)
	}
}
