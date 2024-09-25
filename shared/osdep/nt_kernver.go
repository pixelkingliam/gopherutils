//go:build windows
// +build windows

package osdep

import (
	"fmt"
	"golang.org/x/sys/windows"
)

func GetKernelVersion() (string, error) {
	var info windows.OSVERSIONINFOEX
	info.OSVersionInfoSize = uint32(unsafe.Sizeof(info))
	err := windows.RtlGetVersion(&info)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d.%d.%d", info.MajorVersion, info.MinorVersion, info.BuildNumber), nil
}
