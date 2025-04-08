package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"log"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func findExec(command string)bool{
	if _, err := exec.LookPath(command); err == nil {
        return true
		}
	return false 
}

func search(command string, builtins []string) string {
	command =strings.TrimSpace(command)
	for _, builtin := range builtins {
		if builtin == command {
			return fmt.Sprintf("%s is a shell builtin", strings.TrimSpace(builtin))
		}
	}

	if path, err := exec.LookPath(command); err == nil {
        return fmt.Sprintf("%s is %s", command, path)
		} 
	return fmt.Sprintf("%s: not found",strings.TrimSpace(command))
}

func main() {
	var builtins = []string{"exit","echo","type"}
	//fmt.Println(paths)
	// Uncomment this block to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		homwDir:=os.Getenv("HOME")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command = strings.TrimSpace(command)
		commandSlice:=strings.Split(command," ")
		if commandSlice[0]=="exit"{
			os.Exit(0)
		}else if strings.Count(command, "echo")==1 && strings.Index(command, "echo")==0{
			fmt.Println(strings.TrimSpace(command[4:]))
		}else if commandSlice[0]=="type"{
			fmt.Println(search(commandSlice[1],builtins))
		}else if findExec(commandSlice[0]){
			out, err := exec.Command(commandSlice[0]).Output()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s", out)
		}else if commandSlice[0]=="cd"{
			if len(commandSlice)==1 || commandSlice[1]=="~" {
				_=os.Chdir(homwDir)
			}else{
				err:=os.Chdir(commandSlice[1])
				if err !=nil{
					fmt.Printf("cd: %s: No such file or directory\n",commandSlice[1])
				}
			}

		}else{
			fmt.Println(command[:len(command)] + ": command not found")
		}
	}
}