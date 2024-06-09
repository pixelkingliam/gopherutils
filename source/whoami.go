package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	res, err := user.Current()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(res.Name)
}
