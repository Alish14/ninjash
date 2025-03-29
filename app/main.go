package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		switch command {
		case "exit 0":
			os.Exit(0)
		case "echo":
			fmt.Println(command[:len(command)-1])
		default:
			fmt.Println(command[:len(command)-1] + ": command not found")
		}

		// fmt.Println(command[:len(command)-1] + ": command not found")
	}
}
