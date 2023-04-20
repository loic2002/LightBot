package main

import (
	"flag"
	"fmt"

	"github.com/loic2002/LightBot/internal/config"
	"github.com/loic2002/LightBot/internal/inits"
)

const defaultConfigName = "./config/config.json"

var flagConfig = flag.String("c", defaultConfigName, "The location of the config file.")


// Main() is the entry point of the program
func main() {

	flag.Parse()


	fmt.Println("Program started !")
	fmt.Println("Version 1 !")

	fmt.Println("Read file config")
	err := config.ReadConfig(*flagConfig)

	if err != nil {
		panic(err)
	}

	fmt.Println("Start bot !")
	inits.InitDiscordBot()

	fmt.Println("Program end !")


}
