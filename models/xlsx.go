package models

/*-------------------------------------------------------------------
* @copyright 2015 有量(上海)信息技术有限公司
* @author liyouyou <youyou.li78@gmail.com>
#Time-stamp: <liyouyou 2015-08-17 10:21:27>
* @doc
* xlsx.go
* @end
* Created : 2015/08/14 08:52:19 liyouyou

-------------------------------------------------------------------*/

import (
	"fmt"
	"log"

	"github.com/tealeg/xlsx"
)

func XlsxDecode(file string) {
	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		log.Printf("[ERROR] open xls file:%s failed:%v\n", file, err)
		return
	}
	for _, sheet := range xlFile.Sheets {
		log.Printf("sheet:%+v\n", sheet.Name)
		switch sheet.Name {
		case "支出":
			expense := &Expense{}
			expense.DecodeAndSave(sheet)
		case "收入":
			income := &Income{}
			income.DecodeAndSave(sheet)
		case "转账":
			transfer := &Transfer{}
			transfer.DecodeAndSave(sheet)
		case "借入借出":
			lend := &Lend{}
			lend.DecodeAndSave(sheet)
		case "收款还款":
			lend := &Lend{}
			lend.DecodeAndSaveRepay(sheet)
		}
		for _, row := range sheet.Rows {
			for k, cell := range row.Cells {
				fmt.Printf("cell:%v, %s\n", k, cell.String())
			}
		}
	}
}
