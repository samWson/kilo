package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/sys/unix"
)

func enableRawMode() {
	termios, err := unix.IoctlGetTermios(unix.Stdin, unix.TIOCGETA)
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	termios.Lflag = termios.Lflag &^ unix.ECHO

	err = unix.IoctlSetTermios(unix.Stdin, unix.TIOCSETA, termios)
	if err != nil {
		fmt.Println("Error: failed to set raw mode")
		os.Exit(1)
	}
}

func main() {
	enableRawMode()

	for {
		c := make([]byte, 1)

		bytesRead, err := os.Stdin.Read(c)
		if bytesRead != 1 {
			break
		}

		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println("Error")
			os.Exit(1)
		}

		if string(c) == "q" {
			break
		}
	}

	os.Exit(0)
}
