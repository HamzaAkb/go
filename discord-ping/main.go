package main

import (
	"fmt"

	"github.com/HamzaAkb/go/discord-ping/bot"
	"github.com/HamzaAkb/go/discord-ping/config"
)

func main(){
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}