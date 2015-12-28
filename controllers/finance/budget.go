package finance

/*-------------------------------------------------------------------
* @copyright 2015 有量(上海)信息技术有限公司
* @author liyouyou <youyou.li78@gmail.com>
#Time-stamp: <2015-12-28 00:14:58>
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

	budgetBigSortList, err := this.listSortBudget(int(year))
	if err != nil {
		this.TplNames = "error.html"
	} else {
		this.Data["Year"] = year
		this.Data["BudgetList"] = budgetBigSortList
		this.TplNames = "finance/budget_list.html"
	}
}

func (this *Budget) Save() {
	var err error
	fmt.Printf("data:%+v\n", this.Ctx.Request.PostForm)
	fmt.Printf("year:%+v\n", this.GetString("year"))
	fmt.Printf("year:%+v\n", this.GetStrings("sort_id"))

	year, _ := this.GetInt("year")
	bid_list := this.GetStrings("bid")
	sid_list := this.GetStrings("sid")
	bigBudget_arr := this.GetStrings("b_budgetnum")
	smallBudget_arr := this.GetStrings("s_budgetnum")

	year_budget := make([]models.Budget, 0)
	budget := models.Budget{
		Year: int(year),
	}
	for index, id := range sid_list {
		budget.SortId = id
		budget.SortType = "smallsort"
		budget.BudgetNum = lib.Be_int(smallBudget_arr[index])
		year_budget = append(year_budget, budget)
	}

	for index, id := range bid_list {
		budget.SortId = id
		budget.SortType = "bigsort"
		budget.BudgetNum = lib.Be_int(bigBudget_arr[index])
		year_budget = append(year_budget, budget)
	}

	err = budget.Save(int(year), year_budget)

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

	year, _ := this.GetInt("year")

	budgetBigSortList, err := this.listSortBudget(int(year))
	if err != nil {
		this.TplNames = "error.html"
	} else {
		this.Data["Year"] = year
		this.Data["BudgetList"] = budgetBigSortList
		this.TplNames = "finance/budget_edit.html"
	}
}

/*================================================================================*/

func (this *Budget) listSortBudget(year int) (budgetBigSortList []BudgetBigSort,
	err error) {

	budget := &models.Budget{}

	big := &models.BigSort{}
	bigSortList, err := big.List(nil)
	if err != nil {
		return
	}

	budget_list, err := budget.List(bson.M{"year": year})
	if err != nil {
		return
	}

	budgetBigSortList = make([]BudgetBigSort, 0)
	budgetBigSort := BudgetBigSort{}
	for _, bigSort := range bigSortList {
		budgetBigSort.Bid = bigSort.Bid
		budgetBigSort.SortName = bigSort.SortName
		budgetBigSort.BudgetNum = this.getBudgetNumById(bigSort.Bid,
			"bigsort", budget_list)

		budgetSmallSortList := make([]BudgetSmallSort, 0)
		budgetSmallSort := BudgetSmallSort{}
		small := &models.SmallSort{}
		smallSortList, _ := small.List(bson.M{"bid": bigSort.Bid})
		for _, smallSort := range smallSortList {
			budgetSmallSort.Sid = smallSort.Sid
			budgetSmallSort.SortName = smallSort.SortName
			budgetSmallSort.BudgetNum = this.getBudgetNumById(smallSort.Sid,
				"smallsort", budget_list)

			budgetSmallSortList = append(budgetSmallSortList, budgetSmallSort)
		}

		budgetBigSort.BudgetSmallSortList = budgetSmallSortList
		budgetBigSortList = append(budgetBigSortList, budgetBigSort)
	}

	return
}

func (this *Budget) getBudgetNumById(id, sort_type string,
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
