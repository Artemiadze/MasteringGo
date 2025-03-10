package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

/*
Функция создает переменную кодировщика JSON с именем encodeJSON;
эта переменная связана с именем файла,
в который будут помещаться данные. Вызов функции Encode() кодирует данные
и сохраняет их в нужный файл.
*/
func saveToJSON(filename *os.File, key interface{}) {
	encodeJSON := json.NewEncoder(filename)
	err := encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

/*
В этом коде определена структурная переменная, в которой хранятся данные,
которые нужно сохранить в формате JSON с помощью функции saveToJSON().
Поскольку используется os.Stdout, то данные выводятся на экран, а не сохра-
няются в файл
*/
func main() {
	myRecord := Record{
		Name:    "Mihalis",
		Surname: "Tsoukalos",
		Tel: []Telephone{Telephone{Mobile: true, Number: "1234-567"},
			Telephone{Mobile: true, Number: "1234-abcd"},
			Telephone{Mobile: false, Number: "abcc-567"},
		},
	}

	saveToJSON(os.Stdout, myRecord)
}
