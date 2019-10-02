package main

import "fmt"

var files map[string][]byte = {
	"foo" : []byte{
		// proof : size of executable increases linearly with number of bytes here. (512K entries tested)
		0xAA,
		0x55
	},
}

func main() {
	fmt.Println(len(files["foo"]))
}
