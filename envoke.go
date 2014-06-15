package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
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

var EMPTY_STRING_FUNC func() string

var REGEXP_FUNCTION_NOT_DEFINED_ERROR *regexp.Regexp

func init() {
	REGEXP_FUNCTION_NOT_DEFINED_ERROR = regexp.MustCompile(`function "([^"]+)" not defined$`)

	EMPTY_STRING_FUNC = func() string {
		return ""
	}
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
		return c.envokeReader(os.Stdin, env)
	} else {
		return c.envokeFile(env)
	}
}

func (c config) envokeReader(stdin io.Reader, env environment) error {
	buf, err := ioutil.ReadAll(stdin)

	if nil != err {
		return err
	}

	return c.envokeString(string(buf), env)
}

func (c config) envokeFile(env environment) error {
	fileBytes, err := ioutil.ReadFile(c.filename)
	if nil != err {
		return err
	}

	fileText := string(fileBytes)

	return c.envokeString(fileText, env)
}

func (c config) envokeString(templateStr string, env environment) error {
	var err error
	// Convert env to funcs, to avoid .VAR syntax
	envFuncs := make(template.FuncMap, len(env))
	for name, val := range env {
		envFuncs[name] = (func(str string) func() string {
			return func() string {
				return str
			}
		})(val)
	}

addEnvFuncs:

	t := template.New("envokeTemplate")
	t.Funcs(envFuncs)
	t.Delims(c.leftDelim, c.rightDelim)
	t, err = t.Parse(templateStr)
	if nil != err {
		missingFunctionNameMatch := REGEXP_FUNCTION_NOT_DEFINED_ERROR.FindStringSubmatch(err.Error())
		if missingFunctionNameMatch != nil {
			if !c.strict {
				// Recur!
				envFuncs[missingFunctionNameMatch[1]] = EMPTY_STRING_FUNC
				goto addEnvFuncs
			} else {
				// Non-strict mode. Fail with missing variable
				return fmt.Errorf(`Undefined variable: "%s"`, missingFunctionNameMatch[1])
			}
		}
		return err
	}

	var buff bytes.Buffer
	err = t.Execute(&buff, env)

	if nil != err {
		return err
	}

	output := buff.String()
	fmt.Fprint(c.output, output)
	return nil
}
