package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	. "police-util/model"
	"strings"
	"time"
)

var TEMPLATE = "2006-01-02 15::04:05"

//创建 huobi Excel
func CreateExcel(file *excelize.File, sheetName, symbol string, data []HuobiKline) {
	sheetIndex := file.NewSheet(sheetName)
	file.SetCellValue(sheetName, "a1", "交易对")
	file.SetCellValue(sheetName, "b1", "开盘价")
	file.SetCellValue(sheetName, "c1", "收盘价")
	file.SetCellValue(sheetName, "d1", "最高价")
	file.SetCellValue(sheetName, "e1", "最低价")
	file.SetCellValue(sheetName, "f1", "K线生成时间")
	file.SetCellValue(sheetName, "g1", "K线类型")

	for i, v := range data {
		file.SetCellValue(sheetName, execCell("a%d", i+2), symbol)
		file.SetCellValue(sheetName, execCell("b%d", i+2), v.Open)
		file.SetCellValue(sheetName, execCell("c%d", i+2), v.Close)
		file.SetCellValue(sheetName, execCell("d%d", i+2), v.High)
		file.SetCellValue(sheetName, execCell("e%d", i+2), v.Low)
		file.SetCellValue(sheetName, execCell("f%d", i+2), time.Unix(v.Id, 0).Format(TEMPLATE))
		file.SetCellValue(sheetName, execCell("g%d", i+2), "日线")
	}
	file.SetActiveSheet(sheetIndex)
}

//创建 binance Excel
func CreateBinanceExcel(file *excelize.File, sheetName, symbol string, data [][]interface{}) {
	sheetIndex := file.NewSheet(sheetName)
	file.SetCellValue(sheetName, "a1", "交易对")
	file.SetCellValue(sheetName, "b1", "开盘价")
	file.SetCellValue(sheetName, "c1", "收盘价")
	file.SetCellValue(sheetName, "d1", "最高价")
	file.SetCellValue(sheetName, "e1", "最低价")
	file.SetCellValue(sheetName, "f1", "K线生成时间")
	file.SetCellValue(sheetName, "g1", "K线类型")
	for i, v := range data {
		file.SetCellValue(sheetName, execCell("a%d", i+2), symbol)
		file.SetCellValue(sheetName, execCell("b%d", i+2), v[1])
		file.SetCellValue(sheetName, execCell("c%d", i+2), v[4])
		file.SetCellValue(sheetName, execCell("d%d", i+2), v[2])
		file.SetCellValue(sheetName, execCell("e%d", i+2), v[3])
		t := int64(v[6].(float64))
		file.SetCellValue(sheetName, execCell("f%d", i+2), time.Unix(t/1000, 0).Format(TEMPLATE))
		file.SetCellValue(sheetName, execCell("g%d", i+2), "日线")
	}
	file.SetActiveSheet(sheetIndex)
}

//创建 okex Excel
func CreateOkexExcel(file *excelize.File, sheetName, symbol string, data [][]interface{}) {
	sheetIndex := file.NewSheet(sheetName)
	file.SetCellValue(sheetName, "a1", "交易对")
	file.SetCellValue(sheetName, "b1", "开盘价")
	file.SetCellValue(sheetName, "c1", "收盘价")
	file.SetCellValue(sheetName, "d1", "最高价")
	file.SetCellValue(sheetName, "e1", "最低价")
	file.SetCellValue(sheetName, "f1", "K线生成时间")
	file.SetCellValue(sheetName, "g1", "K线类型")
	for i, v := range data {
		file.SetCellValue(sheetName, execCell("a%d", i+2), symbol)
		file.SetCellValue(sheetName, execCell("b%d", i+2), v[1])
		file.SetCellValue(sheetName, execCell("c%d", i+2), v[4])
		file.SetCellValue(sheetName, execCell("d%d", i+2), v[2])
		file.SetCellValue(sheetName, execCell("e%d", i+2), v[3])
		file.SetCellValue(sheetName, execCell("f%d", i+2), strings.Split(v[0].(string), "T")[0])
		file.SetCellValue(sheetName, execCell("g%d", i+2), "日线")
	}
	file.SetActiveSheet(sheetIndex)
}

func execCell(format string, index int) string {
	return fmt.Sprintf(format, index)
}
