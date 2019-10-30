package main

import (
	"cryptocurrency_trading/bitflyer"
	"cryptocurrency_trading/config"
	"cryptocurrency_trading/utils"
	"fmt"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)

	order := &bitflyer.Order{
		ProductCode:     config.Config.ProductCode,
		ChildOrderType:  "LIMIT",
		Side:            "SELL",
		Price:           7000,
		Size:            0.01,
		MinuteToExpires: 1,
		TimeInForce:     "GTC",
	}
	res, _ := apiClient.SendOrder(order)
	fmt.Println(res.ChildOrderAcceptanceID)

	// i := "ChildOrderAcceptanceID"
	// params := map[string]string{
	// 	"product_code":              config.Config.ProductCode,
	// 	"child_order_acceptance_id": i,
	// }
	// r, _ := apiClient.ListOrder(params)
	// fmt.Println(r)
}
