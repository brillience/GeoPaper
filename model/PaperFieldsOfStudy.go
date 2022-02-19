package model

import (
	"log"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

// 采用redis数据库
type PaperFieldsOfStudy struct {
	PaperId        int64
	FieldOfStudyId int64
	Score          float64
}

func (paperFieldsOfStudy *PaperFieldsOfStudy) SetPaperIdToDB() {
	conn := rdbPool.Get()
	defer conn.Close()
	_, err := conn.Do("SADD", strconv.FormatInt(paperFieldsOfStudy.FieldOfStudyId, 10), paperFieldsOfStudy.PaperId)
	if err != nil {
		log.Fatalf("SADD %d error: %s", paperFieldsOfStudy.FieldOfStudyId, err.Error())
	}
}

func (paperFieldsOfStudy *PaperFieldsOfStudy) GetPaperIdFromDB(fieldOfStudyId int64) []int64 {
	conn := rdbPool.Get()
	defer conn.Close()
	reply, err := redis.Int64s(conn.Do("SMEMBERS", strconv.FormatInt(fieldOfStudyId, 10)))
	if err != nil {
		log.Printf("Get %d key not found!", paperFieldsOfStudy.FieldOfStudyId)
	}
	return reply
}
