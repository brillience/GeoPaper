package model

type FieldOfStudyChildren struct {
	FieldOfStudyId      int64 `gorm:"primaryKey,uniqueIndex"`
	ChildFieldOfStudyId int64
}

func (FieldOfStudyChildren) CreateInBatches(items []FieldOfStudyChildren, size int) {
	db.CreateInBatches(items, size)
}

func (FieldOfStudyChildren) QueryById(id int64) *FieldOfStudyChildren {
	item := FieldOfStudyChildren{}
	db.First(&item, id)
	return &item
}
