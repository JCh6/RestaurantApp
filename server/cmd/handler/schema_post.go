package handler

import (
	"log"
	"bytes"
	"net/http"
	"restaurantapp/pkg/schema"
)

func PostSchema(url string) bool {

	if url == "" {
		return false
	}

	body := schema.Get()
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("%s: %s", resp.Status, resp.Body)
	}

	return true

}