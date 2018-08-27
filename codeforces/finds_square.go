package codeforces

import (
"fmt"
)

type Point struct {
	x int
	y int
}

func FindSquare() {

	var n int
	var m int

	fmt.Scan(&n)
	fmt.Scan(&m)

	var s string
	var p1 Point
	var p2 Point
	p1.x = -1
	p1.y = -1
	p2.x = -1
	p2.y = -1
	for i := 0; i < n; i++ {
		fmt.Scan(&s)

		for j, sub := range s {

			if string(sub) == "B" && p1.x == -1 {

				p1.x = i + 1
				p1.y = j + 1

			} else if string(sub) == "B" {

				p2.x = i + 1
				p2.y = j + 1
			}
		}
	}

	if p2.x == -1 || p2.y == -1 {
		fmt.Println(p1.x, p1.y)
	} else {

		fmt.Println(((p1.x+p2.x))/2, ((p1.y+p2.y))/2)
	}

}
