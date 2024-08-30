package main

import (
	"bytes"
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

	termios.Iflag = termios.Iflag &^ (unix.BRKINT | unix.ICRNL | unix.INPCK | unix.IXON)
	termios.Oflag = termios.Oflag &^ unix.OPOST
	termios.Cflag = termios.Cflag | unix.CS8
	termios.Lflag = termios.Lflag &^ (unix.ECHO | unix.ICANON | unix.IEXTEN | unix.ISIG)

	err = unix.IoctlSetTermios(unix.Stdin, unix.TIOCSETA, termios)
	if err != nil {
		return rawModeError("Failed to SET TERMIOS for raw mode")
	}

	return nil
}

func disableRawMode() {
	unix.IoctlSetTermios(unix.Stdin, unix.TIOCSETA, originalTermios)
}

func isCtrlKey(c []byte, key rune) bool {
	return bytes.Runes(c)[0] == (key & 0x1f)
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

		if isCtrlKey(c, 'q') {
			break
		}

		if unicode.IsControl(rune(c[0])) {
			fmt.Printf("%d\r\n", c[0])
		} else {
			fmt.Printf("%d ('%c')\r\n", c[0], c[0])
		}
	}

	os.Exit(0)
}
