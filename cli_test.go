package main

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	assert "github.com/stretchr/testify/assert"
	gl "github.com/tongson/gl"
)

func TestVersionString(t *testing.T) {
	if os.Getenv("GOTEST") == "1" {
		printVersion()
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestVersionString")
	cmd.Env = append(os.Environ(), "GOTEST=1")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	_, ok := err.(*exec.ExitError)
	assert.Equal(t, false, ok)
	expectedString := versionString + "\n"
	assert.Equal(t, expectedString, stdout.String())
}

func TestLogCreation(t *testing.T) {
	lf := "cli_test.json"
	jl := jLog(lf)
	jl.Info().Msg("test")
	assert.Equal(t, true, gl.IsFile(lf))
	os.Remove(lf)
}
