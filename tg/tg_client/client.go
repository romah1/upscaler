package tg_client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"upscaler/tg/tg_server"
)

type Client struct {
	apiEndpoint string
}

func NewClient(endpoint string) *Client {
	return &Client{
		apiEndpoint: endpoint + "/api",
	}
}

func (client *Client) PostUpscalingFailed(body tg_server.UpscalingFailedBody) error {
	marshalledBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	url := client.apiEndpoint + "/upscaling_failed"
	_, err = http.Post(url, "application/json", bytes.NewBuffer(marshalledBody))
	return err
}

func (client *Client) PostUpscalingFinished(body tg_server.UpscalingFinishedBody) error {
	marshalledBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	url := client.apiEndpoint + "/upscaling_finished"
	_, err = http.Post(url, "application/json", bytes.NewBuffer(marshalledBody))
	return err
}
