package tg_server

import "upscaler/tg/tg_bot"

type UpscalingFailedBody struct {
	ChatID tg_bot.ChatID `json:"chat_id"`
	Reason string        `json:"reason"`
}

type UpscalingFinishedBody struct {
	ChatID tg_bot.ChatID `json:"chat_id"`
	URL    string        `json:"url"`
}
