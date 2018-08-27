package spoj


import (
"fmt"
"strconv"
)

func main() {

	var c int
	fmt.Scan(&c)
	vector := make([]string, c)

	for c > 0 {
		c--
		var a, b string
		fmt.Scan(&a)
		fmt.Scan(&b)

		a = Reverse(a)
		b = Reverse(b)

		inta, _ := strconv.Atoi(a)
		intb, _ := strconv.Atoi(b)

		total := inta + intb

		totalReverse := Reverse(strconv.Itoa(total))

		if string(totalReverse[0]) == "0" {

			for len(totalReverse) > 0 {

				if string(totalReverse[0]) == "0" {
					totalReverse = totalReverse[1:]
				} else {
					break
				}
			}
		}
		vector = append(vector, totalReverse)
	}
	for _, i := range vector {
		fmt.Println(i)
	}
}

func Reverse(a string) string {

	var out string
	for p, _ := range a {
		out += string(a[len(a)-1-p])
	}

	return out

}
