package main

import (
	"os"
	"github.com/naggie/goblinpack"
	"fmt"
)

// NOTE could use buffer instead of file, this would enable text/template support
// could even do the whole thing with a template https://stackoverflow.com/questions/25173549/go-templates-range-over-string

func main() {
	targetFile, _ := os.Create("/tmp/testpack.go")
	defer targetFile.Close()

	_, err := fmt.Fprintf(targetFile, goblinpack.DataFileHeader, "testpackage")
	checkErr(err)

	for _, filepath := range os.Args[1:] {

		_, err := fmt.Fprintf(targetFile, `    "%s": `, filepath)
		checkErr(err)

		sourceFile, _ := os.Open(filepath)
		goblinpack.WriteLiteralByteSlice(sourceFile, targetFile)
		sourceFile.Close()

		_, err = targetFile.WriteString(",\n")
		checkErr(err)
	}

	_, err = targetFile.WriteString(goblinpack.DataFileFooter)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
