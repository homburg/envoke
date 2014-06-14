package main

import (
	"fmt"
	"io"
	"os"
	"text/template"
)

type environment map[string]string

type config struct {
	leftDelim  string
	rightDelim string
	filename   string
	strict     bool
	output     io.Writer
}

func newConfig(lDelim, rightDelim, filename string, strict bool) config {
	return config{
		lDelim,
		rightDelim,
		filename,
		strict,
		os.Stdout,
	}
}

func (c config) envoke(env environment) error {
	if c.filename == "-" {
		return c.envokeStdin(env)
	} else {
		return c.envokeFile(env)
	}
}

func (c config) envokeFile(env environment) error {
	t, err := template.ParseFiles(c.filename)
	if nil != err {
		return err
	}

	return t.Execute(c.output, env)
}

func (c config) envokeStdin(env environment) error {
	// Line-buffering?
	return fmt.Errorf("Stdin not supported yet")
}
