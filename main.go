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
		
		if command == ""{
			continue
		}else if strings.HasPrefix(command, "whoami"){
			utils.WhoAMI()
		}else if strings.HasPrefix(command, "pwd"){

			fmt.Println(fs.Currentpath)

		}else if strings.HasPrefix(command, "cd"){

			target := strings.TrimSpace(command[2:])

			if err := fs.ChangeDirectory(target); err != nil{
				fmt.Println(err)
			}

			fmt.Println("Current Directory: ", fs.Currentpath)
		}else if strings.HasPrefix(command, "ls"){

			fs.ListHelper(strings.TrimSpace(command[2:]))

		}else if command == "clear"{
			
			utils.ClearScr()

		}else if command == "exit"{
			break
		}
	}
}