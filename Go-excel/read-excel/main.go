package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
)

func main() {
	//output,err := xlsx.FileToSlice("E:\\Go\\goWorkspace\\src\\awesomeProject\\go_excel\\creat-excel\\demo.xlsx")
	output,err := xlsx.FileToSlice("E:\\Go\\goWorkspace\\src\\awesomeProject\\go_excel\\read-excel\\福田.xlsx")
	if err != nil{
		panic(err.Error())
	}
//只需将文件路径传入上述方法，即可自动读取并返回一个三维切片，我们来读取第一个 sheet 的第二行中的第一个单元格。
    log.Println(output[0][1][1])  //Output: 男

	//遍历输出全部信息
	for rowIndex,row := range output[0] {
		for cellIndex ,cell := range row{
			log.Println(fmt.Sprintf("第%d行，第%d个单元格：%s", rowIndex+1, cellIndex+1, cell))
		}
	}


}
