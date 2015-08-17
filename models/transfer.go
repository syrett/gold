package models

/*-------------------------------------------------------------------
* @copyright 2015 有量(上海)信息技术有限公司
* @author liyouyou <youyou.li78@gmail.com>
#Time-stamp: <liyouyou 2015-08-17 10:24:34>
* @doc
* transfer.go
* @end
* Created : 2015/08/17 01:57:57 liyouyou

-------------------------------------------------------------------*/

import (
	"fmt"
	"log"
	"time"

	"github.com/tealeg/xlsx"
	"gopkg.in/mgo.v2/bson"
)

type Transfer struct {
	OutAccount string //转出账户
	OutAmount  int64  //转出金额
	InAccount  string //转入账户
	InAmount   int64  //转入金额
	Created_at int64
	Note       string
}

func (trans *Transfer) DecodeAndSave(sheet *xlsx.Sheet) {
	s, _ := MgoDB.GetSession()
	defer s.Close()
	c := s.DB("gold").C("transfer")
	c.RemoveAll(bson.M{"created_at": bson.M{"$gt": 0}})

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	for k, row := range sheet.Rows {
		if k > 0 {
			row_ex := &Transfer{}
			for k, cell := range row.Cells {
				fmt.Printf("cell:%v, %s\n", k, cell.String())
				switch k {
				case 0:
					row_ex.OutAccount = cell.String()
				case 2:
					f, _ := cell.Float()
					row_ex.OutAmount = int64(f * 100)
				case 3:
					row_ex.InAccount = cell.String()
				case 5:
					f, _ := cell.Float()
					row_ex.InAmount = int64(f * 100)
				case 6:
					t, _ := time.ParseInLocation("2006\\-01\\-02\\ 15:04:05", cell.String(), location)
					row_ex.Created_at = t.Unix()
				case 7:
					row_ex.Note = cell.String()
				}

			}
			log.Printf("row ex:%+v\n", row_ex)
			c.Insert(row_ex)
		}

	}
}
