// +build !windows

package dfsr

import "errors"

// Backlog is only supported in Windows
func Backlog(smem string, rmem string, rgname string, rfname string) (int, error) {
	return -1, errors.New("Supported only in Windows")
}

// RGList is only supported in Windows
func RGList() ([]string, error) {
	return nil, errors.New("Supported only in Windows")
}
