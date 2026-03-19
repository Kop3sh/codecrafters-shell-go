package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	for {
		// Wait for user input
		fmt.Fprint(os.Stdout, "$ ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		str := strings.TrimSpace(input)
		vals := strings.Split(str, " ")



		p := os.Getenv("PATH") 
		match_command(vals, p)
	}
}

func match_command(vals []string, env_path string)   {

	switch vals[0] {
	case "type":
		args := vals[1]
		

		if args == "echo" || args == "exit" || args == "type" || args == "pwd" {
			fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", vals[1])
		} else if fp, err := exec.LookPath(args); err == nil {
			fmt.Fprintf(os.Stdout, "%s is %s\n", args, fp)
		} else {
			fmt.Fprintf(os.Stdout, "%s: not found\n", args)
		}
	case "echo":
		if len(vals) != 0 {
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(vals[1:], " "))
		}
	case "pwd":
		var wd, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(os.Stdout, "%s\n", wd)
	case "exit":
		os.Exit(0)
	default:
		cmd := exec.Command(vals[0], vals[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stdout, "%s: not found\n", vals[0])
		}
	}
}
