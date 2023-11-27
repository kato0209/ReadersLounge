package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func FetchApi(url, response interface{}) error {
	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, response); err != nil {
		return err
	}

	return nil
}
