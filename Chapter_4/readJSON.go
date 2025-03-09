package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// В этом коде Go определены структурные переменные, в которых будут храниться данные JSON.
type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func loadFromJSON(filename string, key interface{}) error {
	/*вызываем функцию
	json.NewDecoder(), чтобы создать новый декодер JSON, связанный с файлом, а за-
	тем — функцию Decode(), чтобы действительно декодировать содержимое файла
	и поместить его в нужную структуру данных.*/
	in, err := os.Open(filename)
	if err != nil {
		return err
	}

	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}

	filename := arguments[1]

	var myRecord Record
	err := loadFromJSON(filename, &myRecord)
	if err == nil {
		fmt.Println(myRecord)
	} else {
		fmt.Println(err)
	}
}
