package models

/*-------------------------------------------------------------------
* @author liyouyou <youyou.li78@gmail.com>
#Time-stamp: <2015-12-21 23:42:25>
* @doc
* small_sort.go
* @end
* Created : 2015/11/18 13:57:50 liyouyou

-------------------------------------------------------------------*/

import (
	"github.com/syrett/gold/lib"
	"gopkg.in/mgo.v2/bson"
)

type SmallSort struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Sid      string        `bson:"sid"` //small sort id
	Bid      string        `bson:"bid"` //big sort id
	SortName string        `bson:"sortname"`
}

func (small *SmallSort) Save(bid string, sortName string) (sid string,
	err error) {

	s, _ := MgoDB.GetSession()
	defer s.Close()
	c := s.DB("gold").C("smallsort")

	newSmallSort := &SmallSort{
		Bid:      bid,
		SortName: sortName,
	}

	err = c.Find(bson.M{"bid": bid, "sortname": sortName}).One(&newSmallSort)
	if err == nil {
		sid = newSmallSort.Sid
		return
	}

	newSmallSort.Sid = lib.GrandStr(8)
	err = c.Insert(newSmallSort)
	sid = newSmallSort.Sid
	return
}

func (small *SmallSort) List(filter bson.M) (sort_list []SmallSort, err error) {
	s, _ := MgoDB.GetSession()
	defer s.Close()
	c := s.DB(MgoDbName).C(TSmallSort)

	err = c.Find(filter).Sort("id").All(&sort_list)
	return
}
