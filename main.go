package main

import (
	"fmt"
	"goandexcel/util"
)

func main() {
	fmt.Println("BOOKS:")
	util.Books()

	fmt.Println("\nCOURSES:")
	util.Courses()
}
