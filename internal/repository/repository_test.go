package repository

import (
	"reflect"
	"testing"

	"github.com/nurcholisnanda/golang-assignment/internal/dto"
	"github.com/nurcholisnanda/golang-assignment/internal/model"
)

func Test_repository_GetRecords(t *testing.T) {
	db, _ := DBConn()
	seedStudents(db)

	type args struct {
		req *dto.FetchRecordsRequest
	}
	tests := []struct {
		name    string
		r       *repository
		args    args
		want    []model.Record
		wantErr bool
	}{
		{
			name: "record not found",
			r:    NewUserRepoImpl(db),
			args: args{
				req: &dto.FetchRecordsRequest{
					StartDate: "2025-01-01",
					EndDate:   "1 OR 1",
					MinCount:  500,
					MaxCount:  1000,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			r:    NewUserRepoImpl(db),
			args: args{
				req: &dto.FetchRecordsRequest{
					StartDate: "2022-01-01",
					EndDate:   "2024-12-31",
					MinCount:  210,
					MaxCount:  300,
				},
			},
			want: []model.Record{
				{
					ID:         1,
					TotalMarks: 300,
				},
				{
					ID:         3,
					TotalMarks: 230,
				},
				{
					ID:         4,
					TotalMarks: 300,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetRecords(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				for i, v := range got {
					if !reflect.DeepEqual(v.ID, tt.want[i].ID) {
						t.Errorf("repository.GetRecords() = %v, want %v", v.ID, tt.want[i].ID)
					}
					if !reflect.DeepEqual(v.TotalMarks, tt.want[i].TotalMarks) {
						t.Errorf("repository.GetRecords() = %v, want %v", v.TotalMarks, tt.want[i].TotalMarks)
					}
				}
			}
		})
	}
}
