package openexchange

type Client struct {
	apiKey    string
	url       string
	urlLatest string
}

func New(url string, apiKey string) *Client {
	return &Client{
		apiKey:    apiKey,
		url:       url,
		urlLatest: url + "/api/latest.json",
	}
}
