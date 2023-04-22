package pkg

import (
	"fmt"
	tg "github.com/Syfaro/telegram-bot-api"
)

type Bot struct {
	Api *tg.BotAPI
}

func NewBot(tgToken string) (*Bot, error) {
	api, err := tg.NewBotAPI(tgToken)
	if err != nil {
		return nil, err
	}
	return &Bot{
		Api: api,
	}, nil
}

func (bot *Bot) Run() error {
	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.Api.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	go func() {
		for update := range updates {
			if update.Message == nil {
				continue
			}

			chatID := update.Message.Chat.ID
			if update.Message.Photo != nil {
				photo := *update.Message.Photo
				bot.handleUpscaleRequest(chatID, photo[len(photo)-1].FileID)
				continue
			}

			switch update.Message.Text {
			case "/start":
				bot.handleStart(chatID)
			default:
				bot.handleDefault(chatID)
			}
		}
	}()
	return nil
}

func (bot *Bot) handleUpscaleRequest(chatID ChatID, fileID FileID) {
	url, err := bot.Api.GetFileDirectURL(fileID)
	if err != nil {
		bot.sendMessage(chatID, "Upscaling... please wait")
	}
	fmt.Println(url)
	bot.sendMessage(chatID, "Done!")
}

func (bot *Bot) handleStart(chatID ChatID) {
	bot.sendMessage(chatID, "Hi, let's upscale your image")
}

func (bot *Bot) handleDefault(chatID ChatID) {
	bot.sendMessage(chatID, "Unknown command...")
}

func (bot *Bot) sendMessage(chatID ChatID, text string) {
	msg := tg.NewMessage(chatID, text)
	_, err := bot.Api.Send(msg)
	if err != nil {
		fmt.Printf("Failed to send message: %s", err.Error())
	}
}
