package lib

/*-------------------------------------------------------------------
* @copyright 2015 有量(上海)信息技术有限公司
* @author liyouyou <youyou.li78@gmail.com>
#Time-stamp: <liyouyou 2015-08-14 16:22:20>
* @doc
* xlsx.go
* @end
* Created : 2015/08/14 08:17:37 liyouyou

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
		for _, row := range sheet.Rows {
			for k, cell := range row.Cells {
				fmt.Printf("cell:%v, %s\n", k, cell.String())
			}
		}
	}
}
