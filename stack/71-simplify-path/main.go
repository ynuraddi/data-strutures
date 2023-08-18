package main

import "strings"

func main() {
	simplifyPath("/a/./b/../../c/")
}

func simplifyPath(path string) string {
	pathSplit := strings.Split(path, "/")

	stack := make([]string, 0, len(pathSplit)/2)
	for _, v := range pathSplit {
		if v == "" || v == "." || v == ".." && len(stack) == 0 {
			continue
		} else if v == ".." {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}

	return "/" + strings.Join(stack, "/")
}
