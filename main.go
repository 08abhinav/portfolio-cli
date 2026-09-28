package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/08abhinav/portfolio-cli/utils"
)

func main() {
	fs := utils.NewFileSystem()
	reader := bufio.NewReader(os.Stdin)
	promptUser := utils.ColorCyan + "abhinav@portfolio" + utils.ColorReset
	promptPath := utils.ColorPurple + fs.Currentpath + utils.ColorReset
	promptSymbol := utils.ColorBold + utils.ColorBlue + "$" + utils.ColorReset
	for {
		fmt.Printf("%s:%s%s ", promptUser, promptPath, promptSymbol)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Input error: %v\n", err)
			break
		}

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "whoami":
			utils.WhoAMI()

		case "pwd":
			fmt.Println(fs.Currentpath)

		case "cd":
			target := "~"
			if len(args) > 1 {
				target = args[1]
			}

			if err := fs.ChangeDirectory(target); err != nil {
				fmt.Println(err)
			}

		case "ls":
			target := ""
			if len(args) > 1 {
				target = args[1]
			}
			fs.ListHelper(target)

		case "clear":
			utils.ClearScr()

		case "tree":
			target := ""
			if len(args) > 1{
				target = args[1]
			}
			fs.TreeHelper(target)
		case "exit":
			return

		default:
			fmt.Printf("command not found: %s\n", cmd)
		}
	}
}