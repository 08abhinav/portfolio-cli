package utils

import "fmt"

func ClearScr(){
	fmt.Print("\033[H\033[2J")
}

// Two escape sequence meaning:
// \033[H -> move cursor to the top-left
// \033[2J -> clear screen

func WhoAMI(){
	fmt.Println("Abhinav Negi - devops engineer")
}