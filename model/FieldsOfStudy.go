package model

import (
	"strconv"
	"time"
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
	CreateDate       time.Time
}

func (receiver FieldsOfStudy) CreateInBatches(items []*FieldsOfStudy, batchSize int) {
	db.CreateInBatches(items, batchSize)
}

func (receiver FieldsOfStudy) QueryById(id int64) *FieldsOfStudy {
	item := FieldsOfStudy{}
	db.First(&item, id)
	return &item
}

func (receiver FieldsOfStudy) QueryByNormalizedNameAndLevel(normalizedName string,level int) []FieldsOfStudy {
	res := []FieldsOfStudy{}
	db.Where("normalizedName = ？ AND level = ?",normalizedName,strconv.Itoa(level)).Find(&res)
	return res
}
