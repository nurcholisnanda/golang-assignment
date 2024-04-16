package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"github.com/nurcholisnanda/golang-assignment/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConn() (*gorm.DB, error) {
	if _, err := os.Stat("./../../.env"); !os.IsNotExist(err) {
		err := godotenv.Load(os.ExpandEnv("./../../.env"))
		if err != nil {
			log.Fatalf("Error getting env %v\n", err)
		}
	}
	return LocalDatabase()
}

func LocalDatabase() (*gorm.DB, error) {

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME_TEST"),
		os.Getenv("DB_PASS"),
	)

	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	if err := db.Migrator().DropTable(&model.Student{}); err != nil {
		log.Panic(err)
	}

	if err := db.AutoMigrate(&model.Student{}); err != nil {
		log.Panic(err)
	}

	return db, nil
}

func seedStudents(db *gorm.DB) []*model.Student {

	students := []*model.Student{
		{
			Name:  "a",
			Marks: pq.Int32Array([]int32{100, 100, 100}),
		},
		{
			Name:  "b",
			Marks: pq.Int32Array([]int32{60, 60, 60}),
		},
		{
			Name:  "c",
			Marks: pq.Int32Array([]int32{70, 70, 90}),
		},
		{
			Name:  "d",
			Marks: pq.Int32Array([]int32{100, 100, 100}),
		},
	}
	if err := db.Create(students).Error; err != nil {
		log.Panic(err)
	}
	return students
}
