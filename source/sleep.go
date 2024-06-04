package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"gopherutils/shared/gquery"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var options struct {
	}
	args, err := flags.ParseArgs(&options, os.Args[1:])
	if err != nil {
		if errors.Is(err, flags.ErrHelp) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	total := 0.0
	for _, arg := range args {
		modifier := 1
		var toParse string
		if !gquery.IsDigit(arg[len(arg)-1]) {
			toParse = arg[:len(arg)-1]
			switch strings.ToLower(arg)[len(arg)-1] {
			case 's':
				modifier = 1
				break
			case 'm':
				modifier = 60
				break
			case 'h':
				modifier = 60 * 60
				break
			case 'd':
				modifier = 60 * 60 * 24
				break
			default:
				fmt.Printf("Invalid time interval unit `%s`\n", arg)
			}

		} else {
			toParse = arg
		}
		out, err := strconv.ParseFloat(toParse, 64)
		if err != nil {
			fmt.Printf("Invalid time interval `%s`\n", arg)
			return
		}
		total += out * float64(modifier)

	}
	time.Sleep(time.Duration(total) * time.Second)
}
