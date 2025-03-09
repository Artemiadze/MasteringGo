package main

import (
	"encoding/json"
	"fmt"
)

/*
В этой части программы определяются две структуры: Record и Telephone, ко-
торые будут использоваться для хранения данных, помещаемых в запись JSON.
*/
type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func main() {
	/*этой части мы определяем переменную myRecord, которая содержит нужные
	данные. Здесь также продемонстрировано использование функции json.Marshal(),
	которая принимает ссылку на переменную myRecord*/
	myRecord := Record{
		Name:    "Mihalis",
		Surname: "Tsoukalos",
		Tel: []Telephone{Telephone{Mobile: true, Number: "1234-567"},
			Telephone{Mobile: true, Number: "1234-abcd"},
			Telephone{Mobile: false, Number: "abcc-567"},
		}}

	rec, err := json.Marshal(&myRecord)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(rec))

	/*Функция json.Unmarshal() получает входные данные JSON и преобразует их
	в структуру Go. Как и в случае с json.Marshal(), json.Unmarshal() также требует
	указатель в качестве аргумента.*/
	var unRec Record
	err1 := json.Unmarshal(rec, &unRec)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(unRec)
}
