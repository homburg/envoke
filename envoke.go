package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

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

func (c config) envoke() error {
	if c.filename == "-" {
		return c.envokeStdin()
	} else {
		return c.envokeFile()
	}
}

func getEnvironment(environment []string) map[string]string {
	envMap := make(map[string]string, len(environment))
	for _, str := range environment {
		i := strings.SplitN(str, "=", 2)
		envMap[i[0]] = i[1]
	}
	return envMap
}

func (c config) envokeFile() error {
	fmt.Println(c.filename)
	t, err := template.ParseFiles(c.filename)
	fmt.Println(os.Environ())
	if nil != err {
		return err
	}

	return t.Execute(c.output, getEnvironment(os.Environ()))
}

func (c config) envokeStdin() error {
	return fmt.Errorf("Stdin not supported yet")
}
