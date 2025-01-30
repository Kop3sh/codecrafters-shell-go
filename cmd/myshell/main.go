package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
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
		vals := strings.SplitN(str, " ", 2)



		p := os.Getenv("PATH") 
		match_command(vals, p)
	}
}

func match_command(vals []string, env_path string)   {

	switch vals[0] {
	case "type":
		args := vals[1]
		

		if args == "echo" || args == "exit" || args == "type" {
			fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", vals[1])
		} else if  fp, err := match_executable_command(env_path, args); err == nil  {
			fmt.Fprintf(os.Stdout, "%s is %s\n", args, fp)
		} else {
			fmt.Fprintf(os.Stdout, "%s: not found\n", args)
		}
	case "echo":
		if len(vals) != 0 {
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(vals[1:], " "))
		}
	case "exit":
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stdout, "%s: not found\n", strings.Join(vals[0:], " "))
	}
}

func match_executable_command(env_path string, command string) (string, error) {
	paths := strings.Split(env_path, ":")

	for _, p := range paths {
		fp := path.Join(p, command)
		if _, err := os.Stat(fp); err == nil {
			return fp, nil
		}
	}
	return "", errors.New("could not find executable")
}
