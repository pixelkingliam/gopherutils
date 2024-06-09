package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
)

func ttyname(fd uintptr) (string, error) {
	path := fmt.Sprintf("/proc/self/fd/%v", fd) // Linux-specific
	tty, err := os.Readlink(path)
	if err != nil {
		return "", err
	}
	return tty, nil
}
func main() {
	if term.IsTerminal(int(os.Stdin.Fd())) {
		result, err := ttyname(os.Stdin.Fd())
		if err != nil {
			fmt.Println("Unexpect error!", err.Error())
		}
		fmt.Println(result)
	} else {
		fmt.Println("Standard input is not a terminal.")
	}

}
