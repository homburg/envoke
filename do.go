package main

func do(lDelim, rightDelim, filename string, strict bool) error {
	return config{lDelim, rightDelim, filename, strict}.envoke()
}

func (c config) envoke() error {
	if c.filename == "-" {
		return c.envokeStdin()
	} else {
		return c.envokeFile()
	}
}
