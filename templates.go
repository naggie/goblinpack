package goblinpack

DataFileHeader := """
package %s

import (
	\"bytes\"
	\"errors\"

var files = map[string][]byte{
"""

DataFileFooter := """
}

func GetByteSlice(path string) ([]byte, error) {
	data, ok := files[path]; !ok {
		return []byte{}, errors.New(\"Filepath not packed\")
	}
}

func GetReader(path string) (io.Reader, error) {
	data, err := GetByteSlice(path)
	if err {return err}

	return bytes.NewReader(data)
}

"""
