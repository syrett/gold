package models

/*-------------------------------------------------------------------
* @copyright 2015 有量(上海)信息技术有限公司
* @author liyouyou <youyou.li78@gmail.com>
#Time-stamp: <liyouyou 2015-08-17 10:24:24>
* @doc
* income.go
* @end
* Created : 2015/08/14 10:14:26 liyouyou

-------------------------------------------------------------------*/

import (
	"fmt"
	"log"
	"time"

	"github.com/tealeg/xlsx"
	"gopkg.in/mgo.v2/bson"
)

type Income struct {
	BigSort     string //收入大类
	Account     string //收入小类
	Project     string //项目
	PaidAccount string //付款方
	Created_at  int64
	Amount      int64
	Member      string
	Note        string
}

func (in *Income) DecodeAndSave(sheet *xlsx.Sheet) {
	s, _ := MgoDB.GetSession()
	defer s.Close()
	c := s.DB("gold").C("income")
	c.RemoveAll(bson.M{"created_at": bson.M{"$gt": 0}})

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	for k, row := range sheet.Rows {
		if k > 0 {
			row_ex := &Income{}
			for k, cell := range row.Cells {
				fmt.Printf("cell:%v, %s\n", k, cell.String())
				switch k {
				case 0:
					row_ex.BigSort = cell.String()
				case 1:
					row_ex.Account = cell.String()
				case 3:
					row_ex.Project = cell.String()
				case 4:
					row_ex.PaidAccount = cell.String()
				case 5:
					t, _ := time.ParseInLocation("2006\\-01\\-02\\ 15:04:05", cell.String(), location)
					row_ex.Created_at = t.Unix()
				case 6:
					f, _ := cell.Float()
					row_ex.Amount = int64(f * 100)
				case 7:
					row_ex.Member = cell.String()
				case 8:
					row_ex.Note = cell.String()
				}

			}
			log.Printf("row ex:%+v\n", row_ex)
			c.Insert(row_ex)
		}

	}
}
