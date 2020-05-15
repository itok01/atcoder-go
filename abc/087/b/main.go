package main

import "fmt"

func main() {
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)

	count := 0

	for y500 := 0; y500 <= a; y500++ {
		for y100 := 0; y100 <= b; y100++ {
			for y50 := 0; y50 <= c; y50++ {
				if ((y500 * 500) + (y100 * 100) + (y50 * 50)) == x {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
