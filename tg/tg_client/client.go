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

func (client *Client) PostUpscalingFailed(body tg_server.Error) error {
	marshalledBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	url := client.apiEndpoint + "/upscaling_failed"
	_, err = http.Post(url, "application/json", bytes.NewBuffer(marshalledBody))
	return err
}
