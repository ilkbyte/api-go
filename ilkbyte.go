package ilkbyte

import (
	"encoding/json"
	"net/http"
	"time"
)

const BaseURLV1 = "https://api.ilkbyte.com"

type Client struct {
	BaseURL    string
	Access     string
	Secret     string
	HTTPClient *http.Client
}

type Response struct {
	Status  bool        `json:"status"`
	Error   string      `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewClient(access, secret string) *Client {
	return &Client{
		BaseURL: BaseURLV1,
		Access:  access,
		Secret:  secret,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) sendRequest(req *http.Request, v interface{}, params map[string]string) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	q := req.URL.Query()
	q.Add("access", c.Access)
	q.Add("secret", c.Secret)
	if len(params) > 0 {
		for k, v := range params {
			q.Add(k, v)
		}
	}
	req.URL.RawQuery = q.Encode()

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
