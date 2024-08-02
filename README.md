# Kilo

A Go version of [antirez/kilo](https://github.com/antirez/kilo?tab=readme-ov-file).

Made from the tutorial [Build Your Own Text Editor](https://viewsourcecode.org/snaptoken/kilo/).

## TERMIOS Primer

[Wikipedia](https://en.wikipedia.org/wiki/POSIX_terminal_interface#The_termios_data_structure)

Note: all of this is in the context of macOS Sonoma using iTerm2.

### Getting a TERMIOS

The package [golang/x/sys/unix](https://pkg.go.dev/golang.org/x/sys/unix) contains low level
access to operating system primitives. TERMIOS is from the POSIX standard and
the Single UNIX Specification.

The Go function `unix.IoctlGetTermios` is roughtly equivalent to the C function
`tcgetattr`.

```
// fd: the file descriptor e.g. unix.Stdin
// req: will be a constant from the unix package e.g.
func IoctlGetTermios(fd int, req uint) (*Termios, error)
```

The `req` parameter is difficult to find any indication on what constant should
be used. What constants are available will depend on which variety of UNIX the
program is compiled on e.g. `unix.TCGETS` is not defined on macOS.

```
// macOS constants in `unix` package. I've only discovered these by seeing which
// ones successfully compile. Documentation seems non-existant.
TIOCGETA
TIOCSETA
```

https://stackoverflow.com/questions/69693105/golang-unix-tcgets-equivalent-on-mac

### Display issues

The TERMIOS state controls the behaviour of input and output. Changing it will
affect how characters are displayed at the terminal.

Just the action of getting the TERMIOS has some visual impact:

```go
func main() {
  unix.IoctlGetTermios(unix.Stdin, unix.TIOCGETA)

  fmt.Println("End")
}
```

There is the newline from `Println` but also a few extra characters. How they
are interpreted depends on the shell in use:

```
# fish shell
> go run kilo.go
End
   ⏎

# zsh
> go run kilo.go
End
   %
```

The cursor is dropping down the next row in the same column at the end of the
previous row of text.
