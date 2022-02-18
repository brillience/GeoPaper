package util

import (
	"geopaper/model"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/valyala/tsvreader"
)

func PushFieldsOfStudyToDB(filePath string, batchSize int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Open file %s error: %s \n", filePath, err.Error())
	}
	defer file.Close()
	reader := tsvreader.New(file)
	items := []*model.FieldsOfStudy{}
	for reader.Next() {
		FieldOfStudyId := reader.Int64()
		Rank := reader.String()
		NormalizedName := reader.String()
		DisplayName := reader.String()
		MainType := reader.String()
		Level := reader.Int()
		PaperCount := reader.Int64()
		PaperFamilyCount := reader.Int64()
		CitationCount := reader.Int64()
		CreateDate := reader.String()
		item := model.FieldsOfStudy{
			FieldOfStudyId:   FieldOfStudyId,
			Rank:             Rank,
			NormalizedName:   NormalizedName,
			DisplayName:      DisplayName,
			MainType:         MainType,
			Level:            Level,
			PaperCount:       PaperCount,
			PaperFamilyCount: PaperFamilyCount,
			CitationCount:    CitationCount,
			CreateDate:       CreateDate,
		}
		items = append(items, &item)
	}
	if err := reader.Error(); err != nil {
		log.Fatalln(err.Error())
	}
	var fieldsOfStudy model.FieldsOfStudy
	fieldsOfStudy.CreateInBatches(items, 100)
}

func PushFieldsOfStudyChildrenToDB(filePath string, batchSize int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Open file %s error: %s \n", filePath, err.Error())
	}
	defer file.Close()
	reader := tsvreader.New(file)
	items := []*model.FieldOfStudyChildren{}
	for reader.Next() {
		FieldOfStudyId := reader.Int64()
		next := reader.String()
		next = strings.ReplaceAll(next, "\r", "")
		ChildFieldOfStudyId, err1 := strconv.ParseInt(next, 10, 64)
		if err1 != nil {
			log.Fatalln("Parse int64 err:", err1.Error())
		}
		item := model.FieldOfStudyChildren{
			FieldOfStudyId:      FieldOfStudyId,
			ChildFieldOfStudyId: ChildFieldOfStudyId,
		}
		items = append(items, &item)

	}
	if err := reader.Error(); err != nil {
		log.Fatalln(err.Error())
	}
	var fieldChildren model.FieldOfStudyChildren
	fieldChildren.CreateInBatches(items, 100)
}

func PushPaperFieldsOfStudyToDB(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Open file %s error: %s \n", filePath, err.Error())
	}
	defer file.Close()
	reader := tsvreader.New(file)
	var wg sync.WaitGroup
	for reader.Next() {
		PaperId := reader.Int64()
		FieldOfStudyId := reader.Int64()
		next := reader.String()
		next = strings.ReplaceAll(next, "\r", "")
		Score, err1 := strconv.ParseFloat(next, 64)
		if err1 != nil {
			log.Fatalln("Parse int64 err:", err1.Error())
		}
		item := model.PaperFieldsOfStudy{
			PaperId:        PaperId,
			FieldOfStudyId: FieldOfStudyId,
			Score:          Score,
		}
		wg.Add(1)
		go func(i model.PaperFieldsOfStudy) {
			defer wg.Done()
			i.SetPaperIdToDB()

		}(item)
	}
	if err := reader.Error(); err != nil {
		log.Fatalln(err.Error())
	}
	wg.Wait()
}
