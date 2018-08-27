package spoj

import "fmt"


func main() {

	var a int
	for {
		fmt.Scan(&a)
		if a == 42 {
			break
		}
		fmt.Println(a)
	}

}

