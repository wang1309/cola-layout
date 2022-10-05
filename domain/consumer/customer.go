package customer

import (
	"github.com/wang1309/cola-layout/domain/consumer/gateway"
	customerDao "github.com/wang1309/cola-layout/infrastructure/repository/customer"
)

/*
* 只实现 customer 领域相关的功能，不具备跨领域的能力
 */

type Customer struct {
	registeredCapital   int64
	customerGatewayImpl gateway.CustomerGatewayI
}

func NewCustomer() *Customer {
	return &Customer{customerGatewayImpl: customerDao.NewCustomerGatewayImpl()}
}

func (c *Customer) SetRegisteredCapital(registeredCapital int64) {
	c.registeredCapital = registeredCapital
}

func (c *Customer) isBigCompany() bool {
	return c.registeredCapital > 10000000
}

func (c *Customer) GetByName(name string) string {
	return c.customerGatewayImpl.GetByName("customer service " + name)
}
