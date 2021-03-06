package main

import (
	"fmt"
	"github.com/naggie/goblinpack"
	"os"
	"path"
)

// NOTE could use buffer instead of file, this would enable text/template support
// could even do the whole thing with a template https://stackoverflow.com/questions/25173549/go-templates-range-over-string

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: goblinpack <module path> <files...>")
		os.Exit(0)
	}

	name := path.Base(os.Args[1])
	moduleDir := os.Args[1]
	targetFilepath := path.Join(moduleDir, "data.go")
	decodersFilepath := path.Join(moduleDir, "decoders.go")
	dataFilepaths := os.Args[2:]

	err := os.MkdirAll(moduleDir, 0755)
	checkErr(err)

	decodersFile, _ := os.Create(decodersFilepath)
	_, err = fmt.Fprintf(decodersFile, goblinpack.Decoders, name)
	checkErr(err)
	decodersFile.Close()

	targetFile, _ := os.Create(targetFilepath)
	defer targetFile.Close()

	_, err = fmt.Fprintf(targetFile, goblinpack.DataFileHeader, name)
	checkErr(err)

	for _, filepath := range dataFilepaths {

		fmt.Println(filepath)

		_, err := fmt.Fprintf(targetFile, `	"%s": `, filepath)
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
