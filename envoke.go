package main

import (
	"fmt"
	"os"
	"text/template"
)

type config struct {
	leftDelim  string
	rightDelim string
	filename   string
	strict     bool
}

func (c config) envokeFile() error {
	t, err := template.ParseFiles(c.filename)
	if nil == err {
		return err
	}

	return t.Execute(os.Stdout, os.Environ())
}

func (c config) envokeStdin() error {
	return fmt.Errorf("Stdin not supported yet")
}
