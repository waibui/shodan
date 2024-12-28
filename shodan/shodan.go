package shodan

const BaseURL = "https://api.shodan.io"

type Client struct {
	apiKey string
}

// Returns a new shodan client
func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}
