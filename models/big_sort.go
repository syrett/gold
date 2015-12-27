package models

/*-------------------------------------------------------------------
* @author liyouyou <youyou.li78@gmail.com>
#Time-stamp: <2015-12-21 23:41:32>
* @doc
* big_sort.go
* @end
* Created : 2015/11/18 13:41:39 liyouyou

-------------------------------------------------------------------*/

import (
	"github.com/syrett/gold/lib"
	"gopkg.in/mgo.v2/bson"
)

type BigSort struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Bid      string        `bson:"bid"` //bigsortid
	SortName string        `bson:"sortname"`
}

func (big *BigSort) Save(sortName string) (bid string, err error) {
	s, _ := MgoDB.GetSession()
	defer s.Close()
	c := s.DB("gold").C("bigsort")

	newBigSort := &BigSort{
		SortName: sortName}

	err = c.Find(bson.M{"sortname": sortName}).One(&newBigSort)
	if err == nil {
		bid = newBigSort.Bid
		return
	}

	newBigSort.Bid = lib.GrandStr(8)
	err = c.Insert(newBigSort)
	bid = newBigSort.Bid
	return
}

func (big *BigSort) List(filter bson.M) (sort_list []BigSort, err error) {
	s, _ := MgoDB.GetSession()
	defer s.Close()
	c := s.DB(MgoDbName).C(TBigSort)

	err = c.Find(filter).Sort("id").All(&sort_list)
	return
}
