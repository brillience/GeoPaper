package util

import (
	"testing"
)

func TestPushFieldsOfStudyToDB(t *testing.T) {
	var filePath string = "../../MagData/FieldsOfStudy.txt"
	PushFieldsOfStudyToDB(filePath, 100)
}
func TestPushFieldsOfStudyChildrenToDB(t *testing.T) {
	var filePath string = "../../MagData/FieldOfStudyChildren.txt"
	PushFieldsOfStudyChildrenToDB(filePath, 100)
}

func TestPushPaperFieldsOfStudyToDB(t *testing.T) {
	var filePath string = "../../MagData/PaperFieldsOfStudy.txt"
	PushPaperFieldsOfStudyToDB(filePath)
}
