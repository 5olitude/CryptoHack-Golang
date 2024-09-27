package main

import "fmt"

func main() {
	arr := []int{81, 64, 75, 66, 70, 93, 73, 72, 1, 92, 109, 2, 84, 109, 66, 75, 70, 90, 2, 92, 79}
	var flag string
	for _, val := range arr {
		flag += string(val ^ 0x32)
	}
	fmt.Println(flag)
}
