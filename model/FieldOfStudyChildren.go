package model

type FieldOfStudyChildren struct {
	FieldOfStudyId      int64 `gorm:"primaryKey,uniqueIndex"`
	ChildFieldOfStudyId int64
}

func (FieldOfStudyChildren) CreateInBatches(items []*FieldOfStudyChildren, size int) {
	db.CreateInBatches(items, size)
}

// 通过父学科的id查找子学科的id
func (FieldOfStudyChildren) QueryById(id int64) []FieldOfStudyChildren {
	items := []FieldOfStudyChildren{}
	db.Find(&items, id)
	return items
}
