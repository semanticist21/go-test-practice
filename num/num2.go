package num

import "fmt"

func PrintInReverse(num int) {
	for ; num != 0; num-- {
		fmt.Println(num)
	}
}
