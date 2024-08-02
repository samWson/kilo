package main

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func main() {
	unix.IoctlGetTermios(unix.Stdin, unix.TIOCGETA)

	fmt.Println("End")
}
