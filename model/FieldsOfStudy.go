package model

import (
	"strconv"
)

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
	CreateDate       string
}

func (receiver FieldsOfStudy) CreateInBatches(items []*FieldsOfStudy, batchSize int) {
	db.CreateInBatches(items, batchSize)
}

func (receiver FieldsOfStudy) QueryById(id int64) *FieldsOfStudy {
	item := FieldsOfStudy{}
	db.First(&item, id)
	return &item
}

func (receiver FieldsOfStudy) QueryByNormalizedNameAndLevel(normalizedName string, level int) []FieldsOfStudy {
	res := []FieldsOfStudy{}
	db.Where("normalized_name = ? AND level = ?", normalizedName, strconv.Itoa(level)).Find(&res)
	return res
}
