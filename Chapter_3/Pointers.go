package main

import "fmt"

func getPointer(n *int) { // Принимает указатель
	*n = *n * *n
}

func returnPointer(n int) *int {
	/*
	   получает целочисленный параметр и
	   возвращает указатель на целое число
	*/
	v := n * n
	return &v
}

/*Обе функции, getPointer() и returnPointer(), вычисляют квадрат целого
числа, однако используют совершенно разные подходы: getPointer() сохраняет
результат в переданном параметре, а returnPointer() возвращает результат, кото-
рый должен быть сохранен в другой переменной.*/

func main() {
	i := -10
	j := 25
	pI := &i
	pJ := &j
	fmt.Println("pI memory:", pI)
	fmt.Println("pI memory:", pJ)
	fmt.Println("pI value:", *pI)
	fmt.Println("pI value:", *pJ)

	*pI = 123456
	*pI--
	fmt.Println("i:", i)
	fmt.Println("j:", j)
}
