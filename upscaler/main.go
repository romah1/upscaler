package main

import (
	"encoding/json"
	"fmt"
	"os"
	"upscaler/base"
	"upscaler/message_queue/mq_common"
	"upscaler/message_queue/mq_receiver"
	"upscaler/tg/tg_bot"
	"upscaler/tg/tg_client"
	"upscaler/tg/tg_server"
)

func main() {
	receiver, err := mq_receiver.NewReceiver(os.Getenv("QUEUE_URL"), mq_common.QueueParams{
		Name: os.Getenv("QUEUE_NAME"),
	})
	base.CheckErr(err)

	delivery, err := receiver.Receive()
	base.CheckErr(err)

	tgClient := tg_client.NewClient(os.Getenv("TG_SERVER_ENDPOINT"))

	for msg := range delivery {
		var upscaleRequest tg_bot.MQUpscaleRequest
		err := json.Unmarshal(msg.Body, &upscaleRequest)
		if err != nil {
			fmt.Printf("malformed message in queue: %s", err.Error())
			continue
		}
		err = tgClient.PostUpscalingFailed(tg_server.Error{
			ChatID: upscaleRequest.ChatID,
			Reason: "unimplemented",
		})
		if err != nil {
			fmt.Printf("failed reach tg server: %s", err.Error())
			continue
		}
	}
}
