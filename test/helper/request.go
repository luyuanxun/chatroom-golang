package helper

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Request struct {
}

func (t *Request) Do(method string, url string, data io.Reader) ([]byte, error) {
	client := &http.Client{}
	client.Timeout = time.Second * 300
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
