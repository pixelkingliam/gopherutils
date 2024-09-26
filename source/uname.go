package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"gopherutils/shared/osdep"
	"os"
	"runtime"
	"strings"
	"unicode"
)

func main() {
	var options struct {
		KernelName bool `short:"s" long:"kernel-name" description:"Prints the kernel's name'"`               // GNU Compatible
		Hostname   bool `short:"n" long:"nodename" description:"Prints the computer's network name."`        // GNU Compatible
		Release    bool `short:"r" long:"kernel-release" description:"Prints the kernel's release version."` // GNU Compatible
		BuildDate  bool `short:"v" description:"Prints the kernel's build date."`                            // GNU Compatible
		//SafeVArg
	}
	args, err := flags.ParseArgs(&options, os.Args)
	if len(args) != 0 {
		args = args[1:]

	}
	if err != nil {
		if errors.Is(err, flags.ErrHelp) {
			os.Exit(0)
		} else {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}
	//VCode
	if options.KernelName {
		kernelName := runtime.GOOS
		if kernelName == "windows" {
			kernelName = "NT Kernel"
		} else {
			// Capitalize first letter
			kernelName = fmt.Sprintf("%c%s", unicode.ToUpper(rune(kernelName[0])), kernelName[1:])
		}
		fmt.Print(kernelName)
		fmt.Print(" ")
	}
	if options.Hostname {
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Printf("Error getting hostname: %v\n", err)
		}
		fmt.Print(hostname)
		fmt.Print(" ")
	}
	if options.Release {
		kernelVer, err := osdep.GetKernelVersion()
		if err != nil {
			fmt.Printf("Error getting kernel release version: %v\n", err)
		}
		fmt.Print(kernelVer)
		fmt.Print(" ")
	}
	if options.BuildDate {
		if runtime.GOOS == "linux" {
			file, err := os.ReadFile("/proc/version")
			strFile := string(file)
			strFile = strFile[strings.Index(strFile, "#"):]
			if err != nil {
				return
			}
			fmt.Print(strFile)
		} else {
			fmt.Print("unknown")
		}
		fmt.Print("\b ")
	}
	fmt.Println("\b")
}
