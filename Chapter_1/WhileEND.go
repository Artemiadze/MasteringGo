package main

import (
	"fmt"
	"os"
	"strconv"
)

/*Напишите Go-программу, которая считывает целые числа в командной строке до тех пор, пока
не встретит во входных данных слово END.*/

func main() {
	var sum int = 0

	for _, arg := range os.Args[1:] {
		if arg == "END" {
			break
		}

		num, err := strconv.Atoi(arg)
		if err == nil {
			sum += num
		} else {
			fmt.Printf("Ignoring non-numeric argument: %s\n", arg)
		}
	}

	fmt.Printf("Total sum: %d\n", sum)
}
