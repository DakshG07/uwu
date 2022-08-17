package utils

import (
	"runtime"
	"os"
)

func GetEnv(lookup string, fallback string) string {
	if res, ok := os.LookupEnv(lookup); ok {
		return res
	}
	return fallback
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}
