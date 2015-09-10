package controllers

/*-------------------------------------------------------------------
* @author liyouyou <youyou.li78@gmail.com>
#Time-stamp: <liyouyou 2015-08-26 13:19:37>
* @doc
* angularjs.go
* @end
* Created : 2015/08/26 05:14:54 liyouyou

-------------------------------------------------------------------*/

type Angularjs struct {
	GoldController
}

func (this *Angularjs) Show() {
	this.TplNames = "angularjs.html"
}
