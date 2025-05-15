package main

import (
	"fmt"
	"io"
	"os"
)

type fileReader struct {
	path string
	data []byte
	pos  int
}

type dataWriter struct{}

func NewFileReader(fr *fileReader,fileName string) error {
	fr.path = fileName
	fr.pos = 0

	var er error
	fr.data, er = os.ReadFile(fr.path)

	return er
}


func main() {
	fr := fileReader{}
	dw := dataWriter{}
	argsList := os.Args

	er := NewFileReader(&fr, argsList[len(argsList)-1])
	if er != nil {
		fmt.Println("Error while reading file:",er)
		os.Exit(1)
	}

	_,er = io.Copy(dw,&fr)
    if er != nil {
		fmt.Println("Error while writing file:",er)
		os.Exit(1)
	}
	
}


func (fr *fileReader) Read(bs []byte) (int, error) {

	if fr.pos >= len(fr.data) {
		return 0, io.EOF
	}

	n := copy(bs,fr.data[fr.pos:])
	fr.pos += n

	if fr.pos >= n {
		return n, io.EOF
	}

	return n, nil
}

func (dataWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
