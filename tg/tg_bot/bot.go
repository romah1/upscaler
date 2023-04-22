package tg_bot

import (
	"encoding/json"
	"fmt"
	tg "github.com/Syfaro/telegram-bot-api"
	"github.com/rabbitmq/amqp091-go"
	"upscaler/message_queue/mq_common"
	"upscaler/message_queue/mq_publisher"
)

type Bot struct {
	api         *tg.BotAPI
	mqPublisher *mq_publisher.Publisher
}

func NewBot(tgToken string, queueUrl string, queueName string) (*Bot, error) {
	api, err := tg.NewBotAPI(tgToken)
	if err != nil {
		return nil, err
	}

	mqPublisher, err := mq_publisher.NewPublisher(queueUrl, mq_common.QueueParams{
		Name: queueName,
	})
	if err != nil {
		return nil, err
	}

	return &Bot{
		api:         api,
		mqPublisher: mqPublisher,
	}, nil
}

func (bot *Bot) Run() error {
	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.api.GetUpdatesChan(u)
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

func (bot *Bot) CloseConnections() {
	bot.mqPublisher.Close()
}

func (bot *Bot) SendMessage(chatID ChatID, text string) error {
	msg := tg.NewMessage(chatID, text)
	_, err := bot.api.Send(msg)
	return err
}

func (bot *Bot) handleUpscaleRequest(chatID ChatID, fileID FileID) {
	url, err := bot.api.GetFileDirectURL(fileID)
	if err != nil {
		bot.sendInternalServerError(chatID)
		return
	}

	request := MQUpscaleRequest{
		ChatID:   chatID,
		ImageUrl: url,
	}

	body, err := json.Marshal(request)
	if err != nil {
		bot.sendInternalServerError(chatID)
		return
	}

	err = bot.mqPublisher.Publish(amqp091.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	if err != nil {
		bot.sendInternalServerError(chatID)
		return
	}

	bot.sendMessage_(chatID, "Upscaling... please wait")
}

func (bot *Bot) handleStart(chatID ChatID) {
	bot.sendMessage_(chatID, "Hi, let's upscale your image")
}

func (bot *Bot) handleDefault(chatID ChatID) {
	bot.sendMessage_(chatID, "Unknown command...")
}

func (bot *Bot) sendMessage_(chatID ChatID, text string) {
	err := bot.SendMessage(chatID, text)
	if err != nil {
		fmt.Printf("Failed to send message: %s", err.Error())
	}
}

func (bot *Bot) sendInternalServerError(chatID ChatID) {
	bot.sendMessage_(chatID, "Internal Server Error")
}
