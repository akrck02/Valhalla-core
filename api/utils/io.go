package utils

import (
	"errors"
	"os"
)

// ReadFile reads a file and returns its content
//
// [param] filepath | string : path to the file
// [param] content | []byte : content of the file
//
// [return] []byte : content of the file
func SaveFile(filepath string, content []byte) error {

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer file.Close()
	n, err := file.Write(content)

	if err != nil {
		return err
	}

	if n != len(content) {
		return errors.New("error writing file")
	}

	return nil
}

// ReadFile reads a file and returns its content
//
// [param] filepath | string : path to the file
// [param] content | []byte : content of the file
//
// [return] []byte : content of the file
func ReadFile(filepath string) ([]byte, error) {

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileInfo, err := file.Stat()

	if err != nil {
		return nil, err
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)

	if err != nil {
		return nil, err
	}

	return buffer, nil
}
