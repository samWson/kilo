package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/sys/unix"
)

type rawModeError string

func (r rawModeError) Error() string {
	return string(r)
}

func enableRawMode() error {
	termios, err := unix.IoctlGetTermios(unix.Stdin, unix.TIOCGETA)
	if err != nil {
		return rawModeError("Failed to GET TERMIOS for raw mode")
	}

	termios.Lflag = termios.Lflag &^ unix.ECHO

	err = unix.IoctlSetTermios(unix.Stdin, unix.TIOCSETA, termios)
	if err != nil {
		return rawModeError("Failed to SET TERMIOS for raw mode")
	}

	return nil
}

func main() {
	err := enableRawMode()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
