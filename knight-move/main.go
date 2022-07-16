package knightmove

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//    ________________________
// H |__|__|__|__|__|__|__|__|
// G |__|__|__|__|__|__|__|__|
// F |__|__|__|__|__|__|__|__|
// E |__|__|__|__|__|__|__|__|
// D |__|__|__|__|__|__|__|__|
// C |__|__|__|__|__|__|__|__|
// B |__|__|__|__|__|__|__|__|
// A |__|__|__|__|__|__|__|__|
//    1  2  3  4  5  6  7  8

var lines = [8]string{
	"A","B","C","D","E","F","G","H",
}

var cols = [8]string{
	"1","2","3","4","5","6","7","8",
}


func main() {

	ijarr := [8]IJStruct{
		{
			I: 1,
			J: 2,
		},
		{
			I: 1,
			J: -2,
		},
		{
			I: -1,
			J: 2,
		},
		{
			I: -1,
			J: -2,
		},
		{
			I: 2,
			J: 1,
		},
		{
			I: 2,
			J: -1,
		},
		{
			I: -2,
			J: 1,
		},
		{
			I: -2,
			J: -1,
		},
	}
	
	for true {

		fmt.Println("give me a valid chess cell : ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		char := strings.ToUpper(string(line[0]))
		num := string(line[1])

		if charExists := Contains(lines, char); !charExists {
			fmt.Println("the provided line is not a valid line")
			continue
		}

		if colExists := Contains(cols, num); !colExists {
			fmt.Println("the provided line is not a valid column")
			continue
		}

		var result []string

		mainI := FindIndex(lines, char)
		mainJ := FindIndex(cols, num)

		for _, s := range ijarr {

			newI := mainI + s.I
			newJ := mainJ + s.J

			iInRange := newI >= 0 && newI <= 7
			jInRange := newJ >= 0 && newJ <= 7

			if iInRange &&  jInRange {
				newLine := lines[newI]
				newCol := cols[newJ]
				myStrng := newLine + newCol
				result = append(result, myStrng)
			}

		}

		fmt.Println("the knight can move to these squares")
		fmt.Println(result)
		
	}

	fmt.Println("hello world")
}