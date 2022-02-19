package main

import "geopaper/internal/logic"

func main() {
	getPaperIdLgic := logic.RegisterLogic(logic.GetPaperIdLogic)
	getPaperIdLgic.Run("geology")
}
