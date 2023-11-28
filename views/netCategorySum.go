package views

import "fmt"

type NetCategorySum struct {
	Category  string
	NetAmount int
}

func (NetCategorySumObject *NetCategorySum) ToString() string {
	return fmt.Sprintf("%-10s%-10d\n", NetCategorySumObject.Category, NetCategorySumObject.NetAmount)
}
