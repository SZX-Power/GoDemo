package main

import "github.com/tealeg/xlsx"

func main() {
    file, err := xlsx.OpenFile("demo.xlsx")
    if err != nil {
        panic(err.Error())
    }
    style := xlsx.NewStyle()
    style.Font.Color = xlsx.RGB_Dark_Red
    style.Fill.BgColor = xlsx.RGB_Dark_Green
    style.Alignment = xlsx.Alignment{
        Horizontal:   "center",
        Vertical:     "center",
    }
    file.Sheets[0].Rows[0].Cells[0].SetStyle(style)
    err = file.Save("demo.xlsx")
    if err != nil {
        panic(err.Error())
    }
}