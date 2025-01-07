package handlers

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func LoadPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.New("Error loading page")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("Error reading page")
	}

	return string(body), nil
}
