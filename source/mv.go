package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"golang.org/x/term"
	"io"
	"os"
	"path"
)

func main() {
	var options struct {
		NoOverwrite    bool `short:"n" long:"no-clobber" description:"Prevents overwriting existing files."` // GNU Compatible
		ForceOverwrite bool `short:"f" long:"force" description:"Do not prompt before overwriting files."`   // GNU Compatible
	}
	args, err := flags.ParseArgs(&options, os.Args)
	if err != nil {
		if errors.Is(err, flags.ErrHelp) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	args = args[1:]
	if len(args) == 0 {
		fmt.Println("Missing file operand.")
		os.Exit(1)
	}
	if len(args) == 1 {
		fmt.Println("Missing destination file operand after '" + args[0] + "'")
		os.Exit(1)
	}
	if len(args) == 2 {
		src, err := os.Open(args[0])
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				fmt.Printf("Cannot open `%s`: No such file or directory\n", args[0])
			} else {
				fmt.Println(err.Error())
			}

			os.Exit(1)
		}

		stat, err := os.Stat(args[1])
		if stat.IsDir() {
			if len(args[1]) != '/' {
				args[1] = args[1] + "/"
			}
			args[1] = args[1] + path.Base(args[0])
		}
		stat, err = os.Stat(args[1])
		if err == nil && term.IsTerminal(int(os.Stdin.Fd())) && !options.ForceOverwrite {
			if options.NoOverwrite {
				fmt.Printf("File '%s' already exists!\n", args[1])
				os.Exit(1)
			}
			fmt.Printf("File '%s' already exists! Overwrite? [Y/n]\n", args[1])
			reader := bufio.NewReader(os.Stdin)

			res, resSize, _ := reader.ReadRune()
			if !(resSize != 0 && (res == 'y' || res == 'Y')) {
				os.Exit(0)
			}
		}
		dest, err := os.Create(args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		_, err = io.Copy(dest, src)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		stat, err = os.Stat(args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = os.Chmod(args[1], stat.Mode())
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = os.Remove(args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.Exit(0)

	}
	destination := args[len(args)-1]
	stat, err := os.Stat(destination)
	if !stat.IsDir() {
		fmt.Println("Output is not a directory! Consider using only 2 operands.")
		os.Exit(1)
	}
	if len(args[1]) != '/' {
		destination = destination + "/"
	}
	for i := 0; i < len(args)-1; i++ {
		src, err := os.Open(args[i])
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				fmt.Printf("Cannot open `%s`: No such file or directory\n", args[i])
			} else {
				fmt.Println(err.Error())
			}

			os.Exit(1)
		}

		tmpDest := destination + path.Base(args[i])

		stat, err := os.Stat(tmpDest)
		if err == nil && term.IsTerminal(int(os.Stdin.Fd())) && !options.ForceOverwrite {
			if options.NoOverwrite {
				fmt.Printf("File '%s' already exists!\n", tmpDest)
				os.Exit(1)
			}
			fmt.Printf("File '%s' already exists! Overwrite? [Y/n]\n", tmpDest)
			reader := bufio.NewReader(os.Stdin)

			res, resSize, _ := reader.ReadRune()
			if !(resSize != 0 && (res == 'y' || res == 'Y')) {
				os.Exit(0)
			}
		}
		dest, err := os.Create(tmpDest)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		_, err = io.Copy(dest, src)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		stat, err = os.Stat(args[i])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = os.Chmod(tmpDest, stat.Mode())
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = os.Remove(args[i])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("Moved %s into %s\n", args[i], destination)
	}
}
