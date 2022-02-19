package logic

import (
	"geopaper/model"
	"log"
	"os"
	"strconv"
)

func GetPaperIdLogic(field string) {
	log.Println("Starting GetPaperIdLogic!")
	// 定位root学科代码
	var fieldOfStudy model.FieldsOfStudy
	rootFieldsOfStudies := fieldOfStudy.QueryByNormalizedNameAndLevel(field, 0)
	// 递归寻找子级学科代码
	geoFieldsIds := []int64{}
	for _, field := range rootFieldsOfStudies {
		// 将root学科加入geoFieldsIds
		geoFieldsIds = append(geoFieldsIds, field.FieldOfStudyId)
		// 将当前root学科的所有子孙学科加入geoFieldsIds
		geoFieldsIds = append(geoFieldsIds, getChildrensFieldsIds(field.FieldOfStudyId)...)
	}
	savePath := "GeoFieldsIds.txt"
	f, err := os.OpenFile(savePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer f.Close()
	for _, item := range geoFieldsIds {
		f.WriteString(strconv.FormatInt(item, 10) + "\n")
	}
	log.Println("Finish GetPaperIdLogic!")

}

// 递归寻找该父学科的所有子孙学科的id
func getChildrensFieldsIds(fieldId int64) []int64 {
	var childrenStudy model.FieldOfStudyChildren
	fieldOfChildrenStudies := childrenStudy.QueryById(fieldId)
	// 递归出口
	if len(fieldOfChildrenStudies) == 0 {
		return nil
	}
	// 递归体
	resFields := []int64{}
	for _, child := range fieldOfChildrenStudies {
		// 将当前子学科的id添加到res中
		resFields = append(resFields, child.ChildFieldOfStudyId)
		// 递归获取当前子学科的子学科同样追加到res中
		resChildrenFieldIds := getChildrensFieldsIds(child.ChildFieldOfStudyId)
		if resChildrenFieldIds != nil {
			resFields = append(resFields, resChildrenFieldIds...)
		}

	}
	return resFields

}
