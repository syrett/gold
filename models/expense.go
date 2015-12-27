package models

/*-------------------------------------------------------------------
* @copyright 2015 有量(上海)信息技术有限公司
* @author liyouyou <youyou.li78@gmail.com>
#Time-stamp: <2015-12-21 23:47:03>
* @doc
* 支出
* @end
* Created : 2015/08/14 08:29:55 liyouyou

-------------------------------------------------------------------*/
import (
	"fmt"
	"log"
	"time"

	"github.com/syrett/gold/lib"
	"github.com/tealeg/xlsx"
	"gopkg.in/mgo.v2/bson"
)

type Expense struct {
	ExpenseId  string
	BigSort    string //支出大类
	SmallSort  string //支出小类
	Bid        string //支出大类
	Sid        string //支出小类
	Account    string //账号
	Project    string //项目
	Merchant   string //商家
	IsRefund   int    //是否报销
	Created_at int64
	Amount     int64
	Member     string
	Note       string
}

func (ex *Expense) DecodeAndSave(sheet *xlsx.Sheet) {
	s, _ := MgoDB.GetSession()
	defer s.Close()
	c := s.DB("gold").C("expense")
	c.RemoveAll(bson.M{"created_at": bson.M{"$gt": 0}})

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	for k, row := range sheet.Rows {
		if k > 0 {

			row_ex := &Expense{}
			for k, cell := range row.Cells {
				fmt.Printf("cell:%v, %s\n", k, cell.String())
				switch k {
				case 0:
					row_ex.BigSort = cell.String()
				case 1:
					row_ex.SmallSort = cell.String()
				case 2:
					row_ex.Account = cell.String()
				case 4:
					row_ex.Project = cell.String()
				case 5:
					row_ex.Merchant = cell.String()
				case 6:
					switch cell.String() {
					case "非报销":
						row_ex.IsRefund = 0
					case "报销":
						row_ex.IsRefund = 1
					}
				case 7:
					t, _ := time.ParseInLocation("2006\\-01\\-02\\ 15:04:05", cell.String(), location)
					row_ex.Created_at = t.Unix()
				case 8:
					f, _ := cell.Float()
					row_ex.Amount = int64(f * 100)
				case 9:
					row_ex.Member = cell.String()
				case 10:
					row_ex.Note = cell.String()
				}

			}
			log.Printf("row ex:%+v\n", row_ex)
			bigId, _ := ObBigSort.Save(row_ex.BigSort)
			smallId, _ := ObSmallSort.Save(bigId, row_ex.SmallSort)
			row_ex.Bid = bigId
			row_ex.Sid = smallId
			row_ex.ExpenseId = lib.GrandStr(16)
			c.Insert(row_ex)

		}

	}
}
