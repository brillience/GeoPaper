package logic

import (
	"fmt"
	"geopaper/model"
)

func GetPaperIdLogic(field string) {
	// 定位root学科代码
	var fieldOfStudy model.FieldsOfStudy
	fieldsOfStudies := fieldOfStudy.QueryByNormalizedNameAndLevel(field, 0)
	fmt.Println(fieldsOfStudies)
	// 递归寻找次级学科代码，同时查找对应的PaperId更新到数据库

}
