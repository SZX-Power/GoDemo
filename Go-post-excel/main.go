package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	file, err := xlsx.OpenFile("车辆vin.xlsx")
	if err !=nil{
		panic(err.Error())
	}


	for r := 1 ; r<100; r++{
		var vin string
		fmt.Println("请输入vin:" + vin)
		fmt.Scanln(&vin)
		fmt.Println("<你当前输入的vin值是:"+vin+",请确保完整！>")
		/*
			当action为get时候，浏览器用x-www-form-urlencoded的编码方式把form数据转换成一个字串（name1=value1&name2=value2…），
			然后把这个字串append到url后面，用?分割，加载这个新的url。
		 */
		resp, err := http.Post("http://211.94.119.53:19007/openapi/iov/apiplatform/subscription/validDate.json",
			"application/x-www-form-urlencoded",
			//strings.NewReader("token=b1089dff222546b0bedc6e2e94ae27d7&vin=LVBV3JBB2KY048009"))
			strings.NewReader("token=b1089dff222546b0bedc6e2e94ae27d7&vin="+vin))
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		fmt.Println(string(body))

		file.Sheets[0].Rows[r].Cells[0].Value = ""+vin
		file.Sheets[0].Rows[r].Cells[1].Value = ""+string(body)

		err = file.Save("车辆vin.xlsx")
		if err != nil{
			panic(err.Error())
		}
	}


}