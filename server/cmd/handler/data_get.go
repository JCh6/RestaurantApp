package handler

import (
	"encoding/json"
	"net/http"
	"restaurantapp/pkg/models/response"
	"strconv"
	"time"
)

func GetData(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		date := r.URL.Query().Get("date")
		var resp *response.Response
		var params []string

		if date == "" {
			date = strconv.FormatInt(time.Now().Unix(), 10)
		}

		params = append(params, "date")
		params = append(params, date)

		bodyBuyers, err, status := Get(url+"/buyers", params)

		if err != nil {
			resp = response.New(status, err.Error(), bodyBuyers)
			resBodyBytes, _ := json.Marshal(resp)
			w.Write(resBodyBytes)
			return
		}

		if err := InsertToDgraph("Buyer", bodyBuyers); err != nil {
			resp = response.New(status, err.Error(), bodyBuyers)
			resBodyBytes, _ := json.Marshal(resp)
			w.Write(resBodyBytes)
			return
		}

		bodyProducts, err, status := Get(url+"/products", params)

		if err != nil {
			resp = response.New(status, err.Error(), bodyProducts)
			resBodyBytes, _ := json.Marshal(resp)
			w.Write(resBodyBytes)
			return
		}

		if err := InsertToDgraph("Product", bodyProducts); err != nil {
			resp = response.New(status, err.Error(), bodyBuyers)
			resBodyBytes, _ := json.Marshal(resp)
			w.Write(resBodyBytes)
			return
		}

		bodyTransactions, err, status := Get(url+"/transactions", params)

		if err != nil {
			resp = response.New(status, err.Error(), bodyTransactions)
			resBodyBytes, _ := json.Marshal(resp)
			w.Write(resBodyBytes)
			return
		}

		if err := InsertToDgraph("Transaction", bodyTransactions); err != nil {
			resp = response.New(status, err.Error(), bodyTransactions)
			resBodyBytes, _ := json.Marshal(resp)
			w.Write(resBodyBytes)
			return
		}

		resp = response.New(200, "", "OK")
		resBodyBytes, _ := json.Marshal(resp)
		w.Write(resBodyBytes)
	}
}
