package main

import "main/internal/command"

func main() {
	//gocron.Every(1).Hour().Do(command.RockNRoll)
	//<-gocron.Start()
	command.RockNRoll()
}
