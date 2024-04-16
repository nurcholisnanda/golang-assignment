package service

import (
	"log"

	"github.com/nurcholisnanda/golang-assignment/internal/dto"
	"github.com/nurcholisnanda/golang-assignment/internal/repository"
)

// service represents a service implementation.
type service struct {
	repo repository.RepoInterface
}

// ServiceInterface defines the service interface.
//
//go:generate mockgen -source=service.go -destination=mock/service.go -package=mock
type ServiceInterface interface {
	GetRecords(req *dto.FetchRecordsRequest) (*dto.FetchRecordsResponse, error)
}

// NewService creates a new instance of Service.
func NewService(repo repository.RepoInterface) *service {
	return &service{
		repo: repo,
	}
}

// GetRecords retrieves records based on the provided request and returns a response.
func (s *service) GetRecords(req *dto.FetchRecordsRequest) (*dto.FetchRecordsResponse, error) {
	records, err := s.repo.GetRecords(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &dto.FetchRecordsResponse{
		Code:    dto.Success,
		Msg:     "success",
		Records: records,
	}, nil
}
