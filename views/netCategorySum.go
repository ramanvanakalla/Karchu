package views

import (
	"fmt"
	"strings"
)

type NetCategorySum struct {
	Category  string
	NetAmount int
}

type ByNetAmountDesc []NetCategorySum

func (a ByNetAmountDesc) Len() int      { return len(a) }
func (a ByNetAmountDesc) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByNetAmountDesc) Less(i, j int) bool {
	return a[i].NetAmount > a[j].NetAmount
}

func (NetCategorySumObject *NetCategorySum) ToString() string {
	categorySumString := fmt.Sprintf("%s-%d\n", NetCategorySumObject.Category, NetCategorySumObject.NetAmount)
	categorySumString = strings.ReplaceAll(categorySumString, "\n", "")
	return categorySumString
}
