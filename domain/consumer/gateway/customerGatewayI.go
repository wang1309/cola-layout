package gateway

/*
* 定义基础服务的抽象接口
 */

type CustomerGatewayI interface {
	GetById(id string) string
	GetByName(name string) string
}
