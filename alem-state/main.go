package main

import (
	"fmt"
	"os"
)

type City struct {
	headID int
	units  int

	x int
	y int

	p1 int
	p2 int
}

var cities = make(map[int]City, 10)

type Movement struct {
	from    int
	to      int
	whoFrom int
	whoTo   int
	ticks   int
	units   int
}

func main() {
	for {
		var n, m, ID, tick int
		fmt.Scan(&n, &ID, &tick)

		for i := 0; i < n; i++ {
			var head, units, name, x, y, parametr1, parametr2 int
			fmt.Scan(&head, &units, &name, &x, &y, &parametr1, &parametr2)
			fmt.Fprintln(os.Stderr, head, units, name, x, y, parametr1, parametr2)

			cities[name] = City{
				headID: head,
				units:  units,

				x: x,
				y: y,

				p1: parametr1,
				p2: parametr2,
			}
		}

		fmt.Scan(&m)

		for i := 0; i < m; i++ {
			fmt.Scan()
		}

	}

	// for true {
	// 	var n, m, playerID, tick int
	// 	fmt.Scan(&n, &playerID, &tick)
	// 	// read cities
	// 	for i := 0; i < n; i++ {
	// 		line := ""
	// 		fmt.Scan(&line)
	// 		fmt.Fprint(os.Stderr, line, "\n")
	// 	}

	// 	// number of movements
	// 	fmt.Scan(&m)

	// 	// read movements
	// 	for i := 0; i < m; i++ {
	// 		fmt.Scan(&line)
	// 	}

	// 	// use `os.Stderr` to print for debugging
	// 	fmt.Fprintf(os.Stderr, "debug code\n")
	// 	fmt.Println("100 100 200 200")
	// }
}
