package main

import (
	"github.com/astaxie/beego"
	"github.com/syrett/gold/lib"
	"github.com/syrett/gold/models"
	_ "github.com/syrett/gold/routers"
)

// 这是master后改
// 这是master 2 后改
func main() {
	models.MgoDB = lib.NewMongoDB(beego.AppConfig.String("mongo::dial_url"))
	err := models.MgoDB.Connection()
	if err != nil {
		panic(err)
	}
	models.XlsxDecode("/Users/liyouyou/go/src/github.com/syrett/gold/xls/wacai.xlsx")
	beego.Run()
}
