package main

import (
	"os"
	"encoding/hex"
	"io"
)

func main() {
	targetFile, _ := os.Create("testpack.go")
	defer targetFile.Close()

	for _, fp := range os.Args[1:] {
		sourceFile, _ := os.Open(fp)
		WriteLiteralByteSlice(sourceFile, targetFile)
		sourceFile.Close()
	}
}


func WriteLiteralByteSlice(r *os.File, w *os.File) error {
	raw := make([]byte, 1)
	hexCode := make([]byte, 2)
	var err error

	_, err = w.WriteString("[]byte{")
	if err != nil { return err }

	for {
		_, err = r.Read(raw)

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		hex.Encode(hexCode, raw)

		_, err = w.WriteString("0x")
		if err != nil { return err }


		_, err = w.Write(hexCode)
		if err != nil { return err }

		_, err = w.WriteString(",")
		if err != nil { return err }
	}

	_, err = w.WriteString("}")
	if err != nil { return err }

	return nil
}