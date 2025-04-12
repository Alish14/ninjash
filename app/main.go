package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"log"
)

var _ = fmt.Fprint 	

func findExec(command string) bool {
	if _, err := exec.LookPath(command); err == nil {
		return true
	}
	return false
}

func search(command string, builtins []string) string {
	command = strings.TrimSpace(command)
	for _, builtin := range builtins {
		if builtin == command {
			return fmt.Sprintf("%s is a shell builtin", strings.TrimSpace(builtin))
		}
	}

	if path, err := exec.LookPath(command); err == nil {
		return fmt.Sprintf("%s is %s", command, path)
	}
	return fmt.Sprintf("%s: not found", strings.TrimSpace(command))
}

func main() {
	var builtins = []string{"exit", "echo", "type", "cd", "cat"}
	homeDir := os.Getenv("HOME")

	for {
		fmt.Fprint(os.Stdout, "$ ")

		reader := bufio.NewReader(os.Stdin)
		command, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			}
			log.Fatal(err)
		}

		command = strings.TrimSpace(command)
		if command == "" {
			continue
		}

		commandSlice := strings.Fields(command)
		if len(commandSlice) == 0 {
			continue
		}

		switch commandSlice[0] {
		case "exit":
			os.Exit(0)

		case "type":
			if len(commandSlice) > 1 {
				fmt.Println(search(commandSlice[1], builtins))
			} else {
				fmt.Println("type: missing argument")
			}

		case "cd":
			targetDir := homeDir
			if len(commandSlice) > 1 {
				if commandSlice[1] == "~" {
					targetDir = homeDir
				} else {
					targetDir = commandSlice[1]
				}
			}
			if err := os.Chdir(targetDir); err != nil {
				fmt.Printf("cd: %s: No such file or directory\n", targetDir)
			}

		case "echo":
			args := strings.Join(commandSlice[1:], " ")
			if len(args) > 0 {
				if (strings.HasPrefix(args, "'") && strings.HasSuffix(args, "'")) ||
					(strings.HasPrefix(args, "\"") && strings.HasSuffix(args, "\"")) {
					args = args[1 : len(args)-1]
				}
				fmt.Println(args)
			} else {
				fmt.Println()
			}

		case "cat":
			if len(commandSlice) < 2 {
				fmt.Println("cat: missing file operand")
				continue
			}
			for _, filename := range commandSlice[1:] {
				filename=strings.ReplaceAll(filename,"'","")
				filename=strings.ReplaceAll(filename,"\"","")
				file, err := os.Open(filename)
				if err != nil {
					fmt.Printf("cat: %s: No such file or directory\n", filename)
					continue
				}
				_, err = io.Copy(os.Stdout, file)
				if err != nil {
					fmt.Printf("cat: error reading %s\n", filename)
				}
				file.Close()
				fmt.Println()
			}

		default:
			if findExec(commandSlice[0]) {
				cmd := exec.Command(commandSlice[0], commandSlice[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				_= cmd.Run();
			} else {
				fmt.Printf("%s: command not found\n", commandSlice[0])
			}
		}
	}
}