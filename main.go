package main

import (
	"fmt"
	"lem-in/file"
	"os"
)

func main() {
	files := os.Args[1]
	graph := file.ReadFile(files)
	fmt.Printf("%#v\n", graph)
}
