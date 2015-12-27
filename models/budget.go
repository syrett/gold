package models

/*-------------------------------------------------------------------
* @author liyouyou <youyou.li78@gmail.com>
#Time-stamp: <2015-12-23 22:50:07>
* @doc
* budget.go
* @end
* Created : 2015/12/10 14:42:58 liyouyou

-------------------------------------------------------------------*/

import (
	"log"

	"github.com/syrett/gold/lib"
	"gopkg.in/mgo.v2/bson"
)

type Budget struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	BudgetId  string        `bson:"budgetid"`
	Year      int           `bson:"year"`
	SortType  string        `bson:"sorttype"` //smallsort, bigsort
	SortId    string        `bson:"sortid"`
	BudgetNum int64         `bson:"budget"`
}

func (b *Budget) Save(year int, sort_type string, year_budget []Budget) (err error) {
	s, _ := MgoDB.GetSession()
	defer s.Close()
	c := s.DB(MgoDbName).C(TBuget)

	c.RemoveAll(bson.M{"year": year, "sorttype": sort_type})
	for _, budget := range year_budget {
		budget.Year = year
		budget.SortType = sort_type
		budget.BudgetId = lib.GrandStr(16)
		err = c.Insert(budget)
		if err != nil {
			log.Printf("[ERROR] budget save failed:%v, year:%v\n", err, year)
			c.RemoveAll(bson.M{"year": year, "sorttype": sort_type})
			return
		}
	}
	return
}

func (b *Budget) List(filter bson.M) (budget_list []Budget, err error) {
	s, _ := MgoDB.GetSession()
	defer s.Close()
	c := s.DB(MgoDbName).C(TBuget)

	err = c.Find(filter).Sort("sortid").All(&budget_list)
	return
}
