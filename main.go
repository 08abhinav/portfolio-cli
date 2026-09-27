package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"github.com/08abhinav/portfolio-cli/utils"
)
func main(){
	fs := utils.NewFileSystem();
	for{
		fmt.Print("abhinav@portfolio$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil{
			fmt.Errorf("Input error: %w", err)
		}
		
		command = strings.TrimSpace(command)
		switch command{
		case "":
			continue
		case "pwd":
			fmt.Println(fs.Currentpath)
			break
		
		case "clear":
			utils.ClearScr()
			break
		case "exit":
			return
		}
	}
}