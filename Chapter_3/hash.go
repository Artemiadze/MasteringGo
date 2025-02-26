package main

import (
	"fmt"
)

func main() {
	// const
	const s1 = "My String"
	const s2 string = "My String"

	const (
		C1 = "C1C1C1"
		C2 = "C2C2C2"
		C3 = "C3C3C3"
	)

	//Map's
	var iMap = make(map[string]int) // Создание пустой Хеш-таблицы (map)
	iMap["k1"] = 12
	iMap["k2"] = 13
	fmt.Println("iMap:", iMap) // iMap: map[k1:12 k2:13]

	anotherMap := map[string]int{ // Создание непустой Хеш-таблицы
		"k1": 12,
		"k2": 13,
		"k3": 14,
	}

	delete(anotherMap, "k1") // удаление  anotherMap["k1"]

	// Печать таблицы
	for key, value := range anotherMap {
		fmt.Println(key, value) // Печатает красиво - поэлементно
	}

	//метод, который позволяет определить, есть ли в хеш-таблице данный ключ.
	_, ok := iMap["k1"]
	if ok {
		fmt.Println("iMap['k1'] Exists!")
	} else {
		fmt.Println("iMap['k1'] Does NOT exist")
	}
}
