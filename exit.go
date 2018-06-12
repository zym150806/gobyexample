package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("haha")
	defer fmt.Println("!")

	os.Exit(3)
}
