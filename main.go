package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type ttyError struct{}

func (e ttyError) Error() string {
	return ""
}

func main() {
	usage := `envoke fills templates from environment variables.

	Usage:
		
		envoke template.go.src            Print template output, using environment variables as template context.
		envoke '<<' '>>' template.go.src  Print template output, using custom delimiters.
		envoke -f template.go.src         Print template output, without failing for missing environment variables. Default with "".
		envoke                            Print template output, using a template from stdin.
		envoke '<<' '>>' -                Print template output, using a template from stdin, with custom delimiters.
	
	Options:
		-f                                Do not fail on missing environment variables.
		-h, --help                        Print this help message.
`

	f := flag.Bool("f", false, "Do not fail on missing environment variables.")
	h := flag.Bool("h", false, "Print help.")

	flag.Parse()

	if *h {
		fmt.Println(usage)
		return
	}

	for _, a := range flag.Args() {
		if a == "" {
			log.Fatal("Do not put empty arguments.")
		}
	}

	var conf config

	if flag.NArg() == 3 {
		// Envoke template (or stdin) with custom delimiters
		conf = newConfig(flag.Arg(0), flag.Arg(1), flag.Arg(2), *f)
	} else if flag.NArg() == 2 {
		log.Fatal("Only 1 or 3 args, nothing in between.")
	} else if flag.NArg() == 1 {
		// Envoke template file (or stdin) with default delimiters
		conf = newConfig("", "", flag.Arg(0), *f)
	} else if flag.NArg() == 0 {
		// Envoke stdin with default delimiters
		conf = newConfig("", "", "-", *f)
	}

	err := conf.envoke(getEnvironment(os.Environ()))
	switch err := err.(type) {
	case ttyError:
		fmt.Println(usage)
		return
	default:
		log.Fatal(err)
	}
}
