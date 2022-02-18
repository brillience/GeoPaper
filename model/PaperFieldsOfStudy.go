package model

import (
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
)

// 采用redis数据库
type PaperFieldsOfStudy struct {
	PaperId        int64
	FieldOfStudyId int64
	Score          float64
}

func (paperFieldsOfStudy *PaperFieldsOfStudy) SetPaperIdToDB() {
	err := rdb.Set(rdb.Context(), strconv.FormatInt(paperFieldsOfStudy.FieldOfStudyId, 10), paperFieldsOfStudy.PaperId, 0).Err()
	if err != nil {
		log.Fatalf("Set %d error: %s", paperFieldsOfStudy.FieldOfStudyId, err.Error())
	}

}

func (paperFieldsOfStudy *PaperFieldsOfStudy) GetPaperIdFromDB(fieldOfStudyId int64) string {
	s, err := rdb.Get(rdb.Context(), strconv.FormatInt(fieldOfStudyId, 10)).Result()
	switch {
	case err == redis.Nil:
		log.Printf("Get %d key not found!", paperFieldsOfStudy.FieldOfStudyId)

	case err == nil:
		log.Fatalln("Get error: ", err.Error())
	}
	return s

}
