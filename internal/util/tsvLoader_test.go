package util

import (
	"os"
	"strings"
	"testing"
)

func TestPushFieldsOfStudyToDB(t *testing.T) {
	file, err := os.OpenFile("FieldsOfStudy.tsv", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		t.Error(err.Error())
	}
	data := []string{
		strings.Join([]string{"127313418", "5789", "geology", "Geology", "","0", "6600805", "6505292", "36520322", "2016-06-24"}, "\t"),
		strings.Join([]string{"122690726", "10865", "meta", "Meta-", "","2", "1002", "917", "12091", "2016-06-30"}, "\t"),
	}
	file.WriteString(strings.Join(data,"\n")+"\n")
	file.Close()

	PushFieldsOfStudyToDB("FieldsOfStudy.tsv",100)

}
