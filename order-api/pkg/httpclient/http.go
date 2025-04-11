package http

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func DoRequest(options RequestOptions) error {
	client := resty.New()

	req := client.R().SetResult(options.Result)

	// Adiciona query params, se houver
	if options.QueryParams != nil {
		req.SetQueryParams(options.QueryParams)
	}

	// Adiciona headers, se houver
	if options.Headers != nil {
		req.SetHeaders(options.Headers)
	}

	// Define método (GET, POST, etc.)
	var resp *resty.Response
	var err error
	switch options.Method {
	case "GET":
		resp, err = req.Get(options.URL)
	case "POST":
		resp, err = req.Post(options.URL)
	case "PUT":
		resp, err = req.Put(options.URL)
	case "DELETE":
		resp, err = req.Delete(options.URL)
	default:
		return fmt.Errorf("unsupported HTTP method: %s", options.Method)
	}

	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	if resp.IsError() {
		return fmt.Errorf("API error: %s", resp.Status())
	}

	return nil
}
