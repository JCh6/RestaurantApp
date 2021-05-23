package handler

import (
	"bytes"
	"log"
	"net/http"
	"restaurantapp/pkg/schema"
)

func PostSchema(url string) bool {

	if url == "" {
		return false
	}

	body := schema.Get()
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	log.Println("POST: " + req.URL.String())

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal(resp.Status)
	}

	return true

}
