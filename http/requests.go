package http

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"web_scraping/exceptions"
)

func Get(url string, headers map[string]string, params map[string]string) *Response {
	if params != nil && len(params) > 0 {
		if !strings.Contains(url, "?") {
			url += "?"
		}
		for param, value := range params {
			url += param + "=" + value + "&"
		}
		url = url[:len(url)-1]
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	exceptions.HandleError(err)

	if headers != nil && len(headers) > 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	resp, err := client.Do(req)
	exceptions.HandleError(err)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		exceptions.HandleError(err)
	}(resp.Body)

	b, err := io.ReadAll(resp.Body)
	exceptions.HandleError(err)

	return &Response{
		StatusCode: resp.StatusCode,
		Content:    b,
		Text:       string(b),
	}
}

type Response struct {
	StatusCode int
	Content    []byte
	Text       string
}

func (r *Response) Json() map[string]interface{} {
	var jsn map[string]interface{}
	err := json.Unmarshal(r.Content, &jsn)
	exceptions.HandleError(err)

	return jsn
}
