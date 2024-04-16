package repository

import (
	"github.com/nurcholisnanda/golang-assignment/internal/dto"
	"github.com/nurcholisnanda/golang-assignment/internal/model"
	"gorm.io/gorm"
)

// repository represents a database repository implementation.
type repository struct {
	db *gorm.DB
}

// RepoInterface defines the repository interface.
//
//go:generate mockgen -source=repository.go -destination=mock/repository.go -package=mock
type RepoInterface interface {
	GetRecords(req *dto.FetchRecordsRequest) ([]model.Record, error)
}

// NewUserRepoImpl creates a new instance of Userrepository.
func NewUserRepoImpl(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

// GetRecords retrieves records based on the provided request parameters.
func (r *repository) GetRecords(req *dto.FetchRecordsRequest) ([]model.Record, error) {
	var records []model.Record

	// Define the raw SQL query to calculate the total marks and filter records
	query := `
        SELECT id, created_at, SUM(unnested_marks) AS total_marks
        FROM (
            SELECT s.id, s.created_at, unnest(s.marks) AS unnested_marks
            FROM students s
            WHERE s.created_at BETWEEN ? AND ?
        ) AS marks_per_student
        GROUP BY id, created_at
        HAVING SUM(unnested_marks) BETWEEN ? AND ?
        ORDER BY id
    `

	// Execute the raw SQL query with parameters
	if err := r.db.Raw(query, req.StartDate, req.EndDate, req.MinCount, req.MaxCount).Scan(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}
