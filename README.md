# Envoke

[![Build Status](https://travis-ci.org/homburg/envoke.svg?branch=master)](https://travis-ci.org/homburg/envoke)

`envoke` is for templating a file from environment variables with whatever delimiters you could want.

## Template syntax

`envoke` uses Go's [text/template](http://golang.org/pkg/text/template/) package for templating.
All environment variables are converted to template "functions", ie. `{{ API_KEY }}`

## Usage

```bash
# somefile.go.src
$ echo "
somefile.go.src

package main

const API_KEY = {{ API_KEY }}

func main () {
	
}" > somefile.go.src

# envoke!
$ API_KEY=xkembiy envoke somefile.go.src > somefile.go


# Use your own delimiters:
# somefile.html.src
$ echo "<html>
<head>
	<script type="text/javascript">
	  window.API_KEY = [[ API_KEY ]];
	</script>
</head>
<body>
	<h1>{{ hello }} angular!</h1>
</body>
</html>" > somefile.html

$ API_KEY=seeeecret envoke "[[" "]]" somefile.html.src > somefile.html


# Fail on missing variables
$ envoke "[[" "]]" somefile.html || echo 'Nope!'
# ->
# Nope!

# or not
$ envoke -f "[[" "]]" somefile.html || echo 'Nope!'
# ->
# <head>
# ...


# Envoke from stdin
$ cat "Go {{ HOME }}" | envoke
```

## TODO

- [x] Stdin
- [x] Non-strict/force mode
- [ ] Test for
  - [x] Fail on strict mode
  - [ ] Help message

