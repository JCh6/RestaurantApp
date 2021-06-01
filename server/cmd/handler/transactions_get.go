package handler

import (
	"net/http"
	ModelResponse "restaurantapp/pkg/models/response"
	ModelTransaction "restaurantapp/pkg/models/transactions"
	"strconv"
)

func GetTransactions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data interface{}
		w.Header().Set("Content-Type", ModelResponse.ContentType())
		pageNum, errPage := strconv.Atoi(r.URL.Query().Get("page"))
		limit, errLimit := strconv.Atoi(r.URL.Query().Get("limit"))
		buyer := r.URL.Query().Get("buyer")

		if errPage != nil || pageNum <= 0 {
			pageNum = 1
		}

		if errLimit != nil || limit <= 0 {
			limit = 10
		}

		input := map[string]interface{}{
			"buyer":  buyer,
			"first":  limit,
			"offset": (pageNum - 1) * limit,
		}

		err := QueryGQL(ModelTransaction.GetTransactionsByBuyerGQL(), input, &data)

		if err != nil {
			w.Write(ModelResponse.GetResponseBody(502, err.Error(), nil))
			return
		}

		w.Write(ModelResponse.GetResponseBody(200, "", data))

	}
}
