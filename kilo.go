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
			os.Exit(0)
		}

		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			}

			fmt.Println("Error")
			os.Exit(1)
		}

		fmt.Printf("%v", (string(c)))
	}
}
