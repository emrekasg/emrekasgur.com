package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var allowedLanguages = []string{"en", "tr"}

type (
	Response struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

func Resp(message string, data interface{}) *Response {
	return &Response{
		Message: message,
		Data:    data,
	}
}

func WriteResp(res http.ResponseWriter, statusCode int, message string, data interface{}) {
	res.WriteHeader(statusCode)
	httpResponse := Resp(message, data)
	_ = json.NewEncoder(res).Encode(httpResponse)
}

func GetLimitAndOffset(req *http.Request) (int, int, error) {
	limit, offset := req.URL.Query().Get("limit"), req.URL.Query().Get("offset")
	if limit == "" {
		limit = "10"
	}
	if offset == "" {
		offset = "0"
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		fmt.Println(err)
		return 0, 0, err
	}

	intOffset, err := strconv.Atoi(offset)
	if err != nil {
		fmt.Println(err)
		return 0, 0, err
	}

	return intLimit, intOffset, nil
}

func CheckLanguage(language string) bool {
	for _, lang := range allowedLanguages {
		if lang == language {
			return true
		}
	}
	return false
}

func getAllowedLanguages() string {
	var langs string
	for _, lang := range allowedLanguages {
		langs += lang + ", "
	}
	return langs
}
