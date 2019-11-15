package main

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"log"
	"net/http"
	"police-util/excel"
	"strings"
)

func main() {

	//var symbol = []string{"btcusdt", "ethusdt", "bchusdt", "eosusdt"}
	//var url = "http://api.hadax.com/market/history/kline?symbol=%s&period=1day&size=300"
	////创建 excel文件对象
	//file := excelize.NewFile()
	//for _, v := range symbol {
	//	huobiDayKline, err := http.Get(fmt.Sprintf(url, v))
	//	if err == nil {
	//		bytes, e := ioutil.ReadAll(huobiDayKline.Body)
	//		if e == nil {
	//			var huobiResponse model.HuobiResponse
	//			json.Unmarshal(bytes, &huobiResponse)
	//			excel.CreateExcel(file, strings.ToUpper(fmt.Sprintf("%s日线", v)), strings.ToUpper(v), huobiResponse.Data)
	//		}
	//
	//	}
	//}
	//file.SaveAs(fmt.Sprintf("./%s.xlsx", "火币数据"))
	//
	//var binanceSymbol = []string{"BTCUSDT", "ETHUSDT", "BCHABCUSDT", "EOSUSDT"}
	//var binanceUrl = "https://api.binance.com/api/v3/klines?symbol=%s&interval=1d"
	////创建excel对象
	//binanceFile := excelize.NewFile()
	//for _, v := range binanceSymbol {
	//
	//	dialSocketsProxy,err:=proxy.SOCKS5("tcp","127.0.0.1:1086",nil,proxy.Direct)
	//	if err!=nil{
	//		log.Println("代理出错",err)
	//	}
	//	tr:=&http.Transport{Dial:dialSocketsProxy.Dial}
	//	myClient:=&http.Client{Transport:tr}
	//	response,_:=myClient.Get(fmt.Sprintf(binanceUrl, v))
	//	bytes,_:=ioutil.ReadAll(response.Body)
	//	var data [][]interface{}
	//	json.Unmarshal(bytes,&data)
	//	excel.CreateBinanceExcel(binanceFile,strings.ToUpper(fmt.Sprintf("%s日线",v)),strings.ToUpper(v),data)
	//	binanceFile.SaveAs(fmt.Sprintf("./%s.xlsx","币安数据"))
	//
	//}

	var okexSymbol = []string{"BTC-USDT", "ETH-USDT", "BCH-USDT", "EOS-USDT"}
	var okexUrl = "https://www.okex.me/api/spot/v3/instruments/%s/candles?granularity=86400&start=2019-02-01T08:28:48.899Z&end=2019-06-30T09:28:48.899Z"
	//创建excel对象
	okexFile := excelize.NewFile()
	for _, v := range okexSymbol {

		dialSocketsProxy, err := proxy.SOCKS5("tcp", "127.0.0.1:1086", nil, proxy.Direct)
		if err != nil {
			log.Println("代理出错", err)
		}
		tr := &http.Transport{Dial: dialSocketsProxy.Dial}
		myClient := &http.Client{Transport: tr}
		response, _ := myClient.Get(fmt.Sprintf(okexUrl, v))
		bytes, _ := ioutil.ReadAll(response.Body)
		var data [][]interface{}
		json.Unmarshal(bytes, &data)
		excel.CreateOkexExcel(okexFile, strings.ToUpper(fmt.Sprintf("%s日线", v)), strings.ToUpper(v), data)
		okexFile.SaveAs(fmt.Sprintf("./%s.xlsx", "okex数据"))

	}

}
