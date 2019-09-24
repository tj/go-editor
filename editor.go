// Package editor provides a utility to from $EDITOR.
package editor

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

// editor is the default editor used when $EDITOR is not assigned.
var editor = "vim"

func init() {
	if s := os.Getenv("EDITOR"); s != "" {
		editor = s
	}
}

// Read opens the default editor and returns the value.
func Read() ([]byte, error) {
	return ReadEditor(editor)
}

// ReadEditor opens the editor and returns the value.
func ReadEditor(editor string) ([]byte, error) {
	// tmpfile
	f, err := ioutil.TempFile("", "go-editor")
	if err != nil {
		return nil, errors.Wrap(err, "creating tmpfile")
	}
	defer os.Remove(f.Name())

	// open editor
	cmd := exec.Command("sh", "-c", editor+" "+f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return nil, errors.Wrap(err, "executing")
	}

	// read tmpfile
	b, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return nil, errors.Wrap(err, "reading tmpfile")
	}

	return b, nil
}
