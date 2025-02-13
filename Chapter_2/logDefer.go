//просмотр содержимого /tmp/mGo.log — журнального файла,

package main

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE = "/tmp/mGo.log"

/*
здесь defer используется для того, чтобы гаран-
тировать, что второй вызов aLog.Println() будет выполнен непосредственно
перед завершением работы функции one()
*/
func one(aLog *log.Logger) { //<- указатель на логгер
	aLog.Println("-- FUNCTION one ------")
	defer aLog.Println("-- FUNCTION one ------") //Использует defer для того, чтобы в конце выполнения функции записать строку
	// Оператор defer откладывает выполнение указанного выражения до завершения функции,
	//  в которой он находится. Это полезно для логирования входа и выхода из функции.

	for i := 0; i < 10; i++ {
		aLog.Println(i)
	}
}

func two(aLog *log.Logger) { //<- указатель на логгер
	aLog.Println("---- FUNCTION two")
	defer aLog.Println("FUNCTION two ------")

	for i := 10; i > 0; i-- {
		aLog.Println(i)
	}
}

func main() {
	/*Использует флаги:
	os.O_APPEND — добавлять новые записи в конец файла.
	os.O_CREATE — создать файл, если он не существует.
	os.O_WRONLY — открыть файл только для записи.*/
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	/*Создается новый логгер iLog, который:
	1.Записывает логи в файл f.
	2.Добавляет префикс "logDefer ".
	3.Использует стандартные флаги времени (log.LstdFlags), которые добавляют в лог дату и время.*/
	iLog := log.New(f, "logDefer ", log.LstdFlags)

	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")

	one(iLog)
	two(iLog)
}
