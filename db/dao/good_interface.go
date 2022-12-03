package dao

import (
	"wxcloudrun-golang/db/model"
)

// CounterInterface 计数器数据模型接口
type GoodInterface interface {
	GetGoodByName(name string) (*model.GoodModel, error)
	GetAllGood() ([]model.GoodModel, error)
    GetGoodByID(spu string) (*model.GoodModel, error)
}

// CounterInterfaceImp 计数器数据模型实现
type GoodInterfaceImp struct{}

// Imp 实现实例
var ImpGood GoodInterface = &GoodInterfaceImp{}

