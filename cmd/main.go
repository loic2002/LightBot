package main

import (
	"fmt"

	"github.com/loic2002/LightBot/internal/config"
	"github.com/loic2002/LightBot/internal/inits"
)

// Main() is the entry point of the program
func main() {

	fmt.Println("Program started !")

	fmt.Println("Read file config")
	err := config.ReadConfig()

	if err != nil {
		panic(err)
	}

	fmt.Println("Start bot !")
	inits.InitDiscordBot()

	fmt.Println("Program end !")


}