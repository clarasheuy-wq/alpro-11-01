
package main

import "fmt"

func main() {
	var n, lebar int
	fmt.Scan(&n)

	var tersempit int

	for i := 1; i <= n; i++ {
		fmt.Scan(&lebar)

		if i == 1 || lebar < tersempit {
			tersempit = lebar
		}
	}

	fmt.Println(tersempit)
}
