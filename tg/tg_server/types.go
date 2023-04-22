package tg_server

import "upscaler/tg/tg_bot"

type Error struct {
	ChatID tg_bot.ChatID `json:"chat_id"`
	Reason string        `json:"reason"`
}
