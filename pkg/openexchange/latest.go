package openexchange

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
	"net/http"
)

type Response struct {
	Disclaimer string             `json:"disclaimer"`
	License    string             `json:"license"`
	Timestamp  int64              `json:"timestamp"`
	Base       string             `json:"base"`
	Rates      map[string]float64 `json:"rates"`
}

func (c *Client) Latest(ctx context.Context, currencyFrom string, currencyTo string) (f float64, err error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, c.urlLatest, nil)
	if err != nil {
		return 0, err
	}

	query := httpReq.URL.Query()
	query.Add("app_id", c.apiKey)
	query.Add("base", currencyFrom)
	query.Add("symbols", currencyTo)
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return 0, err
	}

	defer func() {
		closeErr := httpResp.Body.Close()
		if err != nil && closeErr == nil {
			err = multierr.Combine(closeErr, err)
		}
	}()

	if httpResp.StatusCode != http.StatusOK {
		return 0, errors.Errorf("status code not ok, status code = %d", httpResp.StatusCode)
	}

	var resp Response
	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return 0, err
	}

	rate, ok := resp.Rates[currencyTo]
	if !ok {
		return 0, errors.Errorf("currency pair dont exists")
	}

	return rate, nil
}
