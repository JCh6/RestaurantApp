package handler

import (
	"net/http"
	ModelBuyer "restaurantapp/pkg/models/buyers"
	ModelResponse "restaurantapp/pkg/models/response"
	ModelTransaction "restaurantapp/pkg/models/transactions"
	"strconv"
)

func GetTransactions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data interface{}
		var buyersInfo interface{}
		var err error

		w.Header().Set("Content-Type", ModelResponse.ContentType())

		txnId := r.URL.Query().Get("id")
		buyer := r.URL.Query().Get("buyer")

		if txnId != "" {
			err = GetTransactionById(txnId, r, &data)
			ids := make([]string, 0)

			if err == nil {
				dMap := data.(map[string]interface{})

				for _, b := range dMap["other"].([]interface{}) {
					val := b.(map[string]interface{})["buyer"]
					ids = append(ids, val.(string))
				}

				input := map[string]interface{}{
					"ids": ids,
				}

				err = QueryGQL(ModelBuyer.GetBuyersByIds(), input, &buyersInfo)

				dMap["buyerInfo"] = buyersInfo
			}

		} else {
			err = GetTransactionByBuyer(buyer, r, &data)
		}

		if err != nil {
			w.Write(ModelResponse.GetResponseBody(502, err.Error(), nil))
			return
		}

		w.Write(ModelResponse.GetResponseBody(200, "", data))
	}
}

func GetTransactionById(id string, r *http.Request, data interface{}) error {
	ip := r.URL.Query().Get("ip")

	input := map[string]interface{}{
		"id": id,
		"ip": ip,
	}

	return QueryGQL(ModelTransaction.GetTransactionByIdGQL(), input, &data)
}

func GetTransactionByBuyer(id string, r *http.Request, data interface{}) error {
	pageNum, errPage := strconv.Atoi(r.URL.Query().Get("page"))
	limit, errLimit := strconv.Atoi(r.URL.Query().Get("limit"))

	if errPage != nil || pageNum <= 0 {
		pageNum = 1
	}

	if errLimit != nil || limit <= 0 {
		limit = 10
	}

	input := map[string]interface{}{
		"buyer":  id,
		"first":  limit,
		"offset": (pageNum - 1) * limit,
	}

	return QueryGQL(ModelTransaction.GetTransactionsByBuyerGQL(), input, &data)
}
