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
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		os.Stdout.Sync()

		input, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				os.Exit(0)
			}
			log.Fatal(err)
		}

		str := strings.TrimSpace(input)
		if str == "" {
			continue
		}
		vals := strings.Split(str, " ")

		match_command(vals)
	}
}

func match_command(vals []string)   {

	switch vals[0] {
	case "type":
		arg := vals[1]
		

		if arg == "echo" || arg == "exit" || arg == "type" || arg == "pwd" {
			fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", arg)
		} else if fp, err := exec.LookPath(arg); err == nil {
			fmt.Fprintf(os.Stdout, "%s is %s\n", arg, fp)
		} else {
			fmt.Fprintf(os.Stdout, "%s: not found\n", arg)
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
