// go run findIPv4.go log.txt - запуск
package main

import (
	"bufio"         // для построчного чтения файлов
	"fmt"           // для вывода сообщений в консоль.
	"io"            // для обработки ошибок при чтении файлов.
	"net"           // для работы с IP-адресами
	"os"            // для получения аргументов командной строки и работы с файлами
	"path/filepath" // для обработки путей к файлам.
	"regexp"        // для работы с регулярными выражениями.
)

// ищет IPv4-адрес в строке с помощью регулярного выражения.
func findIP(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input) // ищет первое совпадение в строке и возвращает найденный IP-адрес.
}

func main() {
	/* 'os.Args' получает аргументы командной строки.
	Если аргументов меньше двух (то есть не указан файл), программа выводит сообщение об использовании и завершает работу.*/
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("usage: %s logFile\n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	// Открытие и чтение файла
	for _, filename := range arguments[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			os.Exit(-1)
		}
		defer f.Close()

		// Построчное чтение файла
		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
				break
			}

			// Извлечение и проверка IP-адресов
			ip := findIP(line)
			trial := net.ParseIP(ip) // проверяет корректность IP-адреса.
			if trial.To4() == nil {
				continue
			}
			fmt.Println(ip)
		}
	}
}
