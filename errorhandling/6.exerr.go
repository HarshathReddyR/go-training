package main

import (
	"errors"
	"fmt"
	"os"
)

type DirectoryError struct {
	Dirname string
	Err     error
}

func (d *DirectoryError) Error() string {
	return "main." + d.Dirname + d.Err.Error()
}

//var ErrFileNotFound = errors.New("not in the root directory")

func main() {
	var d *DirectoryError
	_, err := openFile("test.txt", "root")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	if err != nil {
		if errors.As(err, &d) {
			fmt.Println("custom error found in the chain", d.Dirname)
			return
		}
	}
}

func openFile(fileName string, dirName string) (*os.File, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, &DirectoryError{
			Dirname: "errorhandling",
			Err:     err,
		}
	}
	return f, nil
}
