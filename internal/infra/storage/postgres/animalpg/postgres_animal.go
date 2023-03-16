package animalpg

import (
	"github.com/lib/pq"
	"github.com/vaberof/it-planet-task/internal/infra/storage/postgres/vstlocationpg"
	"time"
)

type Animal struct {
	Id                 int64
	AnimalTypes        pq.Int64Array `gorm:"type:integer[]"`
	Weight             float32
	Length             float32
	Height             float32
	Gender             string
	LifeStatus         string
	ChippingDateTime   time.Time `gorm:"autoCreateTime"`
	ChipperId          int32
	ChippingLocationId int64
	VisitedLocations   []*vstlocationpg.VisitedLocation
	DeathDateTime      *time.Time `gorm:"default:null"`
}
