package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

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
