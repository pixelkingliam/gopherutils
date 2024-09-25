//go:build linux || darwin || unix || freebsd

package osdep

import (
	"syscall"
)

func GetKernelVersion() (string, error) {
	var uts syscall.Utsname
	if err := syscall.Uname(&uts); err != nil {
		return "", err
	}
	release := make([]byte, len(uts.Release))
	for i, v := range uts.Release {
		release[i] = byte(v)
	}
	return string(release), nil
}
