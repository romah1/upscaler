package main

import (
	"fmt"
	"os"
	"upscaler/base"
	"upscaler/tg/tg_bot"
	"upscaler/tg/tg_server"
)

func main() {
	bot, err := tg_bot.NewBot(os.Getenv("TG_TOKEN"), os.Getenv("QUEUE_URL"), os.Getenv("QUEUE_NAME"))
	base.CheckErr(err)
	defer bot.CloseConnections()

	err = bot.Run()
	base.CheckErr(err)

	fmt.Println("Bot is running...")

	engine := tg_server.SetupGinEngine(bot)
	err = engine.Run(":8080")
	base.CheckErr(err)
}
