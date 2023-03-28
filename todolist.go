package main

import (
	"fmt"
)

func main() {
	fmt.Println("	Todolist");
	fmt.Println("---------------------------")
	//list := make([]string, 0)
}

func Add(list []string, ch <-chan string) {
	list = append(list, <-ch)
}