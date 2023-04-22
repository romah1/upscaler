package main

import (
	"encoding/json"
	"net/http"
)

type UpscalerClient struct {
	apiEndpoint string
}

func NewUpscalerClient(endpoint string) *UpscalerClient {
	return &UpscalerClient{
		apiEndpoint: endpoint + "/api",
	}
}

func (client *UpscalerClient) Upscale() error {
	url := client.apiEndpoint + "/upscale"
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return err
	}

	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}

	return nil
}
