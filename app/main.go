package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint



func search(command string, builtins []string) string {
	command =strings.TrimSpace(command)
	for _, builtin := range builtins {
		if builtin == command {
			return fmt.Sprintf("%s is a shell builtin\n", strings.TrimSpace(builtin))
		}
	}

	if path, err := exec.LookPath(command); err == nil {
        return fmt.Sprintf("%s is %s\n", command, path)
		} 
	return fmt.Sprintf("%s: not found",strings.TrimSpace(command))
}

func main() {
	var builtins = []string{"exit","echo","type"}
	path := os.Getenv("PATH")
	paths := strings.Split(path, ":")
	fmt.Println(paths)
	// Uncomment this block to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command = strings.TrimSpace(command)
		if strings.Count(command, "exit")==1 && strings.Index(command, "exit")==0{
			os.Exit(0)
		}else if strings.Count(command, "echo")==1 && strings.Index(command, "echo")==0{
			fmt.Println(strings.TrimSpace(command[4:]))
		}else if strings.Count(command, "type")>=1 && strings.Index(command, "type")==0{
			fmt.Println(search(command[4:],builtins))
		}else{
			fmt.Println(command[:len(command)] + ": command not found")
		}
	}
}