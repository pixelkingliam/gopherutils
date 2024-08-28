package main

import "runtime"

func main() {
	output := runtime.GOARCH
	switch runtime.GOARCH {
	case "amd64":
		output = "x86_64"
		break
	case "386":
		output = "i386"
	case "arm":
		output = "armv8l"
	case "arm64":
		output = "aarch64"
	case "ppc64":
		output = "ppc64"
	case "ppc64le":
		output = "ppc64le"
	case "s390x":
		output = "s390x"
	case "risc64":
		output = "riscv64"
	case "loong64":
		output = "loong64"
	case "mips":
		output = "mips"
	case "mipsle":
		output = "mips"
	case "mips64":
		output = "mips64"
	case "mips64le":
		output = "mips64"
	}
	println(output)
}
