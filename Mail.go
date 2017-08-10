//Константиновский Борис
package main

import (
	"bufio"
	"fmt"
	"os"
)

//Функция вычисляющая количество строчек в файле
func count(name string) int {
	var i int = 0

	file, error1 := os.Open(name)

	if error1 != nil {
		return -1
	}

	f := bufio.NewReader(file)

	for {
		char, _, error2 := f.ReadRune()

		if error2 != nil {
			break
		} else if char == '\n' {
			i++
		}
	}

	i++
	file.Close()

	return i
}

func main() {
	var name string
	var num int = 0

	if len(os.Args) == 1 {
		fmt.Printf("Enter name of file!\n")
		os.Exit(1)
	}

	name = os.Args[1]

	num = count(name)

	if num == -1 {
		fmt.Printf("Cannot find file!\n")
		os.Exit(2)
	} else {
		fmt.Printf("%d", num)
	}
}
