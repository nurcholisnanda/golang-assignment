package model

import (
	"time"

	pq "github.com/lib/pq"
)

type Student struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Marks     pq.Int32Array `gorm:"type:integer[]"`
	CreatedAt time.Time     `gorm:"index"`
}

type Record struct {
	ID         uint
	CreatedAt  time.Time
	TotalMarks int
}
