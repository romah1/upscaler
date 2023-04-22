package main

import (
	"fmt"
	"os"
	"upscaler/base"
	"upscaler/tg/pkg"
)

func main() {
	bot, err := pkg.NewBot(os.Getenv("TG_TOKEN"))
	base.CheckErr(err)

	err = bot.Run()
	base.CheckErr(err)

	fmt.Println("Bot is running...")

	var forever chan struct{}
	<-forever
}
