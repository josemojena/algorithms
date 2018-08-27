package codeforces

import (
	"fmt"
)

func Code() {


	i := 0
	box := make([]int, 0 )
	var c int

	fmt.Scan(&c)

	for i < c {
		i++
		var n int
		fmt.Scan(&n)
		if n == 0 {
			continue
		}
		box[n] = 1
	}

	fmt.Println(cap(box))

}
