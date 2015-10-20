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
