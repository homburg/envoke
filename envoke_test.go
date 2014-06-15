package main

import (
	"bytes"
	"fmt"
	. "github.com/homburg/goconvey/convey"
	"log"
	"testing"
)

func ExampleEnvoke() {
	environment := environment{
		"USER": "thomas",
	}
	err := newConfig("", "", "test.txt.src", true).
		envoke(environment)

	if nil != err {
		log.Fatal(err)
	}
	// Output:
	// Here is me: thomas!
}

func ExampleEnvokeCustomDelimiters() {
	environment := environment{
		"USER": "thomas",
	}
	err := newConfig("[[", "]]", "test_square_brace.txt.src", true).
		envoke(environment)

	if nil != err {
		log.Fatal(err)
	}
	// Output:
	// Here is {{me}}: thomas?
}

func ExampleNonStrictEnvoke() {
	err := newConfig("", "", "test.txt.src", false).
		envoke(environment{})

	if nil != err {
		log.Fatal(err)
	}
	// Output:
	// Here is me: !
}

func TestFailInNonStrictEnvoke(t *testing.T) {
	template := "test.txt.src"

	Convey("Given a envoke template with in strict mode", t, func() {
		conf := newConfig("", "", template, true)

		Convey("When envoking with an empty environment", func() {
			buff := new(bytes.Buffer)
			conf.output = buff
			err := conf.envoke(environment{})

			Convey("We should get an error", func() {
				So(err, ShouldResemble, fmt.Errorf(`Undefined variable: "USER"`))
			})
		})
	})
}
