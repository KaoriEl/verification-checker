package main

import (
	"github.com/jasonlvhit/gocron"
	"main/internal/command"
)

func main() {
	gocron.Every(1).Hour().Do(command.RockNRoll)
	<-gocron.Start()
}
