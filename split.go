package main

import "fmt"

// func split(sum int) (int, int) {
// 	x := sum * 7 / 9
// 	y := sum - x
// 	return x, y
// }

func split(sum int) (x, y int) {
	x = sum * 7 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}