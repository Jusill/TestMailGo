package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	fmt.Print("Please, enter name of file: ")
	fmt.Scanf("%s", &name)
	num = count(name)

	if num == -1 {
		fmt.Println("Cannot find file!")
		os.Exit(2)
	} else {
		fmt.Println(num)
	}
}
