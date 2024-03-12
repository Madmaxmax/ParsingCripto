package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"stepik_go/money"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := money.ReadCurrency()
	err := tpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(" Server Error"))
	}
	params := u.Query()
	searchKey := params.Get("q")

	currency := money.Search{}
	currency.SearchKey = searchKey

	data := money.ReadCurrency()
	currency.Results = money.FindMatches(searchKey, data)
	err = tpl.Execute(w, currency)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
