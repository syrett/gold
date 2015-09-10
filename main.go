package main

import (
	"github.com/astaxie/beego"
	"github.com/syrett/gold/lib"
	"github.com/syrett/gold/models"
	_ "github.com/syrett/gold/routers"
)

func main() {
	models.MgoDB = lib.NewMongoDB(beego.AppConfig.String("mongo::dial_url"))
	err := models.MgoDB.Connection()
	if err != nil {
		panic(err)
	}
	models.XlsxDecode("/Users/liyouyou/go/src/github.com/syrett/gold/xls/wacai.xlsx")
	beego.Run()
}
