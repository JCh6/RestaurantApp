package handler

import (
	"net/http"
	ModelBuyer "restaurantapp/pkg/models/buyers"
	ModelResponse "restaurantapp/pkg/models/response"
	Schema "restaurantapp/pkg/schema"
	"strconv"
)

func GetBuyers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data Schema.QueryBuyer
		w.Header().Set("Content-Type", ModelResponse.ContentType())
		pageNum, errPage := strconv.Atoi(r.URL.Query().Get("page"))
		limit, errLimit := strconv.Atoi(r.URL.Query().Get("limit"))

		if errPage != nil || pageNum <= 0 {
			pageNum = 1
		}

		if errLimit != nil || limit <= 0 {
			limit = 10
		}

		input := map[string]interface{}{
			"first":  limit,
			"offset": (pageNum - 1) * limit,
		}

		err := QueryGQL(ModelBuyer.GetBuyersGQL(), input, &data)

		if err != nil {
			w.Write(ModelResponse.GetResponseBody(502, err.Error(), nil))
			return
		}

		w.Write(ModelResponse.GetResponseBody(200, "", data))
	}
}
