package main

import "fmt"

func main () {
	num := 20

	var result [][]int

	//  init all arrays 
	for i := 0; i < num; i++ {
		var s []int

		for j := 0; j <= i; j++ {
			s = append(s, 0)
		}

		result = append(result, s)
	}

	// value all the border items
	for i, s := range result {
		result[i][0] = 1
		result[i][len(s) - 1] = 1
	}

	// calculate all the inner items
	for i, s := range result {

		for x := range s {

			if x != 0 && x != len(s) - 1 {
				result[i][x] = result[i - 1][ x -1] + result[i - 1][ x ]
			}
		}
	
	}

	// print the results
	for _, s := range result {

		for index, number := range s{
			var str string
			if index != len(s) -1 {
				str = "-"
			}
			fmt.Print(number, str)
		}	
		fmt.Print("\n")
	}

}