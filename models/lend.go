package models

/*-------------------------------------------------------------------
* @copyright 2015 有量(上海)信息技术有限公司
* @author liyouyou <youyou.li78@gmail.com>
#Time-stamp: <liyouyou 2015-08-17 10:21:10>
* @doc
* lend.go
* @end
* Created : 2015/08/17 02:04:56 liyouyou

-------------------------------------------------------------------*/

import (
	"fmt"
	"log"
	"time"

	"github.com/tealeg/xlsx"
	"gopkg.in/mgo.v2/bson"
)

type Lend struct {
	LendType    string //借贷类型:借出:lend, 借入:borrow, 收款: collection, 还款:repayment
	Created_at  int64
	LendAccount string //借贷账户
	Account     string //账户
	Amount      int64
	Interest    int64 //利息
	Note        string
}

func (lend *Lend) DecodeAndSave(sheet *xlsx.Sheet) {
	s, _ := MgoDB.GetSession()
	defer s.Close()
	c := s.DB("gold").C("lend")
	c.RemoveAll(bson.M{"created_at": bson.M{"$gt": 0}})

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	for k, row := range sheet.Rows {
		if k > 0 {
			row_ex := &Lend{}
			for k, cell := range row.Cells {
				fmt.Printf("cell:%v, %s\n", k, cell.String())
				switch k {
				case 0:
					switch cell.String() {
					case "借出":
						row_ex.LendType = "lend"
					case "借入":
						row_ex.LendType = "borrow"
					}
				case 1:
					t, _ := time.ParseInLocation("2006\\-01\\-02\\ 15:04:05", cell.String(), location)
					row_ex.Created_at = t.Unix()
				case 2:
					row_ex.LendAccount = cell.String()
				case 3:
					row_ex.Account = cell.String()
				case 4:
					f, _ := cell.Float()
					row_ex.Amount = int64(f * 100)
				case 5:
					row_ex.Note = cell.String()
				}

			}
			log.Printf("row ex:%+v\n", row_ex)
			c.Insert(row_ex)
		}

	}
}

func (lend *Lend) DecodeAndSaveRepay(sheet *xlsx.Sheet) {
	s, _ := MgoDB.GetSession()
	defer s.Close()
	c := s.DB("gold").C("lend")
	c.RemoveAll(bson.M{"created_at": bson.M{"$gt": 0}})

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	for k, row := range sheet.Rows {
		if k > 0 {
			row_ex := &Lend{}
			for k, cell := range row.Cells {
				fmt.Printf("cell:%v, %s\n", k, cell.String())
				switch k {
				case 0:
					switch cell.String() {
					case "收款":
						row_ex.LendType = "collection"
					case "还款":
						row_ex.LendType = "repayment"
					}
				case 1:
					t, _ := time.ParseInLocation("2006\\-01\\-02\\ 15:04:05", cell.String(), location)
					row_ex.Created_at = t.Unix()
				case 2:
					row_ex.LendAccount = cell.String()
				case 3:
					row_ex.Account = cell.String()
				case 4:
					f, _ := cell.Float()
					row_ex.Amount = int64(f * 100)
				case 5:
					f, _ := cell.Float()
					row_ex.Interest = int64(f * 100)
				case 6:
					row_ex.Note = cell.String()
				}

			}
			log.Printf("row ex:%+v\n", row_ex)
			c.Insert(row_ex)
		}

	}
}
