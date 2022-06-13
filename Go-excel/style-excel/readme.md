**样式设置**
该开源库不仅支持内容的编辑，还支持表格的样式设置，样式统一由结构体 Style 来负责。

```go
type Style struct {
    Border          Border
    Fill            Fill
    Font            Font
    ApplyBorder     bool
    ApplyFill       bool
    ApplyFont       bool
    ApplyAlignment  bool
    Alignment       Alignment
    NamedStyleIndex *int
}
```


拿上述生成的文件为例，假如我要将姓名所在单元格居中，首先要实例化样式对象。

```go
style := xlsx.NewStyle()
```


赋值居中属性。

```go
style.Alignment = xlsx.Alignment{
  Horizontal:   "center",
  Vertical:     "center",
}
```

给第一行第一个单元格设置样式。

```go
file.Sheets[0].Rows[0].Cells[0].SetStyle(style)
```


与修改表格处理逻辑相同，最后保存文件。

```go
err = file.Save("demo.xlsx")
if err != nil {
  panic(err.Error())
}
```


打开预览，可以看到文字已经上下左右居中。

![image-20220613154138419]([GoDemo\Go-excel\style-excel](https://github.com/SZX-Power/GoDemo/edit/main/Go-excel/style-excel/)\readme.assets\image-20220613154138419.png)

同理，可以修改文字颜色和背景，同样通过 style 的属性来设置。

```go
style.Font.Color = xlsx.RGB_Dark_Red
style.Fill.BgColor = xlsx.RGB_Dark_Green
```


![image-20220613154151160](https://github.com/SZX-Power/GoDemo/edit/main/Go-excel/style-excel/readme.assets\image-20220613154151160.png)

其他还有很多属性可以设置，比如合并单元格、字体、大小等等，大家可以自行测试。
