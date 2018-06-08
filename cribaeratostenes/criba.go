package main

import (
	"math"
	"strings"
	"fmt"
	"strconv"
)

func main() {

	var buffer strings.Builder

	var n = 100
	checked := make([]bool, n)

	for i := 0; i < len(checked); i++ {
		checked[i] = false
	}

	rootSquare := int(math.Sqrt(float64(len(checked))))

	for i := 2; i <= rootSquare; i++ {
		if checked[i] == false {
			for k := i * 2; k < n-2; k += i {
				checked[k] = true;
			}
		}
	}

	for k := 2; k < len(checked); k++ {
		if checked[k] == false {
			buffer.WriteString(strconv.Itoa(k)  + "  " )
		}
	}
	fmt.Println(buffer.String())

}
