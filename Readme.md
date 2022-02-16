# 简介
本程序主要用于提取MAG数据中地学的文献对应的PaperId。主要思路如下：
1. 从`FieldsOfStudy.txt`中找到`Geology`的一级学科代码：`127313418`。
2. 一级学科下边还有子学科；子学科下边还有子学科；...；如此递归下去找到所有地学的学科代码。父子学科关系在`FieldOfStudyChrildren.txt`中。
3. 所有关于地学的学科代码拿到后，在`PaperFieldsOfStudy.txt`中找出对应的PaperId。
# 环境
- Golang
- Mysql
- Redis