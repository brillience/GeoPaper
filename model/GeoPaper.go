package model

type GeoPaper struct {
	PaperId                    int64 `gorm:"primaryKey,uniqueIndex"`
	FieldOfStudyId             int64
	FieldOfStudyLevel          int
	FieldOfStudyNormalizedName string
	ParentFieldOfStudyId       int64
}
