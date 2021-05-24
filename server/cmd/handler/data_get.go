package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetData(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("date")
		var params []string

		if date == "" {
			date = strconv.FormatInt(time.Now().Unix(), 10)
		}

		params = append(params, "date")
		params = append(params, date)

		err, body := Get(url+"/products", params)
		log.Println(err)
		w.Write([]byte("welcome: " + body))
	}
}
