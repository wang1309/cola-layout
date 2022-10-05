package api

import "github.com/wang1309/cola-layout/client/dto/data"

type CustomerServiceI interface {
	ListByName(name string) (*data.CustomerDto, error)
}
