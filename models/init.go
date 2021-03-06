package models

/*-------------------------------------------------------------------
* @copyright 2015 有量(上海)信息技术有限公司
* @author liyouyou <youyou.li78@gmail.com>
#Time-stamp: <liyouyou 2015-11-18 22:42:41>
* @doc
* init.go
* @end
* Created : 2015/08/14 09:35:25 liyouyou

-------------------------------------------------------------------*/

import (
	"github.com/syrett/gold/lib"
)

var (
	MgoDB *lib.MongoDB
)

func init() {
	ObBigSort = &BigSort{}
	ObSmallSort = &SmallSort{}
}
