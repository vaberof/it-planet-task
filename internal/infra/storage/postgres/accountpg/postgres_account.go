package accountpg

import (
	"github.com/vaberof/it-planet-task/internal/infra/storage/postgres/animaltypepg"
	"github.com/vaberof/it-planet-task/internal/infra/storage/postgres/locationpg"
	"time"
)

type Account struct {
	Id          int32
	FirstName   string
	LastName    string
	Email       string
	Password    string
	Locations   []*locationpg.Location     `gorm:"foreignKey:ChipperId"`
	AnimalTypes []*animaltypepg.AnimalType `gorm:"foreignKey:ChipperId"`
	DateCreate  time.Time                  `gorm:"autoCreateTime"`
	UpdateDate  time.Time                  `gorm:"autoUpdateTime"`
	DeleteDate  time.Time                  `gorm:"autoDeleteTime" gorm:"index"`
}
