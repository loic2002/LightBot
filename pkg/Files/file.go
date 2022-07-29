package files

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func ReadFile(path, filename string) ([]byte, error) {

	if len(path) == 0 {
		return nil, errors.New("the file path is empty")
	}

	if len(filename) == 0 {
		return nil, errors.New("the filename is empty")
	}

	file, err := ioutil.ReadFile(path+filename)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return file, err
}