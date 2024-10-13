package main

import (
	"errors"
	"fmt"
	"github.com/jaypipes/ghw"
	"github.com/jessevdk/go-flags"
	"github.com/shirou/gopsutil/cpu"
	"gopherutils/shared/osdep"
	"os"
	"runtime"
	"strings"
	"unicode"
)

func main() {

	var options struct {
		KernelName       bool `short:"s" long:"kernel-name" description:"Prints the kernel's name'"`                 // GNU Compatible
		Hostname         bool `short:"n" long:"nodename" description:"Prints the computer's network name."`          // GNU Compatible
		Release          bool `short:"r" long:"kernel-release" description:"Prints the kernel's release version."`   // GNU Compatible
		BuildDate        bool `short:"v" description:"Prints the kernel's build date."`                              // GNU Compatible
		MachineArch      bool `short:"m" long:"machine" description:"Prints the computer's architecture."`           // GNU Compatible
		Processor        bool `short:"p" long:"processor" description:"Prints the computer's CPU."`                  // GNU Incompatible
		HardwarePlatform bool `short:"i" long:"hardware-platform" description:"Prints the computer's model."`        // GNU Incompatible
		OS               bool `short:"o" long:"operating-system" description:"Prints the operating system's name'."` // GNU Incompatible
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
			fmt.Print(strings.Replace(strFile, "\n", "", -1))
		} else {
			fmt.Print("unknown")
		}
		fmt.Print(" ")
	}
	if options.MachineArch {
		fmt.Print(runtime.GOARCH)
		fmt.Print(" ")
	}
	if options.Processor {
		info, err := cpu.Info()
		if err != nil {
			fmt.Printf("Error getting CPU: %v\n", err)
			return
		}

		fmt.Printf(info[0].ModelName)
		fmt.Print(" ")
	}
	if options.HardwarePlatform {
		err := os.Setenv("GHW_DISABLE_WARNINGS", "1")
		if err != nil {
			fmt.Printf("Error preparing ghw")
			return
		}

		system, err := ghw.Product()
		if err != nil {
			fmt.Printf("Error getting system info: %s\n", err)
			return
		}

		fmt.Print(strings.Replace(system.Name, "\n", "", -1))
		fmt.Print(" ")
	}
	if options.OS {
		// try getting OS from /etc/os-release
		stat, err := os.Stat("/etc/os-release")
		if err == nil {
			if !stat.IsDir() {
				file, err := os.ReadFile("/etc/os-release")
				if err != nil {
					fmt.Printf("Could not read os-release")
					return
				}
				osName := strings.Split(strings.Split(string(file), "\n")[0], "=")[1]
				osName = strings.Trim(osName, "\"")
				fmt.Print(osName)
			}
		}
		// try getting OS info from /etc/lsb-release
		stat, err = os.Stat("/etc/lsb-release")
		if err == nil {
			if !stat.IsDir() {
				file, err := os.ReadFile("/etc/lsb-release")
				if err != nil {
					fmt.Printf("Could not read lsb-release")
					return
				}
				osName := strings.Split(strings.Split(string(file), "\n")[0], "=")[1]
				osName = strings.Trim(osName, "\"")
				fmt.Print(osName)
			}
		}
		fmt.Print(" ")
	}

	fmt.Println("\b")
}
