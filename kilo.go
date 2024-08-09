package main

import (
	"fmt"
	"os"
	"reflect"

	"golang.org/x/sys/unix"
)

func main() {
	termios, err := unix.IoctlGetTermios(unix.Stdin, unix.TIOCGETA)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	termiosType := reflect.TypeOf(termios)
	termiosValue := reflect.ValueOf(termios)

	fmt.Printf("Termios: %v\n", termios)
	fmt.Printf("Type: %v\n", termiosType)
	fmt.Printf("Value: %v\n", termiosValue.String())

	fmt.Println("")

	fmt.Println("Termios Type:")
	fmt.Printf("Alignment: %v\n",termiosType.Align())
	fmt.Printf("Field align: %v\n",termiosType.FieldAlign())
	fmt.Printf("Number of methods: %v\n",termiosType.NumMethod())
	fmt.Printf("Name: %v\n",termiosType.Name())
	fmt.Printf("Package path: %v\n",termiosType.PkgPath())
	fmt.Printf("Size: %v\n",termiosType.Size())
	fmt.Printf("String: %v\n",termiosType.String())
	fmt.Printf("Kind: %v\n",termiosType.Kind())
	fmt.Printf("Comparable: %v\n",termiosType.Comparable())

	fmt.Println("")

	fmt.Println("Termios Value:")
	fmt.Printf("Kind: %v\n",termiosValue.Kind())
	fmt.Printf("Pointer: %v\n",termiosValue.Pointer())
	fmt.Printf("Type: %v\n",termiosValue.Type())
	fmt.Printf("Unsafe pointer: %v\n", termiosValue.UnsafeAddr())

	fmt.Println("End")
}
