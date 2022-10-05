package customer

type customerGatewayImpl struct {
}

func NewCustomerGatewayImpl() *customerGatewayImpl {
	return &customerGatewayImpl{}
}

func (c *customerGatewayImpl) GetById(id string) string {
	return "customer"
}

func (c *customerGatewayImpl) GetByName(name string) string {
	return "dao impl" + name
}
