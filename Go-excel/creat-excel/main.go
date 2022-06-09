package main

//创建表格
//创建表格前需要先引入 excel 库，我们以比较热门的 tealeg/xlsx 库为例。
import "github.com/tealeg/xlsx"

func main() {
	//首先创建一个空文件，拿到文件句柄。
	file := xlsx.NewFile()

	//创建一个名为人员信息收集的 sheet。
	sheet, err := file.AddSheet("人员信息收集")
	if err != nil{
		panic(err.Error())
	}

	//然后为该 sheet 创建一行，这行作为我们的表头。
	row := sheet.AddRow()

	//在该行中创建一个单元格。
	cell := row.AddCell()

	//现在给单元格填充内容，因为是表头，暂且叫姓名
	cell.Value = "姓名"

	//如何创建第二个单元格呢？原理相同，此处 cell 变量已定义，再创建新单元格只需赋值即可。
	cell = row.AddCell()
	cell.Value = "性别"

	//表头已经设置好了，可以开始创建第二行来填充内容了，方式与上述无差别。
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "张三"
	cell = row.AddCell()
	cell.Value = "男"

	//表格设置完成后，将该文件保存，文件名可自定义。
	err = file.Save("demo.xlsx")
	if err != nil {
		panic(err)
	}
}
