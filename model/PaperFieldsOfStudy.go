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

func (paperFieldsOfStudy *PaperFieldsOfStudy) SetPaperId() {
	err := rdb.Set(rdb.Context(), strconv.Itoa(int(paperFieldsOfStudy.FieldOfStudyId)), paperFieldsOfStudy.PaperId, 0).Err()
	if err != nil {
		log.Fatalf("Set %d error: %s", paperFieldsOfStudy.FieldOfStudyId, err.Error())
	}

}

func (paperFieldsOfStudy *PaperFieldsOfStudy) GetPaperId(fieldOfStudyId int64) string {
	s, err := rdb.Get(rdb.Context(), strconv.Itoa(int(fieldOfStudyId))).Result()
	switch {
	case err == redis.Nil:
		log.Printf("Get %d key not found!", paperFieldsOfStudy.FieldOfStudyId)

	case err == nil:
		log.Fatalln("Get error: ", err.Error())
	}
	return s

}
