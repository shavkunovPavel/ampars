package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

// струкртура для приема массива урлов
type UrlsType struct {
	Urls *[]Amurl `json:"urls"`
}

type Amurl struct {
	UrlString string `json:"url"`
}

// преобразование request body в json
func make_json(body io.Reader, v interface{}) {
	decoder := json.NewDecoder(body)
	if err := decoder.Decode(v); err != nil {
		fmt.Println("bad posted json")
	}
}

// обработчик запроса
func handleRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	data := new(UrlsType)
	make_json(r.Body, data)

	parsed_objects := process_urls(data)

	output, _ := json.MarshalIndent(parsed_objects, "", "\t")
	w.Write(output)
}
