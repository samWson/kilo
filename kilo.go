package main

import (
	"fmt"
	"io"
	"os"
	"unicode"

	"golang.org/x/sys/unix"
)

type rawModeError string

var originalTermios *unix.Termios

func (r rawModeError) Error() string {
	return string(r)
}

func enableRawMode() error {
	termios, err := unix.IoctlGetTermios(unix.Stdin, unix.TIOCGETA)
	if err != nil {
		return rawModeError("Failed to GET TERMIOS for raw mode")
	}

	originalTermios = termios

	termios.Lflag = termios.Lflag &^ (unix.ECHO | unix.ICANON)

	err = unix.IoctlSetTermios(unix.Stdin, unix.TIOCSETA, termios)
	if err != nil {
		return rawModeError("Failed to SET TERMIOS for raw mode")
	}

	return nil
}

func disableRawMode() {
	unix.IoctlSetTermios(unix.Stdin, unix.TIOCSETA, originalTermios)
}

func main() {
	err := enableRawMode()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer disableRawMode()

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

		if unicode.IsControl(rune(c[0])) {
			fmt.Printf("%d\n", c[0])
		} else {
			fmt.Printf("%d ('%c')\n", c[0], c[0])
		}
	}

	os.Exit(0)
}
