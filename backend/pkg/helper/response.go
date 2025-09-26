package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"status,omitempty"`
	Message    string      `json:"message"`
	Meta       interface{} `json:"meta,omitempty"`
	Data       interface{} `json:"data,omitempty" `
}

func failResponseWriter(w http.ResponseWriter, err error, errStatusCode int) {
	w.Header().Set("Content-Type", "application/json")

	var resp Response
	w.WriteHeader(errStatusCode)
	resp.StatusCode = errStatusCode
	resp.Message = err.Error()
	resp.Data = nil

	responseBytes, _ := json.Marshal(resp)
	w.Write(responseBytes)
}

func successResponseWriter(w http.ResponseWriter, response *Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseBytes, _ := json.Marshal(response)
	w.Write(responseBytes)
}

func WriteResponse(w http.ResponseWriter, err error, response *Response) {
	switch err.(type) {
	case *ErrForbidden:
		failResponseWriter(w, err, http.StatusForbidden)
	case *ErrUnauthorized:
		failResponseWriter(w, err, http.StatusUnauthorized)
	case *ErrNotFound:
		failResponseWriter(w, err, http.StatusNotFound)
	case *ErrBadRequest:
		failResponseWriter(w, err, http.StatusBadRequest)
	case nil:
		successResponseWriter(w, response)
	default:
		failResponseWriter(w, err, http.StatusInternalServerError)
	}
}

func GenerateTotalPage(totalData, limit int64) (totalPage int64) {
	totalPage = totalData / limit
	modTotalPage := totalData % limit
	if modTotalPage > 0 {
		totalPage++
	}

	return totalPage
}

func GetOffsetAndLimit(page, pageSize int64) (offset, limit int64) {
	offset = (page - 1) * pageSize
	limit = pageSize

	return offset, limit
}
