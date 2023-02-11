package main

import (
	"os"

	"github.com/buger/goterm"
	"golang.org/x/term"
)

func init_term() {
	goterm.Clear()
	goterm.MoveCursor(1, 1)

	// Hide cursor
	goterm.Print("\u001b[?25l")
}

func cleanup_term() {
	// Clear term
	print("\u001b[2J")
	
	// Show cursor
	print("\u001b[?25h")
}

func print_brackets(x int, y int) {
	goterm.MoveCursor((x*2)+1, y+1)
	goterm.Print("[")
	goterm.MoveCursorForward(1)
	goterm.Print("]")

	goterm.MoveCursorForward(1)

}

func clear_brackets() {
	goterm.MoveCursorBackward(4)
	goterm.Print(" ")
	goterm.MoveCursorForward(1)
	goterm.Print(" ")
	goterm.Flush()
}

func await_input() string {
	goterm.Flush()

	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	b := make([]byte, 1)
	os.Stdin.Read(b)

	return string(b[:])
}
