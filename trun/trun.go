package trun

import (
	"os/exec"
	"time"
	"bufio"
	"strings"
	"errors"
)

// Command.Code
var (
	DONE = 0
	TIMEOUT = 1
	ERROR = -1
)

type Command struct{
	// output container
	output []string

	// command to executed
	Command string
	// timeout process
	Timeout int

	// exit code
	Code int
	// exit message
	Message string
}

// return all output
func (this *Command) Output() string {
	return strings.Join(this.output, "\n")
}

// run command with args
// args will passing into exec.Command
func (this *Command) Run(args ...string) error {
	// checking for empty command
	if len(this.Command) < 1 {
		msg := "Empty command"
		this.Code = ERROR
		this.Message = msg
		return errors.New(msg)
	}

	// checking for timeout
	if this.Timeout < 1 {
		msg := "Timeout cannot below than 1"
		this.Code = ERROR
		this.Message = msg
		return errors.New(msg)
	}

	// checking for command
	path, errPath := exec.LookPath(this.Command)
	if errPath != nil {
		this.Code = ERROR
		return errors.New("Command not found")
	}

	// TODO: get cmd.StderrPipe
	var err error
	done := make(chan bool, 1)
	go func(){
		cmd := exec.Command(path, args...)
		stdout, errr := cmd.StdoutPipe()
		if err != nil {
			err = errr
			this.Code = ERROR
			this.Message = "Error while stdout"
		}

		if err = cmd.Start(); err != nil {
			this.Code = ERROR
			this.Message = "Error while start"
		}

		// collect all output
		in := bufio.NewScanner(stdout)
		for in.Scan() {
			this.output = append(this.output, in.Text())
		}
		err = cmd.Wait()

		done <- true
	}()

	select {
	// process finish
	case <-done:
		this.Code = DONE
		return err
	
	// timeout
	case <-time.After(time.Second * time.Duration(this.Timeout)):
		this.Message = "TIMEOUT"
		this.Code = TIMEOUT
		return err
	}
}

