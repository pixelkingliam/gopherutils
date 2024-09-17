package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"gopherutils/shared/convert"
	"gopherutils/shared/gquery"
	"os"
	"strings"
)

func main() {

	var options struct {
		Zero           bool   `short:"z" long:"zero" description:"Ends each output line with NUL, instead of newline."`                                                        // GNU Compatible
		CanonMissing   bool   `short:"m" long:"canonicalize-missing" description:"Suppresses error messages associated with missing directories."`                             // GNU Compatible
		CanonExisting  bool   `short:"e" long:"canonicalize-existing" description:"Throws error if any component of the path don't exist."`                                    // GNU Compatible
		NoSymlink      bool   `short:"s" long:"strip" description:"Ignores symlinks."`                                                                                         // GNU Compatible
		NoSymlinkExtra bool   `short:"S" long:"no-symlinks" description:"Same as -s."`                                                                                         // GNU Compatible
		Physical       bool   `short:"P" long:"physical" description:"Resolves symlinks as encountered. (Default)"`                                                            // GNU Compatible
		Logical        bool   `short:"L" long:"logical" description:"Resolves '..' components before symlinks"`                                                                // GNU Compatible
		RelativeDir    string `short:"b" long:"relative-base" description:"Calculates the relative path from one directory to another"`                                        // GNU Compatible
		RelativeDirTo  string `short:"r" long:"relative-to" description:"Calculates the relative path from one directory to another. Processes symlinks in the relative path"` // GNU Compatible
		//VArg
	}
	usingRelativeBase := false
	// relative-to seems to have identical output to relative-base
	usingTo := false
	if gquery.AnyContains(convert.RunifyStrings(os.Args), convert.RunifyString("-b")) || gquery.AnyContains(convert.RunifyStrings(os.Args), convert.RunifyString("--relative-base")) {
		usingRelativeBase = true
	}
	if gquery.AnyContains(convert.RunifyStrings(os.Args), convert.RunifyString("-r")) || gquery.AnyContains(convert.RunifyStrings(os.Args), convert.RunifyString("--relative-to")) {
		usingRelativeBase = true
		usingTo = true
	}
	//options.RelativeDir = fmt.Sprintf(string(deadBeef))
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
	if usingTo {
		options.RelativeDir = options.RelativeDirTo
	}
	if options.NoSymlinkExtra {
		options.NoSymlink = true
	}
	terminator := "\n"
	if options.Zero {
		terminator = "\x00"
	}
	for _, fakePath := range args {
		final := make([]string, 0)
		var skipParentComp = false

		if fakePath[0] != '/' {
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			final = strings.Split(pwd, "/")
		}

		for i, str := range strings.Split(fakePath, "/") {
			if str == "" {
				continue
			}
			if str == ".." || str == "../" {
				if skipParentComp {
					skipParentComp = false
					continue
				}
				if i == 0 {
					continue
				}
				_, err := os.Stat(formPath(final, fakePath[len(fakePath)-1] == '/', false))
				if err != nil && !options.CanonMissing {
					fmt.Printf("Error: %v\n", err)
					return
				}
				final = final[:len(final)-1]
				continue
			}
			path := append(final, str)
			lstat, err := os.Lstat(formPath(path, fakePath[len(fakePath)-1] == '/', false))
			if err != nil {

				if options.CanonExisting {
					fmt.Printf("File or Directory '%s' does not exist!\n", formPath(path, fakePath[len(fakePath)-1] == '/', false))
					return
				}
				final = path
				continue
			}
			if lstat.Mode()&os.ModeSymlink != 0 && !options.NoSymlink {
				target, err := os.Readlink(formPath(path, fakePath[len(fakePath)-1] == '/', false))
				if err != nil {
					fmt.Printf("Error getting symlink location: %v\n", err)
					return
				}
				if strings.Split(fakePath, "/")[i+1] == ".." && options.Logical {
					skipParentComp = true
					continue
				}
				if target[0] == '/' {
					final = strings.Split(target, "/")
				} else {
					final = append(final, strings.Split(target, "/")...)
				}
			} else {
				final = path
			}
		}
		isDir := fakePath[len(fakePath)-1] == '/'
		isRelative := false
		if usingRelativeBase {
			final = relativePathBase(options.RelativeDir, final)
			isDir = false
			isRelative = true
		}
		fmt.Printf("%s%s", formPath(final, isDir, isRelative), terminator)
	}
}
func relativePathBase(dir string, path []string) []string {
	if dir[len(dir)-1] == '/' {
		dir = dir[:len(dir)-1]
	}
	dirComp := strings.Split(dir, "/")
	// Remove shared components
	i := 0
	for i < len(dirComp) && i < len(path) && dirComp[i] == path[i] {
		i++
	}
	dirCompNew := dirComp[i:] // Remaining part of the dir after the common prefix
	pathNew := path[i:]       // Remaining part of the target path after the common prefix
	// Negate <dir>
	for i, _ := range dirCompNew {
		dirCompNew[i] = ".."
	}
	return append(dirCompNew, pathNew...)
}
func formPath(components []string, isDir bool, isRelative bool) string {
	sorted := make([]string, 0)
	for _, str := range components {
		if str != "" {
			sorted = append(sorted, str)
		}
	}
	components = sorted
	result := ""
	if !isRelative {
		result = "/"
	}
	for _, component := range components {
		result += component
		result += "/"
	}
	if isDir {
		result += "/"
	}
	return result[:len(result)-1]
}
