package handler

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string, params []string) (string, error, int) {

	if url == "" {
		return "", errors.New("Empty url"), 400
	}

	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()

	for i := 0; i < len(params); i += 2 {
		q.Add(params[i], params[i+1])
	}

	req.URL.RawQuery = q.Encode()

	log.Println("GET: " + req.URL.String())

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err, 503
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New(resp.Status), resp.StatusCode
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err, 503
	}

	return string(body), nil, resp.StatusCode
}
