package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"restaurantapp/pkg/schema"
)

func PostSchema(url string, body []byte) bool {

	var response schema.Response

	if url == "" {
		return false
	}

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

	jsonBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(jsonBody), &response)

	if err != nil {
		log.Fatal(err)
	}

	if response.Data.Code == "" {
		log.Println(response.Errors[0].Message)
		return false
	}

	return true
}
