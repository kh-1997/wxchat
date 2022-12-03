package model

// GoodModel 计数器模型
type GoodModel struct {
	Id       int32   `gorm:"column:id" json:"id"`
	Tag      string  `gorm:"column:tag" json:"tag"`
	Title    string  `gorm:"column:title" json:"title"`
	Thumb    string  `gorm:"column:thumb" json:"thumb"`
	Spu      string  `gorm:"column:spu" json:"spu"`
	Remark   string  `gorm:"column:remark" json:"remark"`
	Code     string  `gorm:"column:code" json:"code"`
	Primary  string  `gorm:"column:primary" json:"primary"`
	Maxprice float64 `gorm:"column:price" json:"max_price"`
	Minprice float64 `gorm:"column:price" json:"min_price"`

}
