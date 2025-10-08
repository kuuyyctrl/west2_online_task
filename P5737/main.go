package main

import "fmt"

func yerun(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func main() {
	arr := make([]int, 0)
	var a, b, c, cot int
	c = 1
	fmt.Scan(&a, &b)
	cot = 0
	for a <= b {
		if yerun(a) {
			cot += 1
			arr = append(arr, a)
			c = 4
		}
		a += c
	}
	fmt.Println(cot)
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Print("\n")
}
