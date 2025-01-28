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
		str := input[:len(input)-1]
		vals := strings.Fields(str)

		if err != nil {
			log.Fatal(err)
		}

		switch vals[0] {
		case "echo":
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(vals[1:], " "))
		case "exit":
			fmt.Fprintf(os.Stdout, "exit")
			return
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", str)
		}
	}
}
