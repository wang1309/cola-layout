package executor

import (
	"github.com/wang1309/cola-layout/client/dto/data"
	customer "github.com/wang1309/cola-layout/domain/consumer"
)

type CustomerServiceImpl struct {
	customer *customer.Customer
}

func NewCustomerServiceImpl() *CustomerServiceImpl {
	return &CustomerServiceImpl{
		customer: customer.NewCustomer(),
	}
}

func (c *CustomerServiceImpl) ListByName(name string) (*data.CustomerDto, error) {
	return &data.CustomerDto{CustomerName: c.customer.GetByName(name)}, nil
}
