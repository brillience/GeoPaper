package logic

type Logic func(string)

// 传入学科名字; 获取该学科的PaperId，更新到数据库
func (f Logic) Run(field string) {
	f(field)
}

func RegisterLogic(l func(string)) Logic {
	return Logic(l)
}
