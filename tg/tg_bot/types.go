package tg_bot

type ChatID = int64
type FileID = string

type MQUpscaleRequest struct {
	ChatID   ChatID `json:"chat_id"`
	ImageUrl string `json:"image_url"`
}
