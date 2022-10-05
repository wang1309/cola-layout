package customer

import "github.com/wang1309/cola-layout/application/customer/executor"

func GetName() string {
	customerService := executor.NewCustomerServiceImpl()
	res, _ := customerService.ListByName("WR")
	return res.CustomerName
}
