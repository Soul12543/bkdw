package bkdw

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"golang.org/x/net/html/charset"
)

func GetBookInfo(id string) (io.Reader, error) {
	_, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("Invalid Number")
	}

	url := fmt.Sprintf("https://www.52bqg.org/book_%s", id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	return body, nil
}
