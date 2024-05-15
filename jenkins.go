package jenkins

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
}

func New(rawURL string) (*Client, error) {
	baseURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	client := resty.New().EnableTrace()

	if baseURL.User != nil {
		username := baseURL.User.Username()
		password, _ := baseURL.User.Password()
		client.SetBasicAuth(username, password)
		baseURL.User = nil
	}

	client.SetBaseURL(baseURL.String())

	client.SetHeader("Content-Type", "application/json")

	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		if r.StatusCode() != http.StatusOK {
			return fmt.Errorf("unexpected http status code received: %s", r.String())
		}

		if r.Body() == nil {
			return fmt.Errorf("malformed response body received: %v", r.Result())
		}

		return nil
	})

	c := &Client{
		client,
	}

	return c, nil
}

type Info struct {
	URL string `json:"url"`
}

// Get Basic Information About Jenkins
func (c *Client) Info() (*Info, error) {
	r := &Info{}

	if _, err := c.client.R().SetResult(r).Get("/api/json"); err != nil {
		return nil, err
	}

	return r, nil
}
