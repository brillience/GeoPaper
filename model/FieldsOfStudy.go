package model

import "time"

type FieldsOfStudy struct {
	FieldOfStudyId   int64 `gorm:"primaryKey,uniqueIndex"`
	Rank             string
	NormalizedName   string
	DisplayName      string
	MainType         string
	Level            int
	PaperCount       int64
	PaperFamilyCount int64
	CitationCount    int64
	CreateDate       time.Time
}

func (FieldsOfStudy) CreateInBatches(items []FieldsOfStudy, size int) {
	db.CreateInBatches(items, size)
}

func (FieldsOfStudy) QueryById(id int64) *FieldsOfStudy {
	item := FieldsOfStudy{}
	db.First(&item, id)
	return &item
}
