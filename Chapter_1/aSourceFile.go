package main

import (
	"fmt"

	"github.com/mactsouk/go/simpleGitHub" // заранее скачать go get -v github.com/mactsouk/go/simpleGitHub
)

func main() {
	fmt.Println("This is my first program on Golang!")

	fmt.Println(simpleGitHub.AddTwo(5, 6))
}

//go build aSourceFile.go && ./aSourceFile
//or
//go run aSourceFile.go
