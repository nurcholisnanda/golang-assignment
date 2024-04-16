package dto

import "github.com/nurcholisnanda/golang-assignment/internal/model"

type ResponseCode int

const (
	Success       ResponseCode = 0
	ErrNotFound   ResponseCode = 1
	ErrBadRequest ResponseCode = 2
	ErrUnknown    ResponseCode = 99
)

type FetchRecordsResponse struct {
	Code    ResponseCode   `json:"code"`
	Msg     string         `json:"msg"`
	Records []model.Record `json:"records,omitempty"`
}
