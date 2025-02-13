package main

import (
	"fmt"
	"sort"
)

func main() {
	//создание срезов (2 варианта)
	aSliceLiteral := []int{1, 2, 3, 4, 5}
	integer := make([]int, 20) // пустой срез на 20 элементов (на 20 нулей) (можно расширить при желании)

	for _, i := range aSliceLiteral {
		fmt.Print(i, " ")
	}
	fmt.Println()

	aSliceLiteral = nil          // очистка существ. среза
	integer = append(integer, 1) // добавление элемента (при необходимости автоматически увеличит размер среза)
	integer = append(integer, 44)

	// доступ к элементам среза
	fmt.Println("slice integer[0]:", integer[0])                   // Первый элемент
	fmt.Println("last element of slice:", integer[len(integer)-1]) // Последний элемент
	fmt.Println("slice integer:", integer)                         //Весь срез

	s1 := make([]int, 5)
	reSlice := s1[1:3] // Создание вторичного среза
	fmt.Println("slice s1:", s1)
	fmt.Println("slice reSlice:", reSlice)
	reSlice[0] = -100
	reSlice[1] = 123456
	fmt.Println("slice s1:", s1)
	fmt.Println("slice reSlice:", reSlice)

	// cap() - помогает узнать ёмкость среза (т.е. обьём памяти)
	fmt.Printf("Cap: %d, Length: %d\n", cap(reSlice), len(reSlice))

	// copy() - копирование срезов
	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}
	fmt.Println("a6:", a6) // [-10 1 2 3 4 5]
	fmt.Println("a4:", a4) //  [-1 -2 -3 -4]
	copy(a6, a4)
	fmt.Println("a6:", a6) // [-1 -2 -3 -4 4 5]
	fmt.Println("a4:", a4) // [-1 -2 -3 -4]
	fmt.Println()

	//Сортировка срезов с помощью sort.Slice()
	fmt.Println("slice s1:", s1)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	fmt.Println("Sorted slice s1:", s1)

}
