package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const tableGoodName = "goods"

// GetCounter 查询Counter
func (imp *GoodInterfaceImp) GetGoodByName(name string) (*model.GoodModel, error) {
	var err error
	var counter = new(model.GoodModel)

	cli := db.Get()
	err = cli.Table(tableGoodName).Where("name like ?", "").First(counter).Error

	return counter, err
}

// GetCounter 查询Counter
func (imp *GoodInterfaceImp) GetAllGood() ([]model.GoodModel, error) {
	var err error
	var counter []model.GoodModel

	cli := db.Get()
	err = cli.Table(tableName).Find(&counter).Error

	return counter, err
}

