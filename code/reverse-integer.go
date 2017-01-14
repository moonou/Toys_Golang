package code

import (
	"fmt"
	"strconv"
)

func Reverse_integer() {
	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {
	str := strconv.Itoa(x)
	restr := ""
	flag := false
	for _, v := range str {
		if v == 45 {
			flag = true
			continue
		}
		restr = string(v) + restr
	}
	if flag {
		restr = "-" + restr
	}
	result, err := strconv.Atoi(restr)

	if err != nil || result > 2147483647 || result < -2147483647 {
		return 0
	}
	return result
}
