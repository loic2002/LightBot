package config

import (
	"encoding/json"
	"fmt"

	files "github.com/loic2002/LightBot/pkg/Files"
)

var (
	Token string
	DefaultLocalization string
	Prefix string

	config *configStruct
)

type configStruct struct {
	Token string `json:"Token"`
	Prefix string `json:"Prefix"`
	DefaultLocalization string `json:"DefaultLocalization"`

}

func ReadConfig() error {

	data, err := files.ReadFile("./../config/","config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = json.Unmarshal(data, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	Prefix = config.Prefix
	DefaultLocalization = config.DefaultLocalization


	return nil

}
