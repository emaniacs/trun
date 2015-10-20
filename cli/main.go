package main

import (
	"os"
	"fmt"
	"strconv"
	"github.com/emaniacs/trun"
)

func main(){
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Not enough argument.")
		fmt.Fprintf(os.Stderr, "%s <timeout> command...\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "%s 5 ping google.com\n", os.Args[0])
		os.Exit(255)
	}
	timeout, err := strconv.Atoi(os.Args[1])
	if err != nil || timeout < 1 {
		fmt.Fprintln(os.Stderr, "Invalid timeout")
		os.Exit(255)
	}
	command := os.Args[2]
	args := os.Args[3:]

	cmd := new(trun.Command)
	cmd.Command = command
	cmd.Timeout = timeout

	if err := cmd.Run(args...); err == nil {
		var status string
		switch cmd.Code {
		case trun.ERROR:
			status = "Error"
		case trun.DONE:
			status = "Done"
		case trun.TIMEOUT:
			status = "Timeout"
		}
		fmt.Println("Status:", status)
		fmt.Println("Message:", cmd.Message)
		fmt.Println("Output:")
		fmt.Println(cmd.Output())
	} else {
		fmt.Fprintln(os.Stderr, "Error on run:", err)
		os.Exit(255)
	}

	os.Exit(0)
}
