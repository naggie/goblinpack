package main

import (
	"os"
	"github.com/naggie/goblinpack"
)

// NOTE could use buffer instead of file, this would enable text/template support
// could even do the whole thing with a template https://stackoverflow.com/questions/25173549/go-templates-range-over-string

func main() {
	targetFile, _ := os.Create("testpack.go")
	defer targetFile.Close()

	for _, fp := range os.Args[1:] {
		sourceFile, _ := os.Open(fp)
		goblinpack.WriteLiteralByteSlice(sourceFile, targetFile)
		sourceFile.Close()
	}
}

