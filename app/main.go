package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func search(command string, builtins []string) bool {
	for _, builtin := range builtins {
		if builtin == command {
			return true
		}
	}
	return false
}

func main() {
	var builtins = []string{"exit", "echo", "type"}
	// Uncomment this block to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command = strings.TrimSpace(command)
		if strings.Count(command, "exit")==1 && strings.Index(command, "exit")==0{
			os.Exit(0)
		}else if strings.Count(command, "echo")==1 && strings.Index(command, "echo")==0{
			fmt.Println(command[5:])
		}else if strings.Count(command, "type")==1 && strings.Index(command, "type")==0{
			if search(command[5:], builtins){
				fmt.Println("shell built-in command")
			}else{
				fmt.Println(command[:len(command)] + ": command not found")
			}
		}else{
			fmt.Println(command[:len(command)] + ": command not found")
		}
	}
}
