package handler

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string, params []string) (bool, string) {

	if url == "" {
		return false, ""
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
		log.Println(err)
		return false, ""
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println(resp.Status)
		return false, ""
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return false, ""
	}

	log.Println(string(body))

	return true, string(body)
}
