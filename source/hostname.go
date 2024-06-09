package main

import "os"

func main() {
	res, err := os.Hostname()
	if err != nil {
		os.Exit(1)
	}
	println(res)

}
