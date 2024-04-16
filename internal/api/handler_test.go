package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nurcholisnanda/golang-assignment/internal/dto"
	"github.com/nurcholisnanda/golang-assignment/internal/service/mock"
	"go.uber.org/mock/gomock"
)

func Test_handler_GetRecords(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSvc := mock.NewMockServiceInterface(ctrl)
	mockSvcRecord := mockSvc.EXPECT()

	h := &handler{
		svc: mockSvc,
	}

	tests := []struct {
		name   string
		status int
		url    string
		mocks  *gomock.Call
	}{
		{
			name:   "error get records from service",
			status: http.StatusBadRequest,
			url:    "/records?startDate=1900-01-01&endDate=2024-04-12&minCount=0&maxCount=500",
			mocks:  mockSvcRecord.GetRecords(gomock.Any()).Return(nil, errors.New("any error")),
		},
		{
			name:   "error record not found from service",
			status: http.StatusNotFound,
			url:    "/records?startDate=1900-01-01&endDate=2024-04-12&minCount=0&maxCount=500",
			mocks:  mockSvcRecord.GetRecords(gomock.Any()).Return(nil, errors.New("record not found")),
		},
		{
			name:   "error minCount query",
			status: http.StatusBadRequest,
			url:    "/records?startDate=1900-01-01&endDate=2024-04-12&minCount=a&maxCount=500",
		},
		{
			name:   "error maxCount query",
			status: http.StatusBadRequest,
			url:    "/records?startDate=1900-01-01&endDate=2024-04-12&minCount=0&maxCount=b",
		},
		{
			name:   "success",
			status: http.StatusOK,
			url:    "/records?startDate=1900-01-01&endDate=2024-04-12&minCount=0&maxCount=500",
			mocks:  mockSvcRecord.GetRecords(gomock.Any()).Return(&dto.FetchRecordsResponse{}, nil),
		},
	}
	for _, tt := range tests {
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		ctx.Request, _ = http.NewRequest("GET", tt.url, nil)
		h.GetRecords(ctx)
		if ok := reflect.DeepEqual(ctx.Writer.Status(), tt.status); !ok {
			t.Errorf("http status = %v, want %v", ctx.Writer.Status(), tt.status)
		}
	}
}
