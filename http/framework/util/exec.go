package util

import "os"

func GetExecDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return dir + "/"
}
