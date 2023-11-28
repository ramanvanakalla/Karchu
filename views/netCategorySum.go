package views

import (
	"fmt"
	"strings"
)

type NetCategorySum struct {
	Category  string
	NetAmount int
}

func (NetCategorySumObject *NetCategorySum) ToString() string {
	categorySumString := fmt.Sprintf("%-75s%-10d\n", NetCategorySumObject.Category, NetCategorySumObject.NetAmount)
	categorySumString = strings.ReplaceAll(categorySumString, "\n", "")
	return categorySumString
}
