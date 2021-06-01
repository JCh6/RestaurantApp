package handler

import (
	"net/http"
	Config "restaurantapp/pkg/config"
	ModelResponse "restaurantapp/pkg/models/response"
	Schema "restaurantapp/pkg/schema"
	"strconv"
	"time"
)

func GetData(url string, alterUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data interface{}
		config := Config.Get()

		date := r.URL.Query().Get("date")

		w.Header().Set("Content-Type", ModelResponse.ContentType())

		if PostSchema(alterUrl, Schema.DeleteData()) {

			for _, endPoint := range config {
				var params []string

				if date == "" {
					date = strconv.FormatInt(time.Now().Unix(), 10)
				}

				params = append(params, endPoint.Param)
				params = append(params, date)

				body, err, status := Get(url+endPoint.Route, params)

				if err != nil {
					w.Write(ModelResponse.GetResponseBody(status, err.Error(), nil))
					return
				}

				if err := InsertToDgraph(endPoint.Name, body); err != nil {
					w.Write(ModelResponse.GetResponseBody(502, err.Error(), nil))
					return
				}
			}

			err := QueryGQL(Schema.Count(), nil, &data)

			if err != nil {
				w.Write(ModelResponse.GetResponseBody(502, err.Error(), nil))
				return
			}

			w.Write(ModelResponse.GetResponseBody(200, "", data))

		}
	}
}
