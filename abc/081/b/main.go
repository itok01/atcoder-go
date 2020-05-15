package main

import "fmt"

func main() {
	var n int
	var a []uint

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var _a uint
		fmt.Scan(&_a)

		a = append(a, _a)
	}

	count := 0

	for {
		for i := 0; i < len(a); i++ {
			if a[i]%2 == 0 {
				a[i] /= 2
			} else {
				fmt.Println(count)
				return
			}
		}
		count++
	}
}
