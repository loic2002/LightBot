package json

import (
	"encoding/json"
	"errors"
	"fmt"
)

func JsonToStruc(data []byte, data2 interface{}) (error) {

	if data == nil && len(data) != 0 {
		return errors.New("the data is empty")
	}

	if data2 == nil {
		return errors.New("the interface is empty")
	}

	err := json.Unmarshal(data, data2)
	
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	
	return nil

}

func StrucToJson(data *struct{}) ([]byte, error) {

	if data == nil {
		return nil, errors.New("the data is empty")
	}

	file, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return file, nil

}