package dao

import (
	"wxcloudrun-golang/db/model"
)

// CounterInterface 计数器数据模型接口
type CounterInterface interface {
	GetCounter(id int32) (*model.CounterModel, error)
	UpsertCounter(counter *model.CounterModel) error
	InsertCounter(counter *model.CounterModel) error
	ClearCounter(id int32) error
	GetOrder(name string) ([]model.CounterModel, error)
	GetOrderById(order string) (model.CounterModel, error)
}

// CounterInterfaceImp 计数器数据模型实现
type CounterInterfaceImp struct{}

// Imp 实现实例
var Imp CounterInterface = &CounterInterfaceImp{}
