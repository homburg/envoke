// Envoke
//
// Fill template from environment variable
// Usage:
//   # Fill with default template syntax and print to stdout
//   envoke template.file
//
//   # Fill with alternative delimiters and print to stdout
//	 envoke '<<' '>>' template.file
//
//   # Fill template from stdin
//   envoke
//
//   # Fill template from stdin
//   envoke -
//
//   # Fill with alternative template delimiters from stdin
//   envoke '<<' '>>' -
//   envoke '<<' '>>'
package main
