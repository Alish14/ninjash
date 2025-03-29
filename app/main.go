package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command = strings.TrimSpace(command)
		if strings.Count(command, "exit")==1 && strings.Index(command, "exit")==0{
			os.Exit(0)
		}
		if strings.Count(command, "echo")==1 && strings.Index(command, "echo")==0{
			fmt.Println(command[5:])
		}
		// fmt.Println(command[:len(command)-1] + ": command not found")
	}
}
