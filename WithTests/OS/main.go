package main

import (
	"fmt"
	"runtime"
)

func return_OS(in string) {
	switch in {
	case "windows":
		fmt.Println("You are using Windows")
	case "linux":
		fmt.Println("You are using Linux")
	case "darwin":
		fmt.Println("You are usig MacOS")
	}
}

func main() {
	os := runtime.GOOS
	return_OS(os)
}
