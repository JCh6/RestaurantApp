package handler

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string, params []string) (error, string) {

	if url == "" {
		return errors.New("Empty url"), ""
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
		return err, ""
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(resp.Status), ""
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err, ""
	}

	return nil, string(body)
}
