package finance

/*-------------------------------------------------------------------
* @copyright 2015 有量(上海)信息技术有限公司
* @author liyouyou <youyou.li78@gmail.com>
#Time-stamp: <2015-12-23 23:00:53>
* @doc
* budget.go
* @end
* Created : 2015/12/10 14:51:47 liyouyou

-------------------------------------------------------------------*/

import (
	"fmt"

	"github.com/syrett/gold/controllers"
	"github.com/syrett/gold/lib"
	"github.com/syrett/gold/models"
	"gopkg.in/mgo.v2/bson"
)

type Budget struct {
	controllers.Finance
}

type BudgetSmallSort struct {
	Sid       string
	SortName  string
	BudgetNum int64
}

type BudgetBigSort struct {
	Bid                 string
	SortName            string
	BudgetNum           int64
	BudgetSmallSortList []BudgetSmallSort
}

func (this *Budget) List() {
	year, _ := this.GetInt("year")

	budget := &models.Budget{}

	big := &models.BigSort{}
	bigSortList, _ := big.List(nil)
	budget_list, _ := budget.List(bson.M{"year": int(year)})

	budgetBigSortList := make([]BudgetBigSort, 0)
	budgetBigSort := BudgetBigSort{}
	for _, bigSort := range bigSortList {
		budgetBigSort.Bid = bigSort.Bid
		budgetBigSort.SortName = bigSort.SortName
		budgetBigSort.BudgetNum = this.getBudgetById(bigSort.Bid,
			"bigsort", budget_list)

		budgetSmallSortList := make([]BudgetSmallSort, 0)
		budgetSmallSort := BudgetSmallSort{}
		small := &models.SmallSort{}
		smallSortList, _ := small.List(bson.M{"bid": bigSort.Bid})
		for _, smallSort := range smallSortList {
			budgetSmallSort.Sid = smallSort.Sid
			budgetSmallSort.SortName = smallSort.SortName
			budgetSmallSort.BudgetNum = this.getBudgetById(smallSort.Sid,
				"smallsort", budget_list)

			budgetSmallSortList = append(budgetSmallSortList, budgetSmallSort)
		}

		budgetBigSort.BudgetSmallSortList = budgetSmallSortList
		budgetBigSortList = append(budgetBigSortList, budgetBigSort)
	}

	this.Data["Year"] = year
	this.Data["BudgetList"] = budgetBigSortList
	this.TplNames = "finance/budget_list.html"
}

func (this *Budget) getBudgetById(id, sort_type string,
	budget_list []models.Budget) (budgetNum int64) {

	for _, budget := range budget_list {
		if budget.SortId == id &&
			budget.SortType == sort_type {

			budgetNum = budget.BudgetNum
			return
		}
	}

	return
}

func (this *Budget) Save() {
	var err error
	fmt.Printf("data:%+v\n", this.Ctx.Request.PostForm)
	fmt.Printf("year:%+v\n", this.GetString("year"))
	fmt.Printf("year:%+v\n", this.GetStrings("sort_id"))
	year, _ := this.GetInt("year")
	sort_ids := this.GetStrings("sort_id")
	budget_arr := this.GetStrings("budget")
	sort_type := this.GetString(":sorttype")

	year_budget := make([]models.Budget, 0)
	budget := models.Budget{
		Year:     int(year),
		SortType: sort_type,
	}
	for index, id := range sort_ids {
		budget.SortId = id
		budget.BudgetNum = lib.Be_int(budget_arr[index])
		year_budget = append(year_budget, budget)
	}
	switch sort_type {
	case "smallsort":
		err = budget.Save(int(year), sort_type, year_budget)
	case "bigsort":
		// TODO
	}

	if err != nil {
		this.Data["ERROR"] = err
		this.TplNames = "error.html"
	} else {
		this.Redirect("list?year="+lib.Be_string(year), 302)
	}
}

type BudgetEditData struct {
	Bid           string
	BigSortName   string
	SmallSortList []models.SmallSort
}

func (this *Budget) Edit() {
	//	year, _ := this.GetInt("year")

	sort_type := this.GetString(":sorttype")
	big := &models.BigSort{}
	bigSortList, _ := big.List(nil)
	resultDataList := make([]BudgetEditData, 0)

	for _, bigSort := range bigSortList {
		data := BudgetEditData{
			Bid:         bigSort.Bid,
			BigSortName: bigSort.SortName,
		}
		switch sort_type {
		case "smallsort":
			small := &models.SmallSort{}
			smallSortList, _ := small.List(bson.M{"bid": bigSort.Bid})
			data.SmallSortList = smallSortList
		case "bigsort":
			// TODO
		}
		resultDataList = append(resultDataList, data)
	}
	this.Data["BudgetEditDataList"] = resultDataList

	this.TplNames = "finance/budget_edit.html"
}
