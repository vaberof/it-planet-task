package animaltypepg

import "time"

type AnimalType struct {
	Id         int64
	Type       string
	ChipperId  int32
	DateCreate time.Time `gorm:"autoCreateTime"`
	UpdateDate time.Time `gorm:"autoUpdateTime"`
	DeleteDate time.Time `gorm:"autoDeleteTime" gorm:"index"`
}
