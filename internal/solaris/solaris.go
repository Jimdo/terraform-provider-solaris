package solaris

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

type Config struct {
	Endpoint     string
	ClientID     string
	ClientSecret string
}

type Client struct {
	config     Config
	httpClient *http.Client
}

func NewClient(ctx context.Context, config Config) *Client {
	cc := clientcredentials.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		TokenURL:     fmt.Sprintf("https://%s/oauth/token", config.Endpoint),
	}
	return &Client{config: config, httpClient: cc.Client(ctx)}
}

func (c *Client) sendRequest(req *http.Request, result interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", "jimdo/terraform-provider-solarisbank")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if result == nil {
		return nil
	}

	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		return err
	}

	return nil
}

type GetWebhookResponse struct {
	ID        string `json:"id"`
	EventType string `json:"event_type"`
	URL       string `json:"url"`
}

func (c *Client) GetWebhook(ctx context.Context, ID string) (*GetWebhookResponse, error) {
	url := fmt.Sprintf("https://%s/v1/webhooks/%s", c.config.Endpoint, ID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := GetWebhookResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type CreateWebhookRequest struct {
	EventType string `json:"event_type"`
	URL       string `json:"url"`
}

type CreateWebhookResponse struct {
	ID        string `json:"id"`
	EventType string `json:"event_type"`
	URL       string `json:"url"`
	Secret    string `json:"secret"`
}

func (c *Client) CreateWebhook(ctx context.Context, data *CreateWebhookRequest) (*CreateWebhookResponse, error) {
	url := fmt.Sprintf("https://%s/v1/webhooks", c.config.Endpoint)
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	res := CreateWebhookResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteWebhook(ctx context.Context, ID string) error {
	url := fmt.Sprintf("https://%s/v1/webhooks/%s", c.config.Endpoint, ID)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return err
	}

	if err := c.sendRequest(req, nil); err != nil {
		return err
	}

	return nil
}
