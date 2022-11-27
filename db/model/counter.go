package model

import "time"

// CounterModel 计数器模型
type CounterModel struct {
	Id           int32     `gorm:"column:id" json:"id"`
	Count        int32     `gorm:"column:count" json:"count"`
	CreatedAt    time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updatedAt" json:"updatedAt"`
	User         string    `gorm:"column:user" json:"user"`
	Product      string    `gorm:"column:product" json:"product"`
	Order        string    `gorm:"column:order" json:"order"`
	Title        string    `gorm:"column:title" json:"title"`
	Thumb        string    `gorm:"column:thumb" json:"thumb"`
	PrimaryImage string    `gorm:"column:primary_image" json:"primary_image"`
	Price        int       `gorm:"column:price" json:"price"`
}
