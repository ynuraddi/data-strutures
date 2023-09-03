package main

import (
	"bufio"
	"os"
)

// const input = `
// *****.**
// ***.*.**
// **..****
// **..*...
// ........
// ****.***
// ...*.***
// ******..
// 1
// 3
// ***
// ..*
// ..*
// `

func main() {
	// in := strings.NewReader(input)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
}
