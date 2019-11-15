package model

type HuobiKline struct {
	Id     int64   `json:"id"`     //K线主键
	Open   float64 `json:"open"`   //开盘价
	Close  float64 `json:"close"`  //收盘价
	High   float64 `json:"high"`   //最高价
	Low    float64 `json:"low"`    //最低价
	Vol    float64 `json:"vol"`    //以报价币种计算的交易量
	Amount float64 `json:"amount"` //以基础币种计算的交易量
	Count  float64 `json:"count"`  //交易次数
}

type HuobiResponse struct {
	Status string       `json:"status"`
	Ch     string       `json:"ch"`
	Ts     int64        `json:"ts"`
	Data   []HuobiKline `json:"data"`
}
