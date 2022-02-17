package util

import (
	"geopaper/model"
	"github.com/valyala/tsvreader"
	"log"
	"os"
	"sync"
)

func PushFieldsOfStudyToDB(filePath string, batchSize int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Open file %s error: %s \n", filePath, err.Error())
	}
	defer file.Close()
	reader := tsvreader.New(file)
	items := []*model.FieldsOfStudy{}
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1)
	stopSignal := make(chan int)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stopSignal:
				if len(items)!=0{
					model.FieldsOfStudy{}.CreateInBatches(items, batchSize)
				}
				break
			default:
				if len(items) == batchSize {
					model.FieldsOfStudy{}.CreateInBatches(items, batchSize)
					items = []*model.FieldsOfStudy{}
					mutex.Unlock()
				}
			}
		}
	}()
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
		CreateDate := reader.Date()
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
		if len(items) < batchSize {
			items = append(items, &item)
		} else {
			mutex.Lock()
		}

	}
	if err := reader.Error(); err != nil {
		log.Fatalln(err.Error())
	}
	stopSignal <- 1
	wg.Wait()
}
func PushFieldsOfStudyChildren(filePath string, batchSize int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Open file %s error: %s \n", filePath, err.Error())
	}
	defer file.Close()
	reader := tsvreader.New(file)
	items := []*model.FieldOfStudyChildren{}
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1)
	stopSignal := make(chan int)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stopSignal:
				if len(items)!=0{
					model.FieldOfStudyChildren{}.CreateInBatches(items, batchSize)
				}
				break
			default:
				if len(items) == batchSize {
					model.FieldOfStudyChildren{}.CreateInBatches(items, batchSize)
					items = []*model.FieldOfStudyChildren{}
					mutex.Unlock()
				}
			}
		}
	}()
	for reader.Next() {
		FieldOfStudyId := reader.Int64()
		ChildFieldOfStudyId := reader.Int64()
		item := model.FieldOfStudyChildren{
			FieldOfStudyId:      FieldOfStudyId,
			ChildFieldOfStudyId: ChildFieldOfStudyId,
		}
		if len(items) < batchSize {
			items = append(items, &item)
		} else {
			mutex.Lock()
		}

	}
	if err := reader.Error(); err != nil {
		log.Fatalln(err.Error())
	}
	stopSignal <- 1
	wg.Wait()

}
