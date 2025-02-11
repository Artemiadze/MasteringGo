package main

import (
	"fmt"
	"os"
	"strconv" //Пакет strconv в Go предоставляет функции для преобразования строк в примитивные типы и наоборот.
)

/*Напишите Go-программу, которая вычисляет сумму всех аргументов, переданных в  командной
строки, которые являются действительными числами.*/

func main() {
	var sum float64 = 0

	for _, arg := range os.Args[1:] {
		num, err := strconv.ParseFloat(arg, 64)
		if err == nil {
			sum += num
		} else {
			fmt.Printf("Ignoring non-numeric argument: %s\n", arg)
		}
	}

	fmt.Printf("Total sum: %f\n", sum)
}
