package locationpg

import (
	"time"
)

type Location struct {
	Id         int64
	Latitude   float64
	Longitude  float64
	ChipperId  int32
	DateCreate time.Time `gorm:"autoCreateTime"`
	UpdateDate time.Time `gorm:"autoUpdateTime"`
	DeleteDate time.Time `gorm:"autoDeleteTime" gorm:"index"`
}
