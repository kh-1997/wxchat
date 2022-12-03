package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const tableName = "Counters"

// ClearCounter 清除Counter
func (imp *CounterInterfaceImp) ClearCounter(id int32) error {
	cli := db.Get()
	return cli.Table(tableName).Delete(&model.CounterModel{Id: id}).Error
}

// UpsertCounter 更新/写入counter
func (imp *CounterInterfaceImp) UpsertCounter(counter *model.CounterModel) error {
	cli := db.Get()
	return cli.Table(tableName).Save(counter).Error
}

// UpsertCounter 更新/写入counter
func (imp *CounterInterfaceImp) InsertCounter(counter *model.CounterModel) error {
	cli := db.Get()
	return cli.Table(tableName).Create(counter).Error
}

// GetCounter 查询Counter
func (imp *CounterInterfaceImp) GetCounter(id int32) (*model.CounterModel, error) {
	var err error
	var counter = new(model.CounterModel)

	cli := db.Get()
	err = cli.Table(tableName).Where("id = ?", id).First(counter).Error

	return counter, err
}

// GetCounter 查询Counter
func (imp *CounterInterfaceImp) GetOrder(name string) ([]model.CounterModel, error) {
	var err error
	var counter []model.CounterModel

	cli := db.Get()
	err = cli.Table(tableName).Where("user = ?", name).Order("id desc").Find(&counter).Error

	return counter, err
}

// GetCounter 查询Counter
func (imp *CounterInterfaceImp) GetOrderById(order string) (model.CounterModel, error) {
	var err error
	var counter model.CounterModel

	cli := db.Get()
	err = cli.Table(tableName).Where("order = ?", order).Find(&counter).Error

	return counter, err
}
