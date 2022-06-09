package main

import "github.com/tealeg/xlsx"

func main() {
	file ,err := xlsx.OpenFile("E:\\Go\\goWorkspace\\src\\awesomeProject\\go_excel\\read-excel\\福田.xlsx")
	if err !=nil{
		panic(err.Error())
	}
	file.Sheets[0].Rows[1].Cells[2].Value = "史振兴"

	err = file.Save("E:\\Go\\goWorkspace\\src\\awesomeProject\\go_excel\\read-excel\\福田.xlsx")
	if err !=nil {
		panic(err.Error())
	}

}
