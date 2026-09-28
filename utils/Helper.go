package utils

import "fmt"

func ClearScr(){
	fmt.Print("\033[2J\033[H\033[3J")
}

// Two escape sequence meaning:
// \033[H -> move cursor to the top-left
// \033[2J -> clear screen

func WhoAMI() {
	const (
		ColorBold    = "\033[1m"
		ColorReset   = "\033[0m"
		ColorCyan    = "\033[38;2;80;250;123m"
		ColorPurple  = "\033[38;2;189;147;249m"
	)
	fmt.Printf("%s%sAbhinav Negi%s - %sDevOps Engineer%s\n", ColorBold, ColorCyan, ColorReset, 
		ColorPurple, ColorReset, )
}