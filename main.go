package main

import (
	"geopaper/internal/util"
)

func main() {
	// getPaperIdLgic := logic.RegisterLogic(logic.GetPaperIdLogic)
	// getPaperIdLgic.Run("geology")
	util.PushPaperFieldsOfStudyToDB("MagData/PaperFieldsOfStudy.txt")
}
