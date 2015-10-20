package trun

import (
	"testing"
)

func TestEmptyCommand(t *testing.T) {
	cmd := new(Command)
	err := cmd.Run()

	if err == nil {
		t.Error("Expected run error", err)
	}

	if cmd.Code != ERROR {
		t.Error("Code must return error", cmd.Code)
	}

	if cmd.Message != "Empty command" {
		t.Error("Expected message is 'Empty command'", cmd.Message)
	}
}

func TestInvalidTimeout(t *testing.T) {
	cmd := new(Command)
	cmd.Command = "sh"
	err := cmd.Run()

	if err == nil {
		t.Error("Expected run error", err)
	}
	if cmd.Code != ERROR {
		t.Error("Code must return error", cmd.Code)
	}
	if cmd.Message != "Invalid timeout" {
		t.Error("Expected message is 'Empty command'", cmd.Message)
	}
}

func TestTimeoutCommand(t *testing.T) {
	cmd := new(Command)
	cmd.Command = "sh"
	cmd.Timeout = 1
	err := cmd.Run("-c", "ping google.com")

	if err != nil {
		t.Error("Expected err nil", err)
	}
	if cmd.Code != TIMEOUT {
		t.Error("Code must TIMEOUT", cmd.Code)
	}
	if cmd.Message != "TIMEOUT" {
		t.Error("Expected TIMEOUT message on success TIMEOUT command", cmd.Message)
	}
}

func TestDoneCommand(t *testing.T) {
	cmd := new(Command)
	cmd.Command = "sh"
	cmd.Timeout = 2
	err := cmd.Run("-c", "ping -c 2 google.com")

	if err != nil {
		t.Error("Expected err nil", err)
	}
	if cmd.Code != DONE {
		t.Error("Code must DONE", cmd.Code)
	}
	if cmd.Message != "" {
		t.Error("Expected empty message on success DONE command", cmd.Message)
	}
}

func TestInvalidArgsCommand(t *testing.T) {
	cmd := new(Command)
	cmd.Command = "sh"
	cmd.Timeout = 2
	err := cmd.Run("-c", "unknown command on sh")

	if err == nil {
		t.Error("Expected error got nil", err)
	}

	// code is DONE because error on args on command
	if cmd.Code != DONE {
		t.Error("Code must DONE", cmd.Code)
	}
	if cmd.Message != "" {
		t.Error("Expected empty message on success DONE command", cmd.Message)
	}
}

func TestMultipleCommand(t *testing.T) {
	err1 := make(chan error)
	cmd1 := new(Command)
	go func() {
		cmd1.Command = "sh"
		cmd1.Timeout = 1
		err := cmd1.Run("-c", "ping google.com")
		err1 <- err
	}()

	err2 := make(chan error)
	cmd2 := new(Command)
	go func() {
		cmd2.Command = "sh"
		cmd2.Timeout = 1
		err := cmd2.Run("-c", "ping google.com")
		err2 <- err
	}()

	checker := func(name string, err error, cmdo *Command) {
		if err != nil {
			t.Error("Expected error got nil", err)
		}

		if cmdo.Code != TIMEOUT {
			t.Error("Code must DONE", cmdo.Code)
		}
		if cmdo.Message != "TIMEOUT" {
			t.Error("Expected TIMEOUT message on success TIMEOUT command", cmdo.Message)
		}
	}

	for i := 0; i < 2; i++ {
		select {
		case err := <-err1:
			checker("cmd1", err, cmd1)
		case err := <-err2:
			checker("cmd2", err, cmd2)
		}
	}

}

func CmdRun(){
	cmd := new(Command)
	cmd.Command = "uptime"
	cmd.Timeout = 1
	cmd.Run()
}
func BenchmarkRun1SecondsTimeout(b *testing.B) {
	for n:=0 ; n<b.N; n++ {
		CmdRun()
	}
}
