package views

import "fmt"

type NetCategorySum struct {
	Category  string
	NetAmount int
}

func (NetCategorySumObject *NetCategorySum) ToString() string {
	return fmt.Sprintf("%-30s%-1d\n", NetCategorySumObject.Category, NetCategorySumObject.NetAmount)
}
