package main

import "geopaper/internal/util"

func main() {
	// util.PushFieldsOfStudyToDB("MagData/FieldsOfStudy.txt",100)
//	util.PushFieldsOfStudyChildrenToDB("MagData/FieldOfStudyChildren.txt", 100)
	util.PushPaperFieldsOfStudyToDB("MagData/PaperFieldsOfStudy.txt")
}
