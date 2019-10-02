package goblinpack

import (
	"bufio"
	"encoding/hex"
	"io"
	"os"
)

func WriteLiteralByteSlice(r *os.File, w *os.File) error {
	raw := make([]byte, 1)
	hexCode := make([]byte, 2)
	var err error

	// buffer reads/writes -- essential as we're reading one byte at a time which is
	// otherwise slow (100x slow!) (15000ms vs 130ms)
	rb := bufio.NewReader(r)
	wb := bufio.NewWriter(w)

	_, err = wb.WriteString("[]byte{")
	if err != nil {
		return err
	}

	for {
		_, err = rb.Read(raw)

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		hex.Encode(hexCode, raw)

		_, err = wb.WriteString("0x")
		if err != nil {
			return err
		}

		_, err = wb.Write(hexCode)
		if err != nil {
			return err
		}

		_, err = wb.WriteString(", ")
		if err != nil {
			return err
		}
	}

	// forget last comma + space for go fmt compliance
	wb.Flush()
	_, err = w.Seek(-2, 1)
	if err != nil {
		return err
	}

	_, err = w.WriteString("}")
	if err != nil {
		return err
	}


	return nil
}
