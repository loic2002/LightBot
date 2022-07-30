package files

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func ReadFile(filename string) ([]byte, error) {

	if len(filename) == 0 {
		return nil, errors.New("the filename is empty")
	}

	file, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return file, err
}