package utils

import "fmt"

func ClearScr(){
	fmt.Print("\033[H\033[2J")
}

// Note:
// Two escape sequence meaning:
// \033[H -> move cursor to the top-left
// \033[2J -> clear screen