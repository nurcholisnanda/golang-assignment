package service

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/nurcholisnanda/golang-assignment/internal/dto"
	"github.com/nurcholisnanda/golang-assignment/internal/model"
	"github.com/nurcholisnanda/golang-assignment/internal/repository/mock"
	"go.uber.org/mock/gomock"
)

func Test_service_GetRecords(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepoInterface(ctrl)
	mockRepoRecord := mockRepo.EXPECT()
	now := time.Now()
	req := &dto.FetchRecordsRequest{
		StartDate: "2016-01-26",
		EndDate:   "2026-02-02",
		MinCount:  100,
		MaxCount:  300,
	}
	records := []model.Record{
		{
			ID:         1,
			TotalMarks: 270,
			CreatedAt:  now,
		},
		{
			ID:         2,
			TotalMarks: 150,
			CreatedAt:  now,
		},
	}
	res := &dto.FetchRecordsResponse{
		Code: 0,
		Msg:  "success",
		Records: []model.Record{
			{
				ID:         1,
				CreatedAt:  now,
				TotalMarks: 270,
			},
			{
				ID:         2,
				CreatedAt:  now,
				TotalMarks: 150,
			},
		},
	}

	type args struct {
		req *dto.FetchRecordsRequest
	}

	tests := []struct {
		name    string
		s       *service
		args    args
		want    *dto.FetchRecordsResponse
		wantErr bool
		mocks   *gomock.Call
	}{
		{
			name: "error get records from repo",
			s:    NewService(mockRepo),
			args: args{
				req: req,
			},
			want:    nil,
			wantErr: true,
			mocks:   mockRepoRecord.GetRecords(gomock.Any()).Return(nil, errors.New("any error")),
		},
		{
			name: "success",
			s:    NewService(mockRepo),
			args: args{
				req: req,
			},
			want:    res,
			wantErr: false,
			mocks:   mockRepoRecord.GetRecords(gomock.Any()).Return(records, nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetRecords(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}
