package main

import "strings"

// Convert os.Environ to map[string]string
func getEnvironment(env []string) environment {
	envMap := make(map[string]string, len(env))
	for _, str := range env {
		i := strings.SplitN(str, "=", 2)
		envMap[i[0]] = i[1]
	}
	return envMap
}
