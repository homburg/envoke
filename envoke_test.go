package main

func ExampleEnvoke() {
	environment := environment{
		"USER": "thomas",
	}
	newConfig("", "", "test.txt.src", true).
		envoke(environment)

	// Output:
	// Here is me: thomas!
}

func ExampleNonStrictEnvoke() {
	newConfig("", "", "test.txt.src", false).
		envoke(environment{})

	// Output:
	// Here is me:
}
