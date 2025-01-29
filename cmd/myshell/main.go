package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

		match_command(vals)
	}
}

func match_command(vals []string)   {

	switch vals[0] {
	case "type":
		args := vals[1]
		if args == "echo" || args == "exit" || args == "type" {
			fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", vals[1])
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
