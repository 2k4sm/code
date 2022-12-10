package main

import (
	conv "temptables/conv"
	table "temptables/table"
)

func main() {

	table.DrawTable(conv.CelConv())

	table.DrawTable(conv.FarhenConv())
}
