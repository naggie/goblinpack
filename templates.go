package goblinpack

const DataFileHeader string = `package %s

var files = map[string][]byte{
`

const DataFileFooter string = "}\n"

const Decoders string = `package %s
import (
	"bytes"
	"errors"
)

func GetByteSlice(filepath string) ([]byte, error) {
	if data, ok := files[filepath]; !ok {
		return []byte{}, errors.New("Filepath not packed")
	}
}

func GetReader(filepath string) (io.Reader, error) {
	data, err := GetByteSlice(filepath)
	if err != nil {
		return err
	}

	return bytes.NewReader(data)
}
`
