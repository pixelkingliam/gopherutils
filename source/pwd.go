package main

import (
	"fmt"
	"os"
)

func main() {
	result, _ := os.Getwd()
	fmt.Println(result)
}
